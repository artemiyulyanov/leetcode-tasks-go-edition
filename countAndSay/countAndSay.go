package countAndSay

import "strconv"

func CountAndSay(n int) string {
	if n <= 1 {
		return "1"
	}

	prev := CountAndSay(n - 1)

	cur, ans, count := prev[0], "", 0

	for i := 0; i < len(prev); i++ {
		if prev[i] == cur {
			count += 1
		} else {
			ans += strconv.Itoa(count) + string(cur)
			cur, count = prev[i], 1
		}
	}

	ans += strconv.Itoa(count) + string(cur)

	return ans
}
