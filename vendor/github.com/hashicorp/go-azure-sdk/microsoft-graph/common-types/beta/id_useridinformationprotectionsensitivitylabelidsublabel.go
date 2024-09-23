package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionSensitivityLabelIdSublabelId{}

// UserIdInformationProtectionSensitivityLabelIdSublabelId is a struct representing the Resource ID for a User Id Information Protection Sensitivity Label Id Sublabel
type UserIdInformationProtectionSensitivityLabelIdSublabelId struct {
	UserId              string
	SensitivityLabelId  string
	SensitivityLabelId1 string
}

// NewUserIdInformationProtectionSensitivityLabelIdSublabelID returns a new UserIdInformationProtectionSensitivityLabelIdSublabelId struct
func NewUserIdInformationProtectionSensitivityLabelIdSublabelID(userId string, sensitivityLabelId string, sensitivityLabelId1 string) UserIdInformationProtectionSensitivityLabelIdSublabelId {
	return UserIdInformationProtectionSensitivityLabelIdSublabelId{
		UserId:              userId,
		SensitivityLabelId:  sensitivityLabelId,
		SensitivityLabelId1: sensitivityLabelId1,
	}
}

// ParseUserIdInformationProtectionSensitivityLabelIdSublabelID parses 'input' into a UserIdInformationProtectionSensitivityLabelIdSublabelId
func ParseUserIdInformationProtectionSensitivityLabelIdSublabelID(input string) (*UserIdInformationProtectionSensitivityLabelIdSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionSensitivityLabelIdSublabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionSensitivityLabelIdSublabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdInformationProtectionSensitivityLabelIdSublabelIDInsensitively parses 'input' case-insensitively into a UserIdInformationProtectionSensitivityLabelIdSublabelId
// note: this method should only be used for API response data and not user input
func ParseUserIdInformationProtectionSensitivityLabelIdSublabelIDInsensitively(input string) (*UserIdInformationProtectionSensitivityLabelIdSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionSensitivityLabelIdSublabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionSensitivityLabelIdSublabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdInformationProtectionSensitivityLabelIdSublabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.SensitivityLabelId, ok = input.Parsed["sensitivityLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", input)
	}

	if id.SensitivityLabelId1, ok = input.Parsed["sensitivityLabelId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId1", input)
	}

	return nil
}

// ValidateUserIdInformationProtectionSensitivityLabelIdSublabelID checks that 'input' can be parsed as a User Id Information Protection Sensitivity Label Id Sublabel ID
func ValidateUserIdInformationProtectionSensitivityLabelIdSublabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdInformationProtectionSensitivityLabelIdSublabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Information Protection Sensitivity Label Id Sublabel ID
func (id UserIdInformationProtectionSensitivityLabelIdSublabelId) ID() string {
	fmtString := "/users/%s/informationProtection/sensitivityLabels/%s/sublabels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SensitivityLabelId, id.SensitivityLabelId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Information Protection Sensitivity Label Id Sublabel ID
func (id UserIdInformationProtectionSensitivityLabelIdSublabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
		resourceids.StaticSegment("sublabels", "sublabels", "sublabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId1", "sensitivityLabelId1"),
	}
}

// String returns a human-readable description of this User Id Information Protection Sensitivity Label Id Sublabel ID
func (id UserIdInformationProtectionSensitivityLabelIdSublabelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
		fmt.Sprintf("Sensitivity Label Id 1: %q", id.SensitivityLabelId1),
	}
	return fmt.Sprintf("User Id Information Protection Sensitivity Label Id Sublabel (%s)", strings.Join(components, "\n"))
}
