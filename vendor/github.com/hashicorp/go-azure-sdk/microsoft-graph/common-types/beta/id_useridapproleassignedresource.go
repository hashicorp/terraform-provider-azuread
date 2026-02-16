package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAppRoleAssignedResourceId{}

// UserIdAppRoleAssignedResourceId is a struct representing the Resource ID for a User Id App Role Assigned Resource
type UserIdAppRoleAssignedResourceId struct {
	UserId             string
	ServicePrincipalId string
}

// NewUserIdAppRoleAssignedResourceID returns a new UserIdAppRoleAssignedResourceId struct
func NewUserIdAppRoleAssignedResourceID(userId string, servicePrincipalId string) UserIdAppRoleAssignedResourceId {
	return UserIdAppRoleAssignedResourceId{
		UserId:             userId,
		ServicePrincipalId: servicePrincipalId,
	}
}

// ParseUserIdAppRoleAssignedResourceID parses 'input' into a UserIdAppRoleAssignedResourceId
func ParseUserIdAppRoleAssignedResourceID(input string) (*UserIdAppRoleAssignedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAppRoleAssignedResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAppRoleAssignedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAppRoleAssignedResourceIDInsensitively parses 'input' case-insensitively into a UserIdAppRoleAssignedResourceId
// note: this method should only be used for API response data and not user input
func ParseUserIdAppRoleAssignedResourceIDInsensitively(input string) (*UserIdAppRoleAssignedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAppRoleAssignedResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAppRoleAssignedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAppRoleAssignedResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	return nil
}

// ValidateUserIdAppRoleAssignedResourceID checks that 'input' can be parsed as a User Id App Role Assigned Resource ID
func ValidateUserIdAppRoleAssignedResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAppRoleAssignedResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id App Role Assigned Resource ID
func (id UserIdAppRoleAssignedResourceId) ID() string {
	fmtString := "/users/%s/appRoleAssignedResources/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ServicePrincipalId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id App Role Assigned Resource ID
func (id UserIdAppRoleAssignedResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("appRoleAssignedResources", "appRoleAssignedResources", "appRoleAssignedResources"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
	}
}

// String returns a human-readable description of this User Id App Role Assigned Resource ID
func (id UserIdAppRoleAssignedResourceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
	}
	return fmt.Sprintf("User Id App Role Assigned Resource (%s)", strings.Join(components, "\n"))
}
