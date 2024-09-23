package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{}

// UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId is a struct representing the Resource ID for a User Id App Consent Requests For Approval Id User Consent Request Id Approval Step
type UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId struct {
	UserId               string
	AppConsentRequestId  string
	UserConsentRequestId string
	ApprovalStepId       string
}

// NewUserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID returns a new UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId struct
func NewUserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(userId string, appConsentRequestId string, userConsentRequestId string, approvalStepId string) UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId {
	return UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{
		UserId:               userId,
		AppConsentRequestId:  appConsentRequestId,
		UserConsentRequestId: userConsentRequestId,
		ApprovalStepId:       approvalStepId,
	}
}

// ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID parses 'input' into a UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId
func ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(input string) (*UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepIDInsensitively parses 'input' case-insensitively into a UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId
// note: this method should only be used for API response data and not user input
func ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepIDInsensitively(input string) (*UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ApprovalStepId, ok = input.Parsed["approvalStepId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", input)
	}

	return nil
}

// ValidateUserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID checks that 'input' can be parsed as a User Id App Consent Requests For Approval Id User Consent Request Id Approval Step ID
func ValidateUserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id App Consent Requests For Approval Id User Consent Request Id Approval Step ID
func (id UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId) ID() string {
	fmtString := "/users/%s/appConsentRequestsForApproval/%s/userConsentRequests/%s/approval/steps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AppConsentRequestId, id.UserConsentRequestId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id App Consent Requests For Approval Id User Consent Request Id Approval Step ID
func (id UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("appConsentRequestsForApproval", "appConsentRequestsForApproval", "appConsentRequestsForApproval"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestId"),
		resourceids.StaticSegment("userConsentRequests", "userConsentRequests", "userConsentRequests"),
		resourceids.UserSpecifiedSegment("userConsentRequestId", "userConsentRequestId"),
		resourceids.StaticSegment("approval", "approval", "approval"),
		resourceids.StaticSegment("steps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepId"),
	}
}

// String returns a human-readable description of this User Id App Consent Requests For Approval Id User Consent Request Id Approval Step ID
func (id UserIdAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
		fmt.Sprintf("User Consent Request: %q", id.UserConsentRequestId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("User Id App Consent Requests For Approval Id User Consent Request Id Approval Step (%s)", strings.Join(components, "\n"))
}
