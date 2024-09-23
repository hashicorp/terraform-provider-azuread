package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{}

// MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId is a struct representing the Resource ID for a Me App Consent Requests For Approval Id User Consent Request Id Approval Step
type MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId struct {
	AppConsentRequestId  string
	UserConsentRequestId string
	ApprovalStepId       string
}

// NewMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID returns a new MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId struct
func NewMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(appConsentRequestId string, userConsentRequestId string, approvalStepId string) MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId {
	return MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{
		AppConsentRequestId:  appConsentRequestId,
		UserConsentRequestId: userConsentRequestId,
		ApprovalStepId:       approvalStepId,
	}
}

// ParseMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID parses 'input' into a MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId
func ParseMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(input string) (*MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepIDInsensitively parses 'input' case-insensitively into a MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId
// note: this method should only be used for API response data and not user input
func ParseMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepIDInsensitively(input string) (*MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID checks that 'input' can be parsed as a Me App Consent Requests For Approval Id User Consent Request Id Approval Step ID
func ValidateMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me App Consent Requests For Approval Id User Consent Request Id Approval Step ID
func (id MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId) ID() string {
	fmtString := "/me/appConsentRequestsForApproval/%s/userConsentRequests/%s/approval/steps/%s"
	return fmt.Sprintf(fmtString, id.AppConsentRequestId, id.UserConsentRequestId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me App Consent Requests For Approval Id User Consent Request Id Approval Step ID
func (id MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("appConsentRequestsForApproval", "appConsentRequestsForApproval", "appConsentRequestsForApproval"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestId"),
		resourceids.StaticSegment("userConsentRequests", "userConsentRequests", "userConsentRequests"),
		resourceids.UserSpecifiedSegment("userConsentRequestId", "userConsentRequestId"),
		resourceids.StaticSegment("approval", "approval", "approval"),
		resourceids.StaticSegment("steps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepId"),
	}
}

// String returns a human-readable description of this Me App Consent Requests For Approval Id User Consent Request Id Approval Step ID
func (id MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId) String() string {
	components := []string{
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
		fmt.Sprintf("User Consent Request: %q", id.UserConsentRequestId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Me App Consent Requests For Approval Id User Consent Request Id Approval Step (%s)", strings.Join(components, "\n"))
}
