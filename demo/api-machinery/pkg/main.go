package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func testTicker() {

	wait.Forever(func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}, time.Second)
}
func testTickUnitl(stopCh <-chan struct{}) {

	wait.Until(func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}, 0, stopCh)
}

func testPoll() {
	wait.Poll(time.Second, time.Second*5, func() (done bool, err error) {

		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		//return true, nil
		return false, nil
	})
}

func testPoll1() {
	wait.PollImmediate(time.Second, time.Second*5, func() (done bool, err error) {

		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		//return true, nil
		return false, nil
	})
}

func main() {

	//testTicker()

	//stop := make(chan struct{})
	testTickUnitl(wait.NeverStop)
	//time.Sleep(10 * time.Second)
	//close(stop)
	//fmt.Println("end!")

	//testPoll()

	//testPoll1()

}
