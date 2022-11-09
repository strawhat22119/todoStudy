package serializer

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
	Error  error  `json:"err"`
	Token  string `json:"token"`
}
