package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

type PermissionScopeId struct {
	ApplicationId string
	ScopeID       string
}

func NewPermissionScopeID(applicationId, scopeId string) *PermissionScopeId {
	return &PermissionScopeId{
		ApplicationId: applicationId,
		ScopeID:       scopeId,
	}
}

// ParsePermissionScopeID parses 'input' into an PermissionScopeId
func ParsePermissionScopeID(input string) (*PermissionScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PermissionScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := &PermissionScopeId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.ScopeID, ok = parsed.Parsed["scopeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scopeId", *parsed)
	}

	return id, nil
}

// ValidatePermissionScopeID checks that 'input' can be parsed as an Application ID
func ValidatePermissionScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	id, err := ParsePermissionScopeID(v)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return validation.IsUUID(id.ScopeID, "ID")
}

func (id *PermissionScopeId) ID() string {
	fmtString := "/applications/%s/permissionScopes/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.ScopeID)
}

// Segments returns a slice of Resource ID Segments which comprise this ID
func (id *PermissionScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "00000000-0000-0000-0000-000000000000"),
		resourceids.StaticSegment("permissionScopes", "permissionScopes", "permissionScopes"),
		resourceids.UserSpecifiedSegment("scopeId", "11111111-1111-1111-1111-111111111111"),
	}
}

func (id *PermissionScopeId) String() string {
	return fmt.Sprintf("Permission Scope (Application ID: %q, Scope ID: %q)", id.ApplicationId, id.ScopeID)
}

func (id *PermissionScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.ScopeID, ok = input.Parsed["scopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scopeId", input)
	}

	return nil
}
