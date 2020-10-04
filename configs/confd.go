package configs

import "time"

const (
	LOG_PATH			= "../logs/queueingAPI_log.txt"
	PORT 				= ":5004"
	Threads				= 5
	FailedMessageQueue 	= "failed-mq"
	SleepTime 			= 1 * time.Millisecond
)

var BrokerUrl = ""