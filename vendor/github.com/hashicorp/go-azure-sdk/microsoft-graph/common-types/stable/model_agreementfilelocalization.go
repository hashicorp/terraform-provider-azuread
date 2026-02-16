package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AgreementFileProperties = AgreementFileLocalization{}

type AgreementFileLocalization struct {
	// Read-only. Customized versions of the terms of use agreement in the Microsoft Entra tenant.
	Versions *[]AgreementFileVersion `json:"versions,omitempty"`

	// Fields inherited from AgreementFileProperties

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

func (s AgreementFileLocalization) AgreementFileProperties() BaseAgreementFilePropertiesImpl {
	return BaseAgreementFilePropertiesImpl{
		CreatedDateTime: s.CreatedDateTime,
		DisplayName:     s.DisplayName,
		FileData:        s.FileData,
		FileName:        s.FileName,
		IsDefault:       s.IsDefault,
		IsMajorVersion:  s.IsMajorVersion,
		Language:        s.Language,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s AgreementFileLocalization) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AgreementFileLocalization{}

func (s AgreementFileLocalization) MarshalJSON() ([]byte, error) {
	type wrapper AgreementFileLocalization
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AgreementFileLocalization: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AgreementFileLocalization: %+v", err)
	}

	delete(decoded, "versions")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.agreementFileLocalization"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AgreementFileLocalization: %+v", err)
	}

	return encoded, nil
}
