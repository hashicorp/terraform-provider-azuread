package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDataSecurityAndGovernanceActivityContentActivityId{}

// UserIdDataSecurityAndGovernanceActivityContentActivityId is a struct representing the Resource ID for a User Id Data Security And Governance Activity Content Activity
type UserIdDataSecurityAndGovernanceActivityContentActivityId struct {
	UserId            string
	ContentActivityId string
}

// NewUserIdDataSecurityAndGovernanceActivityContentActivityID returns a new UserIdDataSecurityAndGovernanceActivityContentActivityId struct
func NewUserIdDataSecurityAndGovernanceActivityContentActivityID(userId string, contentActivityId string) UserIdDataSecurityAndGovernanceActivityContentActivityId {
	return UserIdDataSecurityAndGovernanceActivityContentActivityId{
		UserId:            userId,
		ContentActivityId: contentActivityId,
	}
}

// ParseUserIdDataSecurityAndGovernanceActivityContentActivityID parses 'input' into a UserIdDataSecurityAndGovernanceActivityContentActivityId
func ParseUserIdDataSecurityAndGovernanceActivityContentActivityID(input string) (*UserIdDataSecurityAndGovernanceActivityContentActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDataSecurityAndGovernanceActivityContentActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDataSecurityAndGovernanceActivityContentActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDataSecurityAndGovernanceActivityContentActivityIDInsensitively parses 'input' case-insensitively into a UserIdDataSecurityAndGovernanceActivityContentActivityId
// note: this method should only be used for API response data and not user input
func ParseUserIdDataSecurityAndGovernanceActivityContentActivityIDInsensitively(input string) (*UserIdDataSecurityAndGovernanceActivityContentActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDataSecurityAndGovernanceActivityContentActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDataSecurityAndGovernanceActivityContentActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDataSecurityAndGovernanceActivityContentActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ContentActivityId, ok = input.Parsed["contentActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentActivityId", input)
	}

	return nil
}

// ValidateUserIdDataSecurityAndGovernanceActivityContentActivityID checks that 'input' can be parsed as a User Id Data Security And Governance Activity Content Activity ID
func ValidateUserIdDataSecurityAndGovernanceActivityContentActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDataSecurityAndGovernanceActivityContentActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Data Security And Governance Activity Content Activity ID
func (id UserIdDataSecurityAndGovernanceActivityContentActivityId) ID() string {
	fmtString := "/users/%s/dataSecurityAndGovernance/activities/contentActivities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContentActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Data Security And Governance Activity Content Activity ID
func (id UserIdDataSecurityAndGovernanceActivityContentActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("dataSecurityAndGovernance", "dataSecurityAndGovernance", "dataSecurityAndGovernance"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.StaticSegment("contentActivities", "contentActivities", "contentActivities"),
		resourceids.UserSpecifiedSegment("contentActivityId", "contentActivityId"),
	}
}

// String returns a human-readable description of this User Id Data Security And Governance Activity Content Activity ID
func (id UserIdDataSecurityAndGovernanceActivityContentActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Content Activity: %q", id.ContentActivityId),
	}
	return fmt.Sprintf("User Id Data Security And Governance Activity Content Activity (%s)", strings.Join(components, "\n"))
}
