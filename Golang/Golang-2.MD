# Golang 2

1. We will cover
1. Pointers
1. Channels
1. Goroutine and Channel
1. Scaling Goroutine
1. Library support

## Pointers

Pointer, points to the address of a variable. It is used if you want some func/method to update the value of a data type.

https://goplay.tools/snippet/Au6uv9r47AG
```Go
package main

import (
	"fmt"
)

func updateP(i *int) {
	*i = 100
}
func update(i int) {
	i = 200
}
func main() {
	fmt.Println("Pointers")
	i := 10

	fmt.Printf("val of i=%d, add of i=%d\n", i, &i)
	update(i)
	fmt.Printf("val of i=%d, add of i=%d\n", i, &i)
	updateP(&i)
	fmt.Printf("val of i=%d, add of i=%d\n", i, &i)
}
```

## Channels
We can consider channel as queue. But, they are much more than being a queue. They are of type buffered and unbuffered. Most commonly used are unbuffered. 
Once you write data into channel, data is immeadiately available for reading.

```go
c1 := make(chan int) // unbuffered ch of type int
c2 := make(chan int, 10) // buffered ch of type int of cap 10
```

Use `c1<-` to write data and `<-c1` to read data. Example https://goplay.tools/snippet/Gv7y0OMwoWO
```go
package main

import (
	"fmt"
)
func main() {
	fmt.Println("Unbuffered Channels")
	ch := make(chan int)
	// ch <- 10
	// go func() {
	// 	ch <- 10
	// }()
	fmt.Println(<-ch)
}
```

## Goroutine and Channel
We can specify the channel is read only or write only https://goplay.tools/snippet/jFmpiY-yPmw
```go
package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1) 
	pongs := make(chan string, 1)
	ping(pings, "passed message") // Convert to Goroutine and run
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
```

We can read from multiple channel simultaneously. We will use `select / case` key word to read from channel.
Any channel which recieves the data will cause the case part to execute.
https://goplay.tools/snippet/lpHJZmJjbrg
```Go
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        // what will happen if we uncomment this. Non blocking operations
        // default:
        //     fmt.Println("Default)
        }

    }
}
```
We can use `range` to read from a channel. For loop will execute till channel has data. 
https://goplay.tools/snippet/7AH1mTnuriY
```Go
package main

import "fmt"

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }
}
```
We can close channel and stop execution. If we send data on closed channel it panics
https://goplay.tools/snippet/l43umI2hLip
```Go
package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		// time.Sleep(time.Second)
		for j := range jobs {

			fmt.Println("Recevied ", j)
		}
		done <- true
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	// 1.
	// val, open := <-jobs
	// fmt.Println("val ", val, "open ", open)

	fmt.Println("Work finished ", <-done)
	// 2.
	// val, open = <-jobs
	// fmt.Println("val ", val, "open ", open)

	// 3.
	// jobs <- 4
}
```

## Scaling Goroutine
Can we wait for goroutine to complete before we move on to next code?

WaitGroup are used to wait for goroutine to finish the execution. https://goplay.tools/snippet/bePBS51SEGm
```Go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}

```
What happens if we keep on creating goroutines?

We can limit the no of goroutines.  https://goplay.tools/snippet/VtzV2SsmAf9
```Go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(id int, guard chan struct{}) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	<-guard
}

func main() {
	guard := make(chan struct{}, 2)
	for i := 0; i < 10; i++ {
		guard <- struct{}{} // Show without guard
		go worker(i, guard) 
		fmt.Println("No of goroutine running are ", runtime.NumGoroutine())
	}
}
```

## Library support

1. DB 
	1. gorm https://gorm.io/docs/index.html
	1. sqlx https://github.com/jmoiron/sqlx
1. Logs
	1. Default https://pkg.go.dev/log
	1. zerolog https://github.com/rs/zerolog
	1. zap     https://github.com/uber-go/zap
1. Webserver
	1. Default https://go.dev/doc/articles/wiki/
	1. Gin  https://github.com/gin-gonic/gin
1. Scientific computation
	1. Numpy/Scipy -> gonum https://github.com/gonum/gonum
	1. Matplotlib -> plot https://github.com/gonum/plot
	1. Panda -> gota https://github.com/go-gota/gota
	1. ML https://github.com/sjwhitworth/golearn
	1. AI tensorflow https://pkg.go.dev/github.com/tensorflow/tensorflow
	1. AI keras https://pkg.go.dev/github.com/edotau/goFish/keras
	1. AI https://gorgonia.org/
	1. 
