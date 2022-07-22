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

func PresentLabels(labels []*v1.Label) LabelsList {
	var labelsList []Label
	for _, label := range labels {
		lbl := Label{
			ID:        label.ID(),
			CreatedAt: label.CreatedAt(),
			Key:       label.Key(),
			UpdatedAt: label.UpdatedAt(),
			Value:     label.Value(),
			Internal:  label.Internal(),
			HREF:      label.HREF(),
		}
		labelsList = append(labelsList, lbl)
	}
	return labelsList
}
