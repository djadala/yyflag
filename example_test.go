package yyflag_test

import (
	"flag"
	"fmt"
	"time"

	"github.com/djadala/yyflag"
)

func Example() {
	deft := time.Now().UTC().Truncate(24 * time.Hour)

	from := yyflag.New(deft)
	to := yyflag.New(deft.AddDate(0, 0, 1))

	flag.Var(&from, "f", "from date")
	flag.Var(&to, "t", "to date")

	flag.Parse()

	fmt.Println(from.Time())
	fmt.Println(to.Time())
}
