package gen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/strongishllama/cmfive-cli/pkg/gen"
)

func TestNewModel(t *testing.T) {
	setup(t)
	defer teardown(t)

	require.NoError(t, gen.NewModule("payroll"))
	require.NoError(t, gen.NewModel("payroll", "Employee"))
}
