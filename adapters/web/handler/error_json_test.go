package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	require.Equal(t, string([]byte(`{"message":"Hello Json"}`)), string(result))
}
