package database

import (
	"fmt"
)

type TokenList struct {
	Token string `json:"Token"`
}

func CheckToken(jwtToken string) (bool, error) {
	db, err := ConnectDb()
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer CloseDb(db)

	newToken := TokenList{Token: jwtToken}

	if err := db.Table("tokenList").Where("token = ?", newToken.Token).First(&newToken).Error; err != nil {
		if err.Error() == "record not found" {
			return false, nil 
		}
		return false, err 
	}

	return true, nil
}

func AddToken(jwtToken string) error {
	db, err := ConnectDb()
	if err != nil {
		fmt.Println("Error:", err)
        return err
	}
	defer CloseDb(db)

	newToken := TokenList{Token: jwtToken}

	fmt.Println("token adding:", newToken)
	if err := db.Table("tokenList").Create(&newToken).Error; err != nil {
		fmt.Println("error:", err)
		return err
	}

	return nil
}

func DeleteToken(jwtToken string) (bool,error) {
	db, err := ConnectDb()
	if err != nil {
		fmt.Println("Error:", err)
        return false,err
	}
	defer CloseDb(db)

	tokenObject := TokenList{Token: jwtToken}
	if err := db.Table("tokenList").Where("token = ?", jwtToken).Delete(&tokenObject).Error; err != nil {
		fmt.Println("error:", err)
		return false,err
	}

	return true,nil
}