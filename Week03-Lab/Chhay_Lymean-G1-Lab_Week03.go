package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"fmt"
	"os"
)

var ex int
var Input1, Input2 string
var choice int

func main () {
	
	for {
		fmt.Println("1. Exercise 01 \n2. Exercise 02 \n3. Exercise 03 \n4. Exercise 04 \n5. Exercise 05")
		fmt.Print("Please choose exercise: ")
		fmt.Scan(&ex)
		if ex == 1 {
			//Lab 0
			ProofMe()
		} else if ex == 2 {
			//Lab 1
		}
	}
}

func ProofMe () {
	fmt.Println("===== Hashing Plain text =====")
	fmt.Print("Please input 1st plain text: ")
	fmt.Scan(&Input1)
	fmt.Print("Please input 2nd plain text: ")
	fmt.Scan(&Input2)
	fmt.Println("1. MD5 \n2. SHA1 \n3. SHA256 \n4. SHA512 \n5. SHA3 \n6. Exit")
	fmt.Print("Please choose options: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		OutputA := md5.Sum([] byte(Input1))
		OutputB := md5.Sum([] byte(Input2))
		fmt.Printf("1st hash text: %x\n", OutputA)
		fmt.Printf("2nd hash text: %x\n", OutputB)
		if OutputA == OutputB {
			fmt.Println("==> Both of MD5 are MATCH")
		} else if OutputA != OutputB {
			fmt.Println("==> Both of MD5 are NOT MATCH")
		}
	case 2:
		OutputA := sha1.Sum([] byte(Input1))
		OutputB := sha1.Sum([] byte(Input2))
		fmt.Printf("1st hash text: %x\n", OutputA)
		fmt.Printf("2nd hash text: %x\n", OutputB)
		if OutputA == OutputB {
			fmt.Println("==> Both of SHA1 are MATCH")
		} else if OutputA != OutputB {
			fmt.Println("==> Both of SHA1 are NOT MATCH")
		}
	case 3:
		OutputA := sha256.Sum256([] byte(Input1))
		OutputB := sha256.Sum256([] byte(Input2))
		fmt.Printf("1st hash text: %x\n", OutputA)
		fmt.Printf("2nd hash text: %x\n", OutputB)
		if OutputA == OutputB {
			fmt.Println("==> Both of SHA256 are MATCH")
		} else if OutputA != OutputB {
			fmt.Println("==> Both of SHA256 are NOT MATCH")
		}
	case 4:
		OutputA := sha512.Sum512([] byte(Input1))
		OutputB := sha512.Sum512([] byte(Input2))
		fmt.Printf("1st hash text: %x\n", OutputA)
		fmt.Printf("2nd hash text: %x\n", OutputB)
		if OutputA == OutputB {
			fmt.Println("==> Both of SHA512 are MATCH")
		} else if OutputA != OutputB {
			fmt.Println("==> Both of SHA512 are NOT MATCH")
		}
	case 5:
		OutputA := sha3.Sum512([] byte(Input1))
		OutputB := sha3.Sum512([] byte(Input2))
		fmt.Printf("1st hash text: %x\n", OutputA)
		fmt.Printf("2nd hash text: %x\n", OutputB)
		if OutputA == OutputB {
			fmt.Println("==> Both of SHA3 are MATCH !!!")
		} else if OutputA != OutputB {
			fmt.Println("==> Both of SHA3 are NOT MATCH !!!")
		}
	case 6:
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
	}
}

func Lab1 () {

}