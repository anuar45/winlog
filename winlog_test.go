package winlog

import (
	"fmt"
	"testing"
	"time"
)

func TestEvents(t *testing.T) {

	subs, err := Subscribe("Application", "")
	if err != nil {
		t.Error(err)
	}

	for {
		time.Sleep(3 * time.Second)
		eventHandles, err := FetchEvents(subs)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(eventHandles)

		for _, eventHandle := range eventHandles {
			if eventHandle != 0 {
				eventStr, err := RenderEvent(eventHandle)
				if err != nil {
					t.Error(err)
				}
				fmt.Println(eventStr)
			}
		}
	}
}
