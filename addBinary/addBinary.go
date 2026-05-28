package addBinary

import (
	"slices"
	"strconv"
	"strings"
)

func AddBinary(a *string, b *string) string {
	pushTrallingZeros(a, b)

	reverseString(a)
	reverseString(b)

	ans := ""
	potential := 0

	for i := 0; i < len(*a); i++ {
		s := int((*a)[i]-'0') + int((*b)[i]-'0') + potential

		ans += strconv.Itoa(s % 2)
		potential = s / 2
	}

	ans += strconv.Itoa(potential)

	reverseString(&ans)

	return ans
}

func reverseString(s *string) {
	runes := []rune(*s)

	slices.Reverse(runes)

	*s = string(runes)
}

func pushTrallingZeros(a *string, b *string) {
	*a = strings.Repeat("0", max(len(*a), len(*b))-len(*a)) + *a
	*b = strings.Repeat("0", max(len(*a), len(*b))-len(*b)) + *b
}
