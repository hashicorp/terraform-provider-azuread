package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsDeviceScope{}

type UserExperienceAnalyticsDeviceScope struct {
	// Indicates the creation date and time for the custom device scope.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The name of the user experience analytics device Scope configuration.
	DeviceScopeName nullable.Type[string] `json:"deviceScopeName,omitempty"`

	// Indicates whether a device scope is enabled or disabled. When TRUE, the device scope is enabled. When FALSE, the
	// device scope is disabled. Default value is FALSE.
	Enabled *bool `json:"enabled,omitempty"`

	// Indicates whether the device scope configuration is built-in or custom. When TRUE, the device scope configuration is
	// built-in. When FALSE, the device scope configuration is custom. Default value is FALSE.
	IsBuiltIn *bool `json:"isBuiltIn,omitempty"`

	// Indicates the last updated date and time for the custom device scope.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Device scope configuration query operator. Possible values are: equals, notEquals, contains, notContains,
	// greaterThan, lessThan. Default value: equals.
	Operator *DeviceScopeOperator `json:"operator,omitempty"`

	// The unique identifier of the person (admin) who created the device scope configuration.
	OwnerId nullable.Type[string] `json:"ownerId,omitempty"`

	// Device scope configuration parameter. It will be expend in future to add more parameter. Eg: device scope parameter
	// can be OS version, Disk Type, Device manufacturer, device model or Scope tag. Default value: scopeTag.
	Parameter *DeviceScopeParameter `json:"parameter,omitempty"`

	// Indicates the device scope status after the device scope has been enabled. Possible values are: none, computing,
	// insufficientData or completed. Default value is none.
	Status *DeviceScopeStatus `json:"status,omitempty"`

	// The device scope configuration query clause value.
	Value nullable.Type[string] `json:"value,omitempty"`

	// The unique identifier for a user device scope tag Id used for the creation of device scope configuration.
	ValueObjectId nullable.Type[string] `json:"valueObjectId,omitempty"`

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

func (s UserExperienceAnalyticsDeviceScope) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsDeviceScope{}

func (s UserExperienceAnalyticsDeviceScope) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsDeviceScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsDeviceScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsDeviceScope: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsDeviceScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsDeviceScope: %+v", err)
	}

	return encoded, nil
}
