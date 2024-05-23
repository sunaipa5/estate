package database

import (
	"encoding/json"
	"log"
)

type users struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type username struct {
	Username string `json:"Username"`
}

func GetUsers() []username {
	db, err := ConnectDb()
	if err != nil {
		log.Println("Error:", err)
	}

	var getusers []username
	if err := db.Table("users").Find(&getusers).Error; err != nil {
		log.Println("Error: Cannot get table")
	}

	CloseDb(db)

	return getusers
}

func AddUser(jsonData []byte) error {
	db, err := ConnectDb()
	if err != nil {
		log.Println(err)
		return err
	}
	var user users
	if err := json.Unmarshal(jsonData, &user); err != nil {
		log.Println("JSON decode error:", err)
		return err
	}

	if err := db.Table("users").Create(&user).Error; err != nil {
		log.Println("error:", err)
	}

	return nil
}

func CheckUser(userInfo []byte) (int, bool, string) {
	db, err := ConnectDb()
	if err != nil {
		log.Println("Error:", err)
	}

	var user users
	if err := json.Unmarshal(userInfo, &user); err != nil {
		log.Println("JSON decode hatası:", err)
		return 404, false, ""
	}

	result := db.Table("users").Where("username = ? AND password = ?", user.Username, user.Password).First(&user)
	if result.RowsAffected == 0 {
		return 404, false, ""
	} else {
		return 200, true, user.Username
	}

}

func DeleteUser(jsonData []byte) error {
	db, err := ConnectDb()
	if err != nil {
		log.Println("Error:", err)
	}

	var user username
	if err := json.Unmarshal(jsonData, &user); err != nil {
		log.Println("JSON decode hatası:", err)
		return err
	}

	if err := db.Table("users").Where("username = ?", user.Username).Delete(&user).Error; err != nil {
		log.Println("Error: Cannot get table")
	}

	CloseDb(db)

	return nil
}
