package cmfive_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/strongishllama/cmfive-cli/cmfive"
)

func TestNewModule(t *testing.T) {
	require.NoError(t, cmfive.NewModule("test"))
}
