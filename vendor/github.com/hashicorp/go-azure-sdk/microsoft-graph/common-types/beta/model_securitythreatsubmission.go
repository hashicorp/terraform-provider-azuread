package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmission interface {
	Entity
	SecurityThreatSubmission() BaseSecurityThreatSubmissionImpl
}

var _ SecurityThreatSubmission = BaseSecurityThreatSubmissionImpl{}

type BaseSecurityThreatSubmissionImpl struct {
	// Specifies the admin review property that constitutes of who reviewed the user submission, when and what was it
	// identified as.
	AdminReview *SecuritySubmissionAdminReview `json:"adminReview,omitempty"`

	Category *SecuritySubmissionCategory `json:"category,omitempty"`

	// Specifies the source of the submission. The possible values are: microsoft, other, and unkownFutureValue.
	ClientSource *SecuritySubmissionClientSource `json:"clientSource,omitempty"`

	// Specifies the type of content being submitted. The possible values are: email, url, file, app, and unkownFutureValue.
	ContentType *SecuritySubmissionContentType `json:"contentType,omitempty"`

	// Specifies who submitted the email as a threat. Supports $filter = createdBy/email eq 'value'.
	CreatedBy *SecuritySubmissionUserIdentity `json:"createdBy,omitempty"`

	// Specifies when the threat submission was created. Supports $filter = createdDateTime ge 2022-01-01T00:00:00Z and
	// createdDateTime lt 2022-01-02T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Specifies the result of the analysis performed by Microsoft.
	Result *SecuritySubmissionResult `json:"result,omitempty"`

	// Specifies the role of the submitter. Supports $filter = source eq 'value'. The possible values are: administrator,
	// user, and unkownFutureValue.
	Source *SecuritySubmissionSource `json:"source,omitempty"`

	// Indicates whether the threat submission has been analyzed by Microsoft. Supports $filter = status eq 'value'. The
	// possible values are: notStarted, running, succeeded, failed, skipped, and unkownFutureValue.
	Status *SecurityLongRunningOperationStatus `json:"status,omitempty"`

	// Indicates the tenant id of the submitter. Not required when created using a POST operation. It's extracted from the
	// token of the post API call.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

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

func (s BaseSecurityThreatSubmissionImpl) SecurityThreatSubmission() BaseSecurityThreatSubmissionImpl {
	return s
}

func (s BaseSecurityThreatSubmissionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityThreatSubmission = RawSecurityThreatSubmissionImpl{}

// RawSecurityThreatSubmissionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityThreatSubmissionImpl struct {
	securityThreatSubmission BaseSecurityThreatSubmissionImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawSecurityThreatSubmissionImpl) SecurityThreatSubmission() BaseSecurityThreatSubmissionImpl {
	return s.securityThreatSubmission
}

func (s RawSecurityThreatSubmissionImpl) Entity() BaseEntityImpl {
	return s.securityThreatSubmission.Entity()
}

var _ json.Marshaler = BaseSecurityThreatSubmissionImpl{}

func (s BaseSecurityThreatSubmissionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityThreatSubmissionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityThreatSubmissionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityThreatSubmissionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.threatSubmission"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityThreatSubmissionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSecurityThreatSubmissionImplementation(input []byte) (SecurityThreatSubmission, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityThreatSubmission into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.emailThreatSubmission") {
		var out SecurityEmailThreatSubmission
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEmailThreatSubmission: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.fileThreatSubmission") {
		var out SecurityFileThreatSubmission
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFileThreatSubmission: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.urlThreatSubmission") {
		var out SecurityUrlThreatSubmission
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUrlThreatSubmission: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityThreatSubmissionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityThreatSubmissionImpl: %+v", err)
	}

	return RawSecurityThreatSubmissionImpl{
		securityThreatSubmission: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
