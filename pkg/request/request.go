package request

import (
	"fmt"
	"net/url"

	sdk "github.com/openshift-online/ocm-sdk-go"
)

func DeleteRequest(url string, noDryRun bool, connection *sdk.Connection) error {
	request := connection.Delete()
	err := ApplyPathArg(request, url)
	if err != nil {
		return fmt.Errorf("can't parse url '%s': %v\n", url, err)
	}
	if !noDryRun {
		fmt.Printf("DRYRUN: Would have called %v if not passed.\n", request.GetPath())
		return nil
	}
	response, err := request.Send()
	if err != nil {
		return fmt.Errorf("can't send request: %v", err)
	}
	if response.Status() != 204 {
		return fmt.Errorf("operation failed with response status %v", response.Status())
	}
	return nil
}

func PatchRequest(url string, noDryRun bool, connection *sdk.Connection) error {
	request := connection.Patch()
	err := ApplyPathArg(request, url)
	if err != nil {
		return fmt.Errorf("can't parse url '%s': %v\n", url, err)
	}
	if !noDryRun {
		fmt.Printf("DRYRUN: Would have called %v if not passed.\n", request.GetPath())
		return nil
	}
	response, err := request.Send()
	if err != nil {
		return fmt.Errorf("can't send request: %v", err)
	}
	if response.Status() != 204 {
		return fmt.Errorf("operation failed with response status %v", response.Status())
	}
	return nil
}

func ApplyPathArg(request *sdk.Request, value string) error {
	parsed, err := url.Parse(value)
	if err != nil {
		return err
	}
	request.Path(parsed.Path)
	query := parsed.Query()
	for name, values := range query {
		for _, value := range values {
			request.Parameter(name, value)
		}
	}
	return nil
}
