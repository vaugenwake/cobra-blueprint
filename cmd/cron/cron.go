package cron

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func scheduledTask() {
	fmt.Println("[Job 1]: Running scheduled task")
}

func NewCron() {
	fmt.Println("Starting Cron Jobs")

	s := gocron.NewScheduler()
	if err := s.Every(10).Seconds().Do(scheduledTask); err != nil {
		fmt.Printf("Error scheduling job 1, Err: %v", err)
	}
	<-s.Start()
}
