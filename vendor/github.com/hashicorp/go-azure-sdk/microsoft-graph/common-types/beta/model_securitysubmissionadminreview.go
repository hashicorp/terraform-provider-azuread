package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionAdminReview struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies who reviewed the email. The identification is an email ID or other identity strings.
	ReviewBy nullable.Type[string] `json:"reviewBy,omitempty"`

	// Specifies the date time when the review occurred.
	ReviewDateTime nullable.Type[string] `json:"reviewDateTime,omitempty"`

	// Specifies what the review result was. The possible values are: notJunk, spam, phishing, malware, allowedByPolicy,
	// blockedByPolicy, spoof, unknown, noResultAvailable, unknownFutureValue, beingAnalyzed, notSubmittedToMicrosoft,
	// phishingSimulation, allowedDueToOrganizationOverride, blockedDueToOrganizationOverride, allowedDueToUserOverride,
	// blockedDueToUserOverride, itemNotfound, threatsFound, noThreatsFound, domainImpersonation, userImpersonation,
	// brandImpersonation, authenticationFailure, spoofedBlocked, spoofedAllowed, bulk, and reasonLostInTransit. You must
	// use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum:
	// beingAnalyzed, notSubmittedToMicrosoft, phishingSimulation, allowedDueToOrganizationOverride,
	// blockedDueToOrganizationOverride, allowedDueToUserOverride, blockedDueToUserOverride, itemNotfound, threatsFound,
	// noThreatsFound, domainImpersonation, userImpersonation, brandImpersonation, authenticationFailure, spoofedBlocked,
	// spoofedAllowed, bulk, and reasonLostInTransit.
	ReviewResult *SecuritySubmissionResultCategory `json:"reviewResult,omitempty"`
}
