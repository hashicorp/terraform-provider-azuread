package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionResult struct {
	// The submission result category. The possible values are: notJunk, spam, phishing, malware, allowedByPolicy,
	// blockedByPolicy, spoof, unknown, noResultAvailable, unknownFutureValue, beingAnalyzed, notSubmittedToMicrosoft,
	// phishingSimulation, allowedDueToOrganizationOverride, blockedDueToOrganizationOverride, allowedDueToUserOverride,
	// blockedDueToUserOverride, itemNotfound, threatsFound, noThreatsFound, domainImpersonation, userImpersonation,
	// brandImpersonation, authenticationFailure, spoofedBlocked, spoofedAllowed, bulk, and reasonLostInTransit. You must
	// use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum:
	// beingAnalyzed, notSubmittedToMicrosoft, phishingSimulation, allowedDueToOrganizationOverride,
	// blockedDueToOrganizationOverride, allowedDueToUserOverride, blockedDueToUserOverride, itemNotfound, threatsFound,
	// noThreatsFound, domainImpersonation, userImpersonation, brandImpersonation, authenticationFailure, spoofedBlocked,
	// spoofedAllowed, bulk, and reasonLostInTransit.
	Category *SecuritySubmissionResultCategory `json:"category,omitempty"`

	// Specifies the extra details provided by Microsoft to substantiate their analysis result.
	Detail *SecuritySubmissionResultDetail `json:"detail,omitempty"`

	// Specifies the files detected by Microsoft in the submitted emails.
	DetectedFiles *[]SecuritySubmissionDetectedFile `json:"detectedFiles,omitempty"`

	// Specifies the URLs detected by Microsoft in the submitted email.
	DetectedUrls *[]string `json:"detectedUrls,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the setting for user mailbox denoted by a comma-separated string.
	UserMailboxSetting *SecurityUserMailboxSetting `json:"userMailboxSetting,omitempty"`
}
