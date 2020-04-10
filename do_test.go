// Copyright 2020 Alexander Félix. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

func TestNestedStructs(t *testing.T) {
	type Tt struct {
		TtN  string `overwrite:"true"`
		TtN1 string
	}

	type T struct {
		N  int    `overwrite:"true"`
		B  string `overwrite:"true"`
		Tt Tt
	}

	type K struct {
		KN int    `overwrite:"true"`
		KB string `overwrite:"true"`
		T  T
	}

	tDst := &K{
		KN: 1,
		KB: "foo",
		T: T{
			N: 1337,
			B: "foot",
			Tt: Tt{
				TtN1: "polly",
			},
		},
	}

	tSrc := K{
		KN: 11,
		KB: "foos",
		T: T{
			N: 13375,
			B: "foots",
			Tt: Tt{
				TtN: "footz",
			},
		},
	}

	err := Do(tDst, tSrc)
	require.Nil(t, err)
	require.Equal(t, tSrc.KN, tDst.KN)
	require.Equal(t, tSrc.KB, tDst.KB)
	require.Equal(t, tSrc.T.B, tDst.T.B)
	require.Equal(t, tSrc.T.N, tDst.T.N)
	require.Equal(t, tSrc.T.Tt.TtN, tDst.T.Tt.TtN)
	require.NotEqual(t, tSrc.T.Tt.TtN1, tDst.T.Tt.TtN1)
}

func TestMap(t *testing.T) {
	type T struct {
		MSS map[string]string `overwrite:"true"`
		MIB map[int]bool      `overwrite:"true"`
	}

	tDst := &T{
		MSS: map[string]string{"foo": "bar"},
		MIB: map[int]bool{1: true, 2: false},
	}

	tSrc := T{
		MSS: map[string]string{"fooz": "baz"},
		MIB: map[int]bool{1: false, 2: true},
	}

	err := Do(tDst, tSrc)
	require.Nil(t, err)
	require.Equal(t, tSrc.MSS, tDst.MSS)
	require.Equal(t, tSrc.MIB, tDst.MIB)
}

func TestSliceAndArrays(t *testing.T) {
	type T struct {
		NS  []int     `overwrite:"true"`
		BS  []string  `overwrite:"true"`
		BS2 []string  `overwrite:"true"`
		BA  [2]string `overwrite:"true"`
	}

	tDst := &T{
		NS:  make([]int, 0),
		BS:  make([]string, 0),
		BS2: make([]string, 0),
		BA:  [2]string{"foo", "bar"},
	}
	tDst.NS = append(tDst.NS, 2, 3)
	tDst.BS = append(tDst.BS, "monkey", "ape")
	tDst.BS2 = append(tDst.BS, "monkey", "ape")

	tSrc := T{
		NS: make([]int, 0),
		BS: make([]string, 0),
		BA: [2]string{"dim", "sum"},
	}
	tSrc.NS = append(tSrc.NS, 42, 43)
	tSrc.BS = append(tSrc.BS, "kveik", "fox")

	err := Do(tDst, tSrc)
	require.Nil(t, err)
	require.Equal(t, tSrc.NS, tDst.NS)
	require.Equal(t, tSrc.BS, tDst.BS)
	require.Equal(t, tSrc.BS2, tDst.BS2)
	require.Equal(t, tSrc.BA, tDst.BA)
}

func TestSimpleTypes(t *testing.T) {
	type T struct {
		N   int     `overwrite:"true"`
		N1  int8    `overwrite:"true"`
		N2  int16   `overwrite:"true"`
		N3  int32   `overwrite:"true"`
		N4  int64   `overwrite:"true"`
		UN  uint    `overwrite:"true"`
		UN1 uint8   `overwrite:"true"`
		UN2 uint16  `overwrite:"true"`
		UN3 uint32  `overwrite:"true"`
		UN4 uint64  `overwrite:"true"`
		F1  float32 `overwrite:"true"`
		F2  float64 `overwrite:"true"`
		B   string  `overwrite:"true"`
		G   string  `overwrite:"false"`
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
		F1:  1.1,
		F2:  2.1,
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
		F1:  2.2,
		F2:  3.4,
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
	require.Equal(t, tSrc.F1, tDst.F1)
	require.Equal(t, tSrc.F2, tDst.F2)
	require.Equal(t, tSrc.B, tDst.B)
	require.NotEqual(t, tSrc.G, tDst.G)
	require.NotEqual(t, tSrc.J, tDst.J)
	require.NotEqual(t, tSrc.H, tDst.H)
	require.Equal(t, tSrc.L, tDst.L)
}
