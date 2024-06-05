package newsapi

import "time"

func CheckTime(config *Config) bool {
	last_ran := config.Scheduler.LastRan
	// last_ran is in RFC3339 format
	t, err := time.Parse(time.RFC3339, last_ran)
	if err != nil {
		return true
	}
	// check if the interval has passed
	if time.Since(t).Hours() < float64(config.Scheduler.Interval) {
		return false
	}

	return true
}

func UpdateTime(config *Config) {
	config.Scheduler.LastRan = time.Now().Format(time.RFC3339)
}
