// +build debug

package nebutil

import "fmt"

func NebPrint(str string, args ...interface{}) {
	fmt.Printf(str, args...)
}
