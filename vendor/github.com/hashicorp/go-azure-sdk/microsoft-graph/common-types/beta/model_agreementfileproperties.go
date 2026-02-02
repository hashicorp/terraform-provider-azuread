package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementFileProperties interface {
	Entity
	AgreementFileProperties() BaseAgreementFilePropertiesImpl
}

var _ AgreementFileProperties = BaseAgreementFilePropertiesImpl{}

type BaseAgreementFilePropertiesImpl struct {
	// The date time representing when the file was created. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Localized display name of the policy file of an agreement. The localized display name is shown to end users who view
	// the agreement.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Data that represents the terms of use PDF document. Read-only.
	FileData *AgreementFileData `json:"fileData,omitempty"`

	// Name of the agreement file (for example, TOU.pdf). Read-only.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// If none of the languages matches the client preference, indicates whether this is the default agreement file. If none
	// of the files are marked as default, the first one is treated as the default. Read-only.
	IsDefault nullable.Type[bool] `json:"isDefault,omitempty"`

	// Indicates whether the agreement file is a major version update. Major version updates invalidate the agreement's
	// acceptances on the corresponding language.
	IsMajorVersion nullable.Type[bool] `json:"isMajorVersion,omitempty"`

	// The language of the agreement file in the format 'languagecode2-country/regioncode2'. 'languagecode2' is a lowercase
	// two-letter code derived from ISO 639-1, while 'country/regioncode2' is derived from ISO 3166 and usually consists of
	// two uppercase letters, or a BCP-47 language tag. For example, U.S. English is en-US. Read-only.
	Language nullable.Type[string] `json:"language,omitempty"`

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

func (s BaseAgreementFilePropertiesImpl) AgreementFileProperties() BaseAgreementFilePropertiesImpl {
	return s
}

func (s BaseAgreementFilePropertiesImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AgreementFileProperties = RawAgreementFilePropertiesImpl{}

// RawAgreementFilePropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAgreementFilePropertiesImpl struct {
	agreementFileProperties BaseAgreementFilePropertiesImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawAgreementFilePropertiesImpl) AgreementFileProperties() BaseAgreementFilePropertiesImpl {
	return s.agreementFileProperties
}

func (s RawAgreementFilePropertiesImpl) Entity() BaseEntityImpl {
	return s.agreementFileProperties.Entity()
}

var _ json.Marshaler = BaseAgreementFilePropertiesImpl{}

func (s BaseAgreementFilePropertiesImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAgreementFilePropertiesImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAgreementFilePropertiesImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAgreementFilePropertiesImpl: %+v", err)
	}

	delete(decoded, "fileData")
	delete(decoded, "fileName")
	delete(decoded, "isDefault")
	delete(decoded, "language")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.agreementFileProperties"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAgreementFilePropertiesImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAgreementFilePropertiesImplementation(input []byte) (AgreementFileProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AgreementFileProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.agreementFile") {
		var out AgreementFile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AgreementFile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.agreementFileLocalization") {
		var out AgreementFileLocalization
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AgreementFileLocalization: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.agreementFileVersion") {
		var out AgreementFileVersion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AgreementFileVersion: %+v", err)
		}
		return out, nil
	}

	var parent BaseAgreementFilePropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAgreementFilePropertiesImpl: %+v", err)
	}

	return RawAgreementFilePropertiesImpl{
		agreementFileProperties: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
