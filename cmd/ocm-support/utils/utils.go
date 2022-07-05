package utils

import (
	"github.com/nwidger/jsoncolor"
	"os"
)

const MaxRecords = 100

func PrettyPrint(data interface{}) {
	// marshal and pretty print the account(s)
	encoder := jsoncolor.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(data)
}
