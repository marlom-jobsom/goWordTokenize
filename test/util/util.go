package util

import (
	"fmt"
	"time"
	"word-tokenize-in1118/core/communication"
	"word-tokenize-in1118/util"
)

// SumResponsesDuration sums the duration of all responses given
func SumResponsesDuration(responses []communication.Response) time.Duration {
	var total time.Duration
	for _, element := range responses {
		total += element.Duration
	}
	return total
}

// WriteResponsesDuration writes the responses duration in a file
func WriteResponsesDuration(filePath string, responses []communication.Response) {
	for _, response := range responses {
		util.AppendContentFile(filePath, fmt.Sprint(response.Duration.Nanoseconds()))
	}
}
