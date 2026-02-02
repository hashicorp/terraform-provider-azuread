package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAlertEvidence = SecurityTeamsMessageEvidence{}

type SecurityTeamsMessageEvidence struct {
	// The identifier of the campaign that this Teams message is part of.
	CampaignId nullable.Type[string] `json:"campaignId,omitempty"`

	// The channel ID associated with this Teams message.
	ChannelId nullable.Type[string] `json:"channelId,omitempty"`

	// The delivery action of this Teams message. Possible values are: unknown, deliveredAsSpam, delivered, blocked,
	// replaced, unknownFutureValue.
	DeliveryAction *SecurityTeamsMessageDeliveryAction `json:"deliveryAction,omitempty"`

	// The delivery location of this Teams message. Possible values are: unknown, teams, quarantine, failed,
	// unknownFutureValue.
	DeliveryLocation *SecurityTeamsDeliveryLocation `json:"deliveryLocation,omitempty"`

	// The list of file entities that are attached to this Teams message.
	Files *[]SecurityFileEvidence `json:"files,omitempty"`

	// The identifier of the team or group that this message is part of.
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Indicates whether the message is owned by the organization that reported the security detection alert.
	IsExternal nullable.Type[bool] `json:"isExternal,omitempty"`

	// Indicates whether the message is owned by your organization.
	IsOwned nullable.Type[bool] `json:"isOwned,omitempty"`

	// Date and time when the message was last edited. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The direction of the Teams message. The possible values are: unknown, inbound, outbound, intraorg,
	// unknownFutureValue.
	MessageDirection *SecurityAntispamTeamsDirection `json:"messageDirection,omitempty"`

	// The message identifier, unique within the thread.
	MessageId nullable.Type[string] `json:"messageId,omitempty"`

	// Tenant ID (GUID) of the owner of the message.
	OwningTenantId nullable.Type[string] `json:"owningTenantId,omitempty"`

	// Identifier of the message to which the current message is a reply; otherwise, it's the same as the messageId.
	ParentMessageId nullable.Type[string] `json:"parentMessageId,omitempty"`

	// The received date of this message. The Timestamp type represents date and time information using ISO 8601 format and
	// is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ReceivedDateTime nullable.Type[string] `json:"receivedDateTime,omitempty"`

	// The recipients of this Teams message.
	Recipients *[]string `json:"recipients,omitempty"`

	// The SMTP format address of the sender.
	SenderFromAddress nullable.Type[string] `json:"senderFromAddress,omitempty"`

	// The IP address of the sender.
	SenderIP nullable.Type[string] `json:"senderIP,omitempty"`

	// Source of the message; for example, desktop and mobile.
	SourceAppName nullable.Type[string] `json:"sourceAppName,omitempty"`

	// The source ID of this Teams message.
	SourceId nullable.Type[string] `json:"sourceId,omitempty"`

	// The subject of this Teams message.
	Subject nullable.Type[string] `json:"subject,omitempty"`

	// The list of recipients who were detected as suspicious.
	SuspiciousRecipients *[]string `json:"suspiciousRecipients,omitempty"`

	// Identifier of the channel or chat that this message is part of.
	ThreadId nullable.Type[string] `json:"threadId,omitempty"`

	// The Teams message type. Supported values are: Chat, Topic, Space, and Meeting.
	ThreadType nullable.Type[string] `json:"threadType,omitempty"`

	// The URLs contained in this Teams message.
	Urls *[]SecurityUrlEvidence `json:"urls,omitempty"`

	// Fields inherited from SecurityAlertEvidence

	// The date and time when the evidence was created and added to the alert. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Detailed description of the entity role/s in an alert. Values are free-form.
	DetailedRoles *[]string `json:"detailedRoles,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RemediationStatus *SecurityEvidenceRemediationStatus `json:"remediationStatus,omitempty"`

	// Details about the remediation status.
	RemediationStatusDetails nullable.Type[string] `json:"remediationStatusDetails,omitempty"`

	// The role/s that an evidence entity represents in an alert, for example, an IP address that is associated with an
	// attacker has the evidence role Attacker.
	Roles *[]SecurityEvidenceRole `json:"roles,omitempty"`

	// Array of custom tags associated with an evidence instance, for example, to denote a group of devices, high-value
	// assets, etc.
	Tags *[]string `json:"tags,omitempty"`

	Verdict *SecurityEvidenceVerdict `json:"verdict,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityTeamsMessageEvidence) SecurityAlertEvidence() BaseSecurityAlertEvidenceImpl {
	return BaseSecurityAlertEvidenceImpl{
		CreatedDateTime:          s.CreatedDateTime,
		DetailedRoles:            s.DetailedRoles,
		ODataId:                  s.ODataId,
		ODataType:                s.ODataType,
		RemediationStatus:        s.RemediationStatus,
		RemediationStatusDetails: s.RemediationStatusDetails,
		Roles:                    s.Roles,
		Tags:                     s.Tags,
		Verdict:                  s.Verdict,
	}
}

var _ json.Marshaler = SecurityTeamsMessageEvidence{}

func (s SecurityTeamsMessageEvidence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityTeamsMessageEvidence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityTeamsMessageEvidence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityTeamsMessageEvidence: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.teamsMessageEvidence"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityTeamsMessageEvidence: %+v", err)
	}

	return encoded, nil
}
