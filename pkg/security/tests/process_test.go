// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build functionaltests

package tests

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"testing"
	"time"

	"github.com/pkg/errors"

	"github.com/DataDog/datadog-agent/pkg/security/probe"
	"github.com/DataDog/datadog-agent/pkg/security/rules"
)

func TestProcess(t *testing.T) {
	currentUser, err := user.LookupId("0")
	if err != nil {
		t.Fatal(err)
	}

	executable, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}

	ruleDef := &rules.RuleDefinition{
		ID:         "test_rule",
		Expression: fmt.Sprintf(`process.user == "%s" && process.name == "%s" && open.filename == "{{.Root}}/test-process"`, currentUser.Name, path.Base(executable)),
	}

	test, err := newTestModule(nil, []*rules.RuleDefinition{ruleDef}, testOpts{})
	if err != nil {
		t.Fatal(err)
	}
	defer test.Close()

	testFile, _, err := test.Path("test-process")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create(testFile)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFile)

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	_, rule, err := test.GetEvent()
	if err != nil {
		t.Error(err)
	} else {
		if rule.ID != "test_rule" {
			t.Errorf("expected rule 'test-rule' to be triggered, got %s", rule.ID)
		}
	}
}

func TestProcessContext(t *testing.T) {
	ruleDef := &rules.RuleDefinition{
		ID:         "test_rule",
		Expression: fmt.Sprintf(`open.filename == "{{.Root}}/test-process-context" && open.flags & O_CREAT == 0`),
	}

	test, err := newTestModule(nil, []*rules.RuleDefinition{ruleDef}, testOpts{})
	if err != nil {
		t.Fatal(err)
	}
	defer test.Close()

	testFile, _, err := test.Path("test-process-context")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create(testFile)
	if err != nil {
		t.Fatal(err)
	}

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFile)

	t.Run("inode", func(t *testing.T) {
		executable, err := os.Executable()
		if err != nil {
			t.Fatal(err)
		}

		f, err := os.Open(testFile)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		event, _, err := test.GetEvent()
		if err != nil {
			t.Error(err)
		} else {
			if filename, _ := event.GetFieldValue("process.filename"); filename.(string) != executable {
				t.Errorf("not able to find the proper process filename `%v` vs `%s`", filename, executable)
			}

			// not working on centos8 in docker env
			/*if inode := getInode(t, executable); inode != event.Process.Inode {
				t.Errorf("expected inode %d, got %d", event.Process.Inode, inode)
			}*/

			testContainerPath(t, event, "process.container_path")
		}
	})

	t.Run("tty", func(t *testing.T) {
		// not working on centos8
		t.Skip()

		executable := "/usr/bin/cat"
		if _, err := os.Stat(executable); err != nil {
			executable = "/bin/cat"
		}

		cmd := exec.Command("script", "/dev/null", "-c", executable+" "+testFile)
		if _, err := cmd.CombinedOutput(); err != nil {
			t.Error(err)
		}

		event, _, err := test.GetEvent()
		if err != nil {
			t.Error(err)
		} else {
			if filename, _ := event.GetFieldValue("process.filename"); filename.(string) != executable {
				t.Errorf("not able to find the proper process filename `%v` vs `%s`", filename, executable)
			}

			if name, _ := event.GetFieldValue("process.tty_name"); name.(string) == "" {
				t.Error("not able to get a tty name")
			}

			if inode := getInode(t, executable); inode != event.Process.Inode {
				t.Errorf("expected inode %d, got %d", event.Process.Inode, inode)
			}

			testContainerPath(t, event, "process.container_path")
		}
	})
}

func TestProcessExec(t *testing.T) {
	executable := "/usr/bin/touch"
	if resolved, err := os.Readlink(executable); err == nil {
		executable = resolved
	} else {
		if os.IsNotExist(err) {
			executable = "/bin/touch"
		}
	}

	ruleDef := &rules.RuleDefinition{
		ID:         "test_rule",
		Expression: fmt.Sprintf(`exec.filename == "%s"`, executable),
	}

	test, err := newTestModule(nil, []*rules.RuleDefinition{ruleDef}, testOpts{})
	if err != nil {
		t.Fatal(err)
	}
	defer test.Close()

	cmd := exec.Command("sh", "-c", executable+" /dev/null")
	if _, err := cmd.CombinedOutput(); err != nil {
		t.Error(err)
	}

	event, _, err := test.GetEvent()
	if err != nil {
		t.Error(err)
	} else {
		if filename, _ := event.GetFieldValue("exec.filename"); filename.(string) != executable {
			t.Errorf("expected exec filename `%v`, got `%v`", executable, filename)
		}

		if filename, _ := event.GetFieldValue("process.filename"); filename.(string) != executable {
			t.Errorf("expected process filename `%v`, got `%v`", executable, filename)
		}

		testContainerPath(t, event, "exec.container_path")
	}
}

