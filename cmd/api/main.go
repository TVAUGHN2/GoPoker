package main

import (
	"github.com/tvaughn2/GoPoker/api/router"
)

/* TODOs:
 *  1. Add concurrency.
 *  2. Apply idiomatic Golang test harness.
 *	3. Convert fmt to logging framework. (glog for lightweight use w/ ability to set logging level)
 *	4. Create logging interface
 */

const RUN_TESTS bool = true

func main() {

	router.HandleRequests()
}
