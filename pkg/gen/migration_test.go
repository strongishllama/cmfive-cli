package gen_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/strongishllama/cmfive-cli/pkg/gen"
)

func TestNewMigration(t *testing.T) {
	setup(t)
	defer teardown(t)

	require.NoError(t, gen.NewModule("payroll"))
	require.NoError(t, gen.NewMigration("payroll", "InitialMigration"))
}
