package util

import "strings"

// FixCommaBasedStringSliceInput fixes params by combining successive values such that
// they are recovered due to strings split on a comma during CLI parsing.
//
// Background:
// Cobra CLI framework splits the input values on a string slice based on comma.
// https://github.com/spf13/pflag/blob/9a97c102cda95a86cec2345a6f09f55a939babf5/string_slice.go#L85
// Therefore, a custom formatted input is unable to use comma in the format.
// This command joins the split components.
//
// Recommendation:
// Do not use comma in the value to avoid use of such custom parsing logic.
func FixCommaBasedStringSliceInput(params, args []string) []string {
	if len(args) == 0 {
		return params
	}

	var out []string
	for i := 0; i < len(params); i++ {
		combo := false

		// try to see combining with next element matches with an entry in args
		if i < len(params)-1 {
			try := params[i] + "," + params[i+1]
			for _, arg := range args {
				if strings.Contains(arg, try) {
					out = append(out, try)
					combo = true
					break
				}
			}
		}

		if !combo {
			out = append(out, params[i])
		} else {
			i += 1
		}
	}

	return out
}
