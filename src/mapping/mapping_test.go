package mapping

import "testing"

type mock struct {
	Name  string
	Value string
}

func TestMapping(t *testing.T) {

	in := [3]mock{
		mock{
			Name:  "spam",
			Value: "spamval",
		},
		mock{
			Name:  "ham",
			Value: "spamval",
		},
		mock{
			Name:  "egg",
			Value: "spamval",
		},
	}

	out := map[string]mock{
		"spam": mock{
			Name:  "spam",
			Value: "spamval",
		},
		"ham": mock{
			Name:  "ham",
			Value: "spamval",
		},
		"egg": mock{
			Name:  "egg",
			Value: "spamval",
		},
	}

	x, _ := Mapping(in, "Name")

	for k, v := range x {

		if v != out[k] {
			t.Fail()
		}
	}

}
