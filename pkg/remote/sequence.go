package remote

import (
	"sync"
	"time"
)

var value = 0
var mockValue = 0

type Sequence interface {
	GetNext() int
	Reset()
	GetValue() int
}

type SequenceImpl struct {
	wait  time.Duration
	mutex *sync.Mutex
}

func NewSequenceImpl() *SequenceImpl {
	return &SequenceImpl{100 * time.Millisecond, &sync.Mutex{}}
}

func (seq *SequenceImpl) GetValue() int {
	<-time.After(seq.wait)
	return value
}

func (seq *SequenceImpl) GetNext() int {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()
	value++
	<-time.After(seq.wait)
	return value
}

func (seq *SequenceImpl) Reset() {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()
	value = 0
	<-time.After(seq.wait)
}

type SequenceMock struct {
	mutex *sync.Mutex
}

func NewSequenceMock() *SequenceMock {
	return &SequenceMock{&sync.Mutex{}}
}

func (seq *SequenceMock) GetValue() int {
	return mockValue
}

func (seq *SequenceMock) GetNext() int {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()
	mockValue++
	return mockValue
}

func (seq *SequenceMock) Reset() {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()
	mockValue = 0
}
