package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAppConsentRequestsForApprovalId{}

// UserIdAppConsentRequestsForApprovalId is a struct representing the Resource ID for a User Id App Consent Requests For Approval
type UserIdAppConsentRequestsForApprovalId struct {
	UserId              string
	AppConsentRequestId string
}

// NewUserIdAppConsentRequestsForApprovalID returns a new UserIdAppConsentRequestsForApprovalId struct
func NewUserIdAppConsentRequestsForApprovalID(userId string, appConsentRequestId string) UserIdAppConsentRequestsForApprovalId {
	return UserIdAppConsentRequestsForApprovalId{
		UserId:              userId,
		AppConsentRequestId: appConsentRequestId,
	}
}

// ParseUserIdAppConsentRequestsForApprovalID parses 'input' into a UserIdAppConsentRequestsForApprovalId
func ParseUserIdAppConsentRequestsForApprovalID(input string) (*UserIdAppConsentRequestsForApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAppConsentRequestsForApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAppConsentRequestsForApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAppConsentRequestsForApprovalIDInsensitively parses 'input' case-insensitively into a UserIdAppConsentRequestsForApprovalId
// note: this method should only be used for API response data and not user input
func ParseUserIdAppConsentRequestsForApprovalIDInsensitively(input string) (*UserIdAppConsentRequestsForApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAppConsentRequestsForApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAppConsentRequestsForApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAppConsentRequestsForApprovalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.AppConsentRequestId, ok = input.Parsed["appConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", input)
	}

	return nil
}

// ValidateUserIdAppConsentRequestsForApprovalID checks that 'input' can be parsed as a User Id App Consent Requests For Approval ID
func ValidateUserIdAppConsentRequestsForApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAppConsentRequestsForApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id App Consent Requests For Approval ID
func (id UserIdAppConsentRequestsForApprovalId) ID() string {
	fmtString := "/users/%s/appConsentRequestsForApproval/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AppConsentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id App Consent Requests For Approval ID
func (id UserIdAppConsentRequestsForApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("appConsentRequestsForApproval", "appConsentRequestsForApproval", "appConsentRequestsForApproval"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestId"),
	}
}

// String returns a human-readable description of this User Id App Consent Requests For Approval ID
func (id UserIdAppConsentRequestsForApprovalId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
	}
	return fmt.Sprintf("User Id App Consent Requests For Approval (%s)", strings.Join(components, "\n"))
}
