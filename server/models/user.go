package models

type User struct {
	Uid        int    `json:"uid"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Enable     int    `json:"enable"`
	Remove     int    `json:"remove"`
	Verify     int    `json:"verify"`
	Root       int    `json:"root"`
	Updatetime int64  `json:"updatetime"`
	Createtime int64  `json:"createtime"`
}

func (m User) TableName() string {
	return "manager"
}
