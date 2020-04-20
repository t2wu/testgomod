package subpackage2

import (
	"log"

	"github.com/t2wu/testgomod/subpackage"
)

// Bar uses Foo and print
func Bar() {
	subpackage.Foo()
	log.Println("Bar ran")
}
