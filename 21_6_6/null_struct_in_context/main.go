package main

import (
	"context"
	"fmt"
)

//作key，用于context中找value，作为唯一的key
type ctxKey struct{}

func main() {
	ctx := context.WithValue(context.Background(), &ctxKey{}, "Hello empty struct")
	fmt.Println(getValue(ctx))
	// output: Hello empty struct
}

func getValue(ctx context.Context) string {
	return ctx.Value(&ctxKey{}).(string)
}
