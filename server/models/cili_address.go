package models

import "time"

type Address struct {
	Id      int       `json:"id"`
	Url     string    `json:"url"`
	Name    string    `json:"name"`
	Desc    string    `json:"desc"`
	Enable  bool      `json:"enable"`
	Created time.Time `json:"created"`
}
