package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Identity interface {
	Identity() BaseIdentityImpl
}

var _ Identity = BaseIdentityImpl{}

type BaseIdentityImpl struct {
	// The display name of the identity. This property is read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The identifier of the identity. This property is read-only.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIdentityImpl) Identity() BaseIdentityImpl {
	return s
}

var _ Identity = RawIdentityImpl{}

// RawIdentityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIdentityImpl struct {
	identity BaseIdentityImpl
	Type     string
	Values   map[string]interface{}
}

func (s RawIdentityImpl) Identity() BaseIdentityImpl {
	return s.identity
}

func UnmarshalIdentityImplementation(input []byte) (Identity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Identity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.azureCommunicationServicesUserIdentity") {
		var out AzureCommunicationServicesUserIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureCommunicationServicesUserIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.userIdentity") {
		var out CallRecordsUserIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsUserIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.communicationsApplicationIdentity") {
		var out CommunicationsApplicationIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommunicationsApplicationIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.communicationsApplicationInstanceIdentity") {
		var out CommunicationsApplicationInstanceIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommunicationsApplicationInstanceIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.communicationsEncryptedIdentity") {
		var out CommunicationsEncryptedIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommunicationsEncryptedIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.communicationsGuestIdentity") {
		var out CommunicationsGuestIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommunicationsGuestIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.communicationsPhoneIdentity") {
		var out CommunicationsPhoneIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommunicationsPhoneIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.communicationsUserIdentity") {
		var out CommunicationsUserIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommunicationsUserIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.emailIdentity") {
		var out EmailIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.initiator") {
		var out Initiator
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Initiator: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.programResource") {
		var out ProgramResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProgramResource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.provisionedIdentity") {
		var out ProvisionedIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProvisionedIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.provisioningServicePrincipal") {
		var out ProvisioningServicePrincipal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProvisioningServicePrincipal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.provisioningSystem") {
		var out ProvisioningSystem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProvisioningSystem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.submissionUserIdentity") {
		var out SecuritySubmissionUserIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySubmissionUserIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.servicePrincipalIdentity") {
		var out ServicePrincipalIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePrincipalIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharePointIdentity") {
		var out SharePointIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharePointIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sourceProvisionedIdentity") {
		var out SourceProvisionedIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SourceProvisionedIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetProvisionedIdentity") {
		var out TargetProvisionedIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetProvisionedIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkApplicationIdentity") {
		var out TeamworkApplicationIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkApplicationIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkConversationIdentity") {
		var out TeamworkConversationIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkConversationIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkTagIdentity") {
		var out TeamworkTagIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkTagIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamworkUserIdentity") {
		var out TeamworkUserIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamworkUserIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userIdentity") {
		var out UserIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserIdentity: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentityImpl: %+v", err)
	}

	return RawIdentityImpl{
		identity: parent,
		Type:     value,
		Values:   temp,
	}, nil

}
