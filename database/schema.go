package database

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v3"
)

type User struct {
	ID            uuid.UUID    `json:"id"`
	Name          null.String  `json:"name"`
	User_Name     null.String  `json:"user_name"`
	Email         null.String  `json:"email"`
	Phone         null.String  `json:"phone"`
	Sex           null.Int     `json:"sex"`
	Age           null.Int     `json:"age"`
	Info          null.String  `json:"info"`
	Password_Hash []byte       `-`
	User_Type     int          `json:"user_type"`
	Create_Date   jsonTime     `json:"create_date"`
	Operate_Date  jsonTime     `json:"operate_time"`
	Operator      uuid.UUID    `json:"operator"`
	Banned_Date   jsonNullTime `json:"banned_date"`
	Banned        int          `json:"banned"`
}

type UserOption func(*User)
