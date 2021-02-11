package ossec

import (
	"fmt"
	"strings"
	"testing"
)

func TestInitInfo(t *testing.T) {
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

	d, ok := info.Get("DATE")
	fmt.Printf("%s %t\n", d, ok)
	fmt.Printf("%v\n", info)
}
