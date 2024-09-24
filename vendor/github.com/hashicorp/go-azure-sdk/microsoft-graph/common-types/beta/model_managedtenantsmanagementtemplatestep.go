package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagementTemplateStep{}

type ManagedTenantsManagementTemplateStep struct {
	AcceptedVersion    *ManagedTenantsManagementTemplateStepVersion   `json:"acceptedVersion,omitempty"`
	Category           *ManagedTenantsManagementCategory              `json:"category,omitempty"`
	CreatedByUserId    nullable.Type[string]                          `json:"createdByUserId,omitempty"`
	CreatedDateTime    nullable.Type[string]                          `json:"createdDateTime,omitempty"`
	Description        nullable.Type[string]                          `json:"description,omitempty"`
	DisplayName        nullable.Type[string]                          `json:"displayName,omitempty"`
	InformationLinks   *[]ActionUrl                                   `json:"informationLinks,omitempty"`
	LastActionByUserId nullable.Type[string]                          `json:"lastActionByUserId,omitempty"`
	LastActionDateTime nullable.Type[string]                          `json:"lastActionDateTime,omitempty"`
	ManagementTemplate *ManagedTenantsManagementTemplate              `json:"managementTemplate,omitempty"`
	PortalLink         *ActionUrl                                     `json:"portalLink,omitempty"`
	Priority           nullable.Type[int64]                           `json:"priority,omitempty"`
	UserImpact         nullable.Type[string]                          `json:"userImpact,omitempty"`
	Versions           *[]ManagedTenantsManagementTemplateStepVersion `json:"versions,omitempty"`

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

func (s ManagedTenantsManagementTemplateStep) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagementTemplateStep{}

func (s ManagedTenantsManagementTemplateStep) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementTemplateStep
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementTemplateStep: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementTemplateStep: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managementTemplateStep"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementTemplateStep: %+v", err)
	}

	return encoded, nil
}
