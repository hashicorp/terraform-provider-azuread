package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityEmailThreatSubmission = SecurityEmailContentThreatSubmission{}

type SecurityEmailContentThreatSubmission struct {
	// Base64 encoded file content.
	FileContent *string `json:"fileContent,omitempty"`

	// Fields inherited from SecurityEmailThreatSubmission

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

func (s SecurityEmailContentThreatSubmission) SecurityEmailThreatSubmission() BaseSecurityEmailThreatSubmissionImpl {
	return BaseSecurityEmailThreatSubmissionImpl{
		AttackSimulationInfo:         s.AttackSimulationInfo,
		InternetMessageId:            s.InternetMessageId,
		OriginalCategory:             s.OriginalCategory,
		ReceivedDateTime:             s.ReceivedDateTime,
		RecipientEmailAddress:        s.RecipientEmailAddress,
		Sender:                       s.Sender,
		SenderIP:                     s.SenderIP,
		Subject:                      s.Subject,
		TenantAllowOrBlockListAction: s.TenantAllowOrBlockListAction,
		AdminReview:                  s.AdminReview,
		Category:                     s.Category,
		ClientSource:                 s.ClientSource,
		ContentType:                  s.ContentType,
		CreatedBy:                    s.CreatedBy,
		CreatedDateTime:              s.CreatedDateTime,
		Result:                       s.Result,
		Source:                       s.Source,
		Status:                       s.Status,
		TenantId:                     s.TenantId,
		Id:                           s.Id,
		ODataId:                      s.ODataId,
		ODataType:                    s.ODataType,
	}
}

func (s SecurityEmailContentThreatSubmission) SecurityThreatSubmission() BaseSecurityThreatSubmissionImpl {
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

func (s SecurityEmailContentThreatSubmission) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEmailContentThreatSubmission{}

func (s SecurityEmailContentThreatSubmission) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEmailContentThreatSubmission
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEmailContentThreatSubmission: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEmailContentThreatSubmission: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.emailContentThreatSubmission"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEmailContentThreatSubmission: %+v", err)
	}

	return encoded, nil
}
