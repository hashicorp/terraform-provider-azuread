package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDataSecurityAndGovernanceSensitivityLabelId{}

// UserIdDataSecurityAndGovernanceSensitivityLabelId is a struct representing the Resource ID for a User Id Data Security And Governance Sensitivity Label
type UserIdDataSecurityAndGovernanceSensitivityLabelId struct {
	UserId             string
	SensitivityLabelId string
}

// NewUserIdDataSecurityAndGovernanceSensitivityLabelID returns a new UserIdDataSecurityAndGovernanceSensitivityLabelId struct
func NewUserIdDataSecurityAndGovernanceSensitivityLabelID(userId string, sensitivityLabelId string) UserIdDataSecurityAndGovernanceSensitivityLabelId {
	return UserIdDataSecurityAndGovernanceSensitivityLabelId{
		UserId:             userId,
		SensitivityLabelId: sensitivityLabelId,
	}
}

// ParseUserIdDataSecurityAndGovernanceSensitivityLabelID parses 'input' into a UserIdDataSecurityAndGovernanceSensitivityLabelId
func ParseUserIdDataSecurityAndGovernanceSensitivityLabelID(input string) (*UserIdDataSecurityAndGovernanceSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDataSecurityAndGovernanceSensitivityLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDataSecurityAndGovernanceSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDataSecurityAndGovernanceSensitivityLabelIDInsensitively parses 'input' case-insensitively into a UserIdDataSecurityAndGovernanceSensitivityLabelId
// note: this method should only be used for API response data and not user input
func ParseUserIdDataSecurityAndGovernanceSensitivityLabelIDInsensitively(input string) (*UserIdDataSecurityAndGovernanceSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDataSecurityAndGovernanceSensitivityLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDataSecurityAndGovernanceSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDataSecurityAndGovernanceSensitivityLabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.SensitivityLabelId, ok = input.Parsed["sensitivityLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", input)
	}

	return nil
}

// ValidateUserIdDataSecurityAndGovernanceSensitivityLabelID checks that 'input' can be parsed as a User Id Data Security And Governance Sensitivity Label ID
func ValidateUserIdDataSecurityAndGovernanceSensitivityLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDataSecurityAndGovernanceSensitivityLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Data Security And Governance Sensitivity Label ID
func (id UserIdDataSecurityAndGovernanceSensitivityLabelId) ID() string {
	fmtString := "/users/%s/dataSecurityAndGovernance/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SensitivityLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Data Security And Governance Sensitivity Label ID
func (id UserIdDataSecurityAndGovernanceSensitivityLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("dataSecurityAndGovernance", "dataSecurityAndGovernance", "dataSecurityAndGovernance"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
	}
}

// String returns a human-readable description of this User Id Data Security And Governance Sensitivity Label ID
func (id UserIdDataSecurityAndGovernanceSensitivityLabelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
	}
	return fmt.Sprintf("User Id Data Security And Governance Sensitivity Label (%s)", strings.Join(components, "\n"))
}
