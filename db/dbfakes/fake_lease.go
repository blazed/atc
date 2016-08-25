// This file was generated by counterfeiter
package dbfakes

import (
	"sync"

	"github.com/concourse/atc/db"
)

type FakeLease struct {
	AttemptSignStub        func() (bool, error)
	attemptSignMutex       sync.RWMutex
	attemptSignArgsForCall []struct{}
	attemptSignReturns     struct {
		result1 bool
		result2 error
	}
	BreakStub        func() error
	breakMutex       sync.RWMutex
	breakArgsForCall []struct{}
	breakReturns     struct {
		result1 error
	}
	AfterBreakStub        func(func() error)
	afterBreakMutex       sync.RWMutex
	afterBreakArgsForCall []struct {
		arg1 func() error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLease) Acquire() (bool, error) {
	fake.attemptSignMutex.Lock()
	fake.attemptSignArgsForCall = append(fake.attemptSignArgsForCall, struct{}{})
	fake.recordInvocation("Acquire", []interface{}{})
	fake.attemptSignMutex.Unlock()
	if fake.AttemptSignStub != nil {
		return fake.AttemptSignStub()
	} else {
		return fake.attemptSignReturns.result1, fake.attemptSignReturns.result2
	}
}

func (fake *FakeLease) AttemptSignCallCount() int {
	fake.attemptSignMutex.RLock()
	defer fake.attemptSignMutex.RUnlock()
	return len(fake.attemptSignArgsForCall)
}

func (fake *FakeLease) AttemptSignReturns(result1 bool, result2 error) {
	fake.AttemptSignStub = nil
	fake.attemptSignReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeLease) Release() error {
	fake.breakMutex.Lock()
	fake.breakArgsForCall = append(fake.breakArgsForCall, struct{}{})
	fake.recordInvocation("Release", []interface{}{})
	fake.breakMutex.Unlock()
	if fake.BreakStub != nil {
		return fake.BreakStub()
	} else {
		return fake.breakReturns.result1
	}
}

func (fake *FakeLease) BreakCallCount() int {
	fake.breakMutex.RLock()
	defer fake.breakMutex.RUnlock()
	return len(fake.breakArgsForCall)
}

func (fake *FakeLease) BreakReturns(result1 error) {
	fake.BreakStub = nil
	fake.breakReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLease) AfterRelease(arg1 func() error) {
	fake.afterBreakMutex.Lock()
	fake.afterBreakArgsForCall = append(fake.afterBreakArgsForCall, struct {
		arg1 func() error
	}{arg1})
	fake.recordInvocation("AfterRelease", []interface{}{arg1})
	fake.afterBreakMutex.Unlock()
	if fake.AfterBreakStub != nil {
		fake.AfterBreakStub(arg1)
	}
}

func (fake *FakeLease) AfterBreakCallCount() int {
	fake.afterBreakMutex.RLock()
	defer fake.afterBreakMutex.RUnlock()
	return len(fake.afterBreakArgsForCall)
}

func (fake *FakeLease) AfterBreakArgsForCall(i int) func() error {
	fake.afterBreakMutex.RLock()
	defer fake.afterBreakMutex.RUnlock()
	return fake.afterBreakArgsForCall[i].arg1
}

func (fake *FakeLease) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attemptSignMutex.RLock()
	defer fake.attemptSignMutex.RUnlock()
	fake.breakMutex.RLock()
	defer fake.breakMutex.RUnlock()
	fake.afterBreakMutex.RLock()
	defer fake.afterBreakMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLease) recordInvocation(key string, args []interface{}) {
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

var _ db.Lock = new(FakeLease)
