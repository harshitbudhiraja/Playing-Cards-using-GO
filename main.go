package main

import (
	"fmt"
	"strings"
)

func main() {

	d := []string{"Red","Yellow","Green"};
	ans := strings.Join(d,",");

	// fmt.Println(ans);

	// cards := newDeck()

	cards := readFromDisk("newFile");
	
	fmt.Println(cards);


}