package handler

import "fmt"

func showSessions() {
	fmt.Println("******")
	for k, v := range dbSessions {
		fmt.Println(k, v.Id)
	}
	fmt.Println("")
}
