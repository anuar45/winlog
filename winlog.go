package winlog

import (
	"syscall"

	"golang.org/x/sys/windows"
)

func Subscribe(logName, xquery string) (EvtHandle, error) {
	sigEvent, _ := windows.CreateEvent(nil, 0, 0, nil)
	var ecp, qp *uint16
	ecp, _ = syscall.UTF16PtrFromString(logName)
	//qp, _ = syscall.UTF16PtrFromString(xquery)
	subsHandle, err := _EvtSubscribe(0, uintptr(sigEvent), ecp, qp, 0, 0, 0, EvtSubscribeToFutureEvents)
	if err != nil {
		return 0, err
	}

	return subsHandle, nil
}

func FetchEvents(subsHandle EvtHandle) ([]EvtHandle, error) {
	var eventsNumber uint32
	var evtReturned uint32

	eventsNumber = 5

	eventHandles := make([]EvtHandle, eventsNumber)

	err := _EvtNext(subsHandle, eventsNumber, &eventHandles[0], 0, 0, &evtReturned)
	if err != nil {
		return nil, err
	}

	return eventHandles[:evtReturned], nil
}

func RenderEvent(e EvtHandle) (string, error) {
	bufferSize := 1 << 14
	renderBuffer := make([]byte, bufferSize)
	var bufferUsed, propertyCount uint32
	var result string
	err := _EvtRender(0, e, EvtRenderEventXml, uint32(len(renderBuffer)), &renderBuffer[0], &bufferUsed, &propertyCount)
	if err != nil {
		return result, err
	}

	result, err = DecodeUTF16(renderBuffer[:bufferUsed])

	return result, err
}
