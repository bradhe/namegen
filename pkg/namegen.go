package namegen

import (
	"fmt"
	"math/rand"
	"strings"
)

func randomString(str []string) string {
	return str[rand.Intn(len(str))]
}

func randomSuffix() string {
	return fmt.Sprintf("%03d", rand.Intn(999))
}

func Gen() string {
	terms := []string{
		randomString(Adjectives),
		randomString(Nouns),
		randomSuffix(),
	}

	return strings.Join(terms, "-")
}
