// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch07/07_03/sqrt_test.go
package sqrt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)
	require.NoError(t, err)
	require.InDelta(t, 1.414214, val, 0.001)
}
