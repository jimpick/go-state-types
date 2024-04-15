package builtin

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jimpick/go-state-types/abi"
)

func TestGenerateMethodNum(t *testing.T) {

	methodNum, err := GenerateFRCMethodNum("Receive")
	require.NoError(t, err)
	require.Equal(t, methodNum, abi.MethodNum(3726118371))
}
