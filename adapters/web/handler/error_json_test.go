package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "test message"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"test message"}`), result)
}
