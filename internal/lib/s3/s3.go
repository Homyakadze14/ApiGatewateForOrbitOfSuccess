package s3

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/internal/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type S3Storage struct {
	*s3.S3
	bucket *string
	log    *slog.Logger
}

func NewS3Storage(l *slog.Logger, cfg config.S3) *S3Storage {
	const op = "S3Storage.NewS3Storage"

	log := l.With(
		slog.String("op", op),
	)

	log.Error("trying to connect to s3")
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(cfg.ACCESS_KEY, cfg.SECRET_ACCESS_KEY, ""),
		Endpoint:         aws.String(cfg.ENDPOINT),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Error("failed to connect to s3")
		panic(fmt.Errorf("%s: %w", op, err))
	}
	s3Client := s3.New(newSession)
	log.Error("successfully connected to s3")

	return &S3Storage{s3Client, &cfg.BUCKET_NAME, l}
}

func (s *S3Storage) saveToS3(urlCh chan<- string, errCh chan<- error, photo io.ReadSeeker) {
	uid := uuid.New().String()
	_, err := s.PutObject(&s3.PutObjectInput{
		Body:   photo,
		Bucket: s.bucket,
		Key:    aws.String(uid),
	})

	if err != nil {
		errCh <- err
	}

	urlCh <- fmt.Sprintf("%s/%s/%s", s.Endpoint, *s.bucket, uid) + ";"
}

func (s *S3Storage) Save(files []io.ReadSeeker) (string, error) {
	urls := ""

	urlChan := make(chan string)
	errChan := make(chan error)

	defer close(urlChan)
	defer close(errChan)

	for _, file := range files {
		go s.saveToS3(urlChan, errChan, file)
	}

	for i := 0; i < len(files); i++ {
		select {
		case url := <-urlChan:
			urls += url
		case err := <-errChan:
			return "", err
		}
	}

	return urls, nil
}
