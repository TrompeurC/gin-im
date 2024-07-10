package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string     `json:"name" `
	Password      string     `json:"password"`
	Phone         string     `json:"phone" binging:"e164"`
	Email         string     `json:"email" binding:"email"`
	Salt          string     `json:"salt"`
	Identity      string     `json:"identity"`
	ClientIp      string     `json:"clientIp"`
	ClientPort    string     `json:"clientPort"`
	LoginTime     *LocalTime `json:"loginTime"`
	HeartbeatTime *LocalTime `json:"heartbeatTime"`
	LoginOutTime  *LocalTime `json:"loginOutTime"`
	IsLogout      bool       `json:"isLogout"`
	DeviceInfo    string     `json:"deviceInfo"`
}
