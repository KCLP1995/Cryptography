package main

import (
    "fmt"
    "os"
    "encoding/base64"
    "encoding/hex"
)

var num1,num2 int
var com1, com2 int
var bit1, bit2 int
var op1, op2, op int
var con int
var txt string
var text string
var key byte
var en_de int

func main() {

    //Lab 01
    fmt.Print("Please input 2 numbers: ") //Print and Scan for input
    fmt.Scan(&num1, &num2)
    Lab01(num1, num2)

    //Lab 02
    fmt.Print("Please input 2 numbers: ") //Print and Scan for input
    fmt.Scan(&com1, &com2)
    //Lab02(com1, com2)

    //Lab 03 
    fmt.Print("Please input 2 number: ")
    fmt.Scan(&bit1, &bit2)
    Lab03(bit1, bit2)

    //Lab 04
    fmt.Println("1. Add \n2. Sub \n3. Mul \n4. Div \n5. Mod \n6. Exit")
    fmt.Print("Please choose the option: ")
    fmt.Scan(&op)
    fmt.Print("Please input value of 1st number: ")
    fmt.Scan(&op1)
    fmt.Print("Please input value of 2nd number: ")
    fmt.Scan(&op2)
    Lab04()

    //Lab 05
    fmt.Println("1. Text ==> Binary \n2. Text ==> Hexadecimal \n3. Text ==> Base64 \n4. Exit")
    fmt.Print("Please choose the option: ")
    fmt.Scan(&con)
    Lab05(con, txt)

    // Lab 06
	for {
		fmt.Println("1. Encryption \n2. Decryption \n3. Exit")
		fmt.Print("Choose an option: ")
		
		fmt.Scan(&en_de)

		xorEncrypt(en_de)
	}
}

func Lab01(num1, num2 int) {

    fmt.Println("operating 01: ", num1 == 2 , num2 == 2) //Println for output

    num1 += 2
    num2 += 2
    fmt.Println("operating 02: ", num1, num2)

    num1 -= 2
    num2 -= 2
    fmt.Println("operating 03: ", num1, num2) 

    num1 *= 2
    num2 *= 2
    fmt.Println("operating 04: ", num1, num2)

    num1 /= 2
    num2 /= 2
    fmt.Println("operating 05: ", num1, num2) 

    num1 %= 2
    num2 %= 2
    fmt.Println("operating 06: ", num1, num2)
}

func Lab02(com1, com2 int) {

    if com1 >= 0 && com2 >= 0 {
        fmt.Println("True: Both numbers are positive")
    } else if com1 >= 0 && com2 < 0{
        fmt.Println("False: First number is positive and Second number is negative")
    } else if com1 < 0 && com2 >= 0 {
        fmt.Println("False: First number is negative and Second number is positive")
    } else {
        fmt.Println("False: Both numbers are negative")
    }

    if com1 > com2 {
        fmt.Println("True: 1st number is bigger than 2nd number")
    } else if com2 > com1 {
        fmt.Println("True: 2nd number is bigger than 1st number")
    } else {
        fmt.Println("No number is bigger")
    }
    
    if com1 != com2 {
        fmt.Println("True: Both numbers are not equal")
    } else {
        fmt.Println("False: Both numbers are equal")
    }
}

func Lab03(bit1, bit2 int) {

    fmt.Printf("XOR operation: %b = %d \n", bit1 ^ bit2 , bit1 ^ bit2)
    fmt.Printf("NOT operation: %b %b \n", ^bit1 , ^bit2)
    fmt.Printf("OR operation: %b = %d \n", bit1 | bit2 , bit1 | bit2)
    fmt.Printf("AND operation: %b = %d \n", bit1 & bit2 , bit1 & bit2)
    fmt.Printf("Left Shift operation: %b %b \n", bit1 <<2 , bit2 <<2)
    fmt.Printf("Right Shift operation: %b %b ", bit1 >>3 , bit2 >>3)
}

func Lab04() {

    switch op {
    case 1: 
        fmt.Println("The Answer is: ", op1 + op2)
    case 2:
        fmt.Println("The Answer is: ", op1 - op2)
    case 3:
        fmt.Println("The Answer is: ", op1 * op2)
    case 4:
        fmt.Println("The Answer is: ", op1 / op2)
    case 5:
        fmt.Println("The Answer is: ", op1 % op2)
    case 6:
        os.Exit(0)
    }
}

func Lab05(con int, txt string) {

    switch con {
    case 1: 
        fmt.Print("Please input plain text: ")
        fmt.Scan(&txt)
        binaryStr := ""
        for _, c := range txt {
            binaryStr += fmt.Sprintf("%08b ", c) // 8-bit binary
        }
        fmt.Println("Binary Text:", binaryStr)
        return

    case 2:
        fmt.Print("Please input plain text: ")
        fmt.Scan(&txt) 
        hexStr := hex.EncodeToString([]byte(txt))
        fmt.Println("Hexadecimal Text:", hexStr)
        return

    case 3:
        fmt.Print("Please input plain text: ")
        fmt.Scan(&txt)
        base64Str := base64.StdEncoding.EncodeToString([]byte(txt))
        fmt.Println("Base64 Text:", base64Str)
        return

    case 4:
        os.Exit(0)
    }
}

func xorEncrypt(mode int) {
	switch mode {
	case 1:
		fmt.Print("PleasenEnter plaintext to encrypt: ")
		fmt.Scan(&text)
		
		fmt.Print("Enter encryption key (single byte as number 0-255): ")
		fmt.Scan(&key)
		
		// Perform XOR encryption
		result := make([]byte, len(text))
		for i := 0; i < len(text); i++ {
			result[i] = text[i] ^ key
		}
		encrypted := string(result)
		
		fmt.Printf("Encrypted (hex): ")
		for _, b := range []byte(encrypted) {
			fmt.Printf("%02x ", b)
		}
		fmt.Printf("\nEncrypted (string): %s\n", encrypted)
		
	case 2:	
		var encryptedText string
		fmt.Print("Please Enter ciphertext to decrypt: ")
		fmt.Scan(&encryptedText)
		
		fmt.Print("Enter decryption key (single byte as number 0-255): ")
		fmt.Scan(&key)
		
		// Perform XOR decryption (same operation as encryption)
		result := make([]byte, len(encryptedText))
		for i := 0; i < len(encryptedText); i++ {
			result[i] = encryptedText[i] ^ key
		}
		decrypted := string(result)
		
		fmt.Printf("Decrypted: %s\n", decrypted)
		
	case 3:
		fmt.Println("Goodbye!")
		os.Exit(0)
		
	default:
		fmt.Println("Invalid option! Please choose 1, 2, or 3.")
	}
}