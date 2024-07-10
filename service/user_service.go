package service

import (
	"errors"
	"gin-im/db"
	"gin-im/models"
	"gin-im/utils"
	"gorm.io/gorm"
	"math/rand"
)

var UserService = new(userService)

type userService struct {
}

func findUserByPhone(phone string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	tx := db.MysqlDB.Where("phone = ?", phone).First(&user)
	if tx.Error != nil && tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, nil
}
func findUserByName(name string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	tx := db.MysqlDB.Where("name = ?", name).First(&user)
	if tx.Error != nil && tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, nil
}

func (u userService) CreateUser(user models.UserBasic) error {
	if u, _ := findUserByPhone(user.Phone); u != nil {
		return errors.New("手机号已被注册！")
	}
	if u, _ := findUserByName(user.Name); u != nil {
		return errors.New("用户名已被注册！")
	}
	user.Password = utils.MakePassword(user.Password, string(rand.Int31()))
	if tx := db.MysqlDB.Create(&user); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u userService) GetUserList() ([]*models.UserBasic, error) {
	var users []*models.UserBasic
	if err := db.MysqlDB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u userService) DeleteUser(id int) error {
	if err := db.MysqlDB.Delete(&models.UserBasic{}, id).Error; err != nil {
		return err
	}
	return nil
}
