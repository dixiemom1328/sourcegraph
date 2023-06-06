// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package server

import (
	"sync"

	perforce "github.com/sourcegraph/sourcegraph/cmd/gitserver/server/perforce"
)

// MockPerforceService is a mock implementation of the PerforceService
// interface (from the package
// github.com/sourcegraph/sourcegraph/cmd/gitserver/server/perforce) used
// for unit testing.
type MockPerforceService struct {
	// EnqueueChangelistMappingJobFunc is an instance of a mock function
	// object controlling the behavior of the method
	// EnqueueChangelistMappingJob.
	EnqueueChangelistMappingJobFunc *PerforceServiceEnqueueChangelistMappingJobFunc
}

// NewMockPerforceService creates a new mock of the PerforceService
// interface. All methods return zero values for all results, unless
// overwritten.
func NewMockPerforceService() *MockPerforceService {
	return &MockPerforceService{
		EnqueueChangelistMappingJobFunc: &PerforceServiceEnqueueChangelistMappingJobFunc{
			defaultHook: func(*perforce.ChangelistMappingJob) {
				return
			},
		},
	}
}

// NewStrictMockPerforceService creates a new mock of the PerforceService
// interface. All methods panic on invocation, unless overwritten.
func NewStrictMockPerforceService() *MockPerforceService {
	return &MockPerforceService{
		EnqueueChangelistMappingJobFunc: &PerforceServiceEnqueueChangelistMappingJobFunc{
			defaultHook: func(*perforce.ChangelistMappingJob) {
				panic("unexpected invocation of MockPerforceService.EnqueueChangelistMappingJob")
			},
		},
	}
}

// NewMockPerforceServiceFrom creates a new mock of the MockPerforceService
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockPerforceServiceFrom(i perforce.PerforceService) *MockPerforceService {
	return &MockPerforceService{
		EnqueueChangelistMappingJobFunc: &PerforceServiceEnqueueChangelistMappingJobFunc{
			defaultHook: i.EnqueueChangelistMappingJob,
		},
	}
}

// PerforceServiceEnqueueChangelistMappingJobFunc describes the behavior
// when the EnqueueChangelistMappingJob method of the parent
// MockPerforceService instance is invoked.
type PerforceServiceEnqueueChangelistMappingJobFunc struct {
	defaultHook func(*perforce.ChangelistMappingJob)
	hooks       []func(*perforce.ChangelistMappingJob)
	history     []PerforceServiceEnqueueChangelistMappingJobFuncCall
	mutex       sync.Mutex
}

// EnqueueChangelistMappingJob delegates to the next hook function in the
// queue and stores the parameter and result values of this invocation.
func (m *MockPerforceService) EnqueueChangelistMappingJob(v0 *perforce.ChangelistMappingJob) {
	m.EnqueueChangelistMappingJobFunc.nextHook()(v0)
	m.EnqueueChangelistMappingJobFunc.appendCall(PerforceServiceEnqueueChangelistMappingJobFuncCall{v0})
	return
}

// SetDefaultHook sets function that is called when the
// EnqueueChangelistMappingJob method of the parent MockPerforceService
// instance is invoked and the hook queue is empty.
func (f *PerforceServiceEnqueueChangelistMappingJobFunc) SetDefaultHook(hook func(*perforce.ChangelistMappingJob)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// EnqueueChangelistMappingJob method of the parent MockPerforceService
// instance invokes the hook at the front of the queue and discards it.
// After the queue is empty, the default hook function is invoked for any
// future action.
func (f *PerforceServiceEnqueueChangelistMappingJobFunc) PushHook(hook func(*perforce.ChangelistMappingJob)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *PerforceServiceEnqueueChangelistMappingJobFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(*perforce.ChangelistMappingJob) {
		return
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *PerforceServiceEnqueueChangelistMappingJobFunc) PushReturn() {
	f.PushHook(func(*perforce.ChangelistMappingJob) {
		return
	})
}

func (f *PerforceServiceEnqueueChangelistMappingJobFunc) nextHook() func(*perforce.ChangelistMappingJob) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *PerforceServiceEnqueueChangelistMappingJobFunc) appendCall(r0 PerforceServiceEnqueueChangelistMappingJobFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// PerforceServiceEnqueueChangelistMappingJobFuncCall objects describing the
// invocations of this function.
func (f *PerforceServiceEnqueueChangelistMappingJobFunc) History() []PerforceServiceEnqueueChangelistMappingJobFuncCall {
	f.mutex.Lock()
	history := make([]PerforceServiceEnqueueChangelistMappingJobFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// PerforceServiceEnqueueChangelistMappingJobFuncCall is an object that
// describes an invocation of method EnqueueChangelistMappingJob on an
// instance of MockPerforceService.
type PerforceServiceEnqueueChangelistMappingJobFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 *perforce.ChangelistMappingJob
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c PerforceServiceEnqueueChangelistMappingJobFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c PerforceServiceEnqueueChangelistMappingJobFuncCall) Results() []interface{} {
	return []interface{}{}
}
