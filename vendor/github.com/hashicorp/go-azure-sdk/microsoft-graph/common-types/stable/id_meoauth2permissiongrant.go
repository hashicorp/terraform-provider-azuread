package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOAuth2PermissionGrantId{}

// MeOAuth2PermissionGrantId is a struct representing the Resource ID for a Me O Auth 2 Permission Grant
type MeOAuth2PermissionGrantId struct {
	OAuth2PermissionGrantId string
}

// NewMeOAuth2PermissionGrantID returns a new MeOAuth2PermissionGrantId struct
func NewMeOAuth2PermissionGrantID(oAuth2PermissionGrantId string) MeOAuth2PermissionGrantId {
	return MeOAuth2PermissionGrantId{
		OAuth2PermissionGrantId: oAuth2PermissionGrantId,
	}
}

// ParseMeOAuth2PermissionGrantID parses 'input' into a MeOAuth2PermissionGrantId
func ParseMeOAuth2PermissionGrantID(input string) (*MeOAuth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOAuth2PermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOAuth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOAuth2PermissionGrantIDInsensitively parses 'input' case-insensitively into a MeOAuth2PermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseMeOAuth2PermissionGrantIDInsensitively(input string) (*MeOAuth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOAuth2PermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOAuth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOAuth2PermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OAuth2PermissionGrantId, ok = input.Parsed["oAuth2PermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", input)
	}

	return nil
}

// ValidateMeOAuth2PermissionGrantID checks that 'input' can be parsed as a Me O Auth 2 Permission Grant ID
func ValidateMeOAuth2PermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOAuth2PermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me O Auth 2 Permission Grant ID
func (id MeOAuth2PermissionGrantId) ID() string {
	fmtString := "/me/oauth2PermissionGrants/%s"
	return fmt.Sprintf(fmtString, id.OAuth2PermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me O Auth 2 Permission Grant ID
func (id MeOAuth2PermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("oauth2PermissionGrants", "oauth2PermissionGrants", "oauth2PermissionGrants"),
		resourceids.UserSpecifiedSegment("oAuth2PermissionGrantId", "oAuth2PermissionGrantId"),
	}
}

// String returns a human-readable description of this Me O Auth 2 Permission Grant ID
func (id MeOAuth2PermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("O Auth 2 Permission Grant: %q", id.OAuth2PermissionGrantId),
	}
	return fmt.Sprintf("Me O Auth 2 Permission Grant (%s)", strings.Join(components, "\n"))
}
