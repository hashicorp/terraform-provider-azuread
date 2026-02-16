package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityEmailThreatSubmissionPolicy{}

type SecurityEmailThreatSubmissionPolicy struct {
	// Specifies the email address of the sender from which email notifications will be sent to end users to inform them
	// whether an email is spam, phish or clean. The default value is null. Optional for creation.
	CustomizedNotificationSenderEmailAddress *string `json:"customizedNotificationSenderEmailAddress,omitempty"`

	// Specifies the destination where the reported messages from end users land whenever they report something as phish,
	// junk or not junk. The default value is null. Optional for creation.
	CustomizedReportRecipientEmailAddress *string `json:"customizedReportRecipientEmailAddress,omitempty"`

	// Indicates whether end users can report a message as spam, phish or junk directly without a confirmation(popup). The
	// default value is true. Optional for creation.
	IsAlwaysReportEnabledForUsers *bool `json:"isAlwaysReportEnabledForUsers,omitempty"`

	// Indicates whether end users can confirm using a popup before reporting messages as spam, phish or not junk. The
	// default value is true. Optional for creation.
	IsAskMeEnabledForUsers *bool `json:"isAskMeEnabledForUsers,omitempty"`

	// Indicates whether the email notifications sent to end users to inform them if an email is a phish mail, spam or junk
	// is customized or not. The default value is false. Optional for creation.
	IsCustomizedMessageEnabled *bool `json:"isCustomizedMessageEnabled,omitempty"`

	// If enabled, customized message only shows when email is reported as phishing. The default value is false. Optional
	// for creation.
	IsCustomizedMessageEnabledForPhishing *bool `json:"isCustomizedMessageEnabledForPhishing,omitempty"`

	// Indicates whether to use the sender email address set using customizedNotificationSenderEmailAddress for sending
	// email notifications to end users. The default value is false. Optional for creation.
	IsCustomizedNotificationSenderEnabled *bool `json:"isCustomizedNotificationSenderEnabled,omitempty"`

	// Indicates whether end users can move the message from one folder to another based on the action of spam, phish or not
	// junk without actually reporting it. The default value is true. Optional for creation.
	IsNeverReportEnabledForUsers *bool `json:"isNeverReportEnabledForUsers,omitempty"`

	// Indicates whether the branding logo should be used in the email notifications sent to end users. The default value is
	// false. Optional for creation.
	IsOrganizationBrandingEnabled *bool `json:"isOrganizationBrandingEnabled,omitempty"`

	// Indicates whether end users can submit from the quarantine page. The default value is true. Optional for creation.
	IsReportFromQuarantineEnabled *bool `json:"isReportFromQuarantineEnabled,omitempty"`

	// Indicates whether emails reported by end users should be sent to the custom mailbox configured using
	// customizedReportRecipientEmailAddress. The default value is false. Optional for creation.
	IsReportToCustomizedEmailAddressEnabled *bool `json:"isReportToCustomizedEmailAddressEnabled,omitempty"`

	// If enabled, the email is sent to Microsoft for analysis. The default value is false. Required for creation.
	IsReportToMicrosoftEnabled *bool `json:"isReportToMicrosoftEnabled,omitempty"`

	// Indicates whether an email notification is sent to the end user who reported the email when it has been reviewed by
	// the admin. The default value is false. Optional for creation.
	IsReviewEmailNotificationEnabled *bool `json:"isReviewEmailNotificationEnabled,omitempty"`

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

func (s SecurityEmailThreatSubmissionPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEmailThreatSubmissionPolicy{}

func (s SecurityEmailThreatSubmissionPolicy) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEmailThreatSubmissionPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEmailThreatSubmissionPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEmailThreatSubmissionPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.emailThreatSubmissionPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEmailThreatSubmissionPolicy: %+v", err)
	}

	return encoded, nil
}
