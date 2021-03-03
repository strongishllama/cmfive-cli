package gen_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/strongishllama/cmfive-cli/pkg/gen"
)

func TestNewModule(t *testing.T) {
	setup(t)
	defer teardown(t)

	require.NoError(t, gen.NewModule("payroll"))
}
