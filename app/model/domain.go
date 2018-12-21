package model

import "time"

type User struct {
	ID          int64        `json:"id", db:"user_id"`
	FullName    string       `json:"fullName", db:"full_name"`
	UserName    string       `json:"userName",db:"user_name"`
	Password    string       `json:"password",db:"password"`
	Created     time.Time    `json:"created",db:"created"`
	Updated     time.Time    `json:"updated", db:"updated"`
	Roles       []Role       `json:"roles", db:"-"`
	Departments []Department `json:"departments", db:"-"`
}

type Role struct {
	Code string `json:"code", db:"role_code"`
	Name string `json:"name", db:"role_name"`
}

type Department struct {
	Code string `json:"code", db:"role_code"`
	Name string `json:"name", db:"role_name"`
}
