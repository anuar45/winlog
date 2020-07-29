// +build windows

package winlog

import (
	"encoding/xml"
	"syscall"

	"golang.org/x/sys/windows"
)

var bufferSize = 1 << 14

type Winlog struct {
	subs EvtHandle
	buf  []byte
}

func Subscribe(logName, xquery string) (*Winlog, error) {
	var w Winlog

	w.subs, err = subscribe()
	if err != nil {
		return nil, err
	}

	w.buf = make([]byte, bufferSize)

	return &w, nil
}

func subscribe(logName, xquery string) (EvtHandle, error) {
	var logNamePtr, xqueryPtr *uint16

	sigEvent, err := windows.CreateEvent(nil, 0, 0, nil)
	if err != nil {
		return 0, err
	}
	defer windows.CloseHandle(sigEvent)

	logNamePtr, err = syscall.UTF16PtrFromString(logName)
	if err != nil {
		return 0, err
	}

	xqueryPtr, err = syscall.UTF16PtrFromString(xquery)
	if err != nil {
		return 0, err
	}

	subsHandle, err := _EvtSubscribe(0, uintptr(sigEvent), logNamePtr, xqueryPtr, 0, 0, 0, EvtSubscribeToFutureEvents)
	if err != nil {
		return 0, err
	}

	return subsHandle, nil
}

func fetchEventHandles(subsHandle EvtHandle) ([]EvtHandle, error) {
	var eventsNumber uint32
	var evtReturned uint32

	eventsNumber = 5

	eventHandles := make([]EvtHandle, eventsNumber)

	err := _EvtNext(subsHandle, eventsNumber, &eventHandles[0], 0, 0, &evtReturned)
	if err != nil {
		if err == ERROR_INVALID_OPERATION && evtReturned == 0 {
			return nil, ERROR_NO_MORE_ITEMS
		}
		return nil, err
	}

	return eventHandles[:evtReturned], nil
}

func (w *Winlog) Fetch() ([]Event, error) {
	var events []Event

	eventHandles, err := fetchEventHandles(subsHandle)
	if err != nil {
		return nil, err
	}

	for _, eventHandle := range eventHandles {
		if eventHandle != 0 {
			eventXML, err := w.renderEvent(eventHandle)
			if err != nil {
				return nil, err
			}

			event := Event{}
			xml.Unmarshal(eventXML, &event)

			events = append(events, event)
		}
	}

	for i := 0; i < len(eventHandles); i++ {
		err := CloseEvent(eventHandles[i])
		if err != nil {
			return events, err
		}
	}
	return events, nil
}

func (w *Winlog) renderEvent(e EvtHandle) ([]byte, error) {
	var bufferUsed, propertyCount uint32

	err := _EvtRender(0, e, EvtRenderEventXml, uint32(len(renderBuffer)), &w.buf[0], &bufferUsed, &propertyCount)
	if err != nil {
		return nil, err
	}

	return DecodeUTF16(w.buf[:bufferUsed])
}

func queryEventHandles(logName, xquery string) ([]EvtHandle, error) {
	// TODO
	return nil, nil
}

func Query(logName, xquery string) ([]Event, error) {
	// TODO
	return nil, nil
}

func CloseEvent(e EvtHandle) error {
	err := _EvtClose(e)
	if err != nil {
		return err
	}
	return nil
}
