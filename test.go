package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func _main() {
	startTime := time.Now()
	fmt.Println(startTime, cronexpr.MustParse("* * * * *").Next(time.Now().Add(-time.Minute)))
	fmt.Println(startTime == cronexpr.MustParse("* * * * *").Next(time.Now().UTC().Add(-time.Minute)))
}
