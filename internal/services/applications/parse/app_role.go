package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

type AppRoleId struct {
	ApplicationId string
	RoleID        string
}

func NewAppRoleID(applicationId, roleId string) *AppRoleId {
	return &AppRoleId{
		ApplicationId: applicationId,
		RoleID:        roleId,
	}
}

// ParseAppRoleID parses 'input' into an AppRoleId
func ParseAppRoleID(input string) (*AppRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AppRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := &AppRoleId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.RoleID, ok = parsed.Parsed["roleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "roleId", *parsed)
	}

	return id, nil
}

// ValidateAppRoleID checks that 'input' can be parsed as an Application ID
func ValidateAppRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	id, err := ParseAppRoleID(v)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return validation.IsUUID(id.RoleID, "ID")
}

func (id *AppRoleId) ID() string {
	fmtString := "/applications/%s/appRoles/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.RoleID)
}

// Segments returns a slice of Resource ID Segments which comprise this ID
func (id *AppRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "00000000-0000-0000-0000-000000000000"),
		resourceids.StaticSegment("appRoles", "appRoles", "appRoles"),
		resourceids.UserSpecifiedSegment("roleId", "11111111-1111-1111-1111-111111111111"),
	}
}

func (id *AppRoleId) String() string {
	return fmt.Sprintf("App Role (Application ID: %q, Role ID: %q)", id.ApplicationId, id.RoleID)
}

func (id *AppRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.RoleID, ok = input.Parsed["roleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleId", input)
	}

	return nil
}
