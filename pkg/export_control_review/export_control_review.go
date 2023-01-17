package account

import (
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1Auth "github.com/openshift-online/ocm-sdk-go/authorizations/v1"
)

type ExportControlReview struct {
	Restricted bool `json:"restricted"`
}

func PostExportControlReview(username string, conn *sdk.Connection) (*v1Auth.ExportControlReviewResponse, error) {
	exportControlReviewRequest, err := v1Auth.NewExportControlReviewRequest().AccountUsername(username).Build()
	if err != nil {
		return nil, fmt.Errorf("can't build export control request: %w", err)
	}

	exportResponse, err := conn.Authorizations().V1().ExportControlReview().Post().Request(exportControlReviewRequest).Send()
	if err != nil {
		return nil, fmt.Errorf("can't retrieve export control: %w", err)
	}

	responseBody := exportResponse.Response()

	return responseBody, nil
}

func PresentExportControlReview(exportControlReview *v1Auth.ExportControlReviewResponse) *ExportControlReview {
	var export *ExportControlReview = nil

	if exportControlReview != nil {
		export = &ExportControlReview{
			Restricted: exportControlReview.Restricted(),
		}
	}

	return export
}
