package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PartnerSecurityPartnerSecurityAlert{}

type PartnerSecurityPartnerSecurityAlert struct {
	// Represents the activity by a partner and includes details of state transitions, who performed them, and when they
	// occurred.
	ActivityLogs *[]PartnerSecurityActivityLog `json:"activityLogs,omitempty"`

	// A bag of name-value pairs that contain more details about an alert.
	AdditionalDetails *PartnerSecurityAdditionalDataDictionary `json:"additionalDetails,omitempty"`

	// Contains details of the resources affected by the security alert.
	AffectedResources *[]PartnerSecurityAffectedResource `json:"affectedResources,omitempty"`

	// The type of vulnerability that impacts the customer due to this alert. For more information, see Security alerts
	// reference guide.
	AlertType *string `json:"alertType,omitempty"`

	// The modern offer category ID of the subscription.
	CatalogOfferId nullable.Type[string] `json:"catalogOfferId,omitempty"`

	ConfidenceLevel *PartnerSecuritySecurityAlertConfidence `json:"confidenceLevel,omitempty"`

	// The impacted customer tenant associated with the alert.
	CustomerTenantId *string `json:"customerTenantId,omitempty"`

	// The description for each alert.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Time when the alert was detected or created. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	DetectedDateTime *string `json:"detectedDateTime,omitempty"`

	// The display name of the alert.
	DisplayName *string `json:"displayName,omitempty"`

	// Time of the first activity associated with the alert. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	FirstObservedDateTime nullable.Type[string] `json:"firstObservedDateTime,omitempty"`

	// Indicates whether an alert is a test alert.
	IsTest nullable.Type[bool] `json:"isTest,omitempty"`

	// Time of the latest activity associated with the alert. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastObservedDateTime nullable.Type[string] `json:"lastObservedDateTime,omitempty"`

	// The UPN of the partner user who resolved the alert.
	ResolvedBy nullable.Type[string] `json:"resolvedBy,omitempty"`

	// Time when the alert was resolved. The timestamp type represents date and time information using ISO 8601 format and
	// is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ResolvedOnDateTime nullable.Type[string] `json:"resolvedOnDateTime,omitempty"`

	// The reason provided by the partner for addressing the alert. The possible values are: legitimate, ignore, fraud,
	// unknownFutureValue.
	ResolvedReason *PartnerSecuritySecurityAlertResolvedReason `json:"resolvedReason,omitempty"`

	Severity *PartnerSecuritySecurityAlertSeverity `json:"severity,omitempty"`
	Status   *PartnerSecuritySecurityAlertStatus   `json:"status,omitempty"`

	// The subscription associated with the alert for the customer.
	SubscriptionId *string `json:"subscriptionId,omitempty"`

	// The value-added reseller tenant associated with the partner tenant and customer tenant.
	ValueAddedResellerTenantId nullable.Type[string] `json:"valueAddedResellerTenantId,omitempty"`

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

func (s PartnerSecurityPartnerSecurityAlert) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PartnerSecurityPartnerSecurityAlert{}

func (s PartnerSecurityPartnerSecurityAlert) MarshalJSON() ([]byte, error) {
	type wrapper PartnerSecurityPartnerSecurityAlert
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PartnerSecurityPartnerSecurityAlert: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnerSecurityPartnerSecurityAlert: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.partner.security.partnerSecurityAlert"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PartnerSecurityPartnerSecurityAlert: %+v", err)
	}

	return encoded, nil
}
