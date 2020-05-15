package talk

import (
	"fmt"

	"rsc.io/quote"
)

func Say(something string) string {
	val := fmt.Sprintf("Hello from speak %s", something)

	return val
}

func Quote() string {
	return quote.Hello()
}
