package taguri_test

import (
	"testing"

	"github.com/reiver/go-taguri"
)

func TestTagURI_String(t *testing.T) {

	tests := []struct{
		TagURI taguri.TagURI
		Expected string
	}{
		{
			TagURI: taguri.TagURI{},
			Expected: "tag:,:",
		},



		{
			TagURI: taguri.TagURI{
				AuthorityName: "example.com",
			},
			Expected: "tag:example.com,:",
		},
		{
			TagURI: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
			},
			Expected: "tag:joeblow@example.com,:",
		},









		{
			TagURI: taguri.TagURI{
				AuthorityName: "example.com",
				Date: "2024",
			},
			Expected: "tag:example.com,2024:",
		},
		{
			TagURI: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024",
			},
			Expected: "tag:joeblow@example.com,2024:",
		},



		{
			TagURI: taguri.TagURI{
				AuthorityName: "example.com",
				Date: "2024-03",
			},
			Expected: "tag:example.com,2024-03:",
		},
		{
			TagURI: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024-03",
			},
			Expected: "tag:joeblow@example.com,2024-03:",
		},



		{
			TagURI: taguri.TagURI{
				AuthorityName: "example.com",
				Date: "2024-03-08",
			},
			Expected: "tag:example.com,2024-03-08:",
		},
		{
			TagURI: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024-03-08",
			},
			Expected: "tag:joeblow@example.com,2024-03-08:",
		},









		{
			TagURI: taguri.TagURI{
				AuthorityName: "example.com",
				Date: "2024",
				Specific: "pEw5AGHy8DDjz7VoLN7",
			},
			Expected: "tag:example.com,2024:pEw5AGHy8DDjz7VoLN7",
		},
		{
			TagURI: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024",
				Specific: "pEw5AGHy8DDjz7VoLN7",
			},
			Expected: "tag:joeblow@example.com,2024:pEw5AGHy8DDjz7VoLN7",
		},



		{
			TagURI: taguri.TagURI{
				AuthorityName: "example.com",
				Date: "2024-03",
				Specific: "pEw5AGHy8DDjz7VoLN7",
			},
			Expected: "tag:example.com,2024-03:pEw5AGHy8DDjz7VoLN7",
		},
		{
			TagURI: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024-03",
				Specific: "pEw5AGHy8DDjz7VoLN7",
			},
			Expected: "tag:joeblow@example.com,2024-03:pEw5AGHy8DDjz7VoLN7",
		},



		{
			TagURI: taguri.TagURI{
				AuthorityName: "example.com",
				Date: "2024-03-08",
				Specific: "pEw5AGHy8DDjz7VoLN7",
			},
			Expected: "tag:example.com,2024-03-08:pEw5AGHy8DDjz7VoLN7",
		},
		{
			TagURI: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024-03-08",
				Specific: "pEw5AGHy8DDjz7VoLN7",
			},
			Expected: "tag:joeblow@example.com,2024-03-08:pEw5AGHy8DDjz7VoLN7",
		},









		{
			TagURI: taguri.TagURI{
				Fragment: "banana-split",
			},
			Expected: "tag:,:#banana-split",
		},



		{
			TagURI: taguri.TagURI{
				AuthorityName: "example.com",
				Date: "2024-03-08",
				Specific: "pEw5AGHy8DDjz7VoLN7",
				Fragment: "banana-split",
			},
			Expected: "tag:example.com,2024-03-08:pEw5AGHy8DDjz7VoLN7#banana-split",
		},
		{
			TagURI: taguri.TagURI{
				AuthorityName: "joeblow@example.com",
				Date: "2024-03-08",
				Specific: "pEw5AGHy8DDjz7VoLN7",
				Fragment: "banana-split",
			},
			Expected: "tag:joeblow@example.com,2024-03-08:pEw5AGHy8DDjz7VoLN7#banana-split",
		},
	}

	for testNumber, test := range tests {

		actual := test.TagURI.String()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual (serialized) 'tar-uri' is not whatwa  expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("TAG-URI: %#v", test.TagURI)
			continue
		}
	}
}
