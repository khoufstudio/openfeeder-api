package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	res "api-openfeeder/utils"
	"net/http"
	config "api-openfeeder/config"
	models "api-openfeeder/models"
)

var dbConnection *sql.DB

// Router is a function that extract routers
func Router() *mux.Router {

	// eb := dbConnection.Ping()

	// if eb != nil {
	// 	panic(eb.Error())
	// }

	router := mux.NewRouter()
	router.HandleFunc("/", handleHome)
	router.HandleFunc("/helpdesk", getAllHelpdesk).Methods("GET")

	return router
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var jsonExample = map[string]interface{} {
		"status": 200,
		"message": "Selamat datang di API OpenFeeder",
	}
	res.ResponseJSON(w, jsonExample, http.StatusOK)

}


func getAllHelpdesk(w http.ResponseWriter, r *http.Request) {	
	dbConnection := config.ConnectMysql()
	defer dbConnection.Close()

	rows, err := dbConnection.Query("SELECT id, email, nama_lengkap, asal_pt, subjek_keluhan, isi_keluhan FROM helpdesk")

	if err != nil {
		res.ResponseJSON(w, map[string]interface{} {
			"message": "data helpdesk tidak ada",
		}, http.StatusOK)
	}

	var helpdesks []models.Helpdesk

	for rows.Next() {
		var helpdesk models.Helpdesk

		if err := rows.Scan(&helpdesk.ID, &helpdesk.Email, &helpdesk.NamaLengkap, &helpdesk.AsalPT, &helpdesk.SubjekKeluhan, &helpdesk.IsiKeluhan); err != nil {
			res.ResponseJSON(w, map[string]interface{} {
				"message": "error",
			}, http.StatusOK)			
		} else {
			helpdesks = append(helpdesks, helpdesk)
		}		
	}

	res.ResponseJSON(w, helpdesks, http.StatusOK)

}