package parse

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-uuid"
)

type ObjectSubResourceId struct {
	objectId string
	subId    string
	Type     string
}

func NewObjectSubResourceID(objectId, typeId, subId string) ObjectSubResourceId {
	return ObjectSubResourceId{
		objectId: objectId,
		Type:     typeId,
		subId:    subId,
	}
}

func (id ObjectSubResourceId) String() string {
	return fmt.Sprintf("%s/%s/%s", id.objectId, id.Type, id.subId)
}

func ObjectSubResourceID(idString, expectedType string) (*ObjectSubResourceId, error) {
	parts := strings.Split(idString, "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("Object Resource ID should be in the format {objectId}/{type}/{subId} - but got %q", idString)
	}

	id := ObjectSubResourceId{
		objectId: parts[0],
		Type:     parts[1],
		subId:    parts[2],
	}

	if _, err := uuid.ParseUUID(id.objectId); err != nil {
		return nil, fmt.Errorf("Object ID isn't a valid UUID (%q): %+v", id.objectId, err)
	}

	if id.Type == "" {
		return nil, fmt.Errorf("Type in {objectID}/{type}/{subID} should not be empty")
	}

	if id.Type != expectedType {
		return nil, fmt.Errorf("Type in {objectID}/{type}/{subID} was expected to be %s, got %s", expectedType, parts[2])
	}

	if _, err := uuid.ParseUUID(id.subId); err != nil {
		return nil, fmt.Errorf("Object Sub Resource ID isn't a valid UUID (%q): %+v", id.subId, err)
	}

	return &id, nil
}
