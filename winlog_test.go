package winlog

import (
	"fmt"
	"testing"
)

func TestEvents(t *testing.T) {

	subs, err := Subscribe("Application", "")
	if err != nil {
		t.Error(err)
	}

	for {
		events, err := FetchEvents(subs)
		if err != nil {
			switch err {
			case ERROR_NO_MORE_ITEMS:
				fmt.Println("no more events")
			default:
				t.Error(err)
			}
		}
		fmt.Println(events)
	}
}
