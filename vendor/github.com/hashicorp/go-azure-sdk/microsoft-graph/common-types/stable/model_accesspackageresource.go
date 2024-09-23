package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageResource{}

type AccessPackageResource struct {
	// Contains information about the attributes to be collected from the requestor and sent to the resource application.
	Attributes *[]AccessPackageResourceAttribute `json:"attributes,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// A description for the resource.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the resource, such as the application name, group name or site name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Contains the environment information for the resource. This can be set using either the @odata.bind annotation or the
	// environment's originId.Supports $expand.
	Environment *AccessPackageResourceEnvironment `json:"environment,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

	// The unique identifier of the resource in the origin system. For a Microsoft Entra group, this is the identifier of
	// the group.
	OriginId nullable.Type[string] `json:"originId,omitempty"`

	// The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
	OriginSystem nullable.Type[string] `json:"originSystem,omitempty"`

	// Read-only. Nullable. Supports $expand.
	Roles *[]AccessPackageResourceRole `json:"roles,omitempty"`

	// Read-only. Nullable. Supports $expand.
	Scopes *[]AccessPackageResourceScope `json:"scopes,omitempty"`

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

func (s AccessPackageResource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageResource{}

func (s AccessPackageResource) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageResource: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "modifiedDateTime")
	delete(decoded, "roles")
	delete(decoded, "scopes")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageResource: %+v", err)
	}

	return encoded, nil
}
