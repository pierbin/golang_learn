package main

import (
	"context"
	"fmt"
	"time"
)

// Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context.
// The chain of function calls between them must propagate the Context,
// optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue.
// When a Context is canceled, all Contexts derived from it are also canceled.
// The WithCancel, WithDeadline, and WithTimeout functions take a Context (the parent) and return a derived Context (the child) and a CancelFunc.
// Calling the CancelFunc cancels the child and its children, removes the parent's reference to the child, and stops any associated timers.

// Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it.
// The Context should be the first parameter, typically named ctx:
// func DoSomething(ctx context.Context, arg Arg) error {
// ... use ctx ...
// }

// Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to use.

/*
	type CancelFunc func()

	A CancelFunc tells an operation to abandon its work.
	A CancelFunc does not wait for the work to stop.
	A CancelFunc may be called by multiple goroutines simultaneously.
	After the first call, subsequent calls to a CancelFunc do nothing.
*/

/*
	func Background() Context
	ctx, cancel:= context.Background()

	Background returns a non-nil, empty Context.
	It is never canceled, has no values, and has no deadline.
	It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
	It is used to derive other contexts and has an empty context as a return type.
*/

/*
	func TODO() Context
	ctx, cancel := context.TODO()

	TODO returns a non-nil, empty Context. Code should use context.
	TODO when it's unclear which Context to use or it is not yet available.
	because the surrounding function has not yet been extended to accept a Context parameter.

	TODO is to never pass nil context , instead, use of TODO is advised .
*/

/*
	func WithValue(parent Context, key, val interface{}) Context
	ctx := context.WithValue(context.Background(), key, “test”)

	WithValue returns a copy of parent in which the value associated with key is val.
	It creates a new context based on a parent context and adds a value to a given key.
	It allows you to store any type of data inside the context.
	It is not recommended to pass in critical parameters using context value, instead, functions should accept those values in the signature making it explicit.
*/

func main() {
	withCancelUsage()
	withDeadlineUsage()
	withValueUsage()
}

// This example demonstrates the use of a cancelable context to prevent a goroutine leak.
// By the end of the example function, the goroutine started by gen will return without leaking.
// WithCancel returns a copy of parent with a new Done channel.
// The returned context's Done channel is closed when the returned cancel function is called or when the parent context's Done channel is closed, whichever happens first.

// context.WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// ctx, cancel := context.WithCancel(context.Background())
// This parent function is passed in as argument.
// This parent context can either be a background context or a context that was passed into the function.
func withCancelUsage() {
	// gen generates integers in a separate goroutine and sends them to the returned channel.
	// The callers of gen need to cancel the context once they are done consuming generated integers not to leak the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println("cancel usage: ", n)
		if n == 5 {
			break
		}
	}
}

// WithDeadline returns a copy of the parent context with the deadline adjusted to be no later than d.
// If the parent's deadline is already earlier than d, WithDeadline(parent, d) is semantically equivalent to parent.
// The returned context's Done channel is closed when the deadline expires, when the returned cancel function is called,
// or when the parent context's Done channel is closed, whichever happens first.
// When that context gets canceled because of the deadline running out, all the functions that got the context get notified to stop work and return.

// context.WithDeadline(parent Context, d time.Time) (ctx Context, cancel CancelFunc)
// ctx, cancel := context.WithDeadline(context.Background(),time.Now().Add(2 * time.Second))

// Similar to WithDeadline , it cancels the contexts but this cancellation is based upon the time duration.
// This function returns a derived context that gets canceled if the cancel function is called or the timeout duration is exceeded.
// context.WithTimeout(parent Context, timeout time.Duration) (ctx Context, cancel CancelFunc)
// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(150)*time.Millisecond)
func withDeadlineUsage() {
	const shortDuration = 1 * time.Millisecond
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice calling its cancellation function in any case.
	// Failure to do so may keep the context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// This example demonstrates how a value can be passed to the context and also how to retrieve it if it exists.
func withValueUsage() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("color")
	ctx := context.WithValue(context.Background(), k, "PHP")

	f(ctx, k)
	f(ctx, favContextKey("test"))
}

// For all the functions within the request call paths, functions are accepting ctx Context.Context as first argument.
// In this way, all request bound values will be propagated in a consistent fashion.
