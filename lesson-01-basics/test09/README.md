package habits

import (
	"strconv"
)

func IsWeekend(day int) bool {
	switch day {
	case 6, 7:
		return true
	default:
		return false
	}
}

func Average(scores []int) float64 {
	if len(scores) == 0 {
		return 0
	}
	total := 0
	for _, s := range scores {
		total += s
	}
	return float64(total) / float64(len(scores))
}

func CountChars(s string) int {
	return len([]rune(s))
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func SumScores(items []interface{}) (int, error) {
	total := 0
	for _, item := range items {
		str, ok := item.(string)
		if !ok {
			return 0, strconv.ErrSyntax
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		total += num
	}
	return total, nil
}