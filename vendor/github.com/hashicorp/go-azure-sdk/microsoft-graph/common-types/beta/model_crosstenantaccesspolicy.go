package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TenantRelationshipAccessPolicyBase = CrossTenantAccessPolicy{}

type CrossTenantAccessPolicy struct {
	// Used to specify which Microsoft clouds an organization would like to collaborate with. By default, this value is
	// empty. Supported values for this field are: microsoftonline.com, microsoftonline.us, and partner.microsoftonline.cn.
	AllowedCloudEndpoints *[]string `json:"allowedCloudEndpoints,omitempty"`

	// Defines the default configuration for how your organization interacts with external Microsoft Entra organizations.
	Default *CrossTenantAccessPolicyConfigurationDefault `json:"default,omitempty"`

	// Defines partner-specific configurations for external Microsoft Entra organizations.
	Partners *[]CrossTenantAccessPolicyConfigurationPartner `json:"partners,omitempty"`

	// Represents the base policy in the directory for multi-tenant organization settings.
	Templates *PolicyTemplate `json:"templates,omitempty"`

	// Fields inherited from TenantRelationshipAccessPolicyBase

	// The raw JSON definition of the cross-tenant access policy. Deprecated. Do not use.
	Definition *[]string `json:"definition,omitempty"`

	// Fields inherited from PolicyBase

	// Description for this policy. Required.
	Description string `json:"description"`

	// Display name for this policy. Required.
	DisplayName string `json:"displayName"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s CrossTenantAccessPolicy) TenantRelationshipAccessPolicyBase() BaseTenantRelationshipAccessPolicyBaseImpl {
	return BaseTenantRelationshipAccessPolicyBaseImpl{
		Definition:      s.Definition,
		Description:     s.Description,
		DisplayName:     s.DisplayName,
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s CrossTenantAccessPolicy) PolicyBase() BasePolicyBaseImpl {
	return BasePolicyBaseImpl{
		Description:     s.Description,
		DisplayName:     s.DisplayName,
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s CrossTenantAccessPolicy) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s CrossTenantAccessPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CrossTenantAccessPolicy{}

func (s CrossTenantAccessPolicy) MarshalJSON() ([]byte, error) {
	type wrapper CrossTenantAccessPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CrossTenantAccessPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CrossTenantAccessPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.crossTenantAccessPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CrossTenantAccessPolicy: %+v", err)
	}

	return encoded, nil
}
