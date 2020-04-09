package overwrite

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type tags struct {
	overwrite bool
	omitempty bool
}

var (
	tagValueOmitempty = "omitempty"
	tagValueTrue      = "true"
	tagValueFalse     = "false"

	// ErrTagValueWrong values for wrong tag values
	ErrTagValueWrong = errors.New("Wrong tag value")
)

func newTags(input string) (tags, error) {
	tgs := tags{}
	inputTags := strings.Split(input, ",")

	switch len(inputTags) {
	case 2:
		omitempty := inputTags[1]
		if omitempty != tagValueOmitempty {
			return tgs, fmt.Errorf("%s not valid tagvalue for second value %w", omitempty, ErrTagValueWrong)
		}
		tgs.omitempty = omitempty == tagValueOmitempty
		fallthrough
	case 1:
		var err error
		if inputTags[0] != "" {
			tgs.overwrite, err = strconv.ParseBool(inputTags[0])
			if err != nil {
				return tgs, fmt.Errorf("%s %w", err.Error(), ErrTagValueWrong)
			}
		}
	case 0:
		return tgs, nil
	default:
		return tgs, fmt.Errorf("Wrong number of tags, only true/false,omitempty allowed (%s) %w", input, ErrTagValueWrong)
	}
	return tgs, nil
}
