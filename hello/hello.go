package main

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"time"

	"example.com/greetings"
	"rsc.io/quote"
)

func checkOS() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func whenIsSaturday() {
	fmt.Print("When is Saturday?")
	switch today := time.Now().Weekday(); time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In 2 days ")
	default:
		fmt.Println("Too far away..")
	}
}

func switchWithNoCondition() {
	switch {

	case time.Hour < 12:
		fmt.Println("")
	case time.Hour >= 12 && time.Hour < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")

	}
}
func main() {
	log.SetPrefix("Greetings:")
	log.SetFlags(0)
	fmt.Println("My first go program")
	fmt.Println(quote.Glass())
	message, err := greetings.Hello("Deepakala")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(message)
	}

	/*message, err = greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(message)
	}*/
	names := make([]string, 0)
	names = append(names, "Aditi")
	names = append(names, "Aadeesh")
	names = append(names, "Bharath")
	fmt.Println("Names slice = ", names)
	messages, err1 := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err1)
	} else {
		fmt.Println(messages)
	}

	var i int = 9
	fmt.Println("Sqrt(", i, ") = ", math.Sqrt(float64(i)))
	i = 0
	sum := 0
	for ; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	checkOS()
	whenIsSaturday()
	switchWithNoCondition()

}
