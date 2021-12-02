package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// import (
// 	"fmt";
// )

type deck []string

func newDeck() deck {

	cards := deck{};

	suitCards := []string{"Diamond" , "Spade" , "Heart" , "Club"};
	suitValues := []string{"Ace" , "Two" , "Three" , "Four" , "Five" , "Six" , "Seven" , "Eight" , "Nine" , "Ten"};

	for i := range suitCards{
		for j := range suitValues{
			cards = append(cards,suitValues[j] + " of " + suitCards[i]); 
		}
	}


	return cards;

}

func (d deck) printCards() {

	for _,card := range d{
		fmt.Println(card);
	}

}


func (d deck) convertToString() string {

	return strings.Join([]string(d),",");
}

func (d deck) saveToDisk(filename string) error {

	return ioutil.WriteFile(filename,[]byte(d.convertToString()),0666);	

}


func readFromDisk(filename string) deck {

	bs , err := ioutil.ReadFile(filename);

	if err != nil {
		fmt.Print(err);
		os.Exit(1);
	}
	
	s := strings.Split(string(bs),",");

	return deck(s);
}

