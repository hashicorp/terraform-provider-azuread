package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId{}

// UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId is a struct representing the Resource ID for a User Id Data Security And Governance Sensitivity Label Id Sublabel
type UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId struct {
	UserId              string
	SensitivityLabelId  string
	SensitivityLabelId1 string
}

// NewUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID returns a new UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId struct
func NewUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID(userId string, sensitivityLabelId string, sensitivityLabelId1 string) UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId {
	return UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId{
		UserId:              userId,
		SensitivityLabelId:  sensitivityLabelId,
		SensitivityLabelId1: sensitivityLabelId1,
	}
}

// ParseUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID parses 'input' into a UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId
func ParseUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID(input string) (*UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelIDInsensitively parses 'input' case-insensitively into a UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId
// note: this method should only be used for API response data and not user input
func ParseUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelIDInsensitively(input string) (*UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID checks that 'input' can be parsed as a User Id Data Security And Governance Sensitivity Label Id Sublabel ID
func ValidateUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Data Security And Governance Sensitivity Label Id Sublabel ID
func (id UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId) ID() string {
	fmtString := "/users/%s/dataSecurityAndGovernance/sensitivityLabels/%s/sublabels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SensitivityLabelId, id.SensitivityLabelId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Data Security And Governance Sensitivity Label Id Sublabel ID
func (id UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("dataSecurityAndGovernance", "dataSecurityAndGovernance", "dataSecurityAndGovernance"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
		resourceids.StaticSegment("sublabels", "sublabels", "sublabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId1", "sensitivityLabelId1"),
	}
}

// String returns a human-readable description of this User Id Data Security And Governance Sensitivity Label Id Sublabel ID
func (id UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
		fmt.Sprintf("Sensitivity Label Id 1: %q", id.SensitivityLabelId1),
	}
	return fmt.Sprintf("User Id Data Security And Governance Sensitivity Label Id Sublabel (%s)", strings.Join(components, "\n"))
}
