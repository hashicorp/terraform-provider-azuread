package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAppConsentRequestsForApprovalId{}

// MeAppConsentRequestsForApprovalId is a struct representing the Resource ID for a Me App Consent Requests For Approval
type MeAppConsentRequestsForApprovalId struct {
	AppConsentRequestId string
}

// NewMeAppConsentRequestsForApprovalID returns a new MeAppConsentRequestsForApprovalId struct
func NewMeAppConsentRequestsForApprovalID(appConsentRequestId string) MeAppConsentRequestsForApprovalId {
	return MeAppConsentRequestsForApprovalId{
		AppConsentRequestId: appConsentRequestId,
	}
}

// ParseMeAppConsentRequestsForApprovalID parses 'input' into a MeAppConsentRequestsForApprovalId
func ParseMeAppConsentRequestsForApprovalID(input string) (*MeAppConsentRequestsForApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAppConsentRequestsForApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAppConsentRequestsForApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAppConsentRequestsForApprovalIDInsensitively parses 'input' case-insensitively into a MeAppConsentRequestsForApprovalId
// note: this method should only be used for API response data and not user input
func ParseMeAppConsentRequestsForApprovalIDInsensitively(input string) (*MeAppConsentRequestsForApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAppConsentRequestsForApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAppConsentRequestsForApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAppConsentRequestsForApprovalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppConsentRequestId, ok = input.Parsed["appConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", input)
	}

	return nil
}

// ValidateMeAppConsentRequestsForApprovalID checks that 'input' can be parsed as a Me App Consent Requests For Approval ID
func ValidateMeAppConsentRequestsForApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAppConsentRequestsForApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me App Consent Requests For Approval ID
func (id MeAppConsentRequestsForApprovalId) ID() string {
	fmtString := "/me/appConsentRequestsForApproval/%s"
	return fmt.Sprintf(fmtString, id.AppConsentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me App Consent Requests For Approval ID
func (id MeAppConsentRequestsForApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("appConsentRequestsForApproval", "appConsentRequestsForApproval", "appConsentRequestsForApproval"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestId"),
	}
}

// String returns a human-readable description of this Me App Consent Requests For Approval ID
func (id MeAppConsentRequestsForApprovalId) String() string {
	components := []string{
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
	}
	return fmt.Sprintf("Me App Consent Requests For Approval (%s)", strings.Join(components, "\n"))
}
