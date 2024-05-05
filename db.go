package main

import (
	"fmt"

	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDb() (*gorm.DB, error) {
	dbhost, dbname, dbuser, dbpass := getDbVariables()
	dsn := dbuser + ":" + dbpass + "@tcp(" + dbhost + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func closeDb(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Veritabanı bağlantısı kapatılırken hata oluştu!")
	}
	defer sqlDB.Close()
}

type users struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func getUsers() []byte {
	db, err := connectDb()
	if err != nil {
		fmt.Println("Error:", err)
	}

	var getusers []users
	if err := db.Table("users").Find(&getusers).Error; err != nil {
		fmt.Println("Error: Cannot get table")
	}

	jsonData, err := json.Marshal(getusers)
	if err != nil {
		fmt.Println("Error: Cannot create JSON")
	}
	closeDb(db)

	return jsonData
}

func checkUser(userInfo []byte) (int, bool, string) {
	db, err := connectDb()
	if err != nil {
		fmt.Println("Error:", err)
	}

	var user users
	if err := json.Unmarshal(userInfo, &user); err != nil {
		fmt.Println("JSON decode hatası:", err)
		return 404, false, ""
	}

	result := db.Table("users").Where("username = ? AND password = ?", user.Username, user.Password).First(&user)
	if result.RowsAffected == 0 {
		return 404, false, ""
	} else {
		return 200, true, user.Username
	}

}

type Notices struct {
	Title        string `json:"Title"`
	Address      string `json:"Address"`
	Room_number  string `json:"Room_number"`
	Centare      int    `json:"Centare"`
	Floor        int    `json:"Floor"`
	Floor_number int    `json:"Floor_number"`
	Description  string `json:"Description"`
	Page_name    string `json:"Page_name"`
	File_names    string `json:"File_names"`
}

type Response struct {
	TotalCount int        `json:"total_count"`
	Notices []Notices `json:"notices"`
}

func getNotices(pageNumber int) Response {
	db, err := connectDb()
	if err != nil {
		fmt.Println("Error:", err)
	}

	if pageNumber != 0 {
		offset := (pageNumber - 1)*4	
		fmt.Println("of:", offset)

		var count int64
		var notices []Notices
		if err := db.Table("notices").Count(&count).Offset(offset).Limit(4).Find(&notices).Error; err != nil {
			fmt.Println("Error: Cannot get table")
			closeDb(db)
		}
		closeDb(db)

		jsonData := Response{
			TotalCount: int(count),
			Notices: notices,
		}


		return jsonData
	}
   return Response{}
}

func addnotice(jsonData []byte) error {
	db, err := connectDb()
	if err != nil {
		fmt.Println("Error:", err)

	}
	var notice Notices
	if err := json.Unmarshal(jsonData, &notice); err != nil {
		fmt.Println("JSON decode hatası:", err)
		return err
	}

	fmt.Println(notice)
	if err := db.Table("notices").Create(&notice).Error; err != nil {
		fmt.Println("error:", err)
	}

	closeDb(db)
	return nil
}

func getNoticePage(pageName string) []byte {
	db, err := connectDb()
	if err != nil {
		fmt.Println("Error:", err)
	}

	var notices []Notices
	if err := db.Table("notices").Where("page_name = ?", pageName).First(&notices).Error; err != nil {
		fmt.Println("Error: Cannot get table")
		closeDb(db)
	}

	closeDb(db)
	jsonData, _ := json.Marshal(notices)
	return jsonData
}
