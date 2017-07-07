package control

import (
	"fmt"
)

type Work struct {
	data  []byte
}

func NewWork(buff []byte) *Work {
	work := Work {
		data: buff
	}

	return &work
}

func (work *Work) GerWork() ([]byte, error) {
	fmt.Println("Get Work: ", string(work.data))
	return work.data, nil
}
