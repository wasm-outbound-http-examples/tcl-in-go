package main

import (
	"io"
	"net/http"

	"github.com/skx/critical/interpreter"
)

func main() {
	code := `
	# custom command
	puts [httpget "https://httpbin.org/anything"]
	`

	interp, err := interpreter.New(code)
	if err != nil {
		panic(err)
	}

	httpGetFunc := func(i *interpreter.Interpreter, args []string) (string, error) {
		url := args[0]
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return string(body), nil
	}

	interp.RegisterBuiltin("httpget", httpGetFunc)

	_, err = interp.Evaluate()
	if err != nil {
		panic(err)
	}
}
