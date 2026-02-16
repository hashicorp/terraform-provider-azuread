package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookMasterCategoryId{}

// UserIdOutlookMasterCategoryId is a struct representing the Resource ID for a User Id Outlook Master Category
type UserIdOutlookMasterCategoryId struct {
	UserId            string
	OutlookCategoryId string
}

// NewUserIdOutlookMasterCategoryID returns a new UserIdOutlookMasterCategoryId struct
func NewUserIdOutlookMasterCategoryID(userId string, outlookCategoryId string) UserIdOutlookMasterCategoryId {
	return UserIdOutlookMasterCategoryId{
		UserId:            userId,
		OutlookCategoryId: outlookCategoryId,
	}
}

// ParseUserIdOutlookMasterCategoryID parses 'input' into a UserIdOutlookMasterCategoryId
func ParseUserIdOutlookMasterCategoryID(input string) (*UserIdOutlookMasterCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookMasterCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookMasterCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOutlookMasterCategoryIDInsensitively parses 'input' case-insensitively into a UserIdOutlookMasterCategoryId
// note: this method should only be used for API response data and not user input
func ParseUserIdOutlookMasterCategoryIDInsensitively(input string) (*UserIdOutlookMasterCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookMasterCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookMasterCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOutlookMasterCategoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OutlookCategoryId, ok = input.Parsed["outlookCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookCategoryId", input)
	}

	return nil
}

// ValidateUserIdOutlookMasterCategoryID checks that 'input' can be parsed as a User Id Outlook Master Category ID
func ValidateUserIdOutlookMasterCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOutlookMasterCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Outlook Master Category ID
func (id UserIdOutlookMasterCategoryId) ID() string {
	fmtString := "/users/%s/outlook/masterCategories/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookCategoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Outlook Master Category ID
func (id UserIdOutlookMasterCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("masterCategories", "masterCategories", "masterCategories"),
		resourceids.UserSpecifiedSegment("outlookCategoryId", "outlookCategoryId"),
	}
}

// String returns a human-readable description of this User Id Outlook Master Category ID
func (id UserIdOutlookMasterCategoryId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Category: %q", id.OutlookCategoryId),
	}
	return fmt.Sprintf("User Id Outlook Master Category (%s)", strings.Join(components, "\n"))
}
