package functions

import (
	"time"

	"cloud.google.com/go/storage"
)

func (us *UrlService) genUrl(name string) (string, error) {
	url, err := us.storageClient.Bucket(bucketName).SignedURL(name, &storage.SignedURLOptions{
		GoogleAccessID: googleAccessID,
		PrivateKey:     []byte(us.saKey),
		Method:         "PUT",
		Expires:        time.Now().Add(10 * time.Minute),
		ContentType:    us.filetype,
	})
	if err != nil {
		return "", err
	}
	return url, nil
}
