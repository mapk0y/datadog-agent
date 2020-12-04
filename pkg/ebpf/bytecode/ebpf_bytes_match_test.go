// +build linux_bpf,ebpf_bindata

package bytecode

import (
	"fmt"
	"io/ioutil"
	"path"
	"testing"

	"github.com/DataDog/datadog-agent/pkg/ebpf"
	"github.com/DataDog/datadog-agent/pkg/ebpf/bytecode/bindata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEbpfBytesCorrect(t *testing.T) {
	cfg := ebpf.NewDefaultConfig()
	for _, filename := range bindata.AssetNames() {
		t.Run(filename, func(t *testing.T) {
			bs, err := ioutil.ReadFile(path.Join(cfg.BPFDir, filename))
			require.NoError(t, err)

			actualReader, err := GetReader(cfg.BPFDir, filename)
			require.NoError(t, err)
			defer actualReader.Close()

			actual, err := ioutil.ReadAll(actualReader)
			require.NoError(t, err)
			assert.Equal(t, bs, actual, fmt.Sprintf("on-disk file %s and bundled content are different", filename))
		})
	}
}
