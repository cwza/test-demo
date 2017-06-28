package remote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeq_GetValue(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	seq := NewSequenceImpl()
	seq.Reset()
	assert.Equal(t, 0, seq.GetValue())
}

func TestSeq_GetNext(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	seq := NewSequenceImpl()
	defer seq.Reset()
	seq.Reset()
	assert.Equal(t, 1, seq.GetNext())
	assert.Equal(t, 2, seq.GetNext())
	assert.Equal(t, 3, seq.GetNext())
}

func TestSeq_Reset(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	seq := NewSequenceImpl()
	defer seq.Reset()
	seq.GetNext()
	seq.GetNext()
	seq.Reset()
	assert.Equal(t, 0, value)
}

func TestSeqMock_GetValue(t *testing.T) {
	seq := NewSequenceMock()
	seq.Reset()
	assert.Equal(t, 0, seq.GetValue())
}

func TestSeqMock_GetNext(t *testing.T) {
	seq := NewSequenceMock()
	defer seq.Reset()
	seq.Reset()
	assert.Equal(t, 1, seq.GetNext())
	assert.Equal(t, 2, seq.GetNext())
	assert.Equal(t, 3, seq.GetNext())
}

func TestSeqMock_Reset(t *testing.T) {
	seq := NewSequenceMock()
	defer seq.Reset()
	seq.GetNext()
	seq.GetNext()
	seq.Reset()
	assert.Equal(t, 0, value)
}
