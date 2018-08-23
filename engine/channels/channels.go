package channels

import (
	"strings"
	"strconv"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"errors"
)

var channels = make(map[string]chan interface{})

// Count returns the number of channels
func Count() int {
	return len(channels)
}

// Add adds an engine channel, assumes these are created before startup
func Add(name string){
	//todo add size?
	idx := strings.Index(name,":")
	buffSize := 0
	chanName := name

	if idx > 0 {
		bSize, err:= strconv.Atoi(name[idx+1:])
		if err != nil {
			logger.Warnf("invalid channel buffer size '%s', defaulting to buffer size of %d", name[idx+1:], buffSize)
		} else {
			buffSize = bSize
		}

		chanName = name[:idx]
	}

	channels[chanName] = make(chan interface{}, buffSize)
}

// Get gets the named channel
func Get(name string) chan interface{} {
	return channels[name]
}

//Close closes all the channels, assumes it is called on shutdown
func Close()  {
	for _, value := range channels {
		close(value)
	}
	channels = make(map[string]chan interface{})
}

func Publish(channelName string, data interface{}) error {

	ch, exists := channels[channelName]
	if !exists {
		return errors.New("unknown channel: " + channelName)
	}

	ch <- data
	return nil
}

func PublishNoWait(channelName string, data interface{}) (bool, error) {

	ch, exists := channels[channelName]
	if !exists {
		return false, errors.New("unknown channel: " + channelName)
	}

	sent := false
	select {
	case ch <- data:
		sent = true
	default:
		sent = false
	}

	return sent, nil
}