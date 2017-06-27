package service

import (
	"testing"

	"github.com/cwza/test-demo/pkg/remote"
	"github.com/stretchr/testify/assert"
)

func TestSeqService_GetNextByStep(t *testing.T) {
	seqService := NewSeqService(remote.NewSequenceMock())
	defer seqService.Reset()
	seqService.Reset()
	seqService.GetValueByStep(1)
	seqService.GetValueByStep(2)
	assert.Equal(t, 3, seqService.GetValue())
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
