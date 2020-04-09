package overwrite

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTags(t *testing.T) {
	tgs, err := newTags("true,omitempty")
	require.Nil(t, err)
	require.True(t, tgs.overwrite)
	require.True(t, tgs.omitempty)

	tgs, err = newTags("true")
	require.Nil(t, err)
	require.True(t, tgs.overwrite)
	require.False(t, tgs.omitempty)

	tgs, err = newTags("false")
	require.Nil(t, err)
	require.False(t, tgs.overwrite)
	require.False(t, tgs.omitempty)

	tgs, err = newTags("gurka,omitempty")
	require.NotNil(t, err)
	fmt.Println(err.Error())
	require.True(t, errors.Is(err, ErrTagValueWrong))

	tgs, err = newTags("true,false,omitempty")
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrTagValueWrong))

	tgs, err = newTags("")
	require.Nil(t, err)
	require.False(t, tgs.overwrite)
	require.False(t, tgs.omitempty)
}
