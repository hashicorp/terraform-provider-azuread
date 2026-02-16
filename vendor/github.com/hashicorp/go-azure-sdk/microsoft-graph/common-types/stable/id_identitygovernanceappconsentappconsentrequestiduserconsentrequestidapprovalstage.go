package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{}

// IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId is a struct representing the Resource ID for a Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Stage
type IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId struct {
	AppConsentRequestId  string
	UserConsentRequestId string
	ApprovalStageId      string
}

// NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID returns a new IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId struct
func NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID(appConsentRequestId string, userConsentRequestId string, approvalStageId string) IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId {
	return IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{
		AppConsentRequestId:  appConsentRequestId,
		UserConsentRequestId: userConsentRequestId,
		ApprovalStageId:      approvalStageId,
	}
}

// ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID parses 'input' into a IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId
func ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID(input string) (*IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageIDInsensitively(input string) (*IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppConsentRequestId, ok = input.Parsed["appConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", input)
	}

	if id.UserConsentRequestId, ok = input.Parsed["userConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userConsentRequestId", input)
	}

	if id.ApprovalStageId, ok = input.Parsed["approvalStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStageId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID checks that 'input' can be parsed as a Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Stage ID
func ValidateIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Stage ID
func (id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId) ID() string {
	fmtString := "/identityGovernance/appConsent/appConsentRequests/%s/userConsentRequests/%s/approval/stages/%s"
	return fmt.Sprintf(fmtString, id.AppConsentRequestId, id.UserConsentRequestId, id.ApprovalStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Stage ID
func (id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("appConsent", "appConsent", "appConsent"),
		resourceids.StaticSegment("appConsentRequests", "appConsentRequests", "appConsentRequests"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestId"),
		resourceids.StaticSegment("userConsentRequests", "userConsentRequests", "userConsentRequests"),
		resourceids.UserSpecifiedSegment("userConsentRequestId", "userConsentRequestId"),
		resourceids.StaticSegment("approval", "approval", "approval"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("approvalStageId", "approvalStageId"),
	}
}

// String returns a human-readable description of this Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Stage ID
func (id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId) String() string {
	components := []string{
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
		fmt.Sprintf("User Consent Request: %q", id.UserConsentRequestId),
		fmt.Sprintf("Approval Stage: %q", id.ApprovalStageId),
	}
	return fmt.Sprintf("Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Stage (%s)", strings.Join(components, "\n"))
}
