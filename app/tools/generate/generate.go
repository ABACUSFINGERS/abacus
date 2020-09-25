package generate

import (
	"fmt"
	"math/rand"
)

func Uuid() string {
	b := make([]byte, 12)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}