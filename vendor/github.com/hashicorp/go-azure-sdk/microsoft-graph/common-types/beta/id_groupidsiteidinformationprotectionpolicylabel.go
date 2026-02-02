package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdInformationProtectionPolicyLabelId{}

// GroupIdSiteIdInformationProtectionPolicyLabelId is a struct representing the Resource ID for a Group Id Site Id Information Protection Policy Label
type GroupIdSiteIdInformationProtectionPolicyLabelId struct {
	GroupId                      string
	SiteId                       string
	InformationProtectionLabelId string
}

// NewGroupIdSiteIdInformationProtectionPolicyLabelID returns a new GroupIdSiteIdInformationProtectionPolicyLabelId struct
func NewGroupIdSiteIdInformationProtectionPolicyLabelID(groupId string, siteId string, informationProtectionLabelId string) GroupIdSiteIdInformationProtectionPolicyLabelId {
	return GroupIdSiteIdInformationProtectionPolicyLabelId{
		GroupId:                      groupId,
		SiteId:                       siteId,
		InformationProtectionLabelId: informationProtectionLabelId,
	}
}

// ParseGroupIdSiteIdInformationProtectionPolicyLabelID parses 'input' into a GroupIdSiteIdInformationProtectionPolicyLabelId
func ParseGroupIdSiteIdInformationProtectionPolicyLabelID(input string) (*GroupIdSiteIdInformationProtectionPolicyLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionPolicyLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionPolicyLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdInformationProtectionPolicyLabelIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdInformationProtectionPolicyLabelId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdInformationProtectionPolicyLabelIDInsensitively(input string) (*GroupIdSiteIdInformationProtectionPolicyLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionPolicyLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionPolicyLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdInformationProtectionPolicyLabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.InformationProtectionLabelId, ok = input.Parsed["informationProtectionLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "informationProtectionLabelId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdInformationProtectionPolicyLabelID checks that 'input' can be parsed as a Group Id Site Id Information Protection Policy Label ID
func ValidateGroupIdSiteIdInformationProtectionPolicyLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdInformationProtectionPolicyLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Information Protection Policy Label ID
func (id GroupIdSiteIdInformationProtectionPolicyLabelId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/policy/labels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.InformationProtectionLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Information Protection Policy Label ID
func (id GroupIdSiteIdInformationProtectionPolicyLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("policy", "policy", "policy"),
		resourceids.StaticSegment("labels", "labels", "labels"),
		resourceids.UserSpecifiedSegment("informationProtectionLabelId", "informationProtectionLabelId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Information Protection Policy Label ID
func (id GroupIdSiteIdInformationProtectionPolicyLabelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Information Protection Label: %q", id.InformationProtectionLabelId),
	}
	return fmt.Sprintf("Group Id Site Id Information Protection Policy Label (%s)", strings.Join(components, "\n"))
}
