package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAdder := ":4000"
	tr := NewTCPTransport(listenAdder)
	assert.Equal(t, tr.listenAddress, listenAdder)

	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
