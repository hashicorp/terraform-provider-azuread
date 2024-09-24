package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdApprovalIdStepId{}

// UserIdApprovalIdStepId is a struct representing the Resource ID for a User Id Approval Id Step
type UserIdApprovalIdStepId struct {
	UserId         string
	ApprovalId     string
	ApprovalStepId string
}

// NewUserIdApprovalIdStepID returns a new UserIdApprovalIdStepId struct
func NewUserIdApprovalIdStepID(userId string, approvalId string, approvalStepId string) UserIdApprovalIdStepId {
	return UserIdApprovalIdStepId{
		UserId:         userId,
		ApprovalId:     approvalId,
		ApprovalStepId: approvalStepId,
	}
}

// ParseUserIdApprovalIdStepID parses 'input' into a UserIdApprovalIdStepId
func ParseUserIdApprovalIdStepID(input string) (*UserIdApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdApprovalIdStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdApprovalIdStepIDInsensitively parses 'input' case-insensitively into a UserIdApprovalIdStepId
// note: this method should only be used for API response data and not user input
func ParseUserIdApprovalIdStepIDInsensitively(input string) (*UserIdApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdApprovalIdStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdApprovalIdStepId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	if id.ApprovalStepId, ok = input.Parsed["approvalStepId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", input)
	}

	return nil
}

// ValidateUserIdApprovalIdStepID checks that 'input' can be parsed as a User Id Approval Id Step ID
func ValidateUserIdApprovalIdStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdApprovalIdStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Approval Id Step ID
func (id UserIdApprovalIdStepId) ID() string {
	fmtString := "/users/%s/approvals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Approval Id Step ID
func (id UserIdApprovalIdStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("approvals", "approvals", "approvals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
		resourceids.StaticSegment("steps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepId"),
	}
}

// String returns a human-readable description of this User Id Approval Id Step ID
func (id UserIdApprovalIdStepId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("User Id Approval Id Step (%s)", strings.Join(components, "\n"))
}
