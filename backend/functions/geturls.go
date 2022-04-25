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

// type UrlService struct {
// 	storageClient  *storage.Client
// 	saKey          string
// 	bucketName     string
// 	googleAccessID string
// 	filetype       string
// }

// func NewUrlService() *UrlService {
// 	ctx := context.Background()
// 	tempSC, err := storage.NewClient(ctx)
// 	if err != nil {
// 		log.Fatal("Couldn't create Storage client:", err)
// 	}

// 	urlSvc := &UrlService{
// 		storageClient:  tempSC,
// 		saKey:          os.Getenv("SA_KEY"),
// 		bucketName:     "send-this-42ae.appspot.com",
// 		googleAccessID: "send-this-42ae@appspot.gserviceaccount.com",
// 		filetype:       "",
// 	}
// 	return urlSvc
// }

// var (
// 	storageClient *storage.Client
// 	saKey         string
// 	filetype      string
// )

type GeturlsResponse struct {
	UploadUrl    string `json:"upload_url"`
	DownloadName string `json:"download_name"`
}

// func init() {
// 	ctx := context.Background()
// 	var err error
// 	storageClient, err = storage.NewClient(ctx)
// 	if err != nil {
// 		log.Fatal("Couldn't create Storage client:", err)
// 	}
// 	saKey = os.Getenv("SA_KEY")
// }

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

// func genUrl(name string) (string, error) {
// 	url, err := storageClient.Bucket(bucketName).SignedURL(name, &storage.SignedURLOptions{
// 		GoogleAccessID: googleAccessID,
// 		PrivateKey:     []byte(saKey),
// 		Method:         "PUT",
// 		Expires:        time.Now().Add(10 * time.Minute),
// 		ContentType:    filetype,
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	return url, nil
// }

// func DownloadUrl(w http.ResponseWriter, r *http.Request) {
// 	saKey = os.Getenv("SA_KEY")
// 	parts := strings.Split(r.URL.Path, "/")
// 	fileUUID := parts[len(parts)-1]
// 	name := "temp/" + fileUUID
// 	log.Printf("Redirecting to a new download signed URL for ephemeral resource %q\n", name)

// 	downloadURL, err := storageClient.Bucket(bucketName).SignedURL(name, &storage.SignedURLOptions{
// 		GoogleAccessID: googleAccessID,
// 		PrivateKey:     []byte(saKey),
// 		Method:         "GET",
// 		Expires:        time.Now().Add(10 * time.Minute),
// 	})
// 	if err != nil {
// 		log.Printf("error generating secure download url: %v\n", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error, down_short_url"})
// 		return
// 	}
// 	http.Redirect(w, r, downloadURL, http.StatusFound)
// }
