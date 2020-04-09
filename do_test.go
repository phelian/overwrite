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

func TestSimpleTypes(t *testing.T) {
	type T struct {
		N   int    `overwrite:"true"`
		N1  int8   `overwrite:"true"`
		N2  int16  `overwrite:"true"`
		N3  int32  `overwrite:"true"`
		N4  int64  `overwrite:"true"`
		UN  uint   `overwrite:"true"`
		UN1 uint8  `overwrite:"true"`
		UN2 uint16 `overwrite:"true"`
		UN3 uint32 `overwrite:"true"`
		UN4 uint64 `overwrite:"true"`
		B   string `overwrite:"true"`
		G   string `overwrite:"false"`
		J   string
		H   string `overwrite:"true,omitempty"`
		L   bool   `overwrite:"true"`
	}

	tDst := &T{
		N:   1,
		N1:  11,
		N2:  111,
		N3:  1111,
		N4:  11111,
		UN:  1,
		UN1: 11,
		UN2: 111,
		UN3: 1111,
		UN4: 11111,
		B:   "foo",
		G:   "foop",
		J:   "foob",
		H:   "foot",
		L:   false,
	}

	tSrc := T{
		N:   42,
		N1:  41,
		N2:  4111,
		N3:  41111,
		N4:  411111,
		UN:  31,
		UN1: 31,
		UN2: 3111,
		UN3: 31111,
		UN4: 311111,
		B:   "bar",
		G:   "",
		J:   "",
		H:   "",
		L:   true,
	}

	err := Do(tDst, tSrc)
	require.Nil(t, err)
	require.Equal(t, tSrc.N, tDst.N)
	require.Equal(t, tSrc.N1, tDst.N1)
	require.Equal(t, tSrc.N2, tDst.N2)
	require.Equal(t, tSrc.N3, tDst.N3)
	require.Equal(t, tSrc.N4, tDst.N4)
	require.Equal(t, tSrc.UN, tDst.UN)
	require.Equal(t, tSrc.UN1, tDst.UN1)
	require.Equal(t, tSrc.UN2, tDst.UN2)
	require.Equal(t, tSrc.UN3, tDst.UN3)
	require.Equal(t, tSrc.UN4, tDst.UN4)
	require.Equal(t, tSrc.B, tDst.B)
	require.NotEqual(t, tSrc.G, tDst.G)
	require.NotEqual(t, tSrc.J, tDst.J)
	require.NotEqual(t, tSrc.H, tDst.H)
	require.Equal(t, tSrc.L, tDst.L)
}
}
