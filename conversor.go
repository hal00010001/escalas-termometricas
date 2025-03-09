package main

import "fmt"

const ebulicaoK = 373

func main() {
	tempK := ebulicaoK
	tempC := tempK - 273
	fmt.Printf("A temperatura de ebulição da água em K = %d e a temperatura de ebulição da água em ºC = %d \n", tempK, tempC)
}
