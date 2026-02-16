package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionThreatAssessmentRequestIdResultId{}

// UserIdInformationProtectionThreatAssessmentRequestIdResultId is a struct representing the Resource ID for a User Id Information Protection Threat Assessment Request Id Result
type UserIdInformationProtectionThreatAssessmentRequestIdResultId struct {
	UserId                    string
	ThreatAssessmentRequestId string
	ThreatAssessmentResultId  string
}

// NewUserIdInformationProtectionThreatAssessmentRequestIdResultID returns a new UserIdInformationProtectionThreatAssessmentRequestIdResultId struct
func NewUserIdInformationProtectionThreatAssessmentRequestIdResultID(userId string, threatAssessmentRequestId string, threatAssessmentResultId string) UserIdInformationProtectionThreatAssessmentRequestIdResultId {
	return UserIdInformationProtectionThreatAssessmentRequestIdResultId{
		UserId:                    userId,
		ThreatAssessmentRequestId: threatAssessmentRequestId,
		ThreatAssessmentResultId:  threatAssessmentResultId,
	}
}

// ParseUserIdInformationProtectionThreatAssessmentRequestIdResultID parses 'input' into a UserIdInformationProtectionThreatAssessmentRequestIdResultId
func ParseUserIdInformationProtectionThreatAssessmentRequestIdResultID(input string) (*UserIdInformationProtectionThreatAssessmentRequestIdResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionThreatAssessmentRequestIdResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionThreatAssessmentRequestIdResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdInformationProtectionThreatAssessmentRequestIdResultIDInsensitively parses 'input' case-insensitively into a UserIdInformationProtectionThreatAssessmentRequestIdResultId
// note: this method should only be used for API response data and not user input
func ParseUserIdInformationProtectionThreatAssessmentRequestIdResultIDInsensitively(input string) (*UserIdInformationProtectionThreatAssessmentRequestIdResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionThreatAssessmentRequestIdResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionThreatAssessmentRequestIdResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdInformationProtectionThreatAssessmentRequestIdResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ThreatAssessmentRequestId, ok = input.Parsed["threatAssessmentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", input)
	}

	if id.ThreatAssessmentResultId, ok = input.Parsed["threatAssessmentResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentResultId", input)
	}

	return nil
}

// ValidateUserIdInformationProtectionThreatAssessmentRequestIdResultID checks that 'input' can be parsed as a User Id Information Protection Threat Assessment Request Id Result ID
func ValidateUserIdInformationProtectionThreatAssessmentRequestIdResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdInformationProtectionThreatAssessmentRequestIdResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Information Protection Threat Assessment Request Id Result ID
func (id UserIdInformationProtectionThreatAssessmentRequestIdResultId) ID() string {
	fmtString := "/users/%s/informationProtection/threatAssessmentRequests/%s/results/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ThreatAssessmentRequestId, id.ThreatAssessmentResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Information Protection Threat Assessment Request Id Result ID
func (id UserIdInformationProtectionThreatAssessmentRequestIdResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("threatAssessmentRequests", "threatAssessmentRequests", "threatAssessmentRequests"),
		resourceids.UserSpecifiedSegment("threatAssessmentRequestId", "threatAssessmentRequestId"),
		resourceids.StaticSegment("results", "results", "results"),
		resourceids.UserSpecifiedSegment("threatAssessmentResultId", "threatAssessmentResultId"),
	}
}

// String returns a human-readable description of this User Id Information Protection Threat Assessment Request Id Result ID
func (id UserIdInformationProtectionThreatAssessmentRequestIdResultId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Threat Assessment Request: %q", id.ThreatAssessmentRequestId),
		fmt.Sprintf("Threat Assessment Result: %q", id.ThreatAssessmentResultId),
	}
	return fmt.Sprintf("User Id Information Protection Threat Assessment Request Id Result (%s)", strings.Join(components, "\n"))
}
