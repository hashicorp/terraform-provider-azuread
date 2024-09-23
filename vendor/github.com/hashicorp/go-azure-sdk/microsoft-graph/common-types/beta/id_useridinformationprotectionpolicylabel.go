package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionPolicyLabelId{}

// UserIdInformationProtectionPolicyLabelId is a struct representing the Resource ID for a User Id Information Protection Policy Label
type UserIdInformationProtectionPolicyLabelId struct {
	UserId                       string
	InformationProtectionLabelId string
}

// NewUserIdInformationProtectionPolicyLabelID returns a new UserIdInformationProtectionPolicyLabelId struct
func NewUserIdInformationProtectionPolicyLabelID(userId string, informationProtectionLabelId string) UserIdInformationProtectionPolicyLabelId {
	return UserIdInformationProtectionPolicyLabelId{
		UserId:                       userId,
		InformationProtectionLabelId: informationProtectionLabelId,
	}
}

// ParseUserIdInformationProtectionPolicyLabelID parses 'input' into a UserIdInformationProtectionPolicyLabelId
func ParseUserIdInformationProtectionPolicyLabelID(input string) (*UserIdInformationProtectionPolicyLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionPolicyLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionPolicyLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdInformationProtectionPolicyLabelIDInsensitively parses 'input' case-insensitively into a UserIdInformationProtectionPolicyLabelId
// note: this method should only be used for API response data and not user input
func ParseUserIdInformationProtectionPolicyLabelIDInsensitively(input string) (*UserIdInformationProtectionPolicyLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionPolicyLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionPolicyLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdInformationProtectionPolicyLabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.InformationProtectionLabelId, ok = input.Parsed["informationProtectionLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "informationProtectionLabelId", input)
	}

	return nil
}

// ValidateUserIdInformationProtectionPolicyLabelID checks that 'input' can be parsed as a User Id Information Protection Policy Label ID
func ValidateUserIdInformationProtectionPolicyLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdInformationProtectionPolicyLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Information Protection Policy Label ID
func (id UserIdInformationProtectionPolicyLabelId) ID() string {
	fmtString := "/users/%s/informationProtection/policy/labels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.InformationProtectionLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Information Protection Policy Label ID
func (id UserIdInformationProtectionPolicyLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("policy", "policy", "policy"),
		resourceids.StaticSegment("labels", "labels", "labels"),
		resourceids.UserSpecifiedSegment("informationProtectionLabelId", "informationProtectionLabelId"),
	}
}

// String returns a human-readable description of this User Id Information Protection Policy Label ID
func (id UserIdInformationProtectionPolicyLabelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Information Protection Label: %q", id.InformationProtectionLabelId),
	}
	return fmt.Sprintf("User Id Information Protection Policy Label (%s)", strings.Join(components, "\n"))
}
