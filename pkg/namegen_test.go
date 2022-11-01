package namegen

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	t.Run("emits elements in slug format", func(t *testing.T) {
		exp := regexp.MustCompile(`^[a-z]+-[a-z]+-\d{3}$`)

		for i := 0; i < 1000; i++ {
			slug := Gen()
			assert.True(t, exp.MatchString(slug), fmt.Sprintf("slug `%s` does not match regexp", slug))
		}
	})
}
