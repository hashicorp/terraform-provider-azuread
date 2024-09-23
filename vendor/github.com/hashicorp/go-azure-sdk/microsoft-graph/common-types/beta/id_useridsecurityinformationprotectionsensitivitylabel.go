package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdSecurityInformationProtectionSensitivityLabelId{}

// UserIdSecurityInformationProtectionSensitivityLabelId is a struct representing the Resource ID for a User Id Security Information Protection Sensitivity Label
type UserIdSecurityInformationProtectionSensitivityLabelId struct {
	UserId             string
	SensitivityLabelId string
}

// NewUserIdSecurityInformationProtectionSensitivityLabelID returns a new UserIdSecurityInformationProtectionSensitivityLabelId struct
func NewUserIdSecurityInformationProtectionSensitivityLabelID(userId string, sensitivityLabelId string) UserIdSecurityInformationProtectionSensitivityLabelId {
	return UserIdSecurityInformationProtectionSensitivityLabelId{
		UserId:             userId,
		SensitivityLabelId: sensitivityLabelId,
	}
}

// ParseUserIdSecurityInformationProtectionSensitivityLabelID parses 'input' into a UserIdSecurityInformationProtectionSensitivityLabelId
func ParseUserIdSecurityInformationProtectionSensitivityLabelID(input string) (*UserIdSecurityInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdSecurityInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdSecurityInformationProtectionSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdSecurityInformationProtectionSensitivityLabelIDInsensitively parses 'input' case-insensitively into a UserIdSecurityInformationProtectionSensitivityLabelId
// note: this method should only be used for API response data and not user input
func ParseUserIdSecurityInformationProtectionSensitivityLabelIDInsensitively(input string) (*UserIdSecurityInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdSecurityInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdSecurityInformationProtectionSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdSecurityInformationProtectionSensitivityLabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.SensitivityLabelId, ok = input.Parsed["sensitivityLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", input)
	}

	return nil
}

// ValidateUserIdSecurityInformationProtectionSensitivityLabelID checks that 'input' can be parsed as a User Id Security Information Protection Sensitivity Label ID
func ValidateUserIdSecurityInformationProtectionSensitivityLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdSecurityInformationProtectionSensitivityLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Security Information Protection Sensitivity Label ID
func (id UserIdSecurityInformationProtectionSensitivityLabelId) ID() string {
	fmtString := "/users/%s/security/informationProtection/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SensitivityLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Security Information Protection Sensitivity Label ID
func (id UserIdSecurityInformationProtectionSensitivityLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("security", "security", "security"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
	}
}

// String returns a human-readable description of this User Id Security Information Protection Sensitivity Label ID
func (id UserIdSecurityInformationProtectionSensitivityLabelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
	}
	return fmt.Sprintf("User Id Security Information Protection Sensitivity Label (%s)", strings.Join(components, "\n"))
}
