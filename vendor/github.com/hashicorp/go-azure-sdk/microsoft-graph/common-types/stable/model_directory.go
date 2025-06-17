package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Directory{}

type Directory struct {
	// Conceptual container for user and group directory objects.
	AdministrativeUnits *[]AdministrativeUnit `json:"administrativeUnits,omitempty"`

	// Group of related custom security attribute definitions.
	AttributeSets *[]AttributeSet `json:"attributeSets,omitempty"`

	// Schema of a custom security attributes (key-value pairs).
	CustomSecurityAttributeDefinitions *[]CustomSecurityAttributeDefinition `json:"customSecurityAttributeDefinitions,omitempty"`

	// Recently deleted items. Read-only. Nullable.
	DeletedItems *[]DirectoryObject `json:"deletedItems,omitempty"`

	// List of OData IDs for `DeletedItems` to bind to this entity
	DeletedItems_ODataBind *[]string `json:"deletedItems@odata.bind,omitempty"`

	// The credentials of the device's local administrator account backed up to Microsoft Entra ID.
	DeviceLocalCredentials *[]DeviceLocalCredentialInfo `json:"deviceLocalCredentials,omitempty"`

	// Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed
	// protocol.
	FederationConfigurations *[]IdentityProviderBase `json:"federationConfigurations,omitempty"`

	// A container for on-premises directory synchronization functionalities that are available for the organization.
	OnPremisesSynchronization *[]OnPremisesDirectorySynchronization `json:"onPremisesSynchronization,omitempty"`

	PublicKeyInfrastructure *PublicKeyInfrastructureRoot `json:"publicKeyInfrastructure,omitempty"`

	// List of commercial subscriptions that an organization acquired.
	Subscriptions *[]CompanySubscription `json:"subscriptions,omitempty"`

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

func (s Directory) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Directory{}

func (s Directory) MarshalJSON() ([]byte, error) {
	type wrapper Directory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Directory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Directory: %+v", err)
	}

	delete(decoded, "deletedItems")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.directory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Directory: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Directory{}

func (s *Directory) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AdministrativeUnits                *[]AdministrativeUnit                 `json:"administrativeUnits,omitempty"`
		AttributeSets                      *[]AttributeSet                       `json:"attributeSets,omitempty"`
		CustomSecurityAttributeDefinitions *[]CustomSecurityAttributeDefinition  `json:"customSecurityAttributeDefinitions,omitempty"`
		DeletedItems_ODataBind             *[]string                             `json:"deletedItems@odata.bind,omitempty"`
		DeviceLocalCredentials             *[]DeviceLocalCredentialInfo          `json:"deviceLocalCredentials,omitempty"`
		OnPremisesSynchronization          *[]OnPremisesDirectorySynchronization `json:"onPremisesSynchronization,omitempty"`
		PublicKeyInfrastructure            *PublicKeyInfrastructureRoot          `json:"publicKeyInfrastructure,omitempty"`
		Subscriptions                      *[]CompanySubscription                `json:"subscriptions,omitempty"`
		Id                                 *string                               `json:"id,omitempty"`
		ODataId                            *string                               `json:"@odata.id,omitempty"`
		ODataType                          *string                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AdministrativeUnits = decoded.AdministrativeUnits
	s.AttributeSets = decoded.AttributeSets
	s.CustomSecurityAttributeDefinitions = decoded.CustomSecurityAttributeDefinitions
	s.DeletedItems_ODataBind = decoded.DeletedItems_ODataBind
	s.DeviceLocalCredentials = decoded.DeviceLocalCredentials
	s.OnPremisesSynchronization = decoded.OnPremisesSynchronization
	s.PublicKeyInfrastructure = decoded.PublicKeyInfrastructure
	s.Subscriptions = decoded.Subscriptions
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Directory into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["deletedItems"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DeletedItems into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DeletedItems' for 'Directory': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DeletedItems = &output
	}

	if v, ok := temp["federationConfigurations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling FederationConfigurations into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentityProviderBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentityProviderBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'FederationConfigurations' for 'Directory': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.FederationConfigurations = &output
	}

	return nil
}
