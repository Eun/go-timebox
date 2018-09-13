package timebox_test

import (
	"fmt"
	"time"

	timebox "github.com/Eun/go-timebox"
	"github.com/tevino/abool"
)

func ExampleTimebox() {
	// run time.Sleep(time.Minute), if the call takes longer than a second continue with execution
	timebox.Timebox(time.Second, time.Sleep, time.Minute)
}

func ExampleTimebox_inlineFunction() {
	result, err := timebox.Timebox(time.Second, func() (int, error) {
		// Some heavy operation
		return 42, nil
	})
	// check if Timebox() has an error
	if err != nil {
		panic(err)
	}
	// check if the function has an error
	if result[1] != nil {
		panic(result[1])
	}
	fmt.Printf("The answer is %d\n", result[0].(int))
}

func ExampleTimebox_checkCancelInsideTheFunction() {
	canceled := abool.New()
	_, err := timebox.Timebox(time.Second, func() (int, error) {
		time.Sleep(time.Second)
		// did we already cancel?
		if canceled.IsSet() {
			return 0, nil
		}
		time.Sleep(time.Second)
		return 42, nil
	})
	if timebox.IsTimeoutError(err) {
		canceled.Set()
	}
}
