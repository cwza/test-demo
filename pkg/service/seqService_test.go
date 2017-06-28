package service

import (
	"testing"

	"github.com/cwza/test-demo/pkg/remote"
	"github.com/stretchr/testify/assert"
)

func TestSeqService_NewSeqService(t *testing.T) {
	seqService := NewSeqService(nil)
	assert.IsType(t, &remote.SequenceImpl{}, seqService.seq)
	seqService = NewSeqService(remote.NewSequenceMock())
	assert.IsType(t, &remote.SequenceMock{}, seqService.seq)
}

func TestSeqService_GetValue(t *testing.T) {
	seqService := NewSeqService(remote.NewSequenceMock())
	assert.IsType(t, 0, seqService.GetValue())
}

func TestSeqService_GetNextByStep(t *testing.T) {
	seqService := NewSeqService(remote.NewSequenceMock())
	defer seqService.Reset()
	seqService.Reset()
	seqService.GetValueByStep(1)
	assert.Equal(t, 1, seqService.GetValue())
	seqService.GetValueByStep(2)
	assert.Equal(t, 3, seqService.GetValue())
	seqService.GetValueByStep(2)
	assert.Equal(t, 5, seqService.GetValue())
}

func TestSeqService_Reset(t *testing.T) {
	seqService := NewSeqService(remote.NewSequenceMock())
	defer seqService.Reset()
	seqService.GetValueByStep(1)
	seqService.GetValueByStep(2)
	seqService.GetValueByStep(3)
	seqService.Reset()
	assert.Equal(t, 0, seqService.GetValue())
}
