package main

import "fmt"

// check if a number given is odd or even
func main(){

	sl := []int{0,1,2,3,4,5,6,7,8,9,10};

	for _, x := range sl {
		if x % 2 == 1 {
			fmt.Println("odd value at ", x);
		}else{
			fmt.Println("even value at" , x);
		}
	}

}