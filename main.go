package main

import (
	_ "api-openfeeder/config"
	res "api-openfeeder/utils"
	routes "api-openfeeder/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API OpenFeeder")

	// http.HandleFunc("/helpdesk", getHelpdesk)

	// db := config.ConnectMysql()
	// defer db.Close()

	// eb := db.Ping()

	// if eb != nil {
	// 	panic(eb.Error())
	// }

	// fmt.Println("success")
	
	router := routes.Router()

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}
}

func getHelpdesk(w http.ResponseWriter, r *http.Request) {
	var jsonExample = map[string]interface{} {
		"status": 200,
		"message": "Hello world coy",
	}
	res.ResponseJSON(w, jsonExample, http.StatusOK)
}
