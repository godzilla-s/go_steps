package queue

import (
	"fmt"
	"time"
)

type Queue struct {
	shutdown		chan bool
	addWorkChan		chan Work
	workChanList	chan []Work
	workChanSize	int
	timerNotify		chan bool
	controller		Controller
}

func NewQueue(workChanSize int, shutdown chan bool, controller Controller) *Queue {
	queue := Queue {
		shutdown : shutdown,
		addWorkChan: make(chan Work),
		workChanList: make(chan []Work, workChanSize),
		workChanSize: workChanSize,
		timerNotify: make(chan bool),
		controller: controller,
	}

	return &queue
}

func (queue *Queue) addWork(work Work) error {
	fmt.Println("<-add Work")

	queue.addWorkChan <- work

	return nil
}

func workHandler(queue Queue) {
	fmt.Println("Work Handler")

	workCount := 0
	var workList []Work

	flush := func(works []Work) {
		tempWorks := make([]Work, len(works))
		copy(tempWorks, works)
		queue.workChanList <- tempWorks
		
		workList = workList[:0]
		workCount = 0
	}

	for {
		select {
			case work := <-queue.addWorkChan:
				fmt.Println("Get Data:")
				if workCount > 50 {
					fmt.Println("work size need flush")
					flush(workList)
				}

				workList = append(workList, work)
				workCount++
				fmt.Println("work size: ", workCount)
			case <- queue.timerNotify:
				fmt.Println("time notify")
				if workCount > 0 {
					fmt.Println("time to flush workchanlist")
					flush(workList)
				}
			case <- queue.shutdown:
				if workCount > 0 {
					fmt.Println("close to flush workchanlist")
					flush(workList)
				}
				return 
		}
	}
} 

func flushWorkCall(works []Work, queue Queue) {
	var is_error bool
	var err error
	var buff []byte
	var data string

	fmt.Println("flushWorkCall begin")
	for v, err := range works {
		buff, err = v.GetWork()
		if err != nil {
			is_error = ture
		} else {
			fmt.Println("Work:", string(buff))
		}
		data = data + string(buff)
	}

	queue.controller.Do([]byte(data))

	if is_error {
		queue.workChanList <- works
	}
	fmt.Println("flushWorkCall Done")
}

func workLoopTimer(queue Queue) {
	for {
		select {
			case works := <- queue.workChanList:
				fmt.Println("call flushwork")
				flushWorkCall(works, queue)
			case <- time.After(10 * time.Second):
				fmt.Println("timer to notify")
				queue.timerNotify <- ture
		}
	}
}
