package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityThreatSubmission = SecurityUrlThreatSubmission{}

type SecurityUrlThreatSubmission struct {
	// Denotes the webUrl that needs to be submitted.
	WebUrl *string `json:"webUrl,omitempty"`

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

func (s SecurityUrlThreatSubmission) SecurityThreatSubmission() BaseSecurityThreatSubmissionImpl {
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

func (s SecurityUrlThreatSubmission) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityUrlThreatSubmission{}

func (s SecurityUrlThreatSubmission) MarshalJSON() ([]byte, error) {
	type wrapper SecurityUrlThreatSubmission
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityUrlThreatSubmission: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityUrlThreatSubmission: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.urlThreatSubmission"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityUrlThreatSubmission: %+v", err)
	}

	return encoded, nil
}
