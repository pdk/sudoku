## channel bits

### declare a channel

    var myChan chan int

default value is nil

### make a new channel

    myChan = make(chan int)
    myChan = make(chan int, 10)

First example is unbuffered (immediatly blocking). Second blocks when 10 items are inline.

### read/write channels with `<-`

    myChan <- 13
    myVar := <- myChan


### using channels w/o goroutines will just break

https://play.golang.org/p/TjygTAqy3Z5


### `range` reads from a channel until closed

    for x := range myChan {
        ...
    }

### ways to break your program

1. reading/writing a `nil` channel blocks forever
2. closing a `nil` or closed channel causes a `panic`
3. reading a closed channel returns the zero value
4. writing a closed channel causes a `panic`

How to read a potentially closed channel:

    myVar, ok := <- myChan

This will block until either a value is available, or the channel is closed. If we receive a value, `ok` will be true. If the channel was closed, `ok` will be false.

### multiplexing

    for {
        select {
            case v := <- alphaChan:
                ...
            case v := <- betaChan:
                ...
        }
    }

`select` will block until there is an available action on one of the channels. (A `default` can be used when there is nothing available, but that's not often used.)

### multiplexing with closures

    for {
        select {
            case v, ok := <- alphaChan:
                if !ok {
                    alphaChan = nil
                    break
                }
                ...
            case v, ok := <- betaChan:
                if !ok {
                    betaChan = nil
                    break
                }
        }
    }