package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var dolarBlue float64
var dolarOfficial float64
var salary float64
var splitPart float64
var splitDolar float64
var receivedDollar float64
var receivedDollarBlue float64
var whSalary float64
var finalSalary float64

//Creating select dollar or pesos function
func userElection() int {
	reader := bufio.NewReader(os.Stdin)            //Create reader
	option, _ := reader.ReadString('\n')           //Read input
	option = strings.Replace(option, "\n", "", -1) //Delete all hop lines
	option = strings.TrimSuffix(option, "\r")      //Remove \r character for the string

	var option_final int
	var err error

	option_final, err = strconv.Atoi(option) //Conver string to int
	if err != nil {
		fmt.Println(err)
	}

	return option_final
}

//Creating function to enter a float value
func intrValue() float64 {
	reader := bufio.NewReader(os.Stdin)            //Create reader
	option, _ := reader.ReadString('\n')           //Read input
	option = strings.Replace(option, "\n", "", -1) //Delete all hop lines
	option = strings.TrimSuffix(option, "\r")      //Remove \r character for the string

	var option_final float64
	var err error

	option_final, err = strconv.ParseFloat(option, 64) //Conver string to int
	if err != nil {
		fmt.Println(err)
	}

	return option_final
}

//Creating a function to check if we select yes or no
func checkOut() bool {
	reader := bufio.NewReader(os.Stdin)            //Create reader
	option, _ := reader.ReadString('\n')           //Read input
	option = strings.Replace(option, "\n", "", -1) //Delete all hop lines
	option = strings.TrimSuffix(option, "\r")      //Remove \r character for the string

	var check bool

	if option == "Y" || option == "y" {
		check = true
	} else {
		check = false
	}
	return check
}

//Creating a function to print the results based on dollars
func giveResultsDollar() {
	fmt.Print("Enter your salary in dollars (Ex: 1500 or 1500.48): ")
	salary = intrValue()
	fmt.Print("Enter the value of the official dollar (Ex: 87 or 87.5): ")
	dolarOfficial = intrValue()
	fmt.Print("Enter the value of the blue dollar (Ex: 87 or 87.5): ")
	dolarBlue = intrValue()
	fmt.Print("Enter the percentage you will recevie in dollars (Ex: 30): ")
	splitPart = intrValue() / 100

	fmt.Print("Are you sure the values are ok? Y/y = yes. Any other key = no : ")
	checkout := checkOut()

	//if checkout == false
	if !checkout {
		fmt.Printf("\x1bc")
		fmt.Println("---------- Welcome to the salary calculator ----------")
		fmt.Println("GREAT! You are earning in dollars")
		giveResultsDollar()
	}

	grossSalary := dolarOfficial * salary
	netSalary := grossSalary * 0.83
	splitDolar = netSalary * splitPart
	receivedDollar = splitDolar / dolarOfficial
	receivedDollarBlue = receivedDollar * dolarBlue
	whSalary = netSalary - splitDolar
	finalSalary = whSalary + receivedDollarBlue

	fmt.Printf("\x1bc")
	fmt.Println("---------- Welcome to the salary calculator ----------")
	fmt.Println("")
	fmt.Println("Dollar official value:\t\t\t", dolarOfficial, "AR$")
	fmt.Println("Dollar Blue value:\t\t\t", dolarBlue, "AR$")
	fmt.Println("Your gross salary in pesos is:\t\t", math.Round(grossSalary*100)/100, "AR$")
	fmt.Println("Your net salary in pesos is:\t\t", math.Round(netSalary*100)/100, "AR$")
	fmt.Println("Your dollarized pesos are:\t\t", math.Round(splitDolar*100)/100, "AR$")
	fmt.Println("Your received dollars are:\t\t", math.Round(receivedDollar*100)/100, "US$")
	fmt.Println("Your received dollars in blue are:\t", math.Round(receivedDollarBlue*100)/100, "AR$")
	fmt.Println("Your final received pesos are:\t\t", math.Round(whSalary*100)/100, "AR$")
	fmt.Println("Your final received salary is:\t\t", math.Round(finalSalary*100)/100, "AR$")
}

