package support

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)


// EmbeddedFlowManager is a simple manager for embedded flows
type EmbeddedFlowManager struct {
	flowsAreCompressed bool
	embeddedFlows      map[string]string
}

// NewEmbeddedFlowManager creates a new EmbeddedFlowManager
func NewEmbeddedFlowManager(compressed bool, embeddedFlows map[string]string) *EmbeddedFlowManager {

	flows := embeddedFlows

	if embeddedFlows == nil {
		flows = make(map[string]string)
	}

	return &EmbeddedFlowManager{flowsAreCompressed: compressed, embeddedFlows: flows}
}

// GetEmbeddedFlowJSON gets the specified embedded flow
func (mgr *EmbeddedFlowManager) GetEmbeddedFlowJSON(flowID string) []byte {

	flow, ok := mgr.embeddedFlows[flowID]

	if !ok {
		return nil
	}

	if !mgr.flowsAreCompressed {
		return []byte(flow)
	}

	flowBytes, err := decodeAndUnzip(flow)

	if err != nil {
		return nil
	}

	return flowBytes
}

func decodeAndUnzip(encoded string) ([]byte, error) {

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	buf := bytes.NewBuffer(decoded)
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	jsonAsBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return jsonAsBytes, nil
}
