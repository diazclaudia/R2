package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SpiralFibonacci_Ok(t *testing.T) {
	assertions := assert.New(t)

	response := [][]int{{0, 1}, {2, 1}}
	req, err := SpiralFibonacci(2, 2)

	assertions.Equal(response, req)
	assertions.Nil(err)
}

func Test_SpiralFibonacci_Err(t *testing.T) {
	assertions := assert.New(t)

	req, error := SpiralFibonacci(-2, 2)

	assertions.Equal([][]int{[]int(nil), []int(nil)}, req)
	assertions.Equal("not valid rows and cols", error.Error())
}
