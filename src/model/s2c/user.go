package s2c

type Login struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
}