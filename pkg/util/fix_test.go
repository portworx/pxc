package util

import "testing"

// TestFix checks if fixing a comma separated value is working in expected scenarios.
func TestFix(t *testing.T) {
	testTable := []struct {
		name     string
		args     []string
		params   []string
		expected []string
	}{
		{
			name:     "a",
			args:     []string{"-f", "a", "-f", "b,c"},
			params:   []string{"a", "b", "c"},
			expected: []string{"a", "b,c"},
		},
		{
			name:     "b",
			args:     []string{"-f", "a,b", "-f", "c"},
			params:   []string{"a", "b", "c"},
			expected: []string{"a,b", "c"},
		},
		{
			name:     "c",
			args:     []string{"-f", "a", "-f", "b", "-f", "c"},
			params:   []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "d",
			args:     []string{},
			params:   []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, test := range testTable {
		out := FixCommaBasedStringSliceInput(test.params, test.args)

		if len(out) != len(test.expected) {
			t.Fatal("test:", test.name, ", out len is:", len(out), ", expected:", len(test.expected))
		}

		for i := range out {
			if out[i] != test.expected[i] {
				t.Fatal("test:", test.name, ", expected:", test.expected[i], ", received:", out[i])
			}
		}
	}
}
