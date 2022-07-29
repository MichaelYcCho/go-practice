package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGolang(t *testing.T) {
	t.Run("TrimSpace를 통한 공백제거", func(t *testing.T) {

		str := "Mike, Marry, John, Nami"
		//  , 을 기준으로 slicing
		actual := strings.Split(str, ",")

		// trime을 통한 element의 공백제거 (양끝의 공백만 제거하므로 actual단에서는 사용할수없다.)
		for i, v := range actual {
			fmt.Println(actual[i])
			actual[i] = strings.TrimSpace(v)
		}

		expected := []string{"Mike", "Marry", "John", "Nami"}
		// actual과 expected를 비교
		assert.Equal(t, expected, actual)
	})

	t.Run("ReplaceAll을 통한 공백제거", func(t *testing.T) {

		str := "Mike, Marry, John, Nami"
		//  , 을 기준으로 slicing, ReplaceAll을 통해 공백제거

		convert_string := strings.ReplaceAll(str, " ", "")
		actual := strings.Split(convert_string, ",")

		expected := []string{"Mike", "Marry", "John", "Nami"}
		// actual과 expected를 비교
		assert.Equal(t, expected, actual)
	})

}
