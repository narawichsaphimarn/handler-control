package models

import "time"

type HttpRequestForm struct {
	TimeOut       time.Duration
	ContentType   string
	Authorization struct {
		Token string
		Type  string
	}
	ContentEncoding string
	ContentLength   string
	From            string
	Host            string
	Date            string
	Cookie          string
	HTTP2Settings   string
	Origin          string
	Warning         string
	UserAgent       string
	XRequestID      string
}
