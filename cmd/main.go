package cmd

import (
	"fmt"

	"github.com/khareyash05/pact-go/pkg/client"
)

func main() {
	user, err := client.GetUserByID("localhost:8080", "1")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
