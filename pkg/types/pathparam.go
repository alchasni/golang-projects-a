package types

import (
	"regexp"
)

type PathParam string

var (
	pathParamRegexp, _ = regexp.Compile(`:[a-zA-Z_]+`)
)

func (p PathParam) String() string {
	return string(p)
}

func (p PathParam) Normalized() string {
	return pathParamRegexp.FindString(p.String())
}

func (p PathParam) Valid() bool {
	return p.String() == p.Normalized()
}

func (p PathParam) Name() string {
	if p.Valid() {
		return p.String()[1:]
	}
	return ""
}
