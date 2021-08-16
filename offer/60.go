package main

import "fmt"

func dicesProbablity(n int)[]float64{
	tmp := []float64{0.16667,0.16667,0.16667,0.16667,0.16667,0.16667}
	if n == 1{
		return tmp
	}
	return nil
}

func main(){
	result := dicesProbablity(1)
	fmt.Println(result)

}
