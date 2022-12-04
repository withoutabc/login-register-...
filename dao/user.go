package dao

import "goproject1/model"

func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("insert into user(name,password) values (?,?)", u.Username, u.Password)
	return
}
func SearchUserByUsername(name string) (u model.User, err error) {
	row := DB.QueryRow("select * from user where name=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.Username, &u.Password)
	return
}
func ChangePassword(NewPassword string, name string) (err error) {
	_, err = DB.Exec("update user set password=? where name=?", NewPassword, name)
	return
}
func DeleteUser(name string) (err error) {
	_, err = DB.Exec("delete from user where name=?", name)
	return
}
func InsertMessage(m model.Message) (err error) {
	_, err = DB.Exec("insert into message(send_uid, rec_uid, detail) values(?,?,?)", m.SendUID, m.RecUID, m.Detail)
	return
}
