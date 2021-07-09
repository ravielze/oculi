package radix36

import "strings"

var (
	b36 = []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z'}
)

func findBytes(f byte, l, r int) int {
	if r >= l {
		mid := l + (r-l)/2
		if b36[mid] == f {
			return mid
		} else if b36[mid] > f {
			return findBytes(f, l, mid-1)
		} else {
			return findBytes(f, mid+1, r)
		}
	}
	return -1
}

func Validate(val string) bool {
	val = strings.ToUpper(val)
	for _, v := range []byte(val) {
		if idx := findBytes(v, 0, 35); idx == -1 {
			return false
		}
	}
	return true
}
