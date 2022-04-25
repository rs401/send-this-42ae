package functions

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

const (
	bucketName     = "send-this-42ae.appspot.com"
	googleAccessID = "send-this-42ae@appspot.gserviceaccount.com"
)

type GeturlsResponse struct {
	UploadUrl    string `json:"upload_url"`
	DownloadName string `json:"download_name"`
}

func (us *UrlService) geturls(w http.ResponseWriter, r *http.Request) {
	us.saKey = os.Getenv("SA_KEY")
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "wrong http method, post only"})
		return
	}
	us.filetype = r.FormValue("filetype")
	uniqueName := uuid.New().String()
	url, err := us.genUrl("temp/" + uniqueName)
	if err != nil {
		log.Printf("error generating secure urls: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error, gen_url"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"upload_url":    url,
		"download_name": uniqueName,
	})

}
