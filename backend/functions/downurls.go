package functions

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
)

func (us *UrlService) downloadUrl(w http.ResponseWriter, r *http.Request) {
	us.saKey = os.Getenv("SA_KEY")
	parts := strings.Split(r.URL.Path, "/")
	fileUUID := parts[len(parts)-1]
	name := "temp/" + fileUUID
	log.Printf("Redirecting to a new download signed URL for ephemeral resource %q\n", name)

	downloadURL, err := us.storageClient.Bucket(bucketName).SignedURL(name, &storage.SignedURLOptions{
		GoogleAccessID: googleAccessID,
		PrivateKey:     []byte(us.saKey),
		Method:         "GET",
		Expires:        time.Now().Add(10 * time.Minute),
	})
	if err != nil {
		log.Printf("error generating secure download url: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error, down_short_url"})
		return
	}
	http.Redirect(w, r, downloadURL, http.StatusFound)
}
