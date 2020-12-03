package environ_test

import (
	"github.com/google/uuid"
	"github.com/gopub/environ"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestEnvVar(t *testing.T) {
	key := "A1_B2_C3"
	val := uuid.New().String()
	err := os.Setenv(key, val)
	require.NoError(t, err)

	assert.Equal(t, val, environ.String(key, ""))
	valByLowercase := environ.String(strings.ToLower(key), "")
	assert.Equal(t, val, valByLowercase)
	valByDotKey := environ.String(strings.Replace(strings.ToLower(key), "_", ".", -1), "")
	assert.Equal(t, val, valByDotKey)
}
