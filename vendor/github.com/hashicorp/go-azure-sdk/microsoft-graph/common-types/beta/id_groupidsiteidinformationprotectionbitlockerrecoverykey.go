package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId{}

// GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId is a struct representing the Resource ID for a Group Id Site Id Information Protection Bitlocker Recovery Key
type GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId struct {
	GroupId                string
	SiteId                 string
	BitlockerRecoveryKeyId string
}

// NewGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID returns a new GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId struct
func NewGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID(groupId string, siteId string, bitlockerRecoveryKeyId string) GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId {
	return GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId{
		GroupId:                groupId,
		SiteId:                 siteId,
		BitlockerRecoveryKeyId: bitlockerRecoveryKeyId,
	}
}

// ParseGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID parses 'input' into a GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId
func ParseGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID(input string) (*GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyIDInsensitively(input string) (*GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.BitlockerRecoveryKeyId, ok = input.Parsed["bitlockerRecoveryKeyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "bitlockerRecoveryKeyId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID checks that 'input' can be parsed as a Group Id Site Id Information Protection Bitlocker Recovery Key ID
func ValidateGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Information Protection Bitlocker Recovery Key ID
func (id GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/bitlocker/recoveryKeys/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.BitlockerRecoveryKeyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Information Protection Bitlocker Recovery Key ID
func (id GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("bitlocker", "bitlocker", "bitlocker"),
		resourceids.StaticSegment("recoveryKeys", "recoveryKeys", "recoveryKeys"),
		resourceids.UserSpecifiedSegment("bitlockerRecoveryKeyId", "bitlockerRecoveryKeyId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Information Protection Bitlocker Recovery Key ID
func (id GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Bitlocker Recovery Key: %q", id.BitlockerRecoveryKeyId),
	}
	return fmt.Sprintf("Group Id Site Id Information Protection Bitlocker Recovery Key (%s)", strings.Join(components, "\n"))
}
