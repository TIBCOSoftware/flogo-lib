package channels

var channels = make(map[string]chan interface{})

// Count returns the number of channels
func Count() int {
	return len(channels)
}

// Add adds an engine channel, assumes these are created before startup
func Add(name string){
	//todo add size?
	channels[name] = make(chan interface{})
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
}