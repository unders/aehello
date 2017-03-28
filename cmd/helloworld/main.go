package main

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

var initTime = time.Now()

func main() {
	err := errors.New("err just a test")
	fmt.Println(err)

}
