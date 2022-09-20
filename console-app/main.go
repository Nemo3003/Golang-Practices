package main

import(

	"fmt"
	"log"
	"github.com/eiannone/keyboard"
	"strconv"
)

func main(){
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	//anonymous function
	defer func(){
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappuccino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappuccino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	for{
		char, _, err := keyboard.GetSingleKey()
		if err != nil{
			log.Fatal(err)
		}
		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))


		if char == 'q' || char == 'Q'{
			break
		}
	}

	fmt.Println("Program Exiting...")
}