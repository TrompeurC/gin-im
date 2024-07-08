package models

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Identity      string    `json:"identity"`
	ClientIp      string    `json:"client_ip"`
	ClientPort    string    `json:"client_port"`
	LoginTime     time.Time `json:"login_time"`
	HeartbeatTime time.Time `json:"heartbeat_time"`
	LoginOutTime  time.Time `json:"login_out_time"`
	IsLogout      bool      `json:"is_logout"`
	DeviceInfo    string    `json:"device_info"`
}
