package storage

type Storage interface {
	BucketName() string
	SignedUrl(object string) (string, error)
}
