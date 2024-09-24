package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Agreement{}

type Agreement struct {
	// Read-only. Information about acceptances of this agreement.
	Acceptances *[]AgreementAcceptance `json:"acceptances,omitempty"`

	// Display name of the agreement. The display name is used for internal tracking of the agreement but isn't shown to end
	// users who view the agreement. Supports $filter (eq).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Default PDF linked to this agreement.
	File *AgreementFile `json:"file,omitempty"`

	// PDFs linked to this agreement. This property is in the process of being deprecated. Use the file property instead.
	// Supports $expand.
	Files *[]AgreementFileLocalization `json:"files,omitempty"`

	// Indicates whether end users are required to accept this agreement on every device that they access it from. The end
	// user is required to register their device in Microsoft Entra ID, if they haven't already done so. Supports $filter
	// (eq).
	IsPerDeviceAcceptanceRequired nullable.Type[bool] `json:"isPerDeviceAcceptanceRequired,omitempty"`

	// Indicates whether the user has to expand the agreement before accepting. Supports $filter (eq).
	IsViewingBeforeAcceptanceRequired nullable.Type[bool] `json:"isViewingBeforeAcceptanceRequired,omitempty"`

	// Expiration schedule and frequency of agreement for all users. Supports $filter (eq).
	TermsExpiration *TermsExpiration `json:"termsExpiration,omitempty"`

	// The duration after which the user must reaccept the terms of use. The value is represented in ISO 8601 format for
	// durations. Supports $filter (eq).
	UserReacceptRequiredFrequency nullable.Type[string] `json:"userReacceptRequiredFrequency,omitempty"`

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

func (s Agreement) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Agreement{}

func (s Agreement) MarshalJSON() ([]byte, error) {
	type wrapper Agreement
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Agreement: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Agreement: %+v", err)
	}

	delete(decoded, "acceptances")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.agreement"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Agreement: %+v", err)
	}

	return encoded, nil
}
