package generators

import (
	"fmt"
	"github.com/patrick-me/tinyUrl/utils"
	"math/rand"
	"strconv"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type SimpleRandGenerator struct{}

func (g *SimpleRandGenerator) GetRandURL() string {
	val, err := strconv.Atoi(utils.GetEnv("DEFAULT_URL_LEN", "10"))
	if err != nil {
		fmt.Println("Can't parse default url len")
	}
	return g.GetNRandURL(val)
}

func (g *SimpleRandGenerator) GetNRandURL(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
