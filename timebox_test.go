package timebox

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNotAFunction(t *testing.T) {
	i := 0
	_, err := Timebox(time.Second, i)
	require.IsType(t, NotAFunctionError{}, err)
}

func TestArguments(t *testing.T) {
	fn := func(arg int) int {
		return arg * 2
	}

	ret, err := Timebox(time.Second, fn, 2)
	require.NoError(t, err)
	require.Len(t, ret, 1)
	require.Equal(t, 4, ret[0])
}

func TestWithoutArguments(t *testing.T) {
	fn := func() int {
		return 4
	}
	ret, err := Timebox(time.Second, fn)
	require.NoError(t, err)
	require.Len(t, ret, 1)
	require.Equal(t, 4, ret[0])
}

func TestWithInvalidArguments(t *testing.T) {
	t.Run("Invalid Count", func(t *testing.T) {
		fn := func(int) int {
			return 4
		}
		_, err := Timebox(time.Second, fn)
		require.Error(t, err)
	})
	t.Run("Invalid Type", func(t *testing.T) {
		fn := func(string) int {
			return 4
		}
		_, err := Timebox(time.Second, fn, 1)
		require.Error(t, err)
	})
}

func TestTimeout(t *testing.T) {
	start := time.Now()
	_, err := Timebox(time.Second, time.Sleep, time.Second*5)
	require.IsType(t, TimeoutError{}, err)
	require.True(t, time.Since(start).Seconds() < 2)
}
