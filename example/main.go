package main

import (
	"fmt"
	"io/ioutil"

	webassets "github.com/honmaple/go-webassets"
)

func main() {
	assets := webassets.New()
	r, err := assets.Run(
		[]string{
			"scss/main.scss",
		},
		[]string{
			"libscss",
			"cssmin",
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := ioutil.ReadAll(r)
	fmt.Println(string(b))
}
