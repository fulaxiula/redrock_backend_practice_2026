package syntax

import "fmt"

const (
	MaxScore = 100
	minScore = 0
)

func Add(a, b int) int {
	return a + b
}

func add(a, b int) int {
	return a + b
}

func ClassifyNumber(n int) string {
	if n > 0 {
		return "positive"
	} else if n == 0 {
		return "zero"
	} else {
		return "negative"
	}
}

func SumTo(n int) int {
	if n <= 0 {
		return 0
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func Describe(name string, age int) string {
	return fmt.Sprintf("%s 今年 %d 岁", name, age)
}

func Clamp(score int) int {
	if score < minScore {
		return minScore
	}
	if score > MaxScore {
		return MaxScore
	}
	return score
}

func InitialCount() int {
	var count int
	return count
}