package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageResourceRequest{}

type AccessPackageResourceRequest struct {
	Catalog *AccessPackageCatalog `json:"catalog,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The type of the request. Use adminAdd to add a resource, if the caller is an administrator or resource owner,
	// adminUpdate to update a resource, or adminRemove to remove a resource.
	RequestType *AccessPackageRequestType `json:"requestType,omitempty"`

	Resource *AccessPackageResource `json:"resource,omitempty"`

	// The outcome of whether the service was able to add the resource to the catalog. The value is delivered if the
	// resource was added or removed, and deliveryFailed if it couldn't be added or removed. Read-only.
	State *AccessPackageRequestState `json:"state,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessPackageResourceRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageResourceRequest{}

func (s AccessPackageResourceRequest) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageResourceRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageResourceRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageResourceRequest: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "state")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageResourceRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageResourceRequest: %+v", err)
	}

	return encoded, nil
}
