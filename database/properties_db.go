package database

import(
	"log"
	"encoding/json"
	"strings"
)


type Properties struct {
	Id              int    `json:"Id"`
	Title           string `json:"Title"`
	Address         string `json:"Address"`
	Room_number     string `json:"Room_number"`
	Centare         int    `json:"Centare"`
	Floor           int    `json:"Floor"`
	Floor_number    int    `json:"Floor_number"`
	Description     string `json:"Description"`
	Page_name       string `json:"Page_name"`
	File_names      string `json:"File_names"`
	Property_type   string `json:"Property_type"`
	Property_status string `json:"Property_status"`
}

type Response struct {
	TotalCount int          `json:"total_count"`
	Properties []Properties `json:"properties"`
}

func GetProperties(pageNumber int, propertyStatus string, propertyType string) Response {
	db, err := ConnectDb()
	if err != nil {
		log.Println("Error:", err)
	}

	if pageNumber != 0 {
		offset := (pageNumber - 1) * 4
		log.Println("of:", offset)

		var whereConditions []string
		var args []interface{}

		if propertyType != "all" {
			whereConditions = append(whereConditions, "property_type = ?")
			args = append(args, propertyType)
		}

		if propertyStatus != "all" {
			whereConditions = append(whereConditions, "property_status = ?")
			args = append(args, propertyStatus)
		}

		query := db.Table("properties")
		if len(whereConditions) > 0 {
			query = query.Where(strings.Join(whereConditions, " AND "), args...)
		}

		var count int64
		var properties []Properties
		if err := query.Count(&count).Offset(offset).Limit(4).Find(&properties).Error; err != nil {
			log.Println("Error: Cannot get table")
		}
		CloseDb(db)

		jsonData := Response{
			TotalCount: int(count),
			Properties: properties,
		}

		return jsonData
	}
	return Response{}
}

func AddProperty(jsonData []byte) error {
	db, err := ConnectDb()
	if err != nil {
		log.Println("Error:", err)

	}
	var property Properties
	if err := json.Unmarshal(jsonData, &property); err != nil {
		log.Println("JSON decode error:", err)
		return err
	}

	log.Println(property)
	if err := db.Table("properties").Create(&property).Error; err != nil {
		log.Println("error:", err)
	}

	CloseDb(db)
	return nil
}

func GetPropertyPage(pageName string) []byte {
	db, err := ConnectDb()
	if err != nil {
		log.Println("Error:", err)
	}

	var properties []Properties
	if err := db.Table("properties").Where("page_name = ?", pageName).First(&properties).Error; err != nil {
		log.Println("Error: Cannot get table")
		CloseDb(db)
	}

	CloseDb(db)
	jsonData, _ := json.Marshal(properties)
	return jsonData
}

func UpdateProperty(jsonData []byte) error {
	db, err := ConnectDb()
	if err != nil {
		log.Println("Error:", err)
		CloseDb(db)
		return err
	}

	var property Properties
	if err := json.Unmarshal(jsonData, &property); err != nil {
		log.Println("JSON decode error:", err)
		CloseDb(db)
		return err
	}

	log.Println(property.Id, property)
	if err := db.Table("properties").Where("id = ?", property.Id).Updates(&property).Error; err != nil {
		log.Println("Error: Cannot get table")
	}
	CloseDb(db)

	return nil
}
