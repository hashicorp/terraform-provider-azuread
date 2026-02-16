package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &OAuth2PermissionGrantId{}

// OAuth2PermissionGrantId is a struct representing the Resource ID for a O Auth 2 Permission Grant
type OAuth2PermissionGrantId struct {
	OAuth2PermissionGrantId string
}

// NewOAuth2PermissionGrantID returns a new OAuth2PermissionGrantId struct
func NewOAuth2PermissionGrantID(oAuth2PermissionGrantId string) OAuth2PermissionGrantId {
	return OAuth2PermissionGrantId{
		OAuth2PermissionGrantId: oAuth2PermissionGrantId,
	}
}

// ParseOAuth2PermissionGrantID parses 'input' into a OAuth2PermissionGrantId
func ParseOAuth2PermissionGrantID(input string) (*OAuth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OAuth2PermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OAuth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseOAuth2PermissionGrantIDInsensitively parses 'input' case-insensitively into a OAuth2PermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseOAuth2PermissionGrantIDInsensitively(input string) (*OAuth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OAuth2PermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OAuth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *OAuth2PermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OAuth2PermissionGrantId, ok = input.Parsed["oAuth2PermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", input)
	}

	return nil
}

// ValidateOAuth2PermissionGrantID checks that 'input' can be parsed as a O Auth 2 Permission Grant ID
func ValidateOAuth2PermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOAuth2PermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted O Auth 2 Permission Grant ID
func (id OAuth2PermissionGrantId) ID() string {
	fmtString := "/oauth2PermissionGrants/%s"
	return fmt.Sprintf(fmtString, id.OAuth2PermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this O Auth 2 Permission Grant ID
func (id OAuth2PermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("oauth2PermissionGrants", "oauth2PermissionGrants", "oauth2PermissionGrants"),
		resourceids.UserSpecifiedSegment("oAuth2PermissionGrantId", "oAuth2PermissionGrantId"),
	}
}

// String returns a human-readable description of this O Auth 2 Permission Grant ID
func (id OAuth2PermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("O Auth 2 Permission Grant: %q", id.OAuth2PermissionGrantId),
	}
	return fmt.Sprintf("O Auth 2 Permission Grant (%s)", strings.Join(components, "\n"))
}
