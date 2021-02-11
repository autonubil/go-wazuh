package ossec

import (
	"fmt"
	"strings"
	"testing"
)

func TestWodle(t *testing.T) {
	sample := `DIRECTORY="/var/ossec"
NAME="Wazuh"
VERSION="v4.0.4"
REVISION="40011"
DATE="Wed Jan 13 14:53:46 UTC 2021"
TYPE="server"`

	reader := strings.NewReader(sample)
	info, err := ReadInitInfo(reader)
	if err != nil {
		t.Error(err)
		return
	}

	wodle, err := NewQueue("test", WithInitInfo(info))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%v\n", wodle)

	err = wodle.DebugMessage("Started")
	if err != nil {
		t.Error(err)
		return
	}

	err = wodle.DebugMessage("Closed")
	if err != nil {
		t.Error(err)
		return
	}

}
