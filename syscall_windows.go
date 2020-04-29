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
	EvtSubscribeToFutureEvents EvtSubscribeFlag = 1
)

type EvtRenderFlag uint32

// EVT_RENDER_FLAGS enumeration
// https://msdn.microsoft.com/en-us/library/windows/desktop/aa385563(v=vs.85).aspx
const (
	EvtRenderEventXml = 1
)

// EvtQueryFlag defines the values that specify how to return the query results
// and whether you are query against a channel or log file.
type EvtQueryFlag uint32

const (
	// EvtQueryForwardDirection specifies that the events in the query result
	// are ordered from oldest to newest. This is the default.
	EvtQueryForwardDirection EvtQueryFlag = 0x100
	// EvtQueryReverseDirection specifies that the events in the query result
	// are ordered from newest to oldest.
	EvtQueryReverseDirection EvtQueryFlag = 0x200
	// EvtQueryTolerateQueryErrors specifies that EvtQuery should run the query
	// even if the part of the query generates an error (is not well formed).
	EvtQueryTolerateQueryErrors EvtQueryFlag = 0x1000
)
