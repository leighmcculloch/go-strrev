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

func ExampleReverseCombining() {
	r := strrev.ReverseCombining("abcdef\u0301\u031dg")

	fmt.Println(r)

	// Output: gf̝́edcba
}
