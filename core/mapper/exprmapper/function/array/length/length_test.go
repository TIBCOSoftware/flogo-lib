package length

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s = &Length{}

func TestLength(t *testing.T) {
	// Length(arr)
	// arr : input array

	// Should produce "3"
	arr := []string{"1", "2", "3"}
	sub := s.Eval(arr)
	fmt.Printf("Result [%v] should be equal to: 3\n", sub)
	assert.Equal(t, 3, sub)

	// Should produce "5"
	arr2 := []int{1, 2, 3, 4, 5}
	sub = s.Eval(arr2)
	fmt.Printf("Result [%v] should be equal to: 5\n", sub)
	assert.Equal(t, 5, sub)
}
