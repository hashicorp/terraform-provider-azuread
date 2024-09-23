package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAppConsentRequestsForApprovalIdUserConsentRequestId{}

// UserIdAppConsentRequestsForApprovalIdUserConsentRequestId is a struct representing the Resource ID for a User Id App Consent Requests For Approval Id User Consent Request
type UserIdAppConsentRequestsForApprovalIdUserConsentRequestId struct {
	UserId               string
	AppConsentRequestId  string
	UserConsentRequestId string
}

// NewUserIdAppConsentRequestsForApprovalIdUserConsentRequestID returns a new UserIdAppConsentRequestsForApprovalIdUserConsentRequestId struct
func NewUserIdAppConsentRequestsForApprovalIdUserConsentRequestID(userId string, appConsentRequestId string, userConsentRequestId string) UserIdAppConsentRequestsForApprovalIdUserConsentRequestId {
	return UserIdAppConsentRequestsForApprovalIdUserConsentRequestId{
		UserId:               userId,
		AppConsentRequestId:  appConsentRequestId,
		UserConsentRequestId: userConsentRequestId,
	}
}

// ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestID parses 'input' into a UserIdAppConsentRequestsForApprovalIdUserConsentRequestId
func ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestID(input string) (*UserIdAppConsentRequestsForApprovalIdUserConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAppConsentRequestsForApprovalIdUserConsentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAppConsentRequestsForApprovalIdUserConsentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestIDInsensitively parses 'input' case-insensitively into a UserIdAppConsentRequestsForApprovalIdUserConsentRequestId
// note: this method should only be used for API response data and not user input
func ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestIDInsensitively(input string) (*UserIdAppConsentRequestsForApprovalIdUserConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAppConsentRequestsForApprovalIdUserConsentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAppConsentRequestsForApprovalIdUserConsentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAppConsentRequestsForApprovalIdUserConsentRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.AppConsentRequestId, ok = input.Parsed["appConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", input)
	}

	if id.UserConsentRequestId, ok = input.Parsed["userConsentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userConsentRequestId", input)
	}

	return nil
}

// ValidateUserIdAppConsentRequestsForApprovalIdUserConsentRequestID checks that 'input' can be parsed as a User Id App Consent Requests For Approval Id User Consent Request ID
func ValidateUserIdAppConsentRequestsForApprovalIdUserConsentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id App Consent Requests For Approval Id User Consent Request ID
func (id UserIdAppConsentRequestsForApprovalIdUserConsentRequestId) ID() string {
	fmtString := "/users/%s/appConsentRequestsForApproval/%s/userConsentRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AppConsentRequestId, id.UserConsentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id App Consent Requests For Approval Id User Consent Request ID
func (id UserIdAppConsentRequestsForApprovalIdUserConsentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("appConsentRequestsForApproval", "appConsentRequestsForApproval", "appConsentRequestsForApproval"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestId"),
		resourceids.StaticSegment("userConsentRequests", "userConsentRequests", "userConsentRequests"),
		resourceids.UserSpecifiedSegment("userConsentRequestId", "userConsentRequestId"),
	}
}

// String returns a human-readable description of this User Id App Consent Requests For Approval Id User Consent Request ID
func (id UserIdAppConsentRequestsForApprovalIdUserConsentRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
		fmt.Sprintf("User Consent Request: %q", id.UserConsentRequestId),
	}
	return fmt.Sprintf("User Id App Consent Requests For Approval Id User Consent Request (%s)", strings.Join(components, "\n"))
}
