package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageResource{}

type AccessPackageResource struct {
	// Contains the environment information for the resource. This environment can be set using either the @odata.bind
	// annotation or the environment's originId. Supports $expand.
	AccessPackageResourceEnvironment *AccessPackageResourceEnvironment `json:"accessPackageResourceEnvironment,omitempty"`

	// Read-only. Nullable. Supports $expand.
	AccessPackageResourceRoles *[]AccessPackageResourceRole `json:"accessPackageResourceRoles,omitempty"`

	// Read-only. Nullable. Supports $expand.
	AccessPackageResourceScopes *[]AccessPackageResourceScope `json:"accessPackageResourceScopes,omitempty"`

	// The name of the user or application that first added this resource. Read-only.
	AddedBy nullable.Type[string] `json:"addedBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	AddedOn nullable.Type[string] `json:"addedOn,omitempty"`

	// Contains information about the attributes to be collected from the requestor and sent to the resource application.
	Attributes *[]AccessPackageResourceAttribute `json:"attributes,omitempty"`

	// A description for the resource.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the resource, such as the application name, group name, or site name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// True if the resource is not yet available for assignment. Read-only.
	IsPendingOnboarding nullable.Type[bool] `json:"isPendingOnboarding,omitempty"`

	// The unique identifier of the resource in the origin system. In the case of a Microsoft Entra group, originId is the
	// identifier of the group. Supports $filter (eq).
	OriginId nullable.Type[string] `json:"originId,omitempty"`

	// The type of the resource in the origin system, such as SharePointOnline, AadApplication, or AadGroup. Supports
	// $filter (eq).
	OriginSystem nullable.Type[string] `json:"originSystem,omitempty"`

	// The type of the resource, such as Application if it is a Microsoft Entra connected application, or SharePoint Online
	// Site for a SharePoint Online site.
	ResourceType nullable.Type[string] `json:"resourceType,omitempty"`

	// A unique resource locator for the resource, such as the URL for signing a user into an application.
	Url nullable.Type[string] `json:"url,omitempty"`

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

	delete(decoded, "accessPackageResourceRoles")
	delete(decoded, "accessPackageResourceScopes")
	delete(decoded, "addedBy")
	delete(decoded, "addedOn")
	delete(decoded, "isPendingOnboarding")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageResource: %+v", err)
	}

	return encoded, nil
}
