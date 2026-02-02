package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageResourceRequest{}

type AccessPackageResourceRequest struct {
	AccessPackageResource *AccessPackageResource `json:"accessPackageResource,omitempty"`

	// The unique ID of the access package catalog.
	CatalogId nullable.Type[string] `json:"catalogId,omitempty"`

	ExecuteImmediately nullable.Type[bool] `json:"executeImmediately,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// If set, doesn't add the resource.
	IsValidationOnly nullable.Type[bool] `json:"isValidationOnly,omitempty"`

	// The requestor's justification for adding or removing the resource.
	Justification nullable.Type[string] `json:"justification,omitempty"`

	// The outcome of whether the service was able to add the resource to the catalog. The value is Delivered if the
	// resource was added or removed. Read-Only.
	RequestState nullable.Type[string] `json:"requestState,omitempty"`

	RequestStatus nullable.Type[string] `json:"requestStatus,omitempty"`

	// Use AdminAdd to add a resource, if the caller is an administrator or resource owner, AdminUpdate to update a
	// resource, or AdminRemove to remove a resource.
	RequestType nullable.Type[string] `json:"requestType,omitempty"`

	// Read-only. Nullable. Supports $expand.
	Requestor *AccessPackageSubject `json:"requestor,omitempty"`

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

	delete(decoded, "requestor")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageResourceRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageResourceRequest: %+v", err)
	}

	return encoded, nil
}
