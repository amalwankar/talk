package talk

import (
	"fmt"
)

func Say(something string) string {
	val := fmt.Sprintf("Hello from talk %s", something)

	return val
}

func Quote() string {
	return "test" //quote.Hello()
}
