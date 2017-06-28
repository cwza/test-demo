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
	seq := remote.NewSequenceMock()
	seq.Reset()
	seqService := NewSeqService(seq)
	assert.Equal(t, 0, seqService.GetValue())
}

func TestSeqService_GetNextByStep(t *testing.T) {
	seq := remote.NewSequenceMock()
	defer seq.Reset()
	seq.Reset()
	seqService := NewSeqService(seq)
	seqService.GetValueByStep(1)
	assert.Equal(t, 1, seq.GetValue())
	seqService.GetValueByStep(2)
	assert.Equal(t, 3, seq.GetValue())
	seqService.GetValueByStep(2)
	assert.Equal(t, 5, seq.GetValue())
}

func TestSeqService_Reset(t *testing.T) {
	seq := remote.NewSequenceMock()
	defer seq.Reset()
	seqService := NewSeqService(seq)
	seqService.GetValueByStep(1)
	seqService.GetValueByStep(2)
	seqService.GetValueByStep(3)
	seqService.Reset()
	assert.Equal(t, 0, seq.GetValue())
}
