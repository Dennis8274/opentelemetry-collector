// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/data"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestSpanLink_MoveTo(t *testing.T) {
	ms := generateTestSpanLink()
	dest := NewSpanLink()
	ms.MoveTo(dest)
	assert.Equal(t, NewSpanLink(), ms)
	assert.Equal(t, generateTestSpanLink(), dest)
}

func TestSpanLink_CopyTo(t *testing.T) {
	ms := NewSpanLink()
	orig := NewSpanLink()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestSpanLink()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestSpanLink_TraceID(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, pcommon.TraceID(data.TraceID([16]byte{})), ms.TraceID())
	testValTraceID := pcommon.TraceID(data.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetTraceID(testValTraceID)
	assert.Equal(t, testValTraceID, ms.TraceID())
}

func TestSpanLink_SpanID(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, pcommon.SpanID(data.SpanID([8]byte{})), ms.SpanID())
	testValSpanID := pcommon.SpanID(data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetSpanID(testValSpanID)
	assert.Equal(t, testValSpanID, ms.SpanID())
}

func TestSpanLink_TraceState(t *testing.T) {
	ms := NewSpanLink()
	internal.FillTestTraceState(internal.TraceState(ms.TraceState()))
	assert.Equal(t, pcommon.TraceState(internal.GenerateTestTraceState()), ms.TraceState())
}

func TestSpanLink_Attributes(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, pcommon.NewMap(), ms.Attributes())
	internal.FillTestMap(internal.Map(ms.Attributes()))
	assert.Equal(t, pcommon.Map(internal.GenerateTestMap()), ms.Attributes())
}

func TestSpanLink_DroppedAttributesCount(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, uint32(0), ms.DroppedAttributesCount())
	ms.SetDroppedAttributesCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedAttributesCount())
}

func generateTestSpanLink() SpanLink {
	tv := NewSpanLink()
	fillTestSpanLink(tv)
	return tv
}

func fillTestSpanLink(tv SpanLink) {
	tv.orig.TraceId = data.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	tv.orig.SpanId = data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1})
	internal.FillTestTraceState(internal.NewTraceState(&tv.orig.TraceState))
	internal.FillTestMap(internal.NewMap(&tv.orig.Attributes))
	tv.orig.DroppedAttributesCount = uint32(17)
}