package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdLicenseDetailId{}

// UserIdLicenseDetailId is a struct representing the Resource ID for a User Id License Detail
type UserIdLicenseDetailId struct {
	UserId           string
	LicenseDetailsId string
}

// NewUserIdLicenseDetailID returns a new UserIdLicenseDetailId struct
func NewUserIdLicenseDetailID(userId string, licenseDetailsId string) UserIdLicenseDetailId {
	return UserIdLicenseDetailId{
		UserId:           userId,
		LicenseDetailsId: licenseDetailsId,
	}
}

// ParseUserIdLicenseDetailID parses 'input' into a UserIdLicenseDetailId
func ParseUserIdLicenseDetailID(input string) (*UserIdLicenseDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdLicenseDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdLicenseDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdLicenseDetailIDInsensitively parses 'input' case-insensitively into a UserIdLicenseDetailId
// note: this method should only be used for API response data and not user input
func ParseUserIdLicenseDetailIDInsensitively(input string) (*UserIdLicenseDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdLicenseDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdLicenseDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdLicenseDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.LicenseDetailsId, ok = input.Parsed["licenseDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "licenseDetailsId", input)
	}

	return nil
}

// ValidateUserIdLicenseDetailID checks that 'input' can be parsed as a User Id License Detail ID
func ValidateUserIdLicenseDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdLicenseDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id License Detail ID
func (id UserIdLicenseDetailId) ID() string {
	fmtString := "/users/%s/licenseDetails/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.LicenseDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id License Detail ID
func (id UserIdLicenseDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("licenseDetails", "licenseDetails", "licenseDetails"),
		resourceids.UserSpecifiedSegment("licenseDetailsId", "licenseDetailsId"),
	}
}

// String returns a human-readable description of this User Id License Detail ID
func (id UserIdLicenseDetailId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("License Details: %q", id.LicenseDetailsId),
	}
	return fmt.Sprintf("User Id License Detail (%s)", strings.Join(components, "\n"))
}
