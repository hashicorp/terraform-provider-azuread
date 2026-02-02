package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdSettingStorageQuotaServiceId{}

// UserIdSettingStorageQuotaServiceId is a struct representing the Resource ID for a User Id Setting Storage Quota Service
type UserIdSettingStorageQuotaServiceId struct {
	UserId                         string
	ServiceStorageQuotaBreakdownId string
}

// NewUserIdSettingStorageQuotaServiceID returns a new UserIdSettingStorageQuotaServiceId struct
func NewUserIdSettingStorageQuotaServiceID(userId string, serviceStorageQuotaBreakdownId string) UserIdSettingStorageQuotaServiceId {
	return UserIdSettingStorageQuotaServiceId{
		UserId:                         userId,
		ServiceStorageQuotaBreakdownId: serviceStorageQuotaBreakdownId,
	}
}

// ParseUserIdSettingStorageQuotaServiceID parses 'input' into a UserIdSettingStorageQuotaServiceId
func ParseUserIdSettingStorageQuotaServiceID(input string) (*UserIdSettingStorageQuotaServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdSettingStorageQuotaServiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdSettingStorageQuotaServiceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdSettingStorageQuotaServiceIDInsensitively parses 'input' case-insensitively into a UserIdSettingStorageQuotaServiceId
// note: this method should only be used for API response data and not user input
func ParseUserIdSettingStorageQuotaServiceIDInsensitively(input string) (*UserIdSettingStorageQuotaServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdSettingStorageQuotaServiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdSettingStorageQuotaServiceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdSettingStorageQuotaServiceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ServiceStorageQuotaBreakdownId, ok = input.Parsed["serviceStorageQuotaBreakdownId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serviceStorageQuotaBreakdownId", input)
	}

	return nil
}

// ValidateUserIdSettingStorageQuotaServiceID checks that 'input' can be parsed as a User Id Setting Storage Quota Service ID
func ValidateUserIdSettingStorageQuotaServiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdSettingStorageQuotaServiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Setting Storage Quota Service ID
func (id UserIdSettingStorageQuotaServiceId) ID() string {
	fmtString := "/users/%s/settings/storage/quota/services/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ServiceStorageQuotaBreakdownId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Setting Storage Quota Service ID
func (id UserIdSettingStorageQuotaServiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.StaticSegment("storage", "storage", "storage"),
		resourceids.StaticSegment("quota", "quota", "quota"),
		resourceids.StaticSegment("services", "services", "services"),
		resourceids.UserSpecifiedSegment("serviceStorageQuotaBreakdownId", "serviceStorageQuotaBreakdownId"),
	}
}

// String returns a human-readable description of this User Id Setting Storage Quota Service ID
func (id UserIdSettingStorageQuotaServiceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Service Storage Quota Breakdown: %q", id.ServiceStorageQuotaBreakdownId),
	}
	return fmt.Sprintf("User Id Setting Storage Quota Service (%s)", strings.Join(components, "\n"))
}
