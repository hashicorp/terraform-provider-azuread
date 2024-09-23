package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAgreementAcceptanceId{}

// UserIdAgreementAcceptanceId is a struct representing the Resource ID for a User Id Agreement Acceptance
type UserIdAgreementAcceptanceId struct {
	UserId                string
	AgreementAcceptanceId string
}

// NewUserIdAgreementAcceptanceID returns a new UserIdAgreementAcceptanceId struct
func NewUserIdAgreementAcceptanceID(userId string, agreementAcceptanceId string) UserIdAgreementAcceptanceId {
	return UserIdAgreementAcceptanceId{
		UserId:                userId,
		AgreementAcceptanceId: agreementAcceptanceId,
	}
}

// ParseUserIdAgreementAcceptanceID parses 'input' into a UserIdAgreementAcceptanceId
func ParseUserIdAgreementAcceptanceID(input string) (*UserIdAgreementAcceptanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAgreementAcceptanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAgreementAcceptanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAgreementAcceptanceIDInsensitively parses 'input' case-insensitively into a UserIdAgreementAcceptanceId
// note: this method should only be used for API response data and not user input
func ParseUserIdAgreementAcceptanceIDInsensitively(input string) (*UserIdAgreementAcceptanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAgreementAcceptanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAgreementAcceptanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAgreementAcceptanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.AgreementAcceptanceId, ok = input.Parsed["agreementAcceptanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementAcceptanceId", input)
	}

	return nil
}

// ValidateUserIdAgreementAcceptanceID checks that 'input' can be parsed as a User Id Agreement Acceptance ID
func ValidateUserIdAgreementAcceptanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAgreementAcceptanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Agreement Acceptance ID
func (id UserIdAgreementAcceptanceId) ID() string {
	fmtString := "/users/%s/agreementAcceptances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AgreementAcceptanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Agreement Acceptance ID
func (id UserIdAgreementAcceptanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("agreementAcceptances", "agreementAcceptances", "agreementAcceptances"),
		resourceids.UserSpecifiedSegment("agreementAcceptanceId", "agreementAcceptanceId"),
	}
}

// String returns a human-readable description of this User Id Agreement Acceptance ID
func (id UserIdAgreementAcceptanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Agreement Acceptance: %q", id.AgreementAcceptanceId),
	}
	return fmt.Sprintf("User Id Agreement Acceptance (%s)", strings.Join(components, "\n"))
}
