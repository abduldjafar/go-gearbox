package entity

type Body struct {
	Responses struct {
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Datas   interface{} `json:"datas"`
	} `json:"responses"`
}
