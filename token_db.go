package main

import (
	"fmt"
)

type TokenList struct {
	Token string `json:"Token"`
}

func checkToken(jwtToken string) (bool, error) {
	db, err := connectDb()
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer closeDb(db)

	newToken := TokenList{Token: jwtToken}

	if err := db.Table("tokenList").Where("token = ?", newToken.Token).First(&newToken).Error; err != nil {
		if err.Error() == "record not found" {
			return false, nil 
		}
		return false, err 
	}

	return true, nil
}

func addToken(jwtToken string) error {
	db, err := connectDb()
	if err != nil {
		fmt.Println("Error:", err)
        return err
	}
	defer closeDb(db)

	newToken := TokenList{Token: jwtToken}

	fmt.Println("token adding:", newToken)
	if err := db.Table("tokenList").Create(&newToken).Error; err != nil {
		fmt.Println("error:", err)
		return err
	}

	return nil
}

func logout(jwtToken string) (bool,error) {
	db, err := connectDb()
	if err != nil {
		fmt.Println("Error:", err)
        return false,err
	}
	defer closeDb(db)

	tokenObject := TokenList{Token: jwtToken}
	if err := db.Table("tokenList").Where("token = ?", jwtToken).Delete(&tokenObject).Error; err != nil {
		fmt.Println("error:", err)
		return false,err
	}

	return true,nil
}