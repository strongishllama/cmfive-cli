package gen_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/strongishllama/cmfive-cli/pkg/gen"
)

func TestNewAction(t *testing.T) {
	setup(t)
	defer teardown(t)

	require.NoError(t, gen.NewModule("payroll"))
	require.NoError(t, gen.NewAction("payroll", "edit", http.MethodGet))
}
