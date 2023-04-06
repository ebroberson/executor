// Code generated by counterfeiter. DO NOT EDIT.
package containerstorefakes

import (
	"sync"

	diego_logging_client "code.cloudfoundry.org/diego-logging-client"
	"code.cloudfoundry.org/executor"
	"code.cloudfoundry.org/executor/depot/containerstore"
	"code.cloudfoundry.org/lager/v3"
)

type FakeDependencyManager struct {
	DownloadCachedDependenciesStub        func(lager.Logger, []executor.CachedDependency, executor.LogConfig, diego_logging_client.IngressClient) (containerstore.BindMounts, error)
	downloadCachedDependenciesMutex       sync.RWMutex
	downloadCachedDependenciesArgsForCall []struct {
		arg1 lager.Logger
		arg2 []executor.CachedDependency
		arg3 executor.LogConfig
		arg4 diego_logging_client.IngressClient
	}
	downloadCachedDependenciesReturns struct {
		result1 containerstore.BindMounts
		result2 error
	}
	downloadCachedDependenciesReturnsOnCall map[int]struct {
		result1 containerstore.BindMounts
		result2 error
	}
	ReleaseCachedDependenciesStub        func(lager.Logger, []containerstore.BindMountCacheKey) error
	releaseCachedDependenciesMutex       sync.RWMutex
	releaseCachedDependenciesArgsForCall []struct {
		arg1 lager.Logger
		arg2 []containerstore.BindMountCacheKey
	}
	releaseCachedDependenciesReturns struct {
		result1 error
	}
	releaseCachedDependenciesReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func(lager.Logger)
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		arg1 lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDependencyManager) DownloadCachedDependencies(arg1 lager.Logger, arg2 []executor.CachedDependency, arg3 executor.LogConfig, arg4 diego_logging_client.IngressClient) (containerstore.BindMounts, error) {
	var arg2Copy []executor.CachedDependency
	if arg2 != nil {
		arg2Copy = make([]executor.CachedDependency, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.downloadCachedDependenciesMutex.Lock()
	ret, specificReturn := fake.downloadCachedDependenciesReturnsOnCall[len(fake.downloadCachedDependenciesArgsForCall)]
	fake.downloadCachedDependenciesArgsForCall = append(fake.downloadCachedDependenciesArgsForCall, struct {
		arg1 lager.Logger
		arg2 []executor.CachedDependency
		arg3 executor.LogConfig
		arg4 diego_logging_client.IngressClient
	}{arg1, arg2Copy, arg3, arg4})
	stub := fake.DownloadCachedDependenciesStub
	fakeReturns := fake.downloadCachedDependenciesReturns
	fake.recordInvocation("DownloadCachedDependencies", []interface{}{arg1, arg2Copy, arg3, arg4})
	fake.downloadCachedDependenciesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDependencyManager) DownloadCachedDependenciesCallCount() int {
	fake.downloadCachedDependenciesMutex.RLock()
	defer fake.downloadCachedDependenciesMutex.RUnlock()
	return len(fake.downloadCachedDependenciesArgsForCall)
}

func (fake *FakeDependencyManager) DownloadCachedDependenciesCalls(stub func(lager.Logger, []executor.CachedDependency, executor.LogConfig, diego_logging_client.IngressClient) (containerstore.BindMounts, error)) {
	fake.downloadCachedDependenciesMutex.Lock()
	defer fake.downloadCachedDependenciesMutex.Unlock()
	fake.DownloadCachedDependenciesStub = stub
}

func (fake *FakeDependencyManager) DownloadCachedDependenciesArgsForCall(i int) (lager.Logger, []executor.CachedDependency, executor.LogConfig, diego_logging_client.IngressClient) {
	fake.downloadCachedDependenciesMutex.RLock()
	defer fake.downloadCachedDependenciesMutex.RUnlock()
	argsForCall := fake.downloadCachedDependenciesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeDependencyManager) DownloadCachedDependenciesReturns(result1 containerstore.BindMounts, result2 error) {
	fake.downloadCachedDependenciesMutex.Lock()
	defer fake.downloadCachedDependenciesMutex.Unlock()
	fake.DownloadCachedDependenciesStub = nil
	fake.downloadCachedDependenciesReturns = struct {
		result1 containerstore.BindMounts
		result2 error
	}{result1, result2}
}

func (fake *FakeDependencyManager) DownloadCachedDependenciesReturnsOnCall(i int, result1 containerstore.BindMounts, result2 error) {
	fake.downloadCachedDependenciesMutex.Lock()
	defer fake.downloadCachedDependenciesMutex.Unlock()
	fake.DownloadCachedDependenciesStub = nil
	if fake.downloadCachedDependenciesReturnsOnCall == nil {
		fake.downloadCachedDependenciesReturnsOnCall = make(map[int]struct {
			result1 containerstore.BindMounts
			result2 error
		})
	}
	fake.downloadCachedDependenciesReturnsOnCall[i] = struct {
		result1 containerstore.BindMounts
		result2 error
	}{result1, result2}
}

func (fake *FakeDependencyManager) ReleaseCachedDependencies(arg1 lager.Logger, arg2 []containerstore.BindMountCacheKey) error {
	var arg2Copy []containerstore.BindMountCacheKey
	if arg2 != nil {
		arg2Copy = make([]containerstore.BindMountCacheKey, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.releaseCachedDependenciesMutex.Lock()
	ret, specificReturn := fake.releaseCachedDependenciesReturnsOnCall[len(fake.releaseCachedDependenciesArgsForCall)]
	fake.releaseCachedDependenciesArgsForCall = append(fake.releaseCachedDependenciesArgsForCall, struct {
		arg1 lager.Logger
		arg2 []containerstore.BindMountCacheKey
	}{arg1, arg2Copy})
	stub := fake.ReleaseCachedDependenciesStub
	fakeReturns := fake.releaseCachedDependenciesReturns
	fake.recordInvocation("ReleaseCachedDependencies", []interface{}{arg1, arg2Copy})
	fake.releaseCachedDependenciesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDependencyManager) ReleaseCachedDependenciesCallCount() int {
	fake.releaseCachedDependenciesMutex.RLock()
	defer fake.releaseCachedDependenciesMutex.RUnlock()
	return len(fake.releaseCachedDependenciesArgsForCall)
}

func (fake *FakeDependencyManager) ReleaseCachedDependenciesCalls(stub func(lager.Logger, []containerstore.BindMountCacheKey) error) {
	fake.releaseCachedDependenciesMutex.Lock()
	defer fake.releaseCachedDependenciesMutex.Unlock()
	fake.ReleaseCachedDependenciesStub = stub
}

func (fake *FakeDependencyManager) ReleaseCachedDependenciesArgsForCall(i int) (lager.Logger, []containerstore.BindMountCacheKey) {
	fake.releaseCachedDependenciesMutex.RLock()
	defer fake.releaseCachedDependenciesMutex.RUnlock()
	argsForCall := fake.releaseCachedDependenciesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDependencyManager) ReleaseCachedDependenciesReturns(result1 error) {
	fake.releaseCachedDependenciesMutex.Lock()
	defer fake.releaseCachedDependenciesMutex.Unlock()
	fake.ReleaseCachedDependenciesStub = nil
	fake.releaseCachedDependenciesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDependencyManager) ReleaseCachedDependenciesReturnsOnCall(i int, result1 error) {
	fake.releaseCachedDependenciesMutex.Lock()
	defer fake.releaseCachedDependenciesMutex.Unlock()
	fake.ReleaseCachedDependenciesStub = nil
	if fake.releaseCachedDependenciesReturnsOnCall == nil {
		fake.releaseCachedDependenciesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.releaseCachedDependenciesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDependencyManager) Stop(arg1 lager.Logger) {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.StopStub
	fake.recordInvocation("Stop", []interface{}{arg1})
	fake.stopMutex.Unlock()
	if stub != nil {
		fake.StopStub(arg1)
	}
}

func (fake *FakeDependencyManager) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeDependencyManager) StopCalls(stub func(lager.Logger)) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeDependencyManager) StopArgsForCall(i int) lager.Logger {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	argsForCall := fake.stopArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDependencyManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadCachedDependenciesMutex.RLock()
	defer fake.downloadCachedDependenciesMutex.RUnlock()
	fake.releaseCachedDependenciesMutex.RLock()
	defer fake.releaseCachedDependenciesMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDependencyManager) recordInvocation(key string, args []interface{}) {
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

var _ containerstore.DependencyManager = new(FakeDependencyManager)
