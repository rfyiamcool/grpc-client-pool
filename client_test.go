package grpcpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractName(t *testing.T) {
	s := &ServiceClientPool{}
	name := s.ExtractServiceName("/api.v1.election.CandidateSvc/Register")

	assert.Equal(t, name, "/api.v1.election.CandidateSvc")
}

func TestHostServiceNames(t *testing.T) {
	m := NewTargetServiceNames()
	m.Set("127.0.0.1:8080", "/api.v1.risk", "/api.v1.ws")
	assert.Equal(t, m.List()["127.0.0.1:8080"][0], "/api.v1.risk")
	assert.Equal(t, m.List()["127.0.0.1:8080"][1], "/api.v1.ws")
}
