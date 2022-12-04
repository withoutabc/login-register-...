package service

import (
	"goproject1/dao"
	"goproject1/model"
)

func SearchUserByUsername(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUsername(name)
	return
}
func CreateUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}
func ChangePassword(NewPassword string, name string) error {
	err := dao.ChangePassword(NewPassword, name)
	return err
}
func DeleteUser(name string) error {
	err := dao.DeleteUser(name)
	return err
}
func SendMessage(m model.Message) error {
	err := dao.InsertMessage(m)
	return err
}
