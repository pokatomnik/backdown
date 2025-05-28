package main

import "time"

func generateArchiveName() string {
	now := time.Now()
	return now.Format("02.01.2006_15.04.05") + ".zip"
}
