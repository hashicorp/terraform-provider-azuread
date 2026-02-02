package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId{}

// IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId is a struct representing the Resource ID for a Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Step
type IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId struct {
	AppConsentRequestId  string
	UserConsentRequestId string
	ApprovalStepId       string
}

// NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepID returns a new IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId struct
func NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepID(appConsentRequestId string, userConsentRequestId string, approvalStepId string) IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId {
	return IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId{
		AppConsentRequestId:  appConsentRequestId,
		UserConsentRequestId: userConsentRequestId,
		ApprovalStepId:       approvalStepId,
	}
}

// ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepID parses 'input' into a IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId
func ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepID(input string) (*IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepIDInsensitively(input string) (*IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppConsentRequestId, ok = input.Parsed["appConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", input)
	}

	if id.UserConsentRequestId, ok = input.Parsed["userConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userConsentRequestId", input)
	}

	if id.ApprovalStepId, ok = input.Parsed["approvalStepId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepID checks that 'input' can be parsed as a Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Step ID
func ValidateIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Step ID
func (id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId) ID() string {
	fmtString := "/identityGovernance/appConsent/appConsentRequests/%s/userConsentRequests/%s/approval/steps/%s"
	return fmt.Sprintf(fmtString, id.AppConsentRequestId, id.UserConsentRequestId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Step ID
func (id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("appConsent", "appConsent", "appConsent"),
		resourceids.StaticSegment("appConsentRequests", "appConsentRequests", "appConsentRequests"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestId"),
		resourceids.StaticSegment("userConsentRequests", "userConsentRequests", "userConsentRequests"),
		resourceids.UserSpecifiedSegment("userConsentRequestId", "userConsentRequestId"),
		resourceids.StaticSegment("approval", "approval", "approval"),
		resourceids.StaticSegment("steps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepId"),
	}
}

// String returns a human-readable description of this Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Step ID
func (id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStepId) String() string {
	components := []string{
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
		fmt.Sprintf("User Consent Request: %q", id.UserConsentRequestId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Identity Governance App Consent App Consent Request Id User Consent Request Id Approval Step (%s)", strings.Join(components, "\n"))
}
