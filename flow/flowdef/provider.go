package flowdef

// Provider is the interface that describes an object
// that can provide flow definitions from a URI
type Provider interface {

	// GetFlow retrieves the flow definition for the specified URI
	GetFlow(flowURI string) (*Definition, error)
}
