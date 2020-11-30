// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build linux

package probe

const (
	dentryPathKeyNotFound = "error: dentry path key not found"
)

// ErrTruncatedSegment is used to notify that a segment of the path was truncated because it was too long
type ErrTruncatedSegment struct {}

func (err ErrTruncatedSegment) Error() string {
	return "truncated_segment"
}

// ErrTruncatedParents is used to notify that some parents of the path are missing
type ErrTruncatedParents struct {}

func (err ErrTruncatedParents) Error() string {
	return "truncated_parents"
}

// NewDentryResolver returns a new dentry resolver
func NewDentryResolver(probe *Probe) (*DentryResolver, error) {
	return &DentryResolver{
		probe: probe,
	}, nil
}
