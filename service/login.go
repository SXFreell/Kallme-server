package service

import (
	"context"
	"kallme/config"
	"kallme/dao"
	"kallme/model"
	"kallme/utils"
	"regexp"
	"unicode/utf8"
)

func CheckUsernameExist(username string) bool {
	res, _ := dao.CheckDataExist("user", "username", username)
	return res
}

func CheckPasswordMatch(username string, password string) bool {
	var users []model.User
	cursor, _ := dao.GetDataByKey("user", "username", username)
	cursor.All(context.TODO(), &users)
	config.Log.Info("CheckPasswordMatch: ", users)
	for _, user := range users {
		if user.Password == password {
			return true
		}
	}
	return false
}

func InsertUser(user model.User) bool {
	_, err := dao.InsertData("user", user)
	config.Log.Error("InsertUser: ", err)
	return err == nil
}

func GetUserId(username string) string {
	var users []model.User
	cursor, _ := dao.GetDataByKey("user", "username", username)
	cursor.All(context.Background(), &users)
	config.Log.Info("GetUserId: ", users)
	for _, user := range users {
		return user.UUID
	}
	return ""
}

func CheckUsernameValid(username string) bool {
	// 检查长度是否超过16个字符, 小于4个字符
	if utf8.RuneCountInString(username) > 16 || utf8.RuneCountInString(username) < 4 {
		return false
	}

	// 创建一个正则表达式对象，用于匹配输入的字符是否符合要求
	// 仅支持英文、数字, 只能以英文开头
	reg, err := regexp.Compile("^[a-zA-Z][a-zA-Z0-9]*$")
	if err != nil {
		config.Log.Error(err)
	}

	// 使用正则表达式检查输入
	return reg.MatchString(username)
}

func GenerateToken(username string) (string, error) {
	token := utils.GenerateUUID()
	uuid := GetUserId(username)
	tokenData := model.Token{
		UUID:  uuid,
		Token: token,
	}
	res, _ := dao.CheckDataExist("token", "uuid", uuid)
	if !res {
		_, err := dao.InsertData("token", tokenData)
		config.Log.Error("InsertToken: ", err)
		if err != nil {
			return "", err
		}
	} else {
		_, err := dao.UpdateData("token", "uuid", uuid, tokenData)
		config.Log.Error("UpdateToken: ", err)
		if err != nil {
			return "", err
		}
	}
	return token, nil
}
