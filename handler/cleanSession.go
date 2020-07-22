package handler

import (
	"fmt"
	"time"
)

func CleanSessions() {
	fmt.Println("BEFORE CLEAN")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.LastActivity) > (time.Minute * 10) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN")
	showSessions()
}