//Creatinga function to print the results based on pesos
func giveResultPesos() {
	fmt.Print("Enter your salary in pesos (Ex: 95000 or 95000.451): ")
	salary = intrValue()
	fmt.Print("Enter the value of the official dollar (Ex: 87 or 87.5): ")
	dolarOfficial = intrValue()
	fmt.Print("Enter the value of the blue dollar (Ex: 87 or 87.5): ")
	dolarBlue = intrValue()
	fmt.Print("Enter the percentage you will recevie in dollars (Ex: 30): ")
	splitPart = intrValue() / 100

	fmt.Print("Are you sure the values are ok? Y/y = yes. Any other key = no : ")
	checkout := checkOut()

	//if checkout == false
	if !checkout {
		fmt.Printf("\x1bc")
		fmt.Println("---------- Welcome to the salary calculator ----------")
		fmt.Println("GREAT! You are earning in Pesos")
		giveResultPesos()
	}

	netSalary := salary * 0.83
	splitDolar = netSalary * splitPart
	receivedDollar = splitDolar / dolarOfficial
	receivedDollarBlue = receivedDollar * dolarBlue
	whSalary = netSalary - splitDolar
	finalSalary = whSalary + receivedDollarBlue

	fmt.Printf("\x1bc")
	fmt.Println("---------- Welcome to the salary calculator ----------")
	fmt.Println("")
	fmt.Println("Dollar official value:\t\t\t", dolarOfficial, "AR$")
	fmt.Println("Dollar Blue value:\t\t\t", dolarBlue, "AR$")
	fmt.Println("Your gross salary in pesos is:\t\t", salary, "AR$")
	fmt.Println("Your net salary in pesos is:\t\t", math.Round(netSalary*100)/100, "AR$")
	fmt.Println("Your dollarized pesos are:\t\t", math.Round(splitDolar*100)/100, "AR$")
	fmt.Println("Your received dollars are:\t\t", math.Round(receivedDollar*100)/100, "US$")
	fmt.Println("Your received dollars in blue are:\t", math.Round(receivedDollarBlue*100)/100, "AR$")
	fmt.Println("Your final received pesos are:\t\t", math.Round(whSalary*100)/100, "AR$")
	fmt.Println("Your final received salary is:\t\t", math.Round(finalSalary*100)/100, "AR$")
}

func main() {

	fmt.Printf("\x1bc")
	fmt.Println("---------- Welcome to the salary calculator ----------")
	fmt.Print("Enter 1 if you earn in dollars or 2 in you earn in pesos: ")
	option := userElection()

	var salaryType string

	for {
		if option == 1 {
			salaryType = "dollar"
			break
		} else if option == 2 {
			salaryType = "pesos"
			break
		} else {
			fmt.Println("Wrong election, please enter a number between 1 or 2!")
			fmt.Print("Enter a number: ")
			option = userElection()
		}
	}

	switch salaryType {
	case "dollar":
		fmt.Printf("\x1bc")
		fmt.Println("---------- Welcome to the salary calculator ----------")
		fmt.Println("GREAT! You are earning in dollars")
		giveResultsDollar()
		fmt.Println("\nDo you want to modify some value?\n- Enter 'Y/y'to modify the values\n- Press any other key to exit")
		fmt.Print("Enter decision: ")
		checkout := checkOut()
		//checkout == true
		for checkout {
			if checkout {
				fmt.Printf("\x1bc")
				fmt.Println("---------- Welcome to the salary calculator ----------")
				fmt.Println("GREAT! You are earning in dollars")
				giveResultsDollar()
				fmt.Println("\nDo you want to modify some value?\n- Enter 'Y/y'to modify the values\n- Press any other key to exit")
				fmt.Print("Enter decision: ")
				checkout = checkOut()
			}
		}
		fmt.Println("")
		fmt.Println("See you in another opportunity. :)")
		fmt.Println("")

	case "pesos":
		fmt.Printf("\x1bc")
		fmt.Println("---------- Welcome to the salary calculator ----------")
		fmt.Println("GREAT! You are earningn in pesos")
		giveResultPesos()
		fmt.Println("\nDo you want to modify some value?\n- Enter 'Y/y'to modify the values\n- Press any other key to exit")
		fmt.Print("Enter decision: ")
		checkout := checkOut()
		//checkout == true
		for checkout {
			if checkout {
				fmt.Printf("\x1bc")
				fmt.Println("---------- Welcome to the salary calculator ----------")
				fmt.Println("GREAT! You are earning in pesos")
				giveResultPesos()
				fmt.Println("\nDo you want to modify some value?\n- Enter 'Y/y'to modify the values\n- Press any other key to exit")
				fmt.Print("Enter decision: ")
				checkout = checkOut()
			}
		}
		fmt.Println("See you in another opportunity. :)")
		fmt.Println("")
	}
}
