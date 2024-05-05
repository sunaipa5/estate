package main

func getFrontendDist() string {
	return "./frontend/dist/"
}

func getDbVariables() (string, string, string, string) {
	dbhost := "127.0.0.1:3306"
	dbname := "estate"
	dbuser := "dbadmin"
	dbpass := "dbadmin"
	return dbhost, dbname, dbuser, dbpass
}

func getListenAdress() string {
	//ip:port
	address := ":3000"
	return address
}
