package cmfive_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/strongishllama/cmfive-cli/pkg/cmfive"
)

func TestNewMigration(t *testing.T) {
	setup(t)
	defer teardown(t)

	require.NoError(t, cmfive.NewModule("payroll"))
	require.NoError(t, cmfive.NewMigration("payroll", "InitialMigration"))
}
