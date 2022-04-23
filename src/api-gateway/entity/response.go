package entity

// ResponseHTTP is response format ror readable reason
type ResponseHTTP struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message []string    `json:"msg"`
}
