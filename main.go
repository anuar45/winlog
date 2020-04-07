package main

import (
	//winsys "github.com/elastic/beats/winlogbeat/sys"

	"bytes"
	"fmt"

	winsys "github.com/elastic/beats/winlogbeat/sys"
	win "github.com/elastic/beats/winlogbeat/sys/wineventlog"
	"golang.org/x/sys/windows"
	//"github.com/elastic/beats/vendor/golang.org/x/sys/windows"
)

const renderBufferSize = 1 << 14

//type Events []winsys.Event //`xml:"Event"`

func main() {
	buf := make([]byte, renderBufferSize)
	out := new(bytes.Buffer)
	evtChannelName := "Application"
	var bookmark win.EvtHandle
	var lastRecID uint64
	var query string

	signalEvent, err := windows.CreateEvent(nil, 0, 0, nil)
	if err != nil {
		panic(err)
	}
	defer windows.CloseHandle(signalEvent)

	events, err := win.EvtQuery(0, "Application", query, win.EvtQueryChannelPath|win.EvtQueryReverseDirection)

	lastEvents, err := win.EventHandles(events, 1)
	if err != nil {
		panic(err)
	}

	err = win.RenderEventXML(lastEvents[0], buf, out)
	if err != nil {
		panic(err)
	}

	lastEvent, _ := winsys.UnmarshalEventXML(out.Bytes())

	lastRecID = lastEvent.RecordID
	//fmt.Println(out.String())

	for {

		bookmark, err = win.CreateBookmarkFromRecordID(evtChannelName, lastRecID)
		if err != nil {
			panic(err)
		}

		subs, err := win.Subscribe(0, signalEvent, "Application", query, bookmark, win.EvtSubscribeStartAfterBookmark)
		if err != nil {
			panic(err)
		}

		eventHandles, err := win.EventHandles(subs, 5)
		if err != nil {
			panic(err)
		}

		for _, eventRaw := range eventHandles {
			out.Reset()
			err := win.RenderEventXML(eventRaw, buf, out)
			if err != nil {
				panic(err)
			}

			evt, _ := winsys.UnmarshalEventXML(out.Bytes())

			lastRecID = evt.RecordID
			fmt.Println(out.String())
			fmt.Println(evt)
		}

		fmt.Scanln()
	}

	//fmt.Println(out.String())

	//event, _ := winsys.UnmarshalEventXML(out.Bytes())

	//fmt.Println(event)

	//ioutil.WriteFile("c:\\temp\\winlog.xml", byte(events), os.FileMode(777))
}
