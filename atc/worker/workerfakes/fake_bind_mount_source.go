// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	sync "sync"

	garden "code.cloudfoundry.org/garden"
	worker "github.com/concourse/concourse/atc/worker"
)

type FakeBindMountSource struct {
	VolumeOnStub        func(worker.Worker) (garden.BindMount, bool, error)
	volumeOnMutex       sync.RWMutex
	volumeOnArgsForCall []struct {
		arg1 worker.Worker
	}
	volumeOnReturns struct {
		result1 garden.BindMount
		result2 bool
		result3 error
	}
	volumeOnReturnsOnCall map[int]struct {
		result1 garden.BindMount
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBindMountSource) VolumeOn(arg1 worker.Worker) (garden.BindMount, bool, error) {
	fake.volumeOnMutex.Lock()
	ret, specificReturn := fake.volumeOnReturnsOnCall[len(fake.volumeOnArgsForCall)]
	fake.volumeOnArgsForCall = append(fake.volumeOnArgsForCall, struct {
		arg1 worker.Worker
	}{arg1})
	fake.recordInvocation("VolumeOn", []interface{}{arg1})
	fake.volumeOnMutex.Unlock()
	if fake.VolumeOnStub != nil {
		return fake.VolumeOnStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.volumeOnReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBindMountSource) VolumeOnCallCount() int {
	fake.volumeOnMutex.RLock()
	defer fake.volumeOnMutex.RUnlock()
	return len(fake.volumeOnArgsForCall)
}

func (fake *FakeBindMountSource) VolumeOnCalls(stub func(worker.Worker) (garden.BindMount, bool, error)) {
	fake.volumeOnMutex.Lock()
	defer fake.volumeOnMutex.Unlock()
	fake.VolumeOnStub = stub
}

func (fake *FakeBindMountSource) VolumeOnArgsForCall(i int) worker.Worker {
	fake.volumeOnMutex.RLock()
	defer fake.volumeOnMutex.RUnlock()
	argsForCall := fake.volumeOnArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindMountSource) VolumeOnReturns(result1 garden.BindMount, result2 bool, result3 error) {
	fake.volumeOnMutex.Lock()
	defer fake.volumeOnMutex.Unlock()
	fake.VolumeOnStub = nil
	fake.volumeOnReturns = struct {
		result1 garden.BindMount
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindMountSource) VolumeOnReturnsOnCall(i int, result1 garden.BindMount, result2 bool, result3 error) {
	fake.volumeOnMutex.Lock()
	defer fake.volumeOnMutex.Unlock()
	fake.VolumeOnStub = nil
	if fake.volumeOnReturnsOnCall == nil {
		fake.volumeOnReturnsOnCall = make(map[int]struct {
			result1 garden.BindMount
			result2 bool
			result3 error
		})
	}
	fake.volumeOnReturnsOnCall[i] = struct {
		result1 garden.BindMount
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindMountSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.volumeOnMutex.RLock()
	defer fake.volumeOnMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBindMountSource) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ worker.BindMountSource = new(FakeBindMountSource)
