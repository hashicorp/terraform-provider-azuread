package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryObject interface {
	Entity
	DirectoryObject() BaseDirectoryObjectImpl
}

var _ DirectoryObject = BaseDirectoryObjectImpl{}

type BaseDirectoryObjectImpl struct {
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

func (s BaseDirectoryObjectImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return s
}

func (s BaseDirectoryObjectImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DirectoryObject = RawDirectoryObjectImpl{}

// RawDirectoryObjectImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDirectoryObjectImpl struct {
	directoryObject BaseDirectoryObjectImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawDirectoryObjectImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return s.directoryObject
}

func (s RawDirectoryObjectImpl) Entity() BaseEntityImpl {
	return s.directoryObject.Entity()
}

var _ json.Marshaler = BaseDirectoryObjectImpl{}

func (s BaseDirectoryObjectImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDirectoryObjectImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDirectoryObjectImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDirectoryObjectImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.directoryObject"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDirectoryObjectImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDirectoryObjectImplementation(input []byte) (DirectoryObject, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DirectoryObject into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.administrativeUnit") {
		var out AdministrativeUnit
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AdministrativeUnit: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appRoleAssignment") {
		var out AppRoleAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppRoleAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.application") {
		var out Application
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Application: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.certificateAuthorityDetail") {
		var out CertificateAuthorityDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificateAuthorityDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.certificateBasedAuthPki") {
		var out CertificateBasedAuthPki
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificateBasedAuthPki: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.contract") {
		var out Contract
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Contract: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.device") {
		var out Device
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Device: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceTemplate") {
		var out DeviceTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryObjectPartnerReference") {
		var out DirectoryObjectPartnerReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryObjectPartnerReference: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryRole") {
		var out DirectoryRole
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryRole: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directoryRoleTemplate") {
		var out DirectoryRoleTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectoryRoleTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.directorySettingTemplate") {
		var out DirectorySettingTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DirectorySettingTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.endpoint") {
		var out Endpoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Endpoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.extensionProperty") {
		var out ExtensionProperty
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExtensionProperty: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalProfile") {
		var out ExternalProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.federatedTokenValidationPolicy") {
		var out FederatedTokenValidationPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FederatedTokenValidationPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.group") {
		var out Group
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Group: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailbox") {
		var out Mailbox
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Mailbox: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.multiTenantOrganizationMember") {
		var out MultiTenantOrganizationMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MultiTenantOrganizationMember: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.orgContact") {
		var out OrgContact
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OrgContact: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.organization") {
		var out Organization
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Organization: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.organizationalUnit") {
		var out OrganizationalUnit
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OrganizationalUnit: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permissionGrantPreApprovalPolicy") {
		var out PermissionGrantPreApprovalPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionGrantPreApprovalPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.policyBase") {
		var out PolicyBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicyBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.resourceSpecificPermissionGrant") {
		var out ResourceSpecificPermissionGrant
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourceSpecificPermissionGrant: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.servicePrincipal") {
		var out ServicePrincipal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePrincipal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trustedCertificateAuthorityAsEntityBase") {
		var out TrustedCertificateAuthorityAsEntityBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrustedCertificateAuthorityAsEntityBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trustedCertificateAuthorityBase") {
		var out TrustedCertificateAuthorityBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrustedCertificateAuthorityBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.user") {
		var out User
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into User: %+v", err)
		}
		return out, nil
	}

	var parent BaseDirectoryObjectImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDirectoryObjectImpl: %+v", err)
	}

	return RawDirectoryObjectImpl{
		directoryObject: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
