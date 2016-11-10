package runner

import (
	"context"
	"errors"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
)

// PooledRunner is a action runner that queues and runs a action in a worker pool
type PooledRunner struct {
	workerQueue chan chan ActionWorkRequest
	workQueue   chan ActionWorkRequest
	numWorkers  int
	workers     []*ActionWorker
	active      bool

	directRunner *DirectRunner
}

// PooledConfig is the configuration object for a PooledRunner
type PooledConfig struct {
	NumWorkers    int `json:"numWorkers"`
	WorkQueueSize int `json:"workQueueSize"`
}

// NewPooledRunner create a new pooled
func NewPooled(config *PooledConfig) *PooledRunner {

	var pooledRunner PooledRunner
	pooledRunner.directRunner = NewDirect()

	// config via engine config
	pooledRunner.numWorkers = config.NumWorkers
	pooledRunner.workQueue = make(chan ActionWorkRequest, config.WorkQueueSize)

	return &pooledRunner
}

// Start will start the engine, by starting all of its workers
func (runner *PooledRunner) Start() error {

	if !runner.active {

		runner.workerQueue = make(chan chan ActionWorkRequest, runner.numWorkers)

		runner.workers = make([]*ActionWorker, runner.numWorkers)

		for i := 0; i < runner.numWorkers; i++ {
			id := i + 1
			log.Debugf("Starting worker with id '%d'", id)
			worker := NewWorker(id, runner.directRunner, runner.workerQueue)
			runner.workers[i] = &worker
			worker.Start()
		}

		go func() {
			for {
				select {
				case work := <-runner.workQueue:
					log.Debug("Received work request")

					//todo fix, this creates unbounded go routines waiting to be serviced by worker queue
					go func() {
						worker := <-runner.workerQueue

						log.Debug("Dispatching work request")
						worker <- work
					}()
				}
			}
		}()

		runner.active = true
	}

	return nil
}

// Stop will stop the engine, by stopping all of its workers
func (runner *PooledRunner) Stop() error {

	if runner.active {

		runner.active = false

		for _, worker := range runner.workers {
			log.Debug("Stopping worker", worker.ID)
			worker.Stop()
		}
	}

	return nil
}

// Run implements action.Runner.Run
func (runner *PooledRunner) Run(context context.Context, action action.Action, uri string, options interface{}) (code int, data interface{}, err error) {

	if action == nil {
		return 0, nil, errors.New("Action not found")
	}

	if runner.active {

		data := &ActionData{context: context, action: action, uri: uri, options: options, rc: make(chan *ActionResult, 1)}
		work := ActionWorkRequest{ReqType: RtRun, actionData: data}

		runner.workQueue <- work
		log.Debugf("Run Action '%s' queued", uri)

		reply := <-data.rc
		log.Debugf("Run Action '%s' complete", uri)

		return reply.code, reply.data, reply.err
	}

	//Run rejected
	return 0, nil, errors.New("Runner not active")
}
