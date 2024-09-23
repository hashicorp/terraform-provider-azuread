package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdItemIdDocumentSetVersionId{}

// GroupIdSiteIdListIdItemIdDocumentSetVersionId is a struct representing the Resource ID for a Group Id Site Id List Id Item Id Document Set Version
type GroupIdSiteIdListIdItemIdDocumentSetVersionId struct {
	GroupId              string
	SiteId               string
	ListId               string
	ListItemId           string
	DocumentSetVersionId string
}

// NewGroupIdSiteIdListIdItemIdDocumentSetVersionID returns a new GroupIdSiteIdListIdItemIdDocumentSetVersionId struct
func NewGroupIdSiteIdListIdItemIdDocumentSetVersionID(groupId string, siteId string, listId string, listItemId string, documentSetVersionId string) GroupIdSiteIdListIdItemIdDocumentSetVersionId {
	return GroupIdSiteIdListIdItemIdDocumentSetVersionId{
		GroupId:              groupId,
		SiteId:               siteId,
		ListId:               listId,
		ListItemId:           listItemId,
		DocumentSetVersionId: documentSetVersionId,
	}
}

// ParseGroupIdSiteIdListIdItemIdDocumentSetVersionID parses 'input' into a GroupIdSiteIdListIdItemIdDocumentSetVersionId
func ParseGroupIdSiteIdListIdItemIdDocumentSetVersionID(input string) (*GroupIdSiteIdListIdItemIdDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdItemIdDocumentSetVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdItemIdDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdItemIdDocumentSetVersionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdItemIdDocumentSetVersionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdItemIdDocumentSetVersionIDInsensitively(input string) (*GroupIdSiteIdListIdItemIdDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdItemIdDocumentSetVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdItemIdDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdItemIdDocumentSetVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.ListId, ok = input.Parsed["listId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listId", input)
	}

	if id.ListItemId, ok = input.Parsed["listItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemId", input)
	}

	if id.DocumentSetVersionId, ok = input.Parsed["documentSetVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "documentSetVersionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdListIdItemIdDocumentSetVersionID checks that 'input' can be parsed as a Group Id Site Id List Id Item Id Document Set Version ID
func ValidateGroupIdSiteIdListIdItemIdDocumentSetVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdItemIdDocumentSetVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Item Id Document Set Version ID
func (id GroupIdSiteIdListIdItemIdDocumentSetVersionId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/items/%s/documentSetVersions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ListItemId, id.DocumentSetVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Item Id Document Set Version ID
func (id GroupIdSiteIdListIdItemIdDocumentSetVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("documentSetVersions", "documentSetVersions", "documentSetVersions"),
		resourceids.UserSpecifiedSegment("documentSetVersionId", "documentSetVersionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Item Id Document Set Version ID
func (id GroupIdSiteIdListIdItemIdDocumentSetVersionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Document Set Version: %q", id.DocumentSetVersionId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Item Id Document Set Version (%s)", strings.Join(components, "\n"))
}
