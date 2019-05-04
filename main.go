package main

import (
	_ "github.com/lib/pq"
	"github.com/zamariola/go-echo-bootstrap/cmd"
)

func main() {
	cmd.Execute()
}
