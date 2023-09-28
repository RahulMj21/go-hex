package main

import (
	"fmt"

	"github.com/RahulMj21/go-hex/internal/adapters/app/api"
	"github.com/RahulMj21/go-hex/internal/adapters/core/arithmetic"
	"github.com/RahulMj21/go-hex/internal/ports"
)

func main() {
	var cors ports.ArithmeticPort
	cors = arithmetic.NewAdapter()
	app := api.NewAdapter(cors)
	fmt.Println(app.GetAddition(1, 2))
}
