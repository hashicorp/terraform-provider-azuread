package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmission interface {
	Entity
	SecurityThreatSubmission
	SecurityEmailThreatSubmission() BaseSecurityEmailThreatSubmissionImpl
}

var _ SecurityEmailThreatSubmission = BaseSecurityEmailThreatSubmissionImpl{}

type BaseSecurityEmailThreatSubmissionImpl struct {
	// If the email is phishing simulation, this field won't be null.
	AttackSimulationInfo *SecurityAttackSimulationInfo `json:"attackSimulationInfo,omitempty"`

	// Specifies the internet message ID of the email being submitted. This information is present in the email header.
	InternetMessageId nullable.Type[string] `json:"internetMessageId,omitempty"`

	// The original category of the submission. The possible values are: notJunk, spam, phishing, malware and
	// unkownFutureValue.
	OriginalCategory *SecuritySubmissionCategory `json:"originalCategory,omitempty"`

	// Specifies the date and time stamp when the email was received.
	ReceivedDateTime nullable.Type[string] `json:"receivedDateTime,omitempty"`

	// Specifies the email address (in smtp format) of the recipient who received the email.
	RecipientEmailAddress *string `json:"recipientEmailAddress,omitempty"`

	// Specifies the email address of the sender.
	Sender nullable.Type[string] `json:"sender,omitempty"`

	// Specifies the IP address of the sender.
	SenderIP nullable.Type[string] `json:"senderIP,omitempty"`

	// Specifies the subject of the email.
	Subject nullable.Type[string] `json:"subject,omitempty"`

	// It's used to automatically add allows for the components such as URL, file, sender; which are deemed bad by Microsoft
	// so that similar messages in the future can be allowed.
	TenantAllowOrBlockListAction *SecurityTenantAllowOrBlockListAction `json:"tenantAllowOrBlockListAction,omitempty"`

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

func (s BaseSecurityEmailThreatSubmissionImpl) SecurityEmailThreatSubmission() BaseSecurityEmailThreatSubmissionImpl {
	return s
}

func (s BaseSecurityEmailThreatSubmissionImpl) SecurityThreatSubmission() BaseSecurityThreatSubmissionImpl {
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

func (s BaseSecurityEmailThreatSubmissionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityEmailThreatSubmission = RawSecurityEmailThreatSubmissionImpl{}

// RawSecurityEmailThreatSubmissionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityEmailThreatSubmissionImpl struct {
	securityEmailThreatSubmission BaseSecurityEmailThreatSubmissionImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawSecurityEmailThreatSubmissionImpl) SecurityEmailThreatSubmission() BaseSecurityEmailThreatSubmissionImpl {
	return s.securityEmailThreatSubmission
}

func (s RawSecurityEmailThreatSubmissionImpl) SecurityThreatSubmission() BaseSecurityThreatSubmissionImpl {
	return s.securityEmailThreatSubmission.SecurityThreatSubmission()
}

func (s RawSecurityEmailThreatSubmissionImpl) Entity() BaseEntityImpl {
	return s.securityEmailThreatSubmission.Entity()
}

var _ json.Marshaler = BaseSecurityEmailThreatSubmissionImpl{}

func (s BaseSecurityEmailThreatSubmissionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityEmailThreatSubmissionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityEmailThreatSubmissionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityEmailThreatSubmissionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.emailThreatSubmission"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityEmailThreatSubmissionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSecurityEmailThreatSubmissionImplementation(input []byte) (SecurityEmailThreatSubmission, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEmailThreatSubmission into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.emailContentThreatSubmission") {
		var out SecurityEmailContentThreatSubmission
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEmailContentThreatSubmission: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.emailUrlThreatSubmission") {
		var out SecurityEmailUrlThreatSubmission
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEmailUrlThreatSubmission: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityEmailThreatSubmissionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityEmailThreatSubmissionImpl: %+v", err)
	}

	return RawSecurityEmailThreatSubmissionImpl{
		securityEmailThreatSubmission: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
