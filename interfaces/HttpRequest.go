package interfaces

type HttpRequest interface {
	ClientRequest(url string, method string, payload []byte) ([]byte, error)
}
