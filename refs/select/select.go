package selecttd

import (
	"net/http"
	"time"
)


func Racer(urls ...string) string {
	shortestDuration := time.Duration(50*float64(time.Second))
	var fastestUrl string
	for  _, url := range urls {
		current_duration := measureResponseTime(url)
		if  current_duration < shortestDuration {
			shortestDuration = current_duration
			fastestUrl = url
		}
	}
	return fastestUrl
}


func SelectRacer(url1 string, url2 string) string {
	select {
	case <- ping(url1):
		return url1
	case <- ping(url2):
		return url2
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)
	return duration
}

func ping(url string ) chan struct{} {
	chn := make(chan struct{})
	go func() {
		http.Get(url)
		close(chn)
	}()
	return chn
}