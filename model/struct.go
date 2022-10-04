package model

type ResponseBody struct {
	// XMLName  xml.Name `xml:"data"`
	UserID   uint   `xml:"user_id"`
	Fullname string `xml:"fullname"`
	Username string `xml:"username"`
	Password string `xml:"password"`
}

type User struct {
	UserID   uint   `json:"user_id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}
