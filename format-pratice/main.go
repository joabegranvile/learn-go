package formatpratice

import "fmt"

func main() {
	fname := "Dall"
	lname := "Adal"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometime as hyprocite is nothing more than a main in the process of changing."

	userLog := fmt.Sprintf(
		"Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %t, Message: %s",
		fname,
		lname,
		age,
		messageRate,
		isSubscribed,
		message,
	)
	fmt.Println(userLog)
}
