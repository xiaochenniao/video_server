package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}


//数据模型

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}