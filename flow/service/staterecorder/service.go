package staterecorder

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/flow/flowinst"
	"github.com/TIBCOSoftware/flogo-lib/flow/service"
	"github.com/TIBCOSoftware/flogo-lib/util"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)


// RemoteStateRecorder is an implementation of StateRecorder service
// that can access flows via URI
type RemoteStateRecorder struct {
	host    string
	enabled bool
}

// NewRemoteStateRecorder creates a new RemoteStateRecorder
func NewRemoteStateRecorder(config *util.ServiceConfig) *RemoteStateRecorder {

	recorder := &RemoteStateRecorder{enabled: config.Enabled}
	recorder.init(config.Settings)

	return recorder
}

func (sr *RemoteStateRecorder) Name() string {
	return service.ServiceStateRecorder
}

func (sr *RemoteStateRecorder) Enabled() bool {
	return sr.enabled
}

// Start implements util.Managed.Start()
func (sr *RemoteStateRecorder) Start() error {
	// no-op
	return nil
}

// Stop implements util.Managed.Stop()
func (sr *RemoteStateRecorder) Stop() error {
	// no-op
	return nil
}

// Init implements services.StateRecorderService.Init()
func (sr *RemoteStateRecorder) init(settings map[string]string) {

	host, set := settings["host"]
	port, set := settings["port"]

	if !set {
		panic("RemoteStateRecorder: requried setting 'host' not set")
	}

	if strings.Index(host, "http") != 0 {
		sr.host = "http://" + host + ":" + port
	} else {
		sr.host = host + ":" + port
	}

	logger.Debugf("RemoteStateRecorder: StateRecoder Server = %s", sr.host)
}

// RecordSnapshot implements flowinst.StateRecorder.RecordSnapshot
func (sr *RemoteStateRecorder) RecordSnapshot(instance *flowinst.Instance) {

	storeReq := &RecordSnapshotReq{
		ID:           instance.StepID(),
		FlowID:       instance.ID(),
		State:        instance.State(),
		Status:       int(instance.Status()),
		SnapshotData: instance,
	}

	uri := sr.host + "/instances/snapshot"

	logger.Debugf("POST Snapshot: %s\n", uri)

	jsonReq, _ := json.Marshal(storeReq)

	logger.Debug("JSON: ", string(jsonReq))

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	logger.Debug("response Status:", resp.Status)

	if resp.StatusCode >= 300 {
		//error
	}
}

// RecordStep implements flowinst.StateRecorder.RecordStep
func (sr *RemoteStateRecorder) RecordStep(instance *flowinst.Instance) {

	storeReq := &RecordStepReq{
		ID:       instance.StepID(),
		FlowID:   instance.ID(),
		State:    instance.State(),
		Status:   int(instance.Status()),
		StepData: instance.ChangeTracker,
	}

	uri := sr.host + "/instances/steps"

	logger.Debugf("POST Snapshot: %s\n", uri)

	jsonReq, _ := json.Marshal(storeReq)

	logger.Debug("JSON: ", string(jsonReq))

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	logger.Debug("response Status:", resp.Status)

	if resp.StatusCode >= 300 {
		//error
	}
}

// RecordSnapshotReq serializable representation of the RecordSnapshot request
type RecordSnapshotReq struct {
	ID     int    `json:"id"`
	FlowID string `json:"flowID"`
	State  int    `json:"state"`
	Status int    `json:"status"`

	SnapshotData *flowinst.Instance `json:"snapshotData"`
}

// RecordStepReq serializable representation of the RecordStep request
type RecordStepReq struct {
	ID     int    `json:"id"`
	FlowID string `json:"flowID"`
	State  int    `json:"state"`
	Status int    `json:"status"`

	StepData *flowinst.InstanceChangeTracker `json:"stepData"`
}

func DefaultConfig() *util.ServiceConfig {
	return &util.ServiceConfig{Name: service.ServiceStateRecorder, Enabled: true, Settings: map[string]string{"host": ""}}
}
