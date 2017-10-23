package strrev_test

import (
	"fmt"

	"4d63.com/strrev"
)

func ExampleReverse() {
	r := strrev.Reverse("abcde丂g")

	fmt.Println(r)

	// Output: g丂edcba
}
