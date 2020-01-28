package storage

import (
	"context"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
)

type Storage struct {
	BucketName string
	Client     *storage.Client
}

func New(ctx context.Context, bucketName string) *Storage {
	c, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &Storage{
		BucketName: bucketName,
		Client:     c,
	}
}

func (gcs *Storage) Create(ctx context.Context, path, contentType string, data []byte) error {

	w := gcs.Client.Bucket(gcs.BucketName).Object(path).NewWriter(ctx)
	w.ContentType = contentType

	if n, err := w.Write(data); err != nil {
		return err
	} else if n != len(data) {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

func (gcs *Storage) Read(ctx context.Context, path string) ([]byte, error) {

	r, err := gcs.Client.Bucket(gcs.BucketName).Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (gcs *Storage) Update() {

}

func (gcs *Storage) Delete() {

}
