package control

import (
	"fmt"
)

type Controller struct {
}

func NewController() *Controller {
	controller := Controller {}
	
	return &Controller
}

func (controller *Controller) Do(data []byte) error {
	fmt.Println("controller: ", string(data))
	return nil
}
