// +build ignore

package main

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/velour/feedme/webfeed"
)

func main() {
	f, err := webfeed.Read(os.Stdin)
	if err != nil {
		panic(err)
	}
	spew.Dump(f)
}
