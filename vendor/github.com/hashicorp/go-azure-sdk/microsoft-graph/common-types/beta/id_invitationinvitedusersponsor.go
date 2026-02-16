package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &InvitationInvitedUserSponsorId{}

// InvitationInvitedUserSponsorId is a struct representing the Resource ID for a Invitation Invited User Sponsor
type InvitationInvitedUserSponsorId struct {
	DirectoryObjectId string
}

// NewInvitationInvitedUserSponsorID returns a new InvitationInvitedUserSponsorId struct
func NewInvitationInvitedUserSponsorID(directoryObjectId string) InvitationInvitedUserSponsorId {
	return InvitationInvitedUserSponsorId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseInvitationInvitedUserSponsorID parses 'input' into a InvitationInvitedUserSponsorId
func ParseInvitationInvitedUserSponsorID(input string) (*InvitationInvitedUserSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvitationInvitedUserSponsorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvitationInvitedUserSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInvitationInvitedUserSponsorIDInsensitively parses 'input' case-insensitively into a InvitationInvitedUserSponsorId
// note: this method should only be used for API response data and not user input
func ParseInvitationInvitedUserSponsorIDInsensitively(input string) (*InvitationInvitedUserSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvitationInvitedUserSponsorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvitationInvitedUserSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InvitationInvitedUserSponsorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateInvitationInvitedUserSponsorID checks that 'input' can be parsed as a Invitation Invited User Sponsor ID
func ValidateInvitationInvitedUserSponsorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInvitationInvitedUserSponsorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Invitation Invited User Sponsor ID
func (id InvitationInvitedUserSponsorId) ID() string {
	fmtString := "/invitations/invitedUserSponsors/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Invitation Invited User Sponsor ID
func (id InvitationInvitedUserSponsorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("invitations", "invitations", "invitations"),
		resourceids.StaticSegment("invitedUserSponsors", "invitedUserSponsors", "invitedUserSponsors"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Invitation Invited User Sponsor ID
func (id InvitationInvitedUserSponsorId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Invitation Invited User Sponsor (%s)", strings.Join(components, "\n"))
}
