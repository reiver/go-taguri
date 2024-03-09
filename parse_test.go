package taguri_test

import (
	"testing"

	"github.com/reiver/go-taguri"
)

func TestParse(t *testing.T) {

	tests := []struct{
		TagURI string
		Expected taguri.TagURI
	}{
		{
			TagURI: "tag:,:",
			Expected: taguri.TagURI{},
		},
		{
			TagURI: "tag:,:#",
			Expected: taguri.TagURI{},
		},



		{
			TagURI: "tag:example.com,:",
			Expected: taguri.TagURI{
				AuthorityName: "example.com",
			},
		},
		{
			TagURI: "tag:joeblow@example.com,:#",
			Expected: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
			},
		},



		{
			TagURI: "tag:example.com,2024:",
			Expected: taguri.TagURI{
				AuthorityName: "example.com",
				Date: "2024",
			},
		},
		{
			TagURI: "tag:joeblow@example.com,2024:#",
			Expected: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024",
			},
		},



		{
			TagURI: "tag:example.com,2024-03:",
			Expected: taguri.TagURI{
				AuthorityName: "example.com",
				Date: "2024-03",
			},
		},
		{
			TagURI: "tag:joeblow@example.com,2024-03:#",
			Expected: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024-03",
			},
		},



		{
			TagURI: "tag:example.com,2024-03-08:",
			Expected: taguri.TagURI{
				AuthorityName: "example.com",
				Date: "2024-03-08",
			},
		},
		{
			TagURI: "tag:joeblow@example.com,2024-03-08:#",
			Expected: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024-03-08",
			},
		},



		{
			TagURI: "tag:example.com,2024-03-08:note-rRjM2W0g6p",
			Expected: taguri.TagURI{
				AuthorityName: "example.com",
				Date: "2024-03-08",
				Specific: "note-rRjM2W0g6p",
			},
		},
		{
			TagURI: "tag:joeblow@example.com,2024-03-08:note-rRjM2W0g6p#",
			Expected: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024-03-08",
				Specific: "note-rRjM2W0g6p",
			},
		},



		{
			TagURI: "tag:joeblow@example.com,2024-03-08:note-rRjM2W0g6p#banana-split",
			Expected: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024-03-08",
				Specific: "note-rRjM2W0g6p",
				Fragment: "banana-split",
			},
		},
	}

	for testNumber, test := range tests {

		var actual taguri.TagURI

		err := taguri.Parse(&actual, test.TagURI)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("TAG-URI: %q", test.TagURI)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the expected 'parsed-tag-uri' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("TAG-URI: %q", test.TagURI)
				continue
			}
		}
	}
}
