package service

import (
	"github.com/cwza/test-demo/pkg/remote"
)

type SeqService struct {
	seq remote.Sequence
}

func NewSeqService(seq remote.Sequence) SeqService {
	if seq == nil {
		return SeqService{remote.NewSequenceImpl()}
	}
	return SeqService{seq}
}

func (seqService *SeqService) GetNextByStep(step int) int {
	value := 0
	for i := 0; i < step; i++ {
		value = seqService.seq.GetNext()
	}
	return value
}

func (seqService *SeqService) GetValue() int {
	return seqService.seq.GetValue()
}

func (seqService *SeqService) Reset() {
	seqService.seq.Reset()
}
