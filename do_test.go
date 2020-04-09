package overwrite

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypes(t *testing.T) {
	err := Do(nil, "")
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrDstNil), err.Error())

	err = Do("", nil)
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrSrcNil), err.Error())

	err = Do("", 1)
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrNotSameType), err.Error())

	type T struct{ N int }
	err = Do(&T{}, &T{})
	require.Nil(t, err)
}
