package winlog

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestEvents(t *testing.T) {

	f, err := os.Create("c:\\temp\\events.log")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	subs, err := Subscribe("Application", "")
	if err != nil {
		t.Error(err)
	}

	for {
	loop:
		for {
			time.Sleep(3 * time.Second)
			eventHandles, err := FetchEvents(subs)
			if err != nil {
				switch err {
				case ERROR_NO_MORE_ITEMS:
					fmt.Println("check no more items")
					break loop
				default:
					t.Error(err)
				}
			}
			fmt.Println(eventHandles)

			for _, eventHandle := range eventHandles {
				if eventHandle != 0 {
					eventXML, err := RenderEvent(eventHandle)
					if err != nil {
						t.Error(err)
					}

					w.WriteString(string(eventXML) + "\n")

					event := Event{}
					xml.Unmarshal(eventXML, &event)
					fmt.Println(event)

					w.WriteString(fmt.Sprintln(event))
				}
			}
			w.Flush()
		}

	}
}
