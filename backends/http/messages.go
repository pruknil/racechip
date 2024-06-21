package http

type RequestHeader struct {
	FuncNm  string `json:"funcNm"`
	RqUID   string `json:"rqUID"`
	RqDt    string `json:"rqDt"`
	RqAppId string `json:"rqAppId"`
	UserId  string `json:"userId"`
}

type ResponseHeader struct {
	FuncNm     string `json:"funcNm"`
	RqUID      string `json:"rqUID"`
	RsDt       string `json:"rqDt"`
	RsAppId    string `json:"rqAppId"`
	StatusCode string `json:"statusCode"`
}

type ExampleBackendRequest struct {
	RequestHeader             `json:"Header"`
	ExampleBackendBodyRequest `json:"Body"`
}

type ExampleBackendResponse []struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ExampleBackendBodyRequest struct {
	DPKName string `json:"dPKName"`
	Data    string `json:"data"`
}

type ExampleBackendBodyResponse struct {
	DPKName string `json:"dPKName"`
	EData   string `json:"eData"`
}
