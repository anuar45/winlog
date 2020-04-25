package winlog

import "syscall"

// Event log error codes.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms681382(v=vs.85).aspx
const (
	ERROR_INSUFFICIENT_BUFFER syscall.Errno = 122
	ERROR_NO_MORE_ITEMS       syscall.Errno = 259
	RPC_S_INVALID_BOUND       syscall.Errno = 1734
	ERROR_INVALID_OPERATION   syscall.Errno = 4317
)

// EvtSubscribeFlag defines the possible values that specify when to start subscribing to events.
type EvtSubscribeFlag uint32

// EVT_SUBSCRIBE_FLAGS enumeration
// https://msdn.microsoft.com/en-us/library/windows/desktop/aa385588(v=vs.85).aspx
const (
	EvtSubscribeToFutureEvents      EvtSubscribeFlag = 1
	EvtSubscribeStartAtOldestRecord EvtSubscribeFlag = 2
	EvtSubscribeStartAfterBookmark  EvtSubscribeFlag = 3
	EvtSubscribeOriginMask          EvtSubscribeFlag = 0x3
	EvtSubscribeTolerateQueryErrors EvtSubscribeFlag = 0x1000
	EvtSubscribeStrict              EvtSubscribeFlag = 0x10000
)

type EvtRenderFlag uint32

// EVT_RENDER_FLAGS enumeration
// https://msdn.microsoft.com/en-us/library/windows/desktop/aa385563(v=vs.85).aspx
const (
	// Render the event properties specified in the rendering context.
	EvtRenderEventValues EvtRenderFlag = iota
	// Render the event as an XML string. For details on the contents of the
	// XML string, see the Event schema.
	EvtRenderEventXml
	// Render the bookmark as an XML string, so that you can easily persist the
	// bookmark for use later.
	EvtRenderBookmark
)
