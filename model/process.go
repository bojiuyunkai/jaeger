// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package model

// Process describes an instance of an application or service that emits tracing data.
type Process struct {
	ServiceName string    `json:"serviceName"`
	Tags        KeyValues `json:"tags,omitempty"`
}

// NewProcess creates a new Process for given serviceName and tags.
// The tags are sorted in place and kept in the the same array/slice,
// in order to store the Process in a canonical form that can be
// useful when comparing or generating a stable hash code.
func NewProcess(serviceName string, tags []KeyValue) *Process {
	typedTags := KeyValues(tags)
	typedTags.Sort()
	return &Process{ServiceName: serviceName, Tags: typedTags}
}

// Equal compares KeyValue object with another KeyValue.
func (p *Process) Equal(other *Process) bool {
	if p.ServiceName != other.ServiceName {
		return false
	}
	l1, l2 := len(p.Tags), len(other.Tags)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		if !p.Tags[i].Equal(&other.Tags[i]) {
			return false
		}
	}
	return true
}