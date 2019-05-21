// Thread tl prints odd numbers from 1 to 100; Thread t2 prints even numbers
// from 1 to 100. Write code in which the two threads, running concurrently,
// print the numbers from1 to 100 in order.

package main

import (
	"fmt"
	"sync"
)

type EvenOdd struct {
	wg sync.WaitGroup
	printRange, exitCount int
	turnEven bool
	notifyExit, notifyDone, notifyEven, notifyOdd chan bool
}

func NewEvenOdd(printRange int) *EvenOdd {
	this := new(EvenOdd)

	this.printRange = printRange

	// Prepare channels for synchronization.
	// Buffered channel to avoid sleeping-channels/deadlock.
	this.notifyExit = make(chan bool, 1)
	this.notifyDone = make(chan bool ,1)
	this.notifyEven = make(chan bool, 1)
	this.notifyOdd = make(chan bool, 1)

	// Start the print handler and allow the consumer to wait till the
	// end.
	this.wg.Add(1)
	go func() {
		defer this.wg.Done()
		this.printHandler()
	}()

	// Start odd and even printers.
	go this.printOdd()
	go this.printEven()

	return this
}

// Start execution with the desired channel first.
func (this *EvenOdd) Print() {
	this.turnEven = false
	this.notifyOdd <- true
	this.wg.Wait()
}

// Handle even and odd printers.
func (this *EvenOdd) printHandler() {
	for this.exitCount < 2 {
		select {
		case <- this.notifyExit:
			this.exitCount++
		case <- this.notifyDone:
			// Send notifications only if both receivers are
			// available.
			if this.exitCount == 0 {
				// Alternate notifications to the channels.
				this.turnEven = !this.turnEven
				if this.turnEven {
					this.notifyEven <- true
				} else {
					this.notifyOdd <- true
				}
			}
		}
	}

	// Close all the channels and exit.
	close(this.notifyExit)
	close(this.notifyDone)
	close(this.notifyEven)
	close(this.notifyOdd)
}

func (this *EvenOdd) waitEven() {
	<-this.notifyEven
}

func (this *EvenOdd) waitOdd() {
	<-this.notifyOdd
}

func (this *EvenOdd) notify() {
	this.notifyDone <- true
}

func (this *EvenOdd) exiting() {
	this.notifyExit <- true
}

func (this *EvenOdd) printEven() {
	go func() {
		defer this.exiting()
		for i := 2; i <= this.printRange; i += 2 {
			this.waitEven()
			fmt.Print(i, " ")
			this.notify()
		}
	}()
}

func (this *EvenOdd) printOdd() {
	go func() {
		defer this.exiting()
		for i := 1; i <= this.printRange; i += 2 {
			this.waitOdd()
			fmt.Print(i, " ")
			this.notify()
		}
	}()
}

func main() {
	resource := NewEvenOdd(100)
	resource.Print()
}
