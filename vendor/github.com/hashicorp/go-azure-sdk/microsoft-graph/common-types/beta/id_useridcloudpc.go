package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCloudPCId{}

// UserIdCloudPCId is a struct representing the Resource ID for a User Id Cloud PC
type UserIdCloudPCId struct {
	UserId    string
	CloudPCId string
}

// NewUserIdCloudPCID returns a new UserIdCloudPCId struct
func NewUserIdCloudPCID(userId string, cloudPCId string) UserIdCloudPCId {
	return UserIdCloudPCId{
		UserId:    userId,
		CloudPCId: cloudPCId,
	}
}

// ParseUserIdCloudPCID parses 'input' into a UserIdCloudPCId
func ParseUserIdCloudPCID(input string) (*UserIdCloudPCId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCloudPCId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCloudPCId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCloudPCIDInsensitively parses 'input' case-insensitively into a UserIdCloudPCId
// note: this method should only be used for API response data and not user input
func ParseUserIdCloudPCIDInsensitively(input string) (*UserIdCloudPCId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCloudPCId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCloudPCId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCloudPCId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.CloudPCId, ok = input.Parsed["cloudPCId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCId", input)
	}

	return nil
}

// ValidateUserIdCloudPCID checks that 'input' can be parsed as a User Id Cloud PC ID
func ValidateUserIdCloudPCID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCloudPCID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Cloud PC ID
func (id UserIdCloudPCId) ID() string {
	fmtString := "/users/%s/cloudPCs/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CloudPCId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Cloud PC ID
func (id UserIdCloudPCId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("cloudPCs", "cloudPCs", "cloudPCs"),
		resourceids.UserSpecifiedSegment("cloudPCId", "cloudPCId"),
	}
}

// String returns a human-readable description of this User Id Cloud PC ID
func (id UserIdCloudPCId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Cloud PC: %q", id.CloudPCId),
	}
	return fmt.Sprintf("User Id Cloud PC (%s)", strings.Join(components, "\n"))
}
