package label

import (
	"fmt"
	"time"

	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"

	"github.com/openshift-online/ocm-support-cli/pkg/types"
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

func CreateLabel(key string, value string, isInternal bool) (*v1.Label, error) {
	lbl, err := v1.NewLabel().Key(key).Value(value).Internal(isInternal).Build()
	if err != nil {
		return nil, fmt.Errorf("can't create new label: %w", err)
	}
	return lbl, nil
}
