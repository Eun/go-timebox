package timebox

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNotAFunction(t *testing.T) {
	i := 0
	_, err := Timebox(time.Second, i)
	require.True(t, IsNotAFunctionError(err))
	require.Equal(t, NotAFunctionError{}.Error(), err.(NotAFunctionError).Error())
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
	require.True(t, IsTimeoutError(err))
	require.Equal(t, TimeoutError{}.Error(), err.(TimeoutError).Error())
	require.True(t, time.Since(start).Seconds() < 2)
}

func TestInfiniteTimeout(t *testing.T) {
	start := time.Now()
	_, err := Timebox(0, time.Sleep, time.Second)
	require.NoError(t, err)
	require.True(t, time.Since(start).Seconds() >= 1)
}
