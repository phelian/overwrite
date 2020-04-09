package overwrite

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypes(t *testing.T) {
	type T struct{ N int }

	err := Do(nil, "")
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrDstNil), err.Error())

	v := ""
	err = Do(&v, nil)
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrSrcNil), err.Error())

	err = Do(T{}, &T{})
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrDstNotPtr), err.Error())

	err = Do(&T{}, &T{})
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrSrcNotStruct), err.Error())

	err = Do(&v, T{})
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrNotSameType), err.Error())

	err = Do(&T{}, T{})
	require.Nil(t, err)
}
	require.Nil(t, err)
}
