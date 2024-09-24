package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId{}

// GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId is a struct representing the Resource ID for a Group Id Site Id Information Protection Data Loss Prevention Policy
type GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId struct {
	GroupId                    string
	SiteId                     string
	DataLossPreventionPolicyId string
}

// NewGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID returns a new GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId struct
func NewGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID(groupId string, siteId string, dataLossPreventionPolicyId string) GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId {
	return GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId{
		GroupId:                    groupId,
		SiteId:                     siteId,
		DataLossPreventionPolicyId: dataLossPreventionPolicyId,
	}
}

// ParseGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID parses 'input' into a GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId
func ParseGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID(input string) (*GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdInformationProtectionDataLossPreventionPolicyIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdInformationProtectionDataLossPreventionPolicyIDInsensitively(input string) (*GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.DataLossPreventionPolicyId, ok = input.Parsed["dataLossPreventionPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataLossPreventionPolicyId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID checks that 'input' can be parsed as a Group Id Site Id Information Protection Data Loss Prevention Policy ID
func ValidateGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Information Protection Data Loss Prevention Policy ID
func (id GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/dataLossPreventionPolicies/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.DataLossPreventionPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Information Protection Data Loss Prevention Policy ID
func (id GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("dataLossPreventionPolicies", "dataLossPreventionPolicies", "dataLossPreventionPolicies"),
		resourceids.UserSpecifiedSegment("dataLossPreventionPolicyId", "dataLossPreventionPolicyId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Information Protection Data Loss Prevention Policy ID
func (id GroupIdSiteIdInformationProtectionDataLossPreventionPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Data Loss Prevention Policy: %q", id.DataLossPreventionPolicyId),
	}
	return fmt.Sprintf("Group Id Site Id Information Protection Data Loss Prevention Policy (%s)", strings.Join(components, "\n"))
}
