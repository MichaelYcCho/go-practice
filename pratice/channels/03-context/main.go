package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context: ", ctx)
	fmt.Println("context err: ", ctx.Err())
	fmt.Printf("context type: %T\t\n", ctx)
	fmt.Println("=======================================")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context: ", ctx)
	fmt.Println("context err: ", ctx.Err())
	fmt.Printf("context type: %T\t\n", ctx)
	fmt.Println("=======================================")

	cancel()

	fmt.Println("context: ", ctx)
	fmt.Println("context err: ", ctx.Err())
	fmt.Printf("context type: %T\t\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type: %T\t\n", cancel)
	fmt.Println("=======================================")
}
