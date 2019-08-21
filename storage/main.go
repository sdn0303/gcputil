package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"io/ioutil"
	"log"
)

type Storage struct {
	BucketName string
	Client     *storage.Client
	Context    context.Context
}

func New(bucketName string, ctx context.Context) *Storage {
	c, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &Storage{
		BucketName: bucketName,
		Client:     c,
		Context:    ctx,
	}
}

func (gcs *Storage) Put(path, contentType string, data []byte) error {

	w := gcs.Client.Bucket(gcs.BucketName).Object(path).NewWriter(gcs.Context)
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

func (gcs *Storage) Read(path string) ([]byte, error) {

	r, err := gcs.Client.Bucket(gcs.BucketName).Object(path).NewReader(gcs.Context)
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
