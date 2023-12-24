package storage

type Storage interface {
	Contains(short string) bool
	Save(short, origin string)
	Get(short string) (string, error)
}
