package c2s

type LoginData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserData struct {
	Id       int64 `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Img      string `json:"img"`
	Addr     string `json:"addr"`
}

type DeleteData struct {
	Id int64 `json:"id"`
}
