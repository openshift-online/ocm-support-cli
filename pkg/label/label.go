package label

import (
	"ocm-support-cli/pkg/types"
	"time"

	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
)

type Label struct {
	types.Meta
	ID        string
	CreatedAt time.Time
	Key       string
	UpdatedAt time.Time
	Value     string
	Internal  bool
	HREF      string
}

type LabelsList []Label

func PresentLabels(labelResponses []*v1.Label) LabelsList {
	var labelsList []Label
	for _, labelResponse := range labelResponses {
		lbl := Label{
			ID:        labelResponse.ID(),
			CreatedAt: labelResponse.CreatedAt(),
			Key:       labelResponse.Key(),
			UpdatedAt: labelResponse.UpdatedAt(),
			Value:     labelResponse.Value(),
			Internal:  labelResponse.Internal(),
			HREF:      labelResponse.HREF(),
		}
		labelsList = append(labelsList, lbl)
	}
	return labelsList
}
