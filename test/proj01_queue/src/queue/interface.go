package queue

type Controller interface {
	Do(data []byte) error
}

type Work interface {
	GetWork() ([]byte, error)
}
