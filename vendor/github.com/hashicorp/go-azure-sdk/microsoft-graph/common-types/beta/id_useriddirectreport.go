package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDirectReportId{}

// UserIdDirectReportId is a struct representing the Resource ID for a User Id Direct Report
type UserIdDirectReportId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserIdDirectReportID returns a new UserIdDirectReportId struct
func NewUserIdDirectReportID(userId string, directoryObjectId string) UserIdDirectReportId {
	return UserIdDirectReportId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserIdDirectReportID parses 'input' into a UserIdDirectReportId
func ParseUserIdDirectReportID(input string) (*UserIdDirectReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDirectReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDirectReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDirectReportIDInsensitively parses 'input' case-insensitively into a UserIdDirectReportId
// note: this method should only be used for API response data and not user input
func ParseUserIdDirectReportIDInsensitively(input string) (*UserIdDirectReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDirectReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDirectReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDirectReportId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateUserIdDirectReportID checks that 'input' can be parsed as a User Id Direct Report ID
func ValidateUserIdDirectReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDirectReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Direct Report ID
func (id UserIdDirectReportId) ID() string {
	fmtString := "/users/%s/directReports/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Direct Report ID
func (id UserIdDirectReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("directReports", "directReports", "directReports"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this User Id Direct Report ID
func (id UserIdDirectReportId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Id Direct Report (%s)", strings.Join(components, "\n"))
}
