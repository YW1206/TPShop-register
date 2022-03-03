package dao

import (
	"tpshop/model"
	"tpshop/utils"
)

//CheckUserNameAndPassword 根据用户名和密码从数据库中查询一条记录
func CheckUserNameAndPassword(username string, password string) (*model.User_chenjunjie, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User_chenjunjie{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

//CheckUserName 根据用户名和密码从数据库中查询一条记录
func CheckUserName(username string) (*model.User_chenjunjie, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User_chenjunjie{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

//SaveUser 向数据库中插入用户信息
func SaveUser(username string, password string, email string, phone string, rec_phone string) error {
	//写sql语句
	sqlStr := "insert into users(username,password,email,phone,rec_phone) values(?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, email, phone, rec_phone)
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfo() ([]*model.User_chenjunjie, error) {
	//写sql语句
	sql := "select id,username,PASSWORD,email from users"
	//执行
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var users []*model.User_chenjunjie
	for rows.Next() {
		user := &model.User_chenjunjie{}
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users, nil
}