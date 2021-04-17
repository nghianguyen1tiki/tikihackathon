package lang

import "github.com/spf13/cast"

func StringSliceToIntSlice(in []string) []int {
	out := make([]int, len(in))
	for i, str := range in {
		out[i] = cast.ToInt(str)
	}
	return out
}
