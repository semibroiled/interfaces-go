package main

import (
	"fmt"
	"math"
	"time"
)


func main(){
	//Code here

	//shapes

	listArea(rectangle{width: 5, height:5})

	//exercise1
	test1(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test1(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test1(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test1(birthdayMessage{
		recipientName: "Bill Deer",
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})

	//exercise 2
	test2(fullTime{
		name:   "Jack",
		salary: 50000,
	})
	test2(contractor{
		name:         "Bob",
		hourlyPay:    100,
		hoursPerYear: 73,
	})
	test2(contractor{
		name:         "Jill",
		hourlyPay:    872,
		hoursPerYear: 982,
	})

	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test3(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test3(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test3(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test3(e, e)

	//exercise4

	test4(email{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	})
	test4(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "jane@doe.com",
	})
	test4(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "elon@doe.com",
	})
	test4(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555509832",
	})
	test4(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555504444",
	})
	test4(invalid{})
}

//Types and Interfaces
//shapes
type shape interface {
	// interface shape must match with methods area and perimeter
	area() float64
	perimeter() float64
}

//Type Structs Rectangle and Circle
type rectangle struct{
	width, height float64
}

type circle struct{
	radius float64
}

//their methods
func (c circle) perimeter() float64 {
	return 2*math.Pi*c.radius
}

func (c circle) area() float64 {
	return 2*math.Pi*math.Pow(c.radius,2) //import math, call Pi
}

func (r rectangle) area() float64 {
	return r.width*r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width+2*r.height
}

//Call func with interface
func listArea(s shape) {
	fmt.Println(s.area())
}

//Exercise one
/*The birthdayMessage and sendingReport structs have already implemented the getMessage methods.
The getMessage method simply returns a string, and any type that implements the method can be considered a message.

First, add the getMessage() method as a requirement on the method interface.

Second, complete the sendMessage function. It should print a message's message, which it obtains through the interface method. 
Notice that your code doesn't need to worry at all about whether a specific message is a birthdayMessage or a sendingRepo*/
func sendMessage(msg message) {
	fmt.Println(msg.getMessage()) //ah you call the interface method in the newer func
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func test1(m message) {
	sendMessage(m)
	fmt.Println("====================================")
}

//exercise 2

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

// ?
func (c contractor) getSalary() int {
	return c.hourlyPay*c.hoursPerYear
}


// don't touch below this line

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func test2(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("====================================")
}
// exercise 3

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05
	}

	return float64(len(e.body))*0.01 //because len got computed to int, it truncateed 0.01 as int to 0
	//watch out
}

func (e email) print() {
	fmt.Println(e.body)
}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
	toAddress	string
}

func print(p printer) {
	p.print()
}

func test3(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

//exercise 4

func getExpenseReport(e expense) (string, float64) {
	//em, ok := e.(email)
	//if ok {
		//my ansatz is the one in muster
		// ok by itself is typed bool
	//	return em.toAddress, em.cost() //print as em's method
	//}

	//sm, ok := e.(sms)
	//if ok {
	//		return sm.toPhoneNumber, sm.cost()
	//}

	//return "",0.0

	//Type switch approach

	switch v:=e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0

	}
}

// don't touch below this line

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}


type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func test4(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

