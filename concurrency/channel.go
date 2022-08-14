package concurrency

import (
	"fmt"
)

func ChannelExplantion() {

	ninza1, ninza2 := make(chan string), make(chan string)
	go CaptainElect(ninza1, "Ninza1 select")
	go CaptainElect(ninza2, "NinZa2 select")

	select {
	case message := <-ninza1: //Getting message through channel
		fmt.Println(message)
	case message := <-ninza2: // Getting message through channel
		fmt.Println(message)
	}

	// fmt.Println(<-ninza1)
	// fmt.Println(<-ninza2)
}

func CaptainElect(ninza chan string, message string) {
	// Sending the message through channel
	ninza <- message
}

func ChannelExplantion2() {

	ninza1 := make(chan interface{})
	close(ninza1)
	ninza2 := make(chan interface{})
	close(ninza2)

	var ninza1Count, ninza2Count int

	for i := 0; i < 1000; i++ {
		select {
		case <-ninza1:
			ninza1Count++
		case <-ninza2:
			ninza2Count++
		}
	}

	fmt.Printf("NinzaCount1 %d , NinzaCount2 %d\n", ninza1Count, ninza2Count)
}
