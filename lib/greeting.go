package lib

import (
	"fmt"
	"time"
)

func Greeting(username string) {
	if currentTime := time.Now(); currentTime.Hour() < 12 {
		fmt.Println("Good Morning,", username)
	} else if currentTime.Hour() >= 12 && currentTime.Hour() < 17 {
		fmt.Println("Good Afternoon,", username)
	} else if currentTime.Hour() >= 17 && currentTime.Hour() < 20 {
		fmt.Println("Good Evening,", username)
	} else {
		fmt.Println("Good Night,", username)
	}
}
