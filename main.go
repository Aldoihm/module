package main

import (
	"fmt"
	"strings"

	"github.com/Aldoihm/module/slices"
	"rsc.io/quote/v3"
)

func main() {
	list := []string{"EDteam", "gophers", "golang", "sotec", quote.HelloV3()}

	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "h")
	})
	fmt.Println(quote.Concurrency())
}
