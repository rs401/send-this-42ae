package functions

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
)

type UrlService struct {
	storageClient  *storage.Client
	saKey          string
	bucketName     string
	googleAccessID string
	filetype       string
}

func NewUrlService() *UrlService {
	ctx := context.Background()
	tempSC, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal("Couldn't create Storage client:", err)
	}

	urlSvc := &UrlService{
		storageClient:  tempSC,
		saKey:          os.Getenv("SA_KEY"),
		bucketName:     "send-this-42ae.appspot.com",
		googleAccessID: "send-this-42ae@appspot.gserviceaccount.com",
		filetype:       "",
	}
	return urlSvc
}

var urlService *UrlService

func init() {
	urlService = NewUrlService()
}

func Geturls(w http.ResponseWriter, r *http.Request) {
	urlService.geturls(w, r)
}

func DownloadUrl(w http.ResponseWriter, r *http.Request) {
	urlService.downloadUrl(w, r)
}
