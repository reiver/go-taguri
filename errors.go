package taguri

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errInvalidTagURI  = erorr.Error("taguri: invalid tag-URI")
	errNilDestination = erorr.Error("taguri: nil destination")
	errNotTagURI      = erorr.Error("taguri: not tag-URI")
)
