package novas_test

import (
	"github.com/pebbe/novas"

	"fmt"
)

func ExampleEphInfo() {
	fmt.Println(novas.EphInfo())
	// Output:
	// JPL ephemeris DE405. Start JD = 2305424.50  End JD = 2525008.50
}
