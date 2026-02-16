package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EdiscoveryDataSourceContainer = EdiscoveryCustodian{}

type EdiscoveryCustodian struct {
	// Date and time the custodian acknowledged a hold notification.
	AcknowledgedDateTime nullable.Type[string] `json:"acknowledgedDateTime,omitempty"`

	// Identifies whether a custodian's sources were placed on hold during creation.
	ApplyHoldToSources nullable.Type[bool] `json:"applyHoldToSources,omitempty"`

	// Email address of the custodian.
	Email *string `json:"email,omitempty"`

	// Data source entity for SharePoint sites associated with the custodian.
	SiteSources *[]EdiscoverySiteSource `json:"siteSources,omitempty"`

	// Data source entity for groups associated with the custodian.
	UnifiedGroupSources *[]EdiscoveryUnifiedGroupSource `json:"unifiedGroupSources,omitempty"`

	// Data source entity for a the custodian. This is the container for a custodian's mailbox and OneDrive for Business
	// site.
	UserSources *[]EdiscoveryUserSource `json:"userSources,omitempty"`

	// Fields inherited from EdiscoveryDataSourceContainer

	// Created date and time of the dataSourceContainer entity.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Display name of the dataSourceContainer entity.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	HoldStatus         *EdiscoveryDataSourceHoldStatus `json:"holdStatus,omitempty"`
	LastIndexOperation *EdiscoveryCaseIndexOperation   `json:"lastIndexOperation,omitempty"`

	// Last modified date and time of the dataSourceContainer.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Date and time that the dataSourceContainer was released from the case.
	ReleasedDateTime nullable.Type[string] `json:"releasedDateTime,omitempty"`

	// Latest status of the dataSourceContainer. Possible values are: Active, Released.
	Status *EdiscoveryDataSourceContainerStatus `json:"status,omitempty"`

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

func (s EdiscoveryCustodian) EdiscoveryDataSourceContainer() BaseEdiscoveryDataSourceContainerImpl {
	return BaseEdiscoveryDataSourceContainerImpl{
		CreatedDateTime:      s.CreatedDateTime,
		DisplayName:          s.DisplayName,
		HoldStatus:           s.HoldStatus,
		LastIndexOperation:   s.LastIndexOperation,
		LastModifiedDateTime: s.LastModifiedDateTime,
		ReleasedDateTime:     s.ReleasedDateTime,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s EdiscoveryCustodian) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdiscoveryCustodian{}

func (s EdiscoveryCustodian) MarshalJSON() ([]byte, error) {
	type wrapper EdiscoveryCustodian
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdiscoveryCustodian: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryCustodian: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.custodian"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdiscoveryCustodian: %+v", err)
	}

	return encoded, nil
}
