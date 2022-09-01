package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))                 // 01-02-2006 is used for formatting
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // 01-02-2006 15:04:05 Monday is also standard

	createdDate := time.Date(2022, time.July, 02, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
