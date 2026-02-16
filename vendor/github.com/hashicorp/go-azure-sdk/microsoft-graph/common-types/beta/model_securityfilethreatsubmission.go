package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileThreatSubmission interface {
	Entity
	SecurityThreatSubmission
	SecurityFileThreatSubmission() BaseSecurityFileThreatSubmissionImpl
}

var _ SecurityFileThreatSubmission = BaseSecurityFileThreatSubmissionImpl{}

type BaseSecurityFileThreatSubmissionImpl struct {
	// It specifies the file name to be submitted.
	FileName *string `json:"fileName,omitempty"`

	// Fields inherited from SecurityThreatSubmission

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

func (s BaseSecurityFileThreatSubmissionImpl) SecurityFileThreatSubmission() BaseSecurityFileThreatSubmissionImpl {
	return s
}

func (s BaseSecurityFileThreatSubmissionImpl) SecurityThreatSubmission() BaseSecurityThreatSubmissionImpl {
	return BaseSecurityThreatSubmissionImpl{
		AdminReview:     s.AdminReview,
		Category:        s.Category,
		ClientSource:    s.ClientSource,
		ContentType:     s.ContentType,
		CreatedBy:       s.CreatedBy,
		CreatedDateTime: s.CreatedDateTime,
		Result:          s.Result,
		Source:          s.Source,
		Status:          s.Status,
		TenantId:        s.TenantId,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseSecurityFileThreatSubmissionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityFileThreatSubmission = RawSecurityFileThreatSubmissionImpl{}

// RawSecurityFileThreatSubmissionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityFileThreatSubmissionImpl struct {
	securityFileThreatSubmission BaseSecurityFileThreatSubmissionImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawSecurityFileThreatSubmissionImpl) SecurityFileThreatSubmission() BaseSecurityFileThreatSubmissionImpl {
	return s.securityFileThreatSubmission
}

func (s RawSecurityFileThreatSubmissionImpl) SecurityThreatSubmission() BaseSecurityThreatSubmissionImpl {
	return s.securityFileThreatSubmission.SecurityThreatSubmission()
}

func (s RawSecurityFileThreatSubmissionImpl) Entity() BaseEntityImpl {
	return s.securityFileThreatSubmission.Entity()
}

var _ json.Marshaler = BaseSecurityFileThreatSubmissionImpl{}

func (s BaseSecurityFileThreatSubmissionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityFileThreatSubmissionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityFileThreatSubmissionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityFileThreatSubmissionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.fileThreatSubmission"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityFileThreatSubmissionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSecurityFileThreatSubmissionImplementation(input []byte) (SecurityFileThreatSubmission, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityFileThreatSubmission into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.fileContentThreatSubmission") {
		var out SecurityFileContentThreatSubmission
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFileContentThreatSubmission: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.fileUrlThreatSubmission") {
		var out SecurityFileUrlThreatSubmission
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFileUrlThreatSubmission: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityFileThreatSubmissionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityFileThreatSubmissionImpl: %+v", err)
	}

	return RawSecurityFileThreatSubmissionImpl{
		securityFileThreatSubmission: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
