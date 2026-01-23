package main

import "_/home/salehabdulla/Desktop/Coding/GoLangEssentials/GoLangWeb/helpers"

const numPool = 10

func CalcVal(intChan chan int) {
	randomNumber := helpers.RandNum(numPool)
	intChan <- randomNumber

}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalcVal(intChan)

	num := <-intChan
	println(num)

}
