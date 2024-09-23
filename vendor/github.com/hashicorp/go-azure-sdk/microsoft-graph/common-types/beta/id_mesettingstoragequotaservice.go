package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeSettingStorageQuotaServiceId{}

// MeSettingStorageQuotaServiceId is a struct representing the Resource ID for a Me Setting Storage Quota Service
type MeSettingStorageQuotaServiceId struct {
	ServiceStorageQuotaBreakdownId string
}

// NewMeSettingStorageQuotaServiceID returns a new MeSettingStorageQuotaServiceId struct
func NewMeSettingStorageQuotaServiceID(serviceStorageQuotaBreakdownId string) MeSettingStorageQuotaServiceId {
	return MeSettingStorageQuotaServiceId{
		ServiceStorageQuotaBreakdownId: serviceStorageQuotaBreakdownId,
	}
}

// ParseMeSettingStorageQuotaServiceID parses 'input' into a MeSettingStorageQuotaServiceId
func ParseMeSettingStorageQuotaServiceID(input string) (*MeSettingStorageQuotaServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeSettingStorageQuotaServiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeSettingStorageQuotaServiceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeSettingStorageQuotaServiceIDInsensitively parses 'input' case-insensitively into a MeSettingStorageQuotaServiceId
// note: this method should only be used for API response data and not user input
func ParseMeSettingStorageQuotaServiceIDInsensitively(input string) (*MeSettingStorageQuotaServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeSettingStorageQuotaServiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeSettingStorageQuotaServiceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeSettingStorageQuotaServiceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServiceStorageQuotaBreakdownId, ok = input.Parsed["serviceStorageQuotaBreakdownId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serviceStorageQuotaBreakdownId", input)
	}

	return nil
}

// ValidateMeSettingStorageQuotaServiceID checks that 'input' can be parsed as a Me Setting Storage Quota Service ID
func ValidateMeSettingStorageQuotaServiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeSettingStorageQuotaServiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Setting Storage Quota Service ID
func (id MeSettingStorageQuotaServiceId) ID() string {
	fmtString := "/me/settings/storage/quota/services/%s"
	return fmt.Sprintf(fmtString, id.ServiceStorageQuotaBreakdownId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Setting Storage Quota Service ID
func (id MeSettingStorageQuotaServiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.StaticSegment("storage", "storage", "storage"),
		resourceids.StaticSegment("quota", "quota", "quota"),
		resourceids.StaticSegment("services", "services", "services"),
		resourceids.UserSpecifiedSegment("serviceStorageQuotaBreakdownId", "serviceStorageQuotaBreakdownId"),
	}
}

// String returns a human-readable description of this Me Setting Storage Quota Service ID
func (id MeSettingStorageQuotaServiceId) String() string {
	components := []string{
		fmt.Sprintf("Service Storage Quota Breakdown: %q", id.ServiceStorageQuotaBreakdownId),
	}
	return fmt.Sprintf("Me Setting Storage Quota Service (%s)", strings.Join(components, "\n"))
}
