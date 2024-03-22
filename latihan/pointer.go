package main

import "fmt"

func gantiKata(k *string){
	*k = "Hari"
}

func main(){
	k := "Hari-hari"
	gantiKata(&k)
	fmt.Println(k)
}