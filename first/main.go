package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(`Anoop G Nair`)
	fmt.Println(runtime.Version(), runtime.GOARCH, runtime.GOOS)
	hey()
	bye()
}
