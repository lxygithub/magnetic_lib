package models

type Manager struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Enable     int    `json:"enable"`
	Delete     int    `json:"delete"`
	Verify     int    `json:"verify"`
	Root       int    `json:"root"`
	Updatetime int64  `json:"updatetime"`
}

func (m Manager) TableName() string {
	return "manager"
}
