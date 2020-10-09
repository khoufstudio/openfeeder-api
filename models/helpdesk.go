package models

// model untuk helpdesk
type Helpdesk struct {
	ID            int    `json:"id"`
	Email         string `json:"email"`
	NamaLengkap   string `json: "namalengkap"`
	AsalPT        string `json:"asal_pt"`
	SubjekKeluhan string `json:"subjek_keluhan"`
	IsiKeluhan    string `json:"isi_keluhan"`
}
