package main

import (
	"fmt"
	"os"
	"poker/poker"
	"strings"
)

func main() {
	data, err := os.ReadFile("UTG.txt")

	if err != nil {
		fmt.Println(err)
	}

	str := string(data)
	words := strings.Split(str, ",")

	all := []poker.HandType{}
	for _, word := range words {
		hr := poker.NewHandRangeFromStr(word)
		fmt.Println(word, poker.ExpandHandRange(hr))
		all = append(all, poker.ExpandHandRange(hr)...)
	}
	fmt.Println(all)
}
