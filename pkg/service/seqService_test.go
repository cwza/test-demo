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
	assert.Equal(t, 1, seqService.GetNextByStep(1))
	assert.Equal(t, 3, seqService.GetNextByStep(2))
	assert.Equal(t, 5, seqService.GetNextByStep(2))
}

func TestSeqService_Reset(t *testing.T) {
	seq := remote.NewSequenceMock()
	defer seq.Reset()

	seqService := NewSeqService(seq)
	seq.GetNext()
	seq.GetNext()
	seq.GetNext()
	seqService.Reset()

	assert.Equal(t, 0, seq.GetValue())
}
