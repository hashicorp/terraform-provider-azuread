package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootListItemDocumentSetVersionId{}

// UserIdDriveIdRootListItemDocumentSetVersionId is a struct representing the Resource ID for a User Id Drive Id Root List Item Document Set Version
type UserIdDriveIdRootListItemDocumentSetVersionId struct {
	UserId               string
	DriveId              string
	DocumentSetVersionId string
}

// NewUserIdDriveIdRootListItemDocumentSetVersionID returns a new UserIdDriveIdRootListItemDocumentSetVersionId struct
func NewUserIdDriveIdRootListItemDocumentSetVersionID(userId string, driveId string, documentSetVersionId string) UserIdDriveIdRootListItemDocumentSetVersionId {
	return UserIdDriveIdRootListItemDocumentSetVersionId{
		UserId:               userId,
		DriveId:              driveId,
		DocumentSetVersionId: documentSetVersionId,
	}
}

// ParseUserIdDriveIdRootListItemDocumentSetVersionID parses 'input' into a UserIdDriveIdRootListItemDocumentSetVersionId
func ParseUserIdDriveIdRootListItemDocumentSetVersionID(input string) (*UserIdDriveIdRootListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootListItemDocumentSetVersionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootListItemDocumentSetVersionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootListItemDocumentSetVersionIDInsensitively(input string) (*UserIdDriveIdRootListItemDocumentSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootListItemDocumentSetVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootListItemDocumentSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootListItemDocumentSetVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DocumentSetVersionId, ok = input.Parsed["documentSetVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "documentSetVersionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdRootListItemDocumentSetVersionID checks that 'input' can be parsed as a User Id Drive Id Root List Item Document Set Version ID
func ValidateUserIdDriveIdRootListItemDocumentSetVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootListItemDocumentSetVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root List Item Document Set Version ID
func (id UserIdDriveIdRootListItemDocumentSetVersionId) ID() string {
	fmtString := "/users/%s/drives/%s/root/listItem/documentSetVersions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DocumentSetVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root List Item Document Set Version ID
func (id UserIdDriveIdRootListItemDocumentSetVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("documentSetVersions", "documentSetVersions", "documentSetVersions"),
		resourceids.UserSpecifiedSegment("documentSetVersionId", "documentSetVersionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root List Item Document Set Version ID
func (id UserIdDriveIdRootListItemDocumentSetVersionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Document Set Version: %q", id.DocumentSetVersionId),
	}
	return fmt.Sprintf("User Id Drive Id Root List Item Document Set Version (%s)", strings.Join(components, "\n"))
}
