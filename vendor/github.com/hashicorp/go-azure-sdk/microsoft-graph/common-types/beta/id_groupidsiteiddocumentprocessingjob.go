package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdDocumentProcessingJobId{}

// GroupIdSiteIdDocumentProcessingJobId is a struct representing the Resource ID for a Group Id Site Id Document Processing Job
type GroupIdSiteIdDocumentProcessingJobId struct {
	GroupId                 string
	SiteId                  string
	DocumentProcessingJobId string
}

// NewGroupIdSiteIdDocumentProcessingJobID returns a new GroupIdSiteIdDocumentProcessingJobId struct
func NewGroupIdSiteIdDocumentProcessingJobID(groupId string, siteId string, documentProcessingJobId string) GroupIdSiteIdDocumentProcessingJobId {
	return GroupIdSiteIdDocumentProcessingJobId{
		GroupId:                 groupId,
		SiteId:                  siteId,
		DocumentProcessingJobId: documentProcessingJobId,
	}
}

// ParseGroupIdSiteIdDocumentProcessingJobID parses 'input' into a GroupIdSiteIdDocumentProcessingJobId
func ParseGroupIdSiteIdDocumentProcessingJobID(input string) (*GroupIdSiteIdDocumentProcessingJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdDocumentProcessingJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdDocumentProcessingJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdDocumentProcessingJobIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdDocumentProcessingJobId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdDocumentProcessingJobIDInsensitively(input string) (*GroupIdSiteIdDocumentProcessingJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdDocumentProcessingJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdDocumentProcessingJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdDocumentProcessingJobId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.DocumentProcessingJobId, ok = input.Parsed["documentProcessingJobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "documentProcessingJobId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdDocumentProcessingJobID checks that 'input' can be parsed as a Group Id Site Id Document Processing Job ID
func ValidateGroupIdSiteIdDocumentProcessingJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdDocumentProcessingJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Document Processing Job ID
func (id GroupIdSiteIdDocumentProcessingJobId) ID() string {
	fmtString := "/groups/%s/sites/%s/documentProcessingJobs/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.DocumentProcessingJobId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Document Processing Job ID
func (id GroupIdSiteIdDocumentProcessingJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("documentProcessingJobs", "documentProcessingJobs", "documentProcessingJobs"),
		resourceids.UserSpecifiedSegment("documentProcessingJobId", "documentProcessingJobId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Document Processing Job ID
func (id GroupIdSiteIdDocumentProcessingJobId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Document Processing Job: %q", id.DocumentProcessingJobId),
	}
	return fmt.Sprintf("Group Id Site Id Document Processing Job (%s)", strings.Join(components, "\n"))
}
