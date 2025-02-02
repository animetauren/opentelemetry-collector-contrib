// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testdata

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var (
<<<<<<< HEAD:internal/coreinternal/testdata/trace.go
	TestSpanStartTime      = time.Date(2020, 2, 11, 20, 26, 12, 321, time.UTC)
	TestSpanStartTimestamp = pcommon.NewTimestampFromTime(TestSpanStartTime)

	TestSpanEventTime      = time.Date(2020, 2, 11, 20, 26, 13, 123, time.UTC)
	TestSpanEventTimestamp = pcommon.NewTimestampFromTime(TestSpanEventTime)

	TestSpanEndTime      = time.Date(2020, 2, 11, 20, 26, 13, 789, time.UTC)
	TestSpanEndTimestamp = pcommon.NewTimestampFromTime(TestSpanEndTime)
)

func GenerateTracesOneEmptyResourceSpans() ptrace.Traces {
	td := ptrace.NewTraces()
	td.ResourceSpans().AppendEmpty()
	return td
}

func GenerateTracesNoLibraries() ptrace.Traces {
	td := GenerateTracesOneEmptyResourceSpans()
	rs0 := td.ResourceSpans().At(0)
	initResource1(rs0.Resource())
	return td
}

func GenerateTracesOneEmptyInstrumentationLibrary() ptrace.Traces {
	td := GenerateTracesNoLibraries()
	td.ResourceSpans().At(0).ScopeSpans().AppendEmpty()
	return td
}

func GenerateTracesOneSpanNoResource() ptrace.Traces {
	td := GenerateTracesOneEmptyResourceSpans()
	rs0 := td.ResourceSpans().At(0)
	fillSpanOne(rs0.ScopeSpans().AppendEmpty().Spans().AppendEmpty())
	return td
}

func GenerateTracesOneSpan() ptrace.Traces {
	td := GenerateTracesOneEmptyInstrumentationLibrary()
	rs0ils0 := td.ResourceSpans().At(0).ScopeSpans().At(0)
	fillSpanOne(rs0ils0.Spans().AppendEmpty())
	return td
}

func GenerateTracesTwoSpansSameResource() ptrace.Traces {
	td := GenerateTracesOneEmptyInstrumentationLibrary()
	rs0ils0 := td.ResourceSpans().At(0).ScopeSpans().At(0)
	fillSpanOne(rs0ils0.Spans().AppendEmpty())
	fillSpanTwo(rs0ils0.Spans().AppendEmpty())
	return td
}

func GenerateTracesTwoSpansSameResourceOneDifferent() ptrace.Traces {
	td := ptrace.NewTraces()
	rs0 := td.ResourceSpans().AppendEmpty()
	initResource1(rs0.Resource())
	rs0ils0 := rs0.ScopeSpans().AppendEmpty()
	fillSpanOne(rs0ils0.Spans().AppendEmpty())
	fillSpanTwo(rs0ils0.Spans().AppendEmpty())
	rs1 := td.ResourceSpans().AppendEmpty()
	initResource2(rs1.Resource())
	rs1ils0 := rs1.ScopeSpans().AppendEmpty()
	fillSpanThree(rs1ils0.Spans().AppendEmpty())
	return td
}

func GenerateTracesManySpansSameResource(spanCount int) ptrace.Traces {
	td := GenerateTracesOneEmptyInstrumentationLibrary()
	rs0ils0 := td.ResourceSpans().At(0).ScopeSpans().At(0)
	rs0ils0.Spans().EnsureCapacity(spanCount)
=======
	spanStartTimestamp = pcommon.NewTimestampFromTime(time.Date(2020, 2, 11, 20, 26, 12, 321, time.UTC))
	spanEventTimestamp = pcommon.NewTimestampFromTime(time.Date(2020, 2, 11, 20, 26, 13, 123, time.UTC))
	spanEndTimestamp   = pcommon.NewTimestampFromTime(time.Date(2020, 2, 11, 20, 26, 13, 789, time.UTC))
)

func GenerateTraces(spanCount int) ptrace.Traces {
	td := ptrace.NewTraces()
	initResource(td.ResourceSpans().AppendEmpty().Resource())
	ss := td.ResourceSpans().At(0).ScopeSpans().AppendEmpty().Spans()
	ss.EnsureCapacity(spanCount)
>>>>>>> upstream/main:internal/testdata/trace.go
	for i := 0; i < spanCount; i++ {
		switch i % 2 {
		case 0:
			fillSpanOne(ss.AppendEmpty())
		case 1:
			fillSpanTwo(ss.AppendEmpty())
		}
	}
	return td
}

func fillSpanOne(span ptrace.Span) {
	span.SetName("operationA")
	span.SetStartTimestamp(spanStartTimestamp)
	span.SetEndTimestamp(spanEndTimestamp)
	span.SetDroppedAttributesCount(1)
	span.SetTraceID([16]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10})
	span.SetSpanID([8]byte{0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18})
	evs := span.Events()
	ev0 := evs.AppendEmpty()
	ev0.SetTimestamp(spanEventTimestamp)
	ev0.SetName("event-with-attr")
	ev0.Attributes().PutStr("span-event-attr", "span-event-attr-val")
	ev0.SetDroppedAttributesCount(2)
	ev1 := evs.AppendEmpty()
	ev1.SetTimestamp(spanEventTimestamp)
	ev1.SetName("event")
	ev1.SetDroppedAttributesCount(2)
	span.SetDroppedEventsCount(1)
	status := span.Status()
	status.SetCode(ptrace.StatusCodeError)
	status.SetMessage("status-cancelled")
}

func fillSpanTwo(span ptrace.Span) {
	span.SetName("operationB")
	span.SetStartTimestamp(spanStartTimestamp)
	span.SetEndTimestamp(spanEndTimestamp)
	link0 := span.Links().AppendEmpty()
	link0.Attributes().PutStr("span-link-attr", "span-link-attr-val")
	link0.SetDroppedAttributesCount(4)
	link1 := span.Links().AppendEmpty()
	link1.SetDroppedAttributesCount(4)
	span.SetDroppedLinksCount(3)
}
<<<<<<< HEAD:internal/coreinternal/testdata/trace.go

func fillSpanThree(span ptrace.Span) {
	span.SetName("operationC")
	span.SetStartTimestamp(TestSpanStartTimestamp)
	span.SetEndTimestamp(TestSpanEndTimestamp)
	span.Attributes().PutStr("span-attr", "span-attr-val")
	span.SetDroppedAttributesCount(5)
}
=======
>>>>>>> upstream/main:internal/testdata/trace.go
