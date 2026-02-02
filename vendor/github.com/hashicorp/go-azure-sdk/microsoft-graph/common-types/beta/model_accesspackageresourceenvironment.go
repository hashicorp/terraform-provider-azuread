package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageResourceEnvironment{}

type AccessPackageResourceEnvironment struct {
	// Read-only. Required.
	AccessPackageResources *[]AccessPackageResource `json:"accessPackageResources,omitempty"`

	// Connection information of an environment used to connect to a resource.
	ConnectionInfo *ConnectionInfo `json:"connectionInfo,omitempty"`

	// The display name of the user that created this object.
	CreatedBy nullable.Type[string] `json:"createdBy,omitempty"`

	// The date and time that this object was created. The DateTimeOffset type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The description of this object.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of this object.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Determines whether this is default environment or not. It is set to true for all static origin systems, such as
	// Microsoft Entra groups and Microsoft Entra Applications.
	IsDefaultEnvironment nullable.Type[bool] `json:"isDefaultEnvironment,omitempty"`

	// The display name of the entity that last modified this object.
	ModifiedBy nullable.Type[string] `json:"modifiedBy,omitempty"`

	// The date and time that this object was last modified. The DateTimeOffset type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

	// The unique identifier of this environment in the origin system.
	OriginId *string `json:"originId,omitempty"`

	// The type of the resource in the origin system, that is, SharePointOnline. Requires $filter (eq).
	OriginSystem *string `json:"originSystem,omitempty"`

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

func (s AccessPackageResourceEnvironment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageResourceEnvironment{}

func (s AccessPackageResourceEnvironment) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageResourceEnvironment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageResourceEnvironment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageResourceEnvironment: %+v", err)
	}

	delete(decoded, "accessPackageResources")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageResourceEnvironment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageResourceEnvironment: %+v", err)
	}

	return encoded, nil
}
