package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

type FromTemplateId struct {
	TemplateId         string
	ApplicationId      string
	ServicePrincipalId string
}

func NewFromTemplateID(templateId, applicationId, servicePrincipalId string) *FromTemplateId {
	return &FromTemplateId{
		TemplateId:         templateId,
		ApplicationId:      applicationId,
		ServicePrincipalId: servicePrincipalId,
	}
}

// ParseFromTemplateID parses 'input' into an FromTemplateId
func ParseFromTemplateID(input string) (*FromTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FromTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := &FromTemplateId{}

	if id.TemplateId, ok = parsed.Parsed["templateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "templateId", *parsed)
	}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	return id, nil
}

// ValidateFromTemplateID checks that 'input' can be parsed as an Application ID
func ValidateFromTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	id, err := ParseFromTemplateID(v)
	if err != nil {
		errors = append(errors, err)
		return
	}

	if warnings, errors = validation.IsUUID(id.TemplateId, "ID"); len(errors) > 0 {
		return
	}

	if warnings, errors = validation.IsUUID(id.ApplicationId, "ID"); len(errors) > 0 {
		return
	}

	if warnings, errors = validation.IsUUID(id.ServicePrincipalId, "ID"); len(errors) > 0 {
		return
	}

	return
}

func (id *FromTemplateId) ID() string {
	fmtString := "/applicationTemplates/%s/instantiate/%s/%s"
	return fmt.Sprintf(fmtString, id.TemplateId, id.ApplicationId, id.ServicePrincipalId)
}

// Segments returns a slice of Resource ID Segments which comprise this ID
func (id *FromTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applicationTemplates", "applicationTemplates", "applicationTemplates"),
		resourceids.UserSpecifiedSegment("templateId", "00000000-0000-0000-0000-000000000000"),
		resourceids.StaticSegment("instantiate", "instantiate", "instantiate"),
		resourceids.UserSpecifiedSegment("applicationId", "11111111-1111-1111-1111-111111111111"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "22222222-2222-2222-2222-222222222222"),
	}
}

func (id *FromTemplateId) String() string {
	return fmt.Sprintf("Application From Template (Template ID: %q, Application ID: %q, Service Principal ID: %q)", id.TemplateId, id.ApplicationId, id.ServicePrincipalId)
}

func (id *FromTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TemplateId, ok = input.Parsed["templateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "templateId", input)
	}

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	return nil
}
