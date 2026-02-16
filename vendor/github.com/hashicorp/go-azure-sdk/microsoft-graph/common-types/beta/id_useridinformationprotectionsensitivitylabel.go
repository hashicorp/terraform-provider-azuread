package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionSensitivityLabelId{}

// UserIdInformationProtectionSensitivityLabelId is a struct representing the Resource ID for a User Id Information Protection Sensitivity Label
type UserIdInformationProtectionSensitivityLabelId struct {
	UserId             string
	SensitivityLabelId string
}

// NewUserIdInformationProtectionSensitivityLabelID returns a new UserIdInformationProtectionSensitivityLabelId struct
func NewUserIdInformationProtectionSensitivityLabelID(userId string, sensitivityLabelId string) UserIdInformationProtectionSensitivityLabelId {
	return UserIdInformationProtectionSensitivityLabelId{
		UserId:             userId,
		SensitivityLabelId: sensitivityLabelId,
	}
}

// ParseUserIdInformationProtectionSensitivityLabelID parses 'input' into a UserIdInformationProtectionSensitivityLabelId
func ParseUserIdInformationProtectionSensitivityLabelID(input string) (*UserIdInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdInformationProtectionSensitivityLabelIDInsensitively parses 'input' case-insensitively into a UserIdInformationProtectionSensitivityLabelId
// note: this method should only be used for API response data and not user input
func ParseUserIdInformationProtectionSensitivityLabelIDInsensitively(input string) (*UserIdInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdInformationProtectionSensitivityLabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.SensitivityLabelId, ok = input.Parsed["sensitivityLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", input)
	}

	return nil
}

// ValidateUserIdInformationProtectionSensitivityLabelID checks that 'input' can be parsed as a User Id Information Protection Sensitivity Label ID
func ValidateUserIdInformationProtectionSensitivityLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdInformationProtectionSensitivityLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Information Protection Sensitivity Label ID
func (id UserIdInformationProtectionSensitivityLabelId) ID() string {
	fmtString := "/users/%s/informationProtection/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SensitivityLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Information Protection Sensitivity Label ID
func (id UserIdInformationProtectionSensitivityLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
	}
}

// String returns a human-readable description of this User Id Information Protection Sensitivity Label ID
func (id UserIdInformationProtectionSensitivityLabelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
	}
	return fmt.Sprintf("User Id Information Protection Sensitivity Label (%s)", strings.Join(components, "\n"))
}
