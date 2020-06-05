package gst

/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"
*/
import "C"

type Event struct {
	event *C.GstEvent
}

type EventType C.GstEventType

const (
	EventUnknown                EventType = C.GST_EVENT_UNKNOWN
	EventFlushStart             EventType = C.GST_EVENT_FLUSH_START
	EventFlushStop              EventType = C.GST_EVENT_FLUSH_STOP
	EventStreamStart            EventType = C.GST_EVENT_STREAM_START
	EventCaps                   EventType = C.GST_EVENT_CAPS
	EventSegment                EventType = C.GST_EVENT_SEGMENT
	EventTag                    EventType = C.GST_EVENT_TAG
	EventBuffersize             EventType = C.GST_EVENT_BUFFERSIZE
	EventSinkMessage            EventType = C.GST_EVENT_SINK_MESSAGE
	EventEos                    EventType = C.GST_EVENT_EOS
	EventToc                    EventType = C.GST_EVENT_TOC
	EventSegmentDone            EventType = C.GST_EVENT_SEGMENT_DONE
	EventGap                    EventType = C.GST_EVENT_GAP
	EventQos                    EventType = C.GST_EVENT_QOS
	EventSeek                   EventType = C.GST_EVENT_SEEK
	EventNavigation             EventType = C.GST_EVENT_NAVIGATION
	EventLatency                EventType = C.GST_EVENT_LATENCY
	EventStep                   EventType = C.GST_EVENT_STEP
	EventReconfigure            EventType = C.GST_EVENT_RECONFIGURE
	EventTocSelect              EventType = C.GST_EVENT_TOC_SELECT
	EventCustomUpstream         EventType = C.GST_EVENT_CUSTOM_UPSTREAM
	EventCustomDownstream       EventType = C.GST_EVENT_CUSTOM_DOWNSTREAM
	EventCustomDownstreamOob    EventType = C.GST_EVENT_CUSTOM_DOWNSTREAM_OOB
	EventCustomDownstreamSticky EventType = C.GST_EVENT_CUSTOM_DOWNSTREAM_STICKY
	EventCustomBoth             EventType = C.GST_EVENT_CUSTOM_BOTH
	EventCustomBothOob          EventType = C.GST_EVENT_CUSTOM_BOTH_OOB
)

func (e *Event) GetType() EventType {
	ctype := C.X_GST_EVENT_TYPE(e.event)
	t := EventType(ctype)

	return t
}
