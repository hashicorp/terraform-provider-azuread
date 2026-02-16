package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = DeviceTemplate{}

type DeviceTemplate struct {
	// A tenant-defined name for the party that's responsible for provisioning and managing devices on the Microsoft Entra
	// tenant. For example, Tailwind Traders (the manufacturer) makes security cameras that are installed in customer
	// buildings and managed by Lakeshore Retail (the device authority). This value is provided to the customer by the
	// device authority (manufacturer or reseller).
	DeviceAuthority nullable.Type[string] `json:"deviceAuthority,omitempty"`

	// Collection of device objects created based on this template.
	DeviceInstances *[]Device `json:"deviceInstances,omitempty"`

	// Manufacturer name.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// Model name.
	Model nullable.Type[string] `json:"model,omitempty"`

	// Object ID of the mutualTlsOauthConfiguration. This value isn't required if self-signed certificates are used. This
	// value is provided to the customer by the device authority (manufacturer or reseller).
	MutualTlsOauthConfigurationId nullable.Type[string] `json:"mutualTlsOauthConfigurationId,omitempty"`

	// ID (tenant ID for device authority) of the tenant that contains the mutualTlsOauthConfiguration. This value isn't
	// required if self-signed certificates are used. This value is provided to the customer by the device authority
	// (manufacturer or reseller).
	MutualTlsOauthConfigurationTenantId nullable.Type[string] `json:"mutualTlsOauthConfigurationTenantId,omitempty"`

	// Operating system type. Supports $filter (eq, in).
	OperatingSystem nullable.Type[string] `json:"operatingSystem,omitempty"`

	// Collection of directory objects that can manage the device template and the related deviceInstances. Owners can be
	// represented as service principals, users, or applications. An owner has full privileges over the device template and
	// doesn't require other administrator roles to create, update, or delete devices from this template, as well as to add
	// or remove template owners. There can be a maximum of 100 owners on a device template. Supports $expand.
	Owners *[]DirectoryObject `json:"owners,omitempty"`

	// List of OData IDs for `Owners` to bind to this entity
	Owners_ODataBind *[]string `json:"owners@odata.bind,omitempty"`

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

func (s DeviceTemplate) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s DeviceTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceTemplate{}

func (s DeviceTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceTemplate{}

func (s *DeviceTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DeviceAuthority                     nullable.Type[string] `json:"deviceAuthority,omitempty"`
		DeviceInstances                     *[]Device             `json:"deviceInstances,omitempty"`
		Manufacturer                        nullable.Type[string] `json:"manufacturer,omitempty"`
		Model                               nullable.Type[string] `json:"model,omitempty"`
		MutualTlsOauthConfigurationId       nullable.Type[string] `json:"mutualTlsOauthConfigurationId,omitempty"`
		MutualTlsOauthConfigurationTenantId nullable.Type[string] `json:"mutualTlsOauthConfigurationTenantId,omitempty"`
		OperatingSystem                     nullable.Type[string] `json:"operatingSystem,omitempty"`
		Owners_ODataBind                    *[]string             `json:"owners@odata.bind,omitempty"`
		DeletedDateTime                     nullable.Type[string] `json:"deletedDateTime,omitempty"`
		Id                                  *string               `json:"id,omitempty"`
		ODataId                             *string               `json:"@odata.id,omitempty"`
		ODataType                           *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DeviceAuthority = decoded.DeviceAuthority
	s.DeviceInstances = decoded.DeviceInstances
	s.Manufacturer = decoded.Manufacturer
	s.Model = decoded.Model
	s.MutualTlsOauthConfigurationId = decoded.MutualTlsOauthConfigurationId
	s.MutualTlsOauthConfigurationTenantId = decoded.MutualTlsOauthConfigurationTenantId
	s.OperatingSystem = decoded.OperatingSystem
	s.Owners_ODataBind = decoded.Owners_ODataBind
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["owners"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Owners into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Owners' for 'DeviceTemplate': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Owners = &output
	}

	return nil
}
