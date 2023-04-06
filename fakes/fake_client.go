// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/executor"
	"code.cloudfoundry.org/lager/v3"
)

type FakeClient struct {
	AllocateContainersStub        func(lager.Logger, []executor.AllocationRequest) []executor.AllocationFailure
	allocateContainersMutex       sync.RWMutex
	allocateContainersArgsForCall []struct {
		arg1 lager.Logger
		arg2 []executor.AllocationRequest
	}
	allocateContainersReturns struct {
		result1 []executor.AllocationFailure
	}
	allocateContainersReturnsOnCall map[int]struct {
		result1 []executor.AllocationFailure
	}
	CleanupStub        func(lager.Logger)
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct {
		arg1 lager.Logger
	}
	DeleteContainerStub        func(lager.Logger, string) error
	deleteContainerMutex       sync.RWMutex
	deleteContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	deleteContainerReturns struct {
		result1 error
	}
	deleteContainerReturnsOnCall map[int]struct {
		result1 error
	}
	GetBulkMetricsStub        func(lager.Logger) (map[string]executor.Metrics, error)
	getBulkMetricsMutex       sync.RWMutex
	getBulkMetricsArgsForCall []struct {
		arg1 lager.Logger
	}
	getBulkMetricsReturns struct {
		result1 map[string]executor.Metrics
		result2 error
	}
	getBulkMetricsReturnsOnCall map[int]struct {
		result1 map[string]executor.Metrics
		result2 error
	}
	GetContainerStub        func(lager.Logger, string) (executor.Container, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	getContainerReturns struct {
		result1 executor.Container
		result2 error
	}
	getContainerReturnsOnCall map[int]struct {
		result1 executor.Container
		result2 error
	}
	GetFilesStub        func(lager.Logger, string, string) (io.ReadCloser, error)
	getFilesMutex       sync.RWMutex
	getFilesArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}
	getFilesReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	getFilesReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	HealthyStub        func(lager.Logger) bool
	healthyMutex       sync.RWMutex
	healthyArgsForCall []struct {
		arg1 lager.Logger
	}
	healthyReturns struct {
		result1 bool
	}
	healthyReturnsOnCall map[int]struct {
		result1 bool
	}
	ListContainersStub        func(lager.Logger) ([]executor.Container, error)
	listContainersMutex       sync.RWMutex
	listContainersArgsForCall []struct {
		arg1 lager.Logger
	}
	listContainersReturns struct {
		result1 []executor.Container
		result2 error
	}
	listContainersReturnsOnCall map[int]struct {
		result1 []executor.Container
		result2 error
	}
	PingStub        func(lager.Logger) error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct {
		arg1 lager.Logger
	}
	pingReturns struct {
		result1 error
	}
	pingReturnsOnCall map[int]struct {
		result1 error
	}
	RemainingResourcesStub        func(lager.Logger) (executor.ExecutorResources, error)
	remainingResourcesMutex       sync.RWMutex
	remainingResourcesArgsForCall []struct {
		arg1 lager.Logger
	}
	remainingResourcesReturns struct {
		result1 executor.ExecutorResources
		result2 error
	}
	remainingResourcesReturnsOnCall map[int]struct {
		result1 executor.ExecutorResources
		result2 error
	}
	RunContainerStub        func(lager.Logger, *executor.RunRequest) error
	runContainerMutex       sync.RWMutex
	runContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 *executor.RunRequest
	}
	runContainerReturns struct {
		result1 error
	}
	runContainerReturnsOnCall map[int]struct {
		result1 error
	}
	SetHealthyStub        func(lager.Logger, bool)
	setHealthyMutex       sync.RWMutex
	setHealthyArgsForCall []struct {
		arg1 lager.Logger
		arg2 bool
	}
	StopContainerStub        func(lager.Logger, string) error
	stopContainerMutex       sync.RWMutex
	stopContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	stopContainerReturns struct {
		result1 error
	}
	stopContainerReturnsOnCall map[int]struct {
		result1 error
	}
	SubscribeToEventsStub        func(lager.Logger) (executor.EventSource, error)
	subscribeToEventsMutex       sync.RWMutex
	subscribeToEventsArgsForCall []struct {
		arg1 lager.Logger
	}
	subscribeToEventsReturns struct {
		result1 executor.EventSource
		result2 error
	}
	subscribeToEventsReturnsOnCall map[int]struct {
		result1 executor.EventSource
		result2 error
	}
	TotalResourcesStub        func(lager.Logger) (executor.ExecutorResources, error)
	totalResourcesMutex       sync.RWMutex
	totalResourcesArgsForCall []struct {
		arg1 lager.Logger
	}
	totalResourcesReturns struct {
		result1 executor.ExecutorResources
		result2 error
	}
	totalResourcesReturnsOnCall map[int]struct {
		result1 executor.ExecutorResources
		result2 error
	}
	UpdateContainerStub        func(lager.Logger, *executor.UpdateRequest) error
	updateContainerMutex       sync.RWMutex
	updateContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 *executor.UpdateRequest
	}
	updateContainerReturns struct {
		result1 error
	}
	updateContainerReturnsOnCall map[int]struct {
		result1 error
	}
	VolumeDriversStub        func(lager.Logger) ([]string, error)
	volumeDriversMutex       sync.RWMutex
	volumeDriversArgsForCall []struct {
		arg1 lager.Logger
	}
	volumeDriversReturns struct {
		result1 []string
		result2 error
	}
	volumeDriversReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) AllocateContainers(arg1 lager.Logger, arg2 []executor.AllocationRequest) []executor.AllocationFailure {
	var arg2Copy []executor.AllocationRequest
	if arg2 != nil {
		arg2Copy = make([]executor.AllocationRequest, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.allocateContainersMutex.Lock()
	ret, specificReturn := fake.allocateContainersReturnsOnCall[len(fake.allocateContainersArgsForCall)]
	fake.allocateContainersArgsForCall = append(fake.allocateContainersArgsForCall, struct {
		arg1 lager.Logger
		arg2 []executor.AllocationRequest
	}{arg1, arg2Copy})
	stub := fake.AllocateContainersStub
	fakeReturns := fake.allocateContainersReturns
	fake.recordInvocation("AllocateContainers", []interface{}{arg1, arg2Copy})
	fake.allocateContainersMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) AllocateContainersCallCount() int {
	fake.allocateContainersMutex.RLock()
	defer fake.allocateContainersMutex.RUnlock()
	return len(fake.allocateContainersArgsForCall)
}

func (fake *FakeClient) AllocateContainersCalls(stub func(lager.Logger, []executor.AllocationRequest) []executor.AllocationFailure) {
	fake.allocateContainersMutex.Lock()
	defer fake.allocateContainersMutex.Unlock()
	fake.AllocateContainersStub = stub
}

func (fake *FakeClient) AllocateContainersArgsForCall(i int) (lager.Logger, []executor.AllocationRequest) {
	fake.allocateContainersMutex.RLock()
	defer fake.allocateContainersMutex.RUnlock()
	argsForCall := fake.allocateContainersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) AllocateContainersReturns(result1 []executor.AllocationFailure) {
	fake.allocateContainersMutex.Lock()
	defer fake.allocateContainersMutex.Unlock()
	fake.AllocateContainersStub = nil
	fake.allocateContainersReturns = struct {
		result1 []executor.AllocationFailure
	}{result1}
}

func (fake *FakeClient) AllocateContainersReturnsOnCall(i int, result1 []executor.AllocationFailure) {
	fake.allocateContainersMutex.Lock()
	defer fake.allocateContainersMutex.Unlock()
	fake.AllocateContainersStub = nil
	if fake.allocateContainersReturnsOnCall == nil {
		fake.allocateContainersReturnsOnCall = make(map[int]struct {
			result1 []executor.AllocationFailure
		})
	}
	fake.allocateContainersReturnsOnCall[i] = struct {
		result1 []executor.AllocationFailure
	}{result1}
}

func (fake *FakeClient) Cleanup(arg1 lager.Logger) {
	fake.cleanupMutex.Lock()
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.CleanupStub
	fake.recordInvocation("Cleanup", []interface{}{arg1})
	fake.cleanupMutex.Unlock()
	if stub != nil {
		fake.CleanupStub(arg1)
	}
}

func (fake *FakeClient) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeClient) CleanupCalls(stub func(lager.Logger)) {
	fake.cleanupMutex.Lock()
	defer fake.cleanupMutex.Unlock()
	fake.CleanupStub = stub
}

func (fake *FakeClient) CleanupArgsForCall(i int) lager.Logger {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	argsForCall := fake.cleanupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) DeleteContainer(arg1 lager.Logger, arg2 string) error {
	fake.deleteContainerMutex.Lock()
	ret, specificReturn := fake.deleteContainerReturnsOnCall[len(fake.deleteContainerArgsForCall)]
	fake.deleteContainerArgsForCall = append(fake.deleteContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteContainerStub
	fakeReturns := fake.deleteContainerReturns
	fake.recordInvocation("DeleteContainer", []interface{}{arg1, arg2})
	fake.deleteContainerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) DeleteContainerCallCount() int {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return len(fake.deleteContainerArgsForCall)
}

func (fake *FakeClient) DeleteContainerCalls(stub func(lager.Logger, string) error) {
	fake.deleteContainerMutex.Lock()
	defer fake.deleteContainerMutex.Unlock()
	fake.DeleteContainerStub = stub
}

func (fake *FakeClient) DeleteContainerArgsForCall(i int) (lager.Logger, string) {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	argsForCall := fake.deleteContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) DeleteContainerReturns(result1 error) {
	fake.deleteContainerMutex.Lock()
	defer fake.deleteContainerMutex.Unlock()
	fake.DeleteContainerStub = nil
	fake.deleteContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DeleteContainerReturnsOnCall(i int, result1 error) {
	fake.deleteContainerMutex.Lock()
	defer fake.deleteContainerMutex.Unlock()
	fake.DeleteContainerStub = nil
	if fake.deleteContainerReturnsOnCall == nil {
		fake.deleteContainerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteContainerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) GetBulkMetrics(arg1 lager.Logger) (map[string]executor.Metrics, error) {
	fake.getBulkMetricsMutex.Lock()
	ret, specificReturn := fake.getBulkMetricsReturnsOnCall[len(fake.getBulkMetricsArgsForCall)]
	fake.getBulkMetricsArgsForCall = append(fake.getBulkMetricsArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.GetBulkMetricsStub
	fakeReturns := fake.getBulkMetricsReturns
	fake.recordInvocation("GetBulkMetrics", []interface{}{arg1})
	fake.getBulkMetricsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) GetBulkMetricsCallCount() int {
	fake.getBulkMetricsMutex.RLock()
	defer fake.getBulkMetricsMutex.RUnlock()
	return len(fake.getBulkMetricsArgsForCall)
}

func (fake *FakeClient) GetBulkMetricsCalls(stub func(lager.Logger) (map[string]executor.Metrics, error)) {
	fake.getBulkMetricsMutex.Lock()
	defer fake.getBulkMetricsMutex.Unlock()
	fake.GetBulkMetricsStub = stub
}

func (fake *FakeClient) GetBulkMetricsArgsForCall(i int) lager.Logger {
	fake.getBulkMetricsMutex.RLock()
	defer fake.getBulkMetricsMutex.RUnlock()
	argsForCall := fake.getBulkMetricsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) GetBulkMetricsReturns(result1 map[string]executor.Metrics, result2 error) {
	fake.getBulkMetricsMutex.Lock()
	defer fake.getBulkMetricsMutex.Unlock()
	fake.GetBulkMetricsStub = nil
	fake.getBulkMetricsReturns = struct {
		result1 map[string]executor.Metrics
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetBulkMetricsReturnsOnCall(i int, result1 map[string]executor.Metrics, result2 error) {
	fake.getBulkMetricsMutex.Lock()
	defer fake.getBulkMetricsMutex.Unlock()
	fake.GetBulkMetricsStub = nil
	if fake.getBulkMetricsReturnsOnCall == nil {
		fake.getBulkMetricsReturnsOnCall = make(map[int]struct {
			result1 map[string]executor.Metrics
			result2 error
		})
	}
	fake.getBulkMetricsReturnsOnCall[i] = struct {
		result1 map[string]executor.Metrics
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetContainer(arg1 lager.Logger, arg2 string) (executor.Container, error) {
	fake.getContainerMutex.Lock()
	ret, specificReturn := fake.getContainerReturnsOnCall[len(fake.getContainerArgsForCall)]
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.GetContainerStub
	fakeReturns := fake.getContainerReturns
	fake.recordInvocation("GetContainer", []interface{}{arg1, arg2})
	fake.getContainerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeClient) GetContainerCalls(stub func(lager.Logger, string) (executor.Container, error)) {
	fake.getContainerMutex.Lock()
	defer fake.getContainerMutex.Unlock()
	fake.GetContainerStub = stub
}

func (fake *FakeClient) GetContainerArgsForCall(i int) (lager.Logger, string) {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	argsForCall := fake.getContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) GetContainerReturns(result1 executor.Container, result2 error) {
	fake.getContainerMutex.Lock()
	defer fake.getContainerMutex.Unlock()
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 executor.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetContainerReturnsOnCall(i int, result1 executor.Container, result2 error) {
	fake.getContainerMutex.Lock()
	defer fake.getContainerMutex.Unlock()
	fake.GetContainerStub = nil
	if fake.getContainerReturnsOnCall == nil {
		fake.getContainerReturnsOnCall = make(map[int]struct {
			result1 executor.Container
			result2 error
		})
	}
	fake.getContainerReturnsOnCall[i] = struct {
		result1 executor.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetFiles(arg1 lager.Logger, arg2 string, arg3 string) (io.ReadCloser, error) {
	fake.getFilesMutex.Lock()
	ret, specificReturn := fake.getFilesReturnsOnCall[len(fake.getFilesArgsForCall)]
	fake.getFilesArgsForCall = append(fake.getFilesArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetFilesStub
	fakeReturns := fake.getFilesReturns
	fake.recordInvocation("GetFiles", []interface{}{arg1, arg2, arg3})
	fake.getFilesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) GetFilesCallCount() int {
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	return len(fake.getFilesArgsForCall)
}

func (fake *FakeClient) GetFilesCalls(stub func(lager.Logger, string, string) (io.ReadCloser, error)) {
	fake.getFilesMutex.Lock()
	defer fake.getFilesMutex.Unlock()
	fake.GetFilesStub = stub
}

func (fake *FakeClient) GetFilesArgsForCall(i int) (lager.Logger, string, string) {
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	argsForCall := fake.getFilesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) GetFilesReturns(result1 io.ReadCloser, result2 error) {
	fake.getFilesMutex.Lock()
	defer fake.getFilesMutex.Unlock()
	fake.GetFilesStub = nil
	fake.getFilesReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetFilesReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.getFilesMutex.Lock()
	defer fake.getFilesMutex.Unlock()
	fake.GetFilesStub = nil
	if fake.getFilesReturnsOnCall == nil {
		fake.getFilesReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.getFilesReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Healthy(arg1 lager.Logger) bool {
	fake.healthyMutex.Lock()
	ret, specificReturn := fake.healthyReturnsOnCall[len(fake.healthyArgsForCall)]
	fake.healthyArgsForCall = append(fake.healthyArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.HealthyStub
	fakeReturns := fake.healthyReturns
	fake.recordInvocation("Healthy", []interface{}{arg1})
	fake.healthyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) HealthyCallCount() int {
	fake.healthyMutex.RLock()
	defer fake.healthyMutex.RUnlock()
	return len(fake.healthyArgsForCall)
}

func (fake *FakeClient) HealthyCalls(stub func(lager.Logger) bool) {
	fake.healthyMutex.Lock()
	defer fake.healthyMutex.Unlock()
	fake.HealthyStub = stub
}

func (fake *FakeClient) HealthyArgsForCall(i int) lager.Logger {
	fake.healthyMutex.RLock()
	defer fake.healthyMutex.RUnlock()
	argsForCall := fake.healthyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) HealthyReturns(result1 bool) {
	fake.healthyMutex.Lock()
	defer fake.healthyMutex.Unlock()
	fake.HealthyStub = nil
	fake.healthyReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeClient) HealthyReturnsOnCall(i int, result1 bool) {
	fake.healthyMutex.Lock()
	defer fake.healthyMutex.Unlock()
	fake.HealthyStub = nil
	if fake.healthyReturnsOnCall == nil {
		fake.healthyReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.healthyReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeClient) ListContainers(arg1 lager.Logger) ([]executor.Container, error) {
	fake.listContainersMutex.Lock()
	ret, specificReturn := fake.listContainersReturnsOnCall[len(fake.listContainersArgsForCall)]
	fake.listContainersArgsForCall = append(fake.listContainersArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.ListContainersStub
	fakeReturns := fake.listContainersReturns
	fake.recordInvocation("ListContainers", []interface{}{arg1})
	fake.listContainersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) ListContainersCallCount() int {
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	return len(fake.listContainersArgsForCall)
}

func (fake *FakeClient) ListContainersCalls(stub func(lager.Logger) ([]executor.Container, error)) {
	fake.listContainersMutex.Lock()
	defer fake.listContainersMutex.Unlock()
	fake.ListContainersStub = stub
}

func (fake *FakeClient) ListContainersArgsForCall(i int) lager.Logger {
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	argsForCall := fake.listContainersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) ListContainersReturns(result1 []executor.Container, result2 error) {
	fake.listContainersMutex.Lock()
	defer fake.listContainersMutex.Unlock()
	fake.ListContainersStub = nil
	fake.listContainersReturns = struct {
		result1 []executor.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListContainersReturnsOnCall(i int, result1 []executor.Container, result2 error) {
	fake.listContainersMutex.Lock()
	defer fake.listContainersMutex.Unlock()
	fake.ListContainersStub = nil
	if fake.listContainersReturnsOnCall == nil {
		fake.listContainersReturnsOnCall = make(map[int]struct {
			result1 []executor.Container
			result2 error
		})
	}
	fake.listContainersReturnsOnCall[i] = struct {
		result1 []executor.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Ping(arg1 lager.Logger) error {
	fake.pingMutex.Lock()
	ret, specificReturn := fake.pingReturnsOnCall[len(fake.pingArgsForCall)]
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.PingStub
	fakeReturns := fake.pingReturns
	fake.recordInvocation("Ping", []interface{}{arg1})
	fake.pingMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeClient) PingCalls(stub func(lager.Logger) error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = stub
}

func (fake *FakeClient) PingArgsForCall(i int) lager.Logger {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	argsForCall := fake.pingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) PingReturns(result1 error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) PingReturnsOnCall(i int, result1 error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = nil
	if fake.pingReturnsOnCall == nil {
		fake.pingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) RemainingResources(arg1 lager.Logger) (executor.ExecutorResources, error) {
	fake.remainingResourcesMutex.Lock()
	ret, specificReturn := fake.remainingResourcesReturnsOnCall[len(fake.remainingResourcesArgsForCall)]
	fake.remainingResourcesArgsForCall = append(fake.remainingResourcesArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.RemainingResourcesStub
	fakeReturns := fake.remainingResourcesReturns
	fake.recordInvocation("RemainingResources", []interface{}{arg1})
	fake.remainingResourcesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) RemainingResourcesCallCount() int {
	fake.remainingResourcesMutex.RLock()
	defer fake.remainingResourcesMutex.RUnlock()
	return len(fake.remainingResourcesArgsForCall)
}

func (fake *FakeClient) RemainingResourcesCalls(stub func(lager.Logger) (executor.ExecutorResources, error)) {
	fake.remainingResourcesMutex.Lock()
	defer fake.remainingResourcesMutex.Unlock()
	fake.RemainingResourcesStub = stub
}

func (fake *FakeClient) RemainingResourcesArgsForCall(i int) lager.Logger {
	fake.remainingResourcesMutex.RLock()
	defer fake.remainingResourcesMutex.RUnlock()
	argsForCall := fake.remainingResourcesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) RemainingResourcesReturns(result1 executor.ExecutorResources, result2 error) {
	fake.remainingResourcesMutex.Lock()
	defer fake.remainingResourcesMutex.Unlock()
	fake.RemainingResourcesStub = nil
	fake.remainingResourcesReturns = struct {
		result1 executor.ExecutorResources
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RemainingResourcesReturnsOnCall(i int, result1 executor.ExecutorResources, result2 error) {
	fake.remainingResourcesMutex.Lock()
	defer fake.remainingResourcesMutex.Unlock()
	fake.RemainingResourcesStub = nil
	if fake.remainingResourcesReturnsOnCall == nil {
		fake.remainingResourcesReturnsOnCall = make(map[int]struct {
			result1 executor.ExecutorResources
			result2 error
		})
	}
	fake.remainingResourcesReturnsOnCall[i] = struct {
		result1 executor.ExecutorResources
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RunContainer(arg1 lager.Logger, arg2 *executor.RunRequest) error {
	fake.runContainerMutex.Lock()
	ret, specificReturn := fake.runContainerReturnsOnCall[len(fake.runContainerArgsForCall)]
	fake.runContainerArgsForCall = append(fake.runContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 *executor.RunRequest
	}{arg1, arg2})
	stub := fake.RunContainerStub
	fakeReturns := fake.runContainerReturns
	fake.recordInvocation("RunContainer", []interface{}{arg1, arg2})
	fake.runContainerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) RunContainerCallCount() int {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return len(fake.runContainerArgsForCall)
}

func (fake *FakeClient) RunContainerCalls(stub func(lager.Logger, *executor.RunRequest) error) {
	fake.runContainerMutex.Lock()
	defer fake.runContainerMutex.Unlock()
	fake.RunContainerStub = stub
}

func (fake *FakeClient) RunContainerArgsForCall(i int) (lager.Logger, *executor.RunRequest) {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	argsForCall := fake.runContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) RunContainerReturns(result1 error) {
	fake.runContainerMutex.Lock()
	defer fake.runContainerMutex.Unlock()
	fake.RunContainerStub = nil
	fake.runContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) RunContainerReturnsOnCall(i int, result1 error) {
	fake.runContainerMutex.Lock()
	defer fake.runContainerMutex.Unlock()
	fake.RunContainerStub = nil
	if fake.runContainerReturnsOnCall == nil {
		fake.runContainerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runContainerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SetHealthy(arg1 lager.Logger, arg2 bool) {
	fake.setHealthyMutex.Lock()
	fake.setHealthyArgsForCall = append(fake.setHealthyArgsForCall, struct {
		arg1 lager.Logger
		arg2 bool
	}{arg1, arg2})
	stub := fake.SetHealthyStub
	fake.recordInvocation("SetHealthy", []interface{}{arg1, arg2})
	fake.setHealthyMutex.Unlock()
	if stub != nil {
		fake.SetHealthyStub(arg1, arg2)
	}
}

func (fake *FakeClient) SetHealthyCallCount() int {
	fake.setHealthyMutex.RLock()
	defer fake.setHealthyMutex.RUnlock()
	return len(fake.setHealthyArgsForCall)
}

func (fake *FakeClient) SetHealthyCalls(stub func(lager.Logger, bool)) {
	fake.setHealthyMutex.Lock()
	defer fake.setHealthyMutex.Unlock()
	fake.SetHealthyStub = stub
}

func (fake *FakeClient) SetHealthyArgsForCall(i int) (lager.Logger, bool) {
	fake.setHealthyMutex.RLock()
	defer fake.setHealthyMutex.RUnlock()
	argsForCall := fake.setHealthyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) StopContainer(arg1 lager.Logger, arg2 string) error {
	fake.stopContainerMutex.Lock()
	ret, specificReturn := fake.stopContainerReturnsOnCall[len(fake.stopContainerArgsForCall)]
	fake.stopContainerArgsForCall = append(fake.stopContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.StopContainerStub
	fakeReturns := fake.stopContainerReturns
	fake.recordInvocation("StopContainer", []interface{}{arg1, arg2})
	fake.stopContainerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) StopContainerCallCount() int {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	return len(fake.stopContainerArgsForCall)
}

func (fake *FakeClient) StopContainerCalls(stub func(lager.Logger, string) error) {
	fake.stopContainerMutex.Lock()
	defer fake.stopContainerMutex.Unlock()
	fake.StopContainerStub = stub
}

func (fake *FakeClient) StopContainerArgsForCall(i int) (lager.Logger, string) {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	argsForCall := fake.stopContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) StopContainerReturns(result1 error) {
	fake.stopContainerMutex.Lock()
	defer fake.stopContainerMutex.Unlock()
	fake.StopContainerStub = nil
	fake.stopContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) StopContainerReturnsOnCall(i int, result1 error) {
	fake.stopContainerMutex.Lock()
	defer fake.stopContainerMutex.Unlock()
	fake.StopContainerStub = nil
	if fake.stopContainerReturnsOnCall == nil {
		fake.stopContainerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopContainerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SubscribeToEvents(arg1 lager.Logger) (executor.EventSource, error) {
	fake.subscribeToEventsMutex.Lock()
	ret, specificReturn := fake.subscribeToEventsReturnsOnCall[len(fake.subscribeToEventsArgsForCall)]
	fake.subscribeToEventsArgsForCall = append(fake.subscribeToEventsArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.SubscribeToEventsStub
	fakeReturns := fake.subscribeToEventsReturns
	fake.recordInvocation("SubscribeToEvents", []interface{}{arg1})
	fake.subscribeToEventsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) SubscribeToEventsCallCount() int {
	fake.subscribeToEventsMutex.RLock()
	defer fake.subscribeToEventsMutex.RUnlock()
	return len(fake.subscribeToEventsArgsForCall)
}

func (fake *FakeClient) SubscribeToEventsCalls(stub func(lager.Logger) (executor.EventSource, error)) {
	fake.subscribeToEventsMutex.Lock()
	defer fake.subscribeToEventsMutex.Unlock()
	fake.SubscribeToEventsStub = stub
}

func (fake *FakeClient) SubscribeToEventsArgsForCall(i int) lager.Logger {
	fake.subscribeToEventsMutex.RLock()
	defer fake.subscribeToEventsMutex.RUnlock()
	argsForCall := fake.subscribeToEventsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) SubscribeToEventsReturns(result1 executor.EventSource, result2 error) {
	fake.subscribeToEventsMutex.Lock()
	defer fake.subscribeToEventsMutex.Unlock()
	fake.SubscribeToEventsStub = nil
	fake.subscribeToEventsReturns = struct {
		result1 executor.EventSource
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) SubscribeToEventsReturnsOnCall(i int, result1 executor.EventSource, result2 error) {
	fake.subscribeToEventsMutex.Lock()
	defer fake.subscribeToEventsMutex.Unlock()
	fake.SubscribeToEventsStub = nil
	if fake.subscribeToEventsReturnsOnCall == nil {
		fake.subscribeToEventsReturnsOnCall = make(map[int]struct {
			result1 executor.EventSource
			result2 error
		})
	}
	fake.subscribeToEventsReturnsOnCall[i] = struct {
		result1 executor.EventSource
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) TotalResources(arg1 lager.Logger) (executor.ExecutorResources, error) {
	fake.totalResourcesMutex.Lock()
	ret, specificReturn := fake.totalResourcesReturnsOnCall[len(fake.totalResourcesArgsForCall)]
	fake.totalResourcesArgsForCall = append(fake.totalResourcesArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.TotalResourcesStub
	fakeReturns := fake.totalResourcesReturns
	fake.recordInvocation("TotalResources", []interface{}{arg1})
	fake.totalResourcesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) TotalResourcesCallCount() int {
	fake.totalResourcesMutex.RLock()
	defer fake.totalResourcesMutex.RUnlock()
	return len(fake.totalResourcesArgsForCall)
}

func (fake *FakeClient) TotalResourcesCalls(stub func(lager.Logger) (executor.ExecutorResources, error)) {
	fake.totalResourcesMutex.Lock()
	defer fake.totalResourcesMutex.Unlock()
	fake.TotalResourcesStub = stub
}

func (fake *FakeClient) TotalResourcesArgsForCall(i int) lager.Logger {
	fake.totalResourcesMutex.RLock()
	defer fake.totalResourcesMutex.RUnlock()
	argsForCall := fake.totalResourcesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) TotalResourcesReturns(result1 executor.ExecutorResources, result2 error) {
	fake.totalResourcesMutex.Lock()
	defer fake.totalResourcesMutex.Unlock()
	fake.TotalResourcesStub = nil
	fake.totalResourcesReturns = struct {
		result1 executor.ExecutorResources
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) TotalResourcesReturnsOnCall(i int, result1 executor.ExecutorResources, result2 error) {
	fake.totalResourcesMutex.Lock()
	defer fake.totalResourcesMutex.Unlock()
	fake.TotalResourcesStub = nil
	if fake.totalResourcesReturnsOnCall == nil {
		fake.totalResourcesReturnsOnCall = make(map[int]struct {
			result1 executor.ExecutorResources
			result2 error
		})
	}
	fake.totalResourcesReturnsOnCall[i] = struct {
		result1 executor.ExecutorResources
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpdateContainer(arg1 lager.Logger, arg2 *executor.UpdateRequest) error {
	fake.updateContainerMutex.Lock()
	ret, specificReturn := fake.updateContainerReturnsOnCall[len(fake.updateContainerArgsForCall)]
	fake.updateContainerArgsForCall = append(fake.updateContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 *executor.UpdateRequest
	}{arg1, arg2})
	stub := fake.UpdateContainerStub
	fakeReturns := fake.updateContainerReturns
	fake.recordInvocation("UpdateContainer", []interface{}{arg1, arg2})
	fake.updateContainerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) UpdateContainerCallCount() int {
	fake.updateContainerMutex.RLock()
	defer fake.updateContainerMutex.RUnlock()
	return len(fake.updateContainerArgsForCall)
}

func (fake *FakeClient) UpdateContainerCalls(stub func(lager.Logger, *executor.UpdateRequest) error) {
	fake.updateContainerMutex.Lock()
	defer fake.updateContainerMutex.Unlock()
	fake.UpdateContainerStub = stub
}

func (fake *FakeClient) UpdateContainerArgsForCall(i int) (lager.Logger, *executor.UpdateRequest) {
	fake.updateContainerMutex.RLock()
	defer fake.updateContainerMutex.RUnlock()
	argsForCall := fake.updateContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) UpdateContainerReturns(result1 error) {
	fake.updateContainerMutex.Lock()
	defer fake.updateContainerMutex.Unlock()
	fake.UpdateContainerStub = nil
	fake.updateContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) UpdateContainerReturnsOnCall(i int, result1 error) {
	fake.updateContainerMutex.Lock()
	defer fake.updateContainerMutex.Unlock()
	fake.UpdateContainerStub = nil
	if fake.updateContainerReturnsOnCall == nil {
		fake.updateContainerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateContainerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) VolumeDrivers(arg1 lager.Logger) ([]string, error) {
	fake.volumeDriversMutex.Lock()
	ret, specificReturn := fake.volumeDriversReturnsOnCall[len(fake.volumeDriversArgsForCall)]
	fake.volumeDriversArgsForCall = append(fake.volumeDriversArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.VolumeDriversStub
	fakeReturns := fake.volumeDriversReturns
	fake.recordInvocation("VolumeDrivers", []interface{}{arg1})
	fake.volumeDriversMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) VolumeDriversCallCount() int {
	fake.volumeDriversMutex.RLock()
	defer fake.volumeDriversMutex.RUnlock()
	return len(fake.volumeDriversArgsForCall)
}

func (fake *FakeClient) VolumeDriversCalls(stub func(lager.Logger) ([]string, error)) {
	fake.volumeDriversMutex.Lock()
	defer fake.volumeDriversMutex.Unlock()
	fake.VolumeDriversStub = stub
}

func (fake *FakeClient) VolumeDriversArgsForCall(i int) lager.Logger {
	fake.volumeDriversMutex.RLock()
	defer fake.volumeDriversMutex.RUnlock()
	argsForCall := fake.volumeDriversArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) VolumeDriversReturns(result1 []string, result2 error) {
	fake.volumeDriversMutex.Lock()
	defer fake.volumeDriversMutex.Unlock()
	fake.VolumeDriversStub = nil
	fake.volumeDriversReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) VolumeDriversReturnsOnCall(i int, result1 []string, result2 error) {
	fake.volumeDriversMutex.Lock()
	defer fake.volumeDriversMutex.Unlock()
	fake.VolumeDriversStub = nil
	if fake.volumeDriversReturnsOnCall == nil {
		fake.volumeDriversReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.volumeDriversReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allocateContainersMutex.RLock()
	defer fake.allocateContainersMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	fake.getBulkMetricsMutex.RLock()
	defer fake.getBulkMetricsMutex.RUnlock()
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	fake.healthyMutex.RLock()
	defer fake.healthyMutex.RUnlock()
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.remainingResourcesMutex.RLock()
	defer fake.remainingResourcesMutex.RUnlock()
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	fake.setHealthyMutex.RLock()
	defer fake.setHealthyMutex.RUnlock()
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	fake.subscribeToEventsMutex.RLock()
	defer fake.subscribeToEventsMutex.RUnlock()
	fake.totalResourcesMutex.RLock()
	defer fake.totalResourcesMutex.RUnlock()
	fake.updateContainerMutex.RLock()
	defer fake.updateContainerMutex.RUnlock()
	fake.volumeDriversMutex.RLock()
	defer fake.volumeDriversMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ executor.Client = new(FakeClient)
