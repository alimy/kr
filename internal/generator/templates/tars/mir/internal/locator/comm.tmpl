package locator

import (
	"sync"

	"github.com/TarsCloud/TarsGo/tars"
)

var (
	curLocator string
	comm       *tars.Communicator
	initOnce   sync.Once
)

// InitLocator initial communicator with locator
func InitLocator(locator string) {
	if locator != "" {
		curLocator = locator
		if comm != nil {
			comm.SetLocator(locator)
		}
	}
}

// MyCommunicator return a communicator instance
func MyCommunicator() *tars.Communicator {
	initOnce.Do(func() {
		comm = tars.NewCommunicator()
		if curLocator != "" {
			comm.SetLocator(curLocator)
		}
	})
	return comm
}
