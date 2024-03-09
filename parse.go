package taguri

import (
	"strings"
)

func Parse(dst *TagURI, value string) error {
	if nil == dst {
		return errNilDestination
	}

	var str string = value

	{
		const prefix string = "tag:"

		if !strings.HasPrefix(str, prefix) {
			return errNotTagURI
		}

		str = str[len(prefix):]
	}

	{
		var index int = strings.IndexByte(str, ',')
		if index < 0 {
			return errInvalidTagURI
		}

		dst.AuthorityName = str[:index]

		str = str[index+1:]
	}

	{
		var index int = strings.IndexByte(str, ':')
		if index < 0 {
			return errInvalidTagURI
		}

		dst.Date = str[:index]

		str = str[index+1:]
	}

	{
		var index int = strings.IndexByte(str, '#')
		switch {
		case index < 0:
			dst.Specific = str
		default:
			dst.Specific = str[:index]
			dst.Fragment = str[index+1:]
		}
	}

	return nil
}
