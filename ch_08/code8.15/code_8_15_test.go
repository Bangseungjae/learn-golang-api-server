package code8_15

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	str := "1 3"
	nums, err := Parse(str)
	if err != nil {
		fmt.Println("err TestParse")
	}
	expected := []int{1, 3}
	assert.Equal(t, expected, nums)
}
