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

var Aliases = map[string][]string{
	"accountcapability":      {"ac", "acctcapability"},
	"organizationcapability": {"oc", "orgcapability"},
	"subscriptioncapability": {"sc", "subcapability"},

	"accountlabel":      {"al", "acctlabel"},
	"organizationlabel": {"ol", "orglabel"},
	"subscriptionlabel": {"sl", "sublabel"},

	"applicationrolebinding":  {"arb", "approlebinding"},
	"organizationrolebinding": {"orb", "orgrolebinding"},
	"subscriptionrolebinding": {"srb", "subrolebinding"},

	"registrycredentials": {"rcs"},
	"organizations":       {"orgs"},
	"accounts":            {"accts"},
	"subscriptions":       {"subs"},

	"capability":   {"cap"},
	"organization": {"org"},
	"subscription": {"sub"},
	"account":      {"acct"},
}
