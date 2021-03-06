package models

type Address struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Enable  bool   `json:"enable"`
	Created int64  `json:"created"`
}

func (address Address) TableName() string {
	return "cili_engine"
}
