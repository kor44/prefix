package prefix

import (
	"fmt"
	"regexp"
)

var (
	inputRE = regexp.MustCompile("^\\d+$")
)

// FromRange generates list of prefixes for given range.
// 'begin' and 'end' should
// - have the same length
// - contain only digits
// - 'begin' be less than 'end'
// in other case return error
func FromRange(begin, end string) ([]string, error) {
	if len(begin) != len(end) {
		return nil, fmt.Errorf("Begin and end values should have the same length")
	}

	if begin > end {
		return nil, fmt.Errorf("Begin should be less than end value")
	}

	if !inputRE.MatchString(begin) {
		return nil, fmt.Errorf("Begin string contains unexpected symbols")
	}

	if !inputRE.MatchString(end) {
		return nil, fmt.Errorf("End string contains unexpected symbols")
	}

	if begin == end {
		return []string{begin}, nil
	}

	var pos int
	for i := range begin {
		if begin[i] != end[i] {
			break
		}
		pos++
	}

	return helper(pos, []byte(begin), []byte(end)), nil
}

func helper(pos int, begin, end []byte) []string {
	var prefixes []string
	if isAll(begin[pos+1:], '0') && isAll(end[pos+1:], '9') {
		if begin[pos] == '0' && end[pos] == '9' {
			return append(prefixes, string(begin[:pos]))
		}

		for b := begin[pos]; b <= end[pos]; b++ {
			begin[pos] = b
			prefixes = append(prefixes, string(begin[:pos+1]))
		}

		return prefixes
	}

	// begin part
	if isAll(begin[pos+1:], '0') {
		prefixes = append(prefixes, string(begin[:pos+1]))
	} else {
		begin_max := fillBy(pos+1, begin, '9')
		prefixes = append(prefixes, helper(pos+1, begin, begin_max)...)
	}

	// middle part
	for i := begin[pos] + 1; i < end[pos]; i++ {
		begin[pos] = i
		prefixes = append(prefixes, string(begin[:pos+1]))
	}

	// end part
	if isAll(end[pos+1:], '9') {
		prefixes = append(prefixes, string(end[:pos+1]))
	} else {
		end_min := fillBy(pos+1, end, '0')
		prefixes = append(prefixes, helper(pos+1, end_min, end)...)
	}

	return prefixes
}

func isAll(data []byte, b byte) bool {
	for i := range data {
		if data[i] != b {
			return false
		}
	}
	return true
}

func fillBy(pos int, data []byte, b byte) []byte {
	result := make([]byte, len(data))

	copy(result, data)

	for i := pos; i < len(result); i++ {
		result[i] = b
	}

	return result
}
