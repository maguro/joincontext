package joincontext_test

import (
	"context"
	"fmt"
	"time"

	"m4o.io/joincontext"
)

func ExampleJoin() {
	ctx1, cancel1 := context.WithCancel(context.Background())
	defer cancel1()
	ctx2 := context.Background()

	ctx, cancel := joincontext.Join(ctx1, ctx2)
	defer cancel()
	select {
	case <-ctx.Done():
	default:
		fmt.Println("context alive")
	}

	cancel1()

	// give some time to propagate
	time.Sleep(100 * time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	default:
	}

	// Output: context alive
	// context canceled
}
