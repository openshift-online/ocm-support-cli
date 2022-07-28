package utils

import (
	"fmt"
	"os"

	"github.com/nwidger/jsoncolor"
)

const MaxRecords = 100

func PrettyPrint(data interface{}) {
	// marshal and pretty print the account(s)
	encoder := jsoncolor.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println("failed to encode the data")
		return
	}
}
