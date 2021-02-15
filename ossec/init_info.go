package ossec

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/coreos/go-semver/semver"
	"github.com/google/martian/log"
)

var defaultInitFileLocation = "/etc/ossec-init.conf"

// LocalInitInfo contains the init info of the locally installed OSSEC
var LocalInitInfo *InitInfo

func init() {
	LocalInitInfo, _ = NewInitInfo()
}

type stringMap map[string]string

// InitInfo information gathered from ossec-init.conf
type InitInfo struct {
	stringMap
	Directory string          `json:"Directory"`
	Name      string          `json:"Name"`
	Version   *semver.Version `json:"Version"`
	Revision  uint            `json:"Revision"`
	Date      time.Time       `json:"Date"`
	Type      string          `json:"Type"`
}

// NewInitInfo read InitInfo from default location
func NewInitInfo() (*InitInfo, error) {
	file, err := os.Open(defaultInitFileLocation)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ReadInitInfo(file)
}

// ReadInitInfo read InitInfo from file
func ReadInitInfo(file io.Reader) (*InitInfo, error) {
	reSplit := regexp.MustCompile(`^([^=]+)="?([^"]+)"?$`)
	initInfo := InitInfo{
		stringMap: make(map[string]string),
	}
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
		txt := scanner.Text()
		matches := reSplit.FindStringSubmatch(txt)
		if len(matches) == 3 {
			switch matches[1] {
			case "DIRECTORY":
				initInfo.Directory = matches[2]
			case "TYPE":
				initInfo.Type = matches[2]
			case "NAME":
				initInfo.Name = matches[2]
			case "DATE":
				initInfo.Date, _ = time.Parse(time.UnixDate, matches[2])
			case "VERSION":
				if len(matches[2]) > 1 && matches[2][0] == 'v' {
					initInfo.Version = semver.New(matches[2][1:])
				} else {
					initInfo.Version = semver.New(matches[2])
				}
			case "REVISION":
				i, err := strconv.Atoi(matches[2])
				if err == nil {
					initInfo.Revision = uint(i)
				}
			}
			(initInfo).stringMap[matches[1]] = matches[2]
		} else {
			log.Debugf("illegal content in init file line %d (%s)", line, txt)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &initInfo, nil
}

// Get retreive raw data
func (i *InitInfo) Get(key string) (string, bool) {
	s, ok := i.stringMap[key]
	return s, ok
}

/**
DIRECTORY="/var/ossec"
NAME="Wazuh"
VERSION="v4.0.4"
REVISION="40011"
DATE="Wed Jan 13 14:53:46 UTC 2021"
TYPE="server"
*/
