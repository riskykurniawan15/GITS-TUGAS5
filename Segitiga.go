package main

import (
	"fmt"
	"strings"
)

func main(){
	// Risky Kurniawan - ARS University
	fmt.Println("================================")
	fmt.Println("    PROGRAM CETAK SEGITIGA")
	fmt.Println("================================")
	var jenis string
	var jml, index int
	fmt.Print("Masukan jenis segitiga (terbalik / tidak) ? ")
	fmt.Scan(&jenis)
	fmt.Print("Masukan jumlah baris ? ")
	fmt.Scan(&jml)

	jenis = strings.ToLower(jenis)

	if jenis == "terbalik" {
		index = (jml*2)-1
		for i:=jml; i>=1; i-- {
			for j:=0; j<(jml-i); j++ {
				fmt.Print(" ")
			}
			for k:=1; k<=index; k++ {
				fmt.Print("*")
			}
			fmt.Println()
			index-=2
		}
	}else if jenis == "tidak"{
		index = 1
		for i:=1; i<=jml; i++ {
			for j:=(jml-i); j>=0; j-- {
				fmt.Print(" ")	
			}	
			for k:=1; k<=index; k++ {
				fmt.Print("*")	
			}	
			fmt.Println()
			index += 2
		}
	}else{
		fmt.Println("Pilihan tidak tersedia!!")
	}
}