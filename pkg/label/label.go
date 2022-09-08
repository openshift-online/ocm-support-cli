package label

import (
	"fmt"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
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

func GetLabels(query string, isInternal bool, size int, conn *sdk.Connection) ([]*v1.Label, error) {
	search := query
	if isInternal {
		search += " and internal=true"
	}
	labels, err := conn.AccountsMgmt().V1().Labels().List().Size(size).Search(search).Send()
	if err != nil {
		return []*v1.Label{}, fmt.Errorf("can't retrieve labels: %w", err)
	}
	return labels.Items().Slice(), nil
}

func GetLabel(id string, conn *sdk.Connection) (*v1.Label, error) {
	search := fmt.Sprint("id = '", id, "'")
	lblResponse, err := conn.AccountsMgmt().V1().Labels().List().Search(search).Send()
	if err != nil {
		return nil, fmt.Errorf("can't retrieve label: %w", err)
	}
	labels := lblResponse.Items().Slice()
	if len(labels) == 0 {
		return nil, fmt.Errorf("label with id %s not found", id)
	}
	return labels[0], nil
}
