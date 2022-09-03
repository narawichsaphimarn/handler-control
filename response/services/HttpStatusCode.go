package services

var message = map[int]string{
	200: "OK",
	201: "Created",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	413: "Payload Too Large",
	415: "Unsupported Media Type",
	429: "Too Many Requests",
	500: "Internal Server Error",
	502: "Bad Gateway",
	504: "Gateway Timeout",
}

func HttpMessage(_value int) string {
	return message[_value]
}
