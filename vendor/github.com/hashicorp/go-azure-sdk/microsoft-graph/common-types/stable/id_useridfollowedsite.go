package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdFollowedSiteId{}

// UserIdFollowedSiteId is a struct representing the Resource ID for a User Id Followed Site
type UserIdFollowedSiteId struct {
	UserId string
	SiteId string
}

// NewUserIdFollowedSiteID returns a new UserIdFollowedSiteId struct
func NewUserIdFollowedSiteID(userId string, siteId string) UserIdFollowedSiteId {
	return UserIdFollowedSiteId{
		UserId: userId,
		SiteId: siteId,
	}
}

// ParseUserIdFollowedSiteID parses 'input' into a UserIdFollowedSiteId
func ParseUserIdFollowedSiteID(input string) (*UserIdFollowedSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdFollowedSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdFollowedSiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdFollowedSiteIDInsensitively parses 'input' case-insensitively into a UserIdFollowedSiteId
// note: this method should only be used for API response data and not user input
func ParseUserIdFollowedSiteIDInsensitively(input string) (*UserIdFollowedSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdFollowedSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdFollowedSiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdFollowedSiteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	return nil
}

// ValidateUserIdFollowedSiteID checks that 'input' can be parsed as a User Id Followed Site ID
func ValidateUserIdFollowedSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdFollowedSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Followed Site ID
func (id UserIdFollowedSiteId) ID() string {
	fmtString := "/users/%s/followedSites/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SiteId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Followed Site ID
func (id UserIdFollowedSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("followedSites", "followedSites", "followedSites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
	}
}

// String returns a human-readable description of this User Id Followed Site ID
func (id UserIdFollowedSiteId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Site: %q", id.SiteId),
	}
	return fmt.Sprintf("User Id Followed Site (%s)", strings.Join(components, "\n"))
}
