package code8_15

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse가 입력값인 문자열을 스페이스 단위로 단어를 분리한 후 integer로 반환한다.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()
	fields := strings.Fields(input)
	numbers, _ = fields2numbers(fields)
	return numbers, nil
}
func fields2numbers(fields []string) ([]int, error) {
	numbers := make([]int, 0, len(fields))
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
