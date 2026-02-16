package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAppConsentRequestsForApprovalIdUserConsentRequestId{}

// MeAppConsentRequestsForApprovalIdUserConsentRequestId is a struct representing the Resource ID for a Me App Consent Requests For Approval Id User Consent Request
type MeAppConsentRequestsForApprovalIdUserConsentRequestId struct {
	AppConsentRequestId  string
	UserConsentRequestId string
}

// NewMeAppConsentRequestsForApprovalIdUserConsentRequestID returns a new MeAppConsentRequestsForApprovalIdUserConsentRequestId struct
func NewMeAppConsentRequestsForApprovalIdUserConsentRequestID(appConsentRequestId string, userConsentRequestId string) MeAppConsentRequestsForApprovalIdUserConsentRequestId {
	return MeAppConsentRequestsForApprovalIdUserConsentRequestId{
		AppConsentRequestId:  appConsentRequestId,
		UserConsentRequestId: userConsentRequestId,
	}
}

// ParseMeAppConsentRequestsForApprovalIdUserConsentRequestID parses 'input' into a MeAppConsentRequestsForApprovalIdUserConsentRequestId
func ParseMeAppConsentRequestsForApprovalIdUserConsentRequestID(input string) (*MeAppConsentRequestsForApprovalIdUserConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAppConsentRequestsForApprovalIdUserConsentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAppConsentRequestsForApprovalIdUserConsentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAppConsentRequestsForApprovalIdUserConsentRequestIDInsensitively parses 'input' case-insensitively into a MeAppConsentRequestsForApprovalIdUserConsentRequestId
// note: this method should only be used for API response data and not user input
func ParseMeAppConsentRequestsForApprovalIdUserConsentRequestIDInsensitively(input string) (*MeAppConsentRequestsForApprovalIdUserConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAppConsentRequestsForApprovalIdUserConsentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAppConsentRequestsForApprovalIdUserConsentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAppConsentRequestsForApprovalIdUserConsentRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppConsentRequestId, ok = input.Parsed["appConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", input)
	}

	if id.UserConsentRequestId, ok = input.Parsed["userConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userConsentRequestId", input)
	}

	return nil
}

// ValidateMeAppConsentRequestsForApprovalIdUserConsentRequestID checks that 'input' can be parsed as a Me App Consent Requests For Approval Id User Consent Request ID
func ValidateMeAppConsentRequestsForApprovalIdUserConsentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAppConsentRequestsForApprovalIdUserConsentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me App Consent Requests For Approval Id User Consent Request ID
func (id MeAppConsentRequestsForApprovalIdUserConsentRequestId) ID() string {
	fmtString := "/me/appConsentRequestsForApproval/%s/userConsentRequests/%s"
	return fmt.Sprintf(fmtString, id.AppConsentRequestId, id.UserConsentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me App Consent Requests For Approval Id User Consent Request ID
func (id MeAppConsentRequestsForApprovalIdUserConsentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("appConsentRequestsForApproval", "appConsentRequestsForApproval", "appConsentRequestsForApproval"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestId"),
		resourceids.StaticSegment("userConsentRequests", "userConsentRequests", "userConsentRequests"),
		resourceids.UserSpecifiedSegment("userConsentRequestId", "userConsentRequestId"),
	}
}

// String returns a human-readable description of this Me App Consent Requests For Approval Id User Consent Request ID
func (id MeAppConsentRequestsForApprovalIdUserConsentRequestId) String() string {
	components := []string{
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
		fmt.Sprintf("User Consent Request: %q", id.UserConsentRequestId),
	}
	return fmt.Sprintf("Me App Consent Requests For Approval Id User Consent Request (%s)", strings.Join(components, "\n"))
}
