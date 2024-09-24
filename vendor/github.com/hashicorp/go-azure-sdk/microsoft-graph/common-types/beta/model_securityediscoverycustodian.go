package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityDataSourceContainer = SecurityEdiscoveryCustodian{}

type SecurityEdiscoveryCustodian struct {
	// Date and time the custodian acknowledged a hold notification.
	AcknowledgedDateTime nullable.Type[string] `json:"acknowledgedDateTime,omitempty"`

	// Email address of the custodian.
	Email nullable.Type[string] `json:"email,omitempty"`

	// Operation entity that represents the latest indexing for the custodian.
	LastIndexOperation *SecurityEdiscoveryIndexOperation `json:"lastIndexOperation,omitempty"`

	// Data source entity for SharePoint sites associated with the custodian.
	SiteSources *[]SecuritySiteSource `json:"siteSources,omitempty"`

	// Data source entity for groups associated with the custodian.
	UnifiedGroupSources *[]SecurityUnifiedGroupSource `json:"unifiedGroupSources,omitempty"`

	// Data source entity for a custodian. This is the container for a custodian's mailbox and OneDrive for Business site.
	UserSources *[]SecurityUserSource `json:"userSources,omitempty"`

	// Fields inherited from SecurityDataSourceContainer

	// Created date and time of the dataSourceContainer entity.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Display name of the dataSourceContainer entity.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The hold status of the dataSourceContainer. The possible values are: notApplied, applied, applying, removing, partial
	HoldStatus *SecurityDataSourceHoldStatus `json:"holdStatus,omitempty"`

	// Last modified date and time of the dataSourceContainer.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Date and time that the dataSourceContainer was released from the case.
	ReleasedDateTime nullable.Type[string] `json:"releasedDateTime,omitempty"`

	// Latest status of the dataSourceContainer. Possible values are: Active, Released.
	Status *SecurityDataSourceContainerStatus `json:"status,omitempty"`

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

func (s SecurityEdiscoveryCustodian) SecurityDataSourceContainer() BaseSecurityDataSourceContainerImpl {
	return BaseSecurityDataSourceContainerImpl{
		CreatedDateTime:      s.CreatedDateTime,
		DisplayName:          s.DisplayName,
		HoldStatus:           s.HoldStatus,
		LastModifiedDateTime: s.LastModifiedDateTime,
		ReleasedDateTime:     s.ReleasedDateTime,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SecurityEdiscoveryCustodian) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryCustodian{}

func (s SecurityEdiscoveryCustodian) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryCustodian
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryCustodian: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryCustodian: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryCustodian"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryCustodian: %+v", err)
	}

	return encoded, nil
}
