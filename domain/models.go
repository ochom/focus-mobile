package domain

//HttpRequestParam is the payload and attributes required to
//make http request
type HttpRequestParam struct {
	Method  string
	Headers map[string]string
	URL     string
	Payload interface{}
}
