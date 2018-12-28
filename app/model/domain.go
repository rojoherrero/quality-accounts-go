package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type User struct {
	ID          int64        `json:"id,omitempty" db:"user_id"`
	FullName    string       `json:"fullName,omitempty" db:"full_name"`
	UserName    string       `json:"userName,omitempty" db:"user_name"`
	Password    string       `json:"password,omitempty" db:"password"`
	Created     time.Time    `json:"created,omitempty" db:"created"`
	Updated     time.Time    `json:"updated,omitempty" db:"updated"`
	Roles       []Role       `json:"roles,omitempty" db:""`
	Departments []Department `json:"departments,omitempty" db:""`
}

type Role struct {
	Code string `json:"code" db:"role_code"`
	Name string `json:"name" db:"role_name"`
}

type Department struct {
	Code string `json:"code" db:"department_code"`
	Name string `json:"name" db:"department_name"`
}

type PropertyMap map[string]interface{}
type PropertyMapSlice []PropertyMap

func (p PropertyMap) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *PropertyMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}
	var i interface{}
	if e := json.Unmarshal(source, &i); e != nil {
		return e
	}
	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}
