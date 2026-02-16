package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagementTemplateStepVersion{}

type ManagedTenantsManagementTemplateStepVersion struct {
	AcceptedFor        *ManagedTenantsManagementTemplateStep             `json:"acceptedFor,omitempty"`
	ContentMarkdown    nullable.Type[string]                             `json:"contentMarkdown,omitempty"`
	CreatedByUserId    nullable.Type[string]                             `json:"createdByUserId,omitempty"`
	CreatedDateTime    nullable.Type[string]                             `json:"createdDateTime,omitempty"`
	Deployments        *[]ManagedTenantsManagementTemplateStepDeployment `json:"deployments,omitempty"`
	LastActionByUserId nullable.Type[string]                             `json:"lastActionByUserId,omitempty"`
	LastActionDateTime nullable.Type[string]                             `json:"lastActionDateTime,omitempty"`
	Name               nullable.Type[string]                             `json:"name,omitempty"`
	TemplateStep       *ManagedTenantsManagementTemplateStep             `json:"templateStep,omitempty"`
	Version            nullable.Type[int64]                              `json:"version,omitempty"`
	VersionInformation nullable.Type[string]                             `json:"versionInformation,omitempty"`

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

func (s ManagedTenantsManagementTemplateStepVersion) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagementTemplateStepVersion{}

func (s ManagedTenantsManagementTemplateStepVersion) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementTemplateStepVersion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementTemplateStepVersion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementTemplateStepVersion: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managementTemplateStepVersion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementTemplateStepVersion: %+v", err)
	}

	return encoded, nil
}
