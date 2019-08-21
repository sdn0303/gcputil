package errorreporting

import (
	"cloud.google.com/go/errorreporting"
	"context"
	"log"
)

type ErrorReporting struct {
	Client *errorreporting.Client
}

func New(ctx context.Context, projectID, name string) *ErrorReporting {
	c, err := errorreporting.NewClient(ctx, projectID, errorreporting.Config{
		ServiceName: name,
		OnError: func(err error) {
			log.Printf("Could not log error: %v", err)
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return &ErrorReporting{
		Client: c,
	}
}

func (errReporting *ErrorReporting) SendError(err error) {
	if err != nil {
		errReporting.Client.Report(errorreporting.Entry{
			Error: err,
		})
		log.Println(err)
	}
	return
}
