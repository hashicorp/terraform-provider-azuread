package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdApprovalId{}

// UserIdApprovalId is a struct representing the Resource ID for a User Id Approval
type UserIdApprovalId struct {
	UserId     string
	ApprovalId string
}

// NewUserIdApprovalID returns a new UserIdApprovalId struct
func NewUserIdApprovalID(userId string, approvalId string) UserIdApprovalId {
	return UserIdApprovalId{
		UserId:     userId,
		ApprovalId: approvalId,
	}
}

// ParseUserIdApprovalID parses 'input' into a UserIdApprovalId
func ParseUserIdApprovalID(input string) (*UserIdApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdApprovalIDInsensitively parses 'input' case-insensitively into a UserIdApprovalId
// note: this method should only be used for API response data and not user input
func ParseUserIdApprovalIDInsensitively(input string) (*UserIdApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdApprovalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	return nil
}

// ValidateUserIdApprovalID checks that 'input' can be parsed as a User Id Approval ID
func ValidateUserIdApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Approval ID
func (id UserIdApprovalId) ID() string {
	fmtString := "/users/%s/approvals/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ApprovalId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Approval ID
func (id UserIdApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("approvals", "approvals", "approvals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
	}
}

// String returns a human-readable description of this User Id Approval ID
func (id UserIdApprovalId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Approval: %q", id.ApprovalId),
	}
	return fmt.Sprintf("User Id Approval (%s)", strings.Join(components, "\n"))
}
