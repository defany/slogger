package main

import (
	"fmt"
	slerr "github.com/defany/slogger/pkg/err"
)

func t() error {
	return slerr.WithSource(fmt.Errorf("some unknown cool error"))
}

func t2() error {
	return slerr.WithSource(t(), "omagad danila ti chto crazy")
}

func t3() error {
	return slerr.WithSource(t2(), "ochen kruto")
}

func main() {
	fmt.Println(slerr.WithSource(t3()))
}