func TestProcessLineage(t *testing.T) {
	executable := "/usr/bin/touch"
	if resolved, err := os.Readlink(executable); err == nil {
		executable = resolved
	} else {
		if os.IsNotExist(err) {
			executable = "/bin/touch"
		}
	}

	rule := &rules.RuleDefinition{
		ID:         "test_rule",
		Expression: fmt.Sprintf(`exec.filename == "%s"`, executable),
	}

	test, err := newTestModule(nil, []*rules.RuleDefinition{rule}, testOpts{wantProbeEvents: true})
	if err != nil {
		t.Fatal(err)
	}
	defer test.Close()

	cmd := exec.Command(executable, "/dev/null")
	if _, err := cmd.CombinedOutput(); err != nil {
		t.Error(err)
	}

	t.Run("fork", func(t *testing.T) {
		event, err := test.GetProbeEvent(3*time.Second, "fork")
		if err != nil {
			t.Error(err)
		} else {
			if err := testProcessLineageFork(t, event); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("exec", func(t *testing.T) {
		event, _, err := test.GetEvent()
		if err != nil {
			t.Error(err)
		} else {
			if err := testProcessLineageExec(t, event); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("exit", func(t *testing.T) {
		event, err := test.GetProbeEvent(3*time.Second, "exit")
		if err != nil {
			t.Error(err)
		} else {
			if err := testProcessLineageExit(t, event, test); err != nil {
				t.Error(err)
			}
		}
	})
}

func testProcessLineageExec(t *testing.T, event *probe.Event) error {
	// check for the new process context
	cacheEntry := event.ResolveProcessCacheEntry()
	if cacheEntry == nil {
		t.Errorf("expected a process cache entry, got nil")
	} else {
		// make sure the container ID was properly inherited from the parent
		if cacheEntry.Parent == nil {
			t.Errorf("expected a parent, got nil")
		} else {
			if cacheEntry.ID != cacheEntry.Parent.ID {
				t.Errorf("expected container ID %s, got %s", cacheEntry.Parent.ID, cacheEntry.ID)
			}
		}
	}

	testContainerPath(t, event, "process.container_path")
	return nil
}

func testProcessLineageFork(t *testing.T, event *probe.Event) error {
	// we need to make sure that the child entry if properly populated with its parent metadata
	newEntry := event.ResolveProcessCacheEntry()
	if newEntry == nil {
		return errors.Errorf("expected a new process cache entry, got nil")
	} else {
		// fetch the parent of the new entry, it should the test binary itself
		parentEntry := newEntry.Parent

		if parentEntry == nil {
			return errors.Errorf("expected a parent cache entry, got nil")
		} else {
			// checking cookie and pathname str should be enough to make sure that the metadata were properly
			// copied from kernel space (those 2 information are stored in 2 different maps)
			if newEntry.Cookie != parentEntry.Cookie {
				return errors.Errorf("expected cookie %d, %d", parentEntry.Cookie, newEntry.Cookie)
			}
			if newEntry.PathnameStr != parentEntry.PathnameStr {
				return errors.Errorf("expected PathnameStr %s, got %s", parentEntry.PathnameStr, newEntry.PathnameStr)
			}

			// we also need to check the container ID lineage
			if newEntry.ID != parentEntry.ID {
				return errors.Errorf("expected container ID %s, got %s", parentEntry.ID, newEntry.ID)
			}

			// We can't check that the new entry is in the list of the children of its parent because the exit event
			// has probably already been processed (thus the parent list of children has already been updated and the
			// child entry deleted).
		}

		testContainerPath(t, event, "process.container_path")
	}
	return nil
}

func testProcessLineageExit(t *testing.T, event *probe.Event, test *testModule) error {
	// check for the new process context
	cacheEntry := event.ResolveProcessCacheEntry()
	if cacheEntry == nil {
		return errors.Errorf("expected a process cache entry, got nil")
	}

	// make sure that the process cache entry of the process was properly deleted from the cache
	resolvers := test.probe.GetResolvers()
	entry := resolvers.ProcessResolver.Get(event.Process.Pid)
	if entry != nil {
		return errors.Errorf("the process cache entry was not deleted from the user space cache")
	}
	return nil
}
