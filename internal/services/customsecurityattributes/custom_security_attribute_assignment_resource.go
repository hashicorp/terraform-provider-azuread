// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package customsecurityattributes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	sdkClient "github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func customSecurityAttributeAssignmentResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: customSecurityAttributeAssignmentResourceCreateUpdate,
		ReadContext:   customSecurityAttributeAssignmentResourceRead,
		UpdateContext: customSecurityAttributeAssignmentResourceCreateUpdate,
		DeleteContext: customSecurityAttributeAssignmentResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(10 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(10 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		CustomizeDiff: customSecurityAttributeAssignmentDiff,

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			parts := strings.SplitN(id, "/", 2)
			if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
				return fmt.Errorf("expected import ID in format <principal_id>/<attribute_set>, got %q", id)
			}
			if _, errs := validation.IsUUID(parts[0], "principal_id"); len(errs) > 0 {
				return fmt.Errorf("principal_id %q is not a valid UUID", parts[0])
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"principal_id": {
				Description:  "The object ID of the principal (user, group, or service principal) to which the custom security attributes are assigned",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"attribute_set": {
				Description:  "The name of the attribute set",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"attribute": {
				Description: "One or more attribute blocks within this attribute set",
				Type:        pluginsdk.TypeSet,
				Required:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"name": {
							Description:  "The name of the custom security attribute",
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},

						"value": {
							Description: "A single string value for the attribute. Mutually exclusive with values and boolean_value",
							Type:        pluginsdk.TypeString,
							Optional:    true,
							Computed:    true,
						},

						"values": {
							Description: "A list of string values for the attribute (multi-value). Mutually exclusive with value and boolean_value",
							Type:        pluginsdk.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
						},

						"boolean_value": {
							Description: "A boolean value for the attribute. Mutually exclusive with value and values",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

// customSecurityAttributeAssignmentDiff validates that each attribute block sets
// exactly one of value, values, or boolean_value. This is done in CustomizeDiff
// because ExactlyOneOf is not supported inside TypeSet nested blocks in the plugin SDK.
//
// We use GetRawConfig() instead of GetOk() because boolean_value=false is Go's zero
// value and indistinguishable from "not set" via the normal schema API.
func customSecurityAttributeAssignmentDiff(_ context.Context, d *pluginsdk.ResourceDiff, _ interface{}) error {
	raw := d.GetRawConfig()
	if raw.IsNull() || !raw.IsKnown() {
		return nil
	}

	attrVal := raw.GetAttr("attribute")
	if attrVal.IsNull() || !attrVal.IsKnown() {
		return nil
	}

	for it := attrVal.ElementIterator(); it.Next(); {
		_, elem := it.Element()
		if elem.IsNull() || !elem.IsKnown() {
			continue
		}

		nameVal := elem.GetAttr("name")
		name := ""
		if !nameVal.IsNull() && nameVal.IsKnown() {
			name = nameVal.AsString()
		}

		valueVal := elem.GetAttr("value")
		valuesVal := elem.GetAttr("values")
		boolVal := elem.GetAttr("boolean_value")

		// A field is "set" if it is non-null and known in the raw config.
		// This correctly handles boolean_value=false (non-null, known, valid).
		hasValue := !valueVal.IsNull() && valueVal.IsKnown() && valueVal.AsString() != ""
		hasValues := !valuesVal.IsNull() && valuesVal.IsKnown() && valuesVal.LengthInt() > 0
		hasBool := !boolVal.IsNull() && boolVal.IsKnown()

		setCount := 0
		if hasValue {
			setCount++
		}
		if hasValues {
			setCount++
		}
		if hasBool {
			setCount++
		}

		if setCount == 0 {
			return fmt.Errorf("attribute %q: exactly one of value, values, or boolean_value must be set", name)
		}
		if setCount > 1 {
			return fmt.Errorf("attribute %q: only one of value, values, or boolean_value may be set", name)
		}
	}

	return nil
}

// expandCustomSecurityAttributes converts the Terraform schema into the JSON payload
// the Graph API expects. The Graph API represents custom security attributes as:
//
//	{
//	  "customSecurityAttributes": {
//	    "<attributeSet>": {
//	      "@odata.type": "#microsoft.graph.customSecurityAttributeValue",
//	      "<name>": "singleValue",
//	      "<name>@odata.type": "#Collection(String)",
//	      "<name>": ["val1", "val2"],
//	      "<name>": true
//	    }
//	  }
//	}
func expandCustomSecurityAttributes(attributeSet string, input []interface{}) (map[string]interface{}, error) {
	setPayload := map[string]interface{}{
		"@odata.type": "#microsoft.graph.customSecurityAttributeValue",
	}

	for _, rawAttr := range input {
		attrMap, ok := rawAttr.(map[string]interface{})
		if !ok {
			continue
		}

		name := attrMap["name"].(string)
		if name == "" {
			return nil, fmt.Errorf("attribute name must not be empty in attribute_set %q", attributeSet)
		}

		singleVal := attrMap["value"].(string)
		multiValRaw, _ := attrMap["values"].([]interface{})
		boolVal, _ := attrMap["boolean_value"].(bool)

		// Route to the correct Graph API type.
		// CustomizeDiff guarantees exactly one field is meaningfully set.
		// boolean_value=false is valid and must be sent explicitly.
		switch {
		case len(multiValRaw) > 0:
			vals := make([]string, 0, len(multiValRaw))
			for _, v := range multiValRaw {
				vals = append(vals, v.(string))
			}
			setPayload[name+"@odata.type"] = "#Collection(String)"
			setPayload[name] = vals

		case singleVal != "":
			setPayload[name] = singleVal

		default:
			// boolean_value — send true or false explicitly (including false)
			setPayload[name] = boolVal
		}
	}

	return map[string]interface{}{
		attributeSet: setPayload,
	}, nil
}

// flattenAttributes converts the attribute set map from the Graph API back into
// the flat list of attribute blocks used in the schema.
func flattenAttributes(setVal map[string]interface{}) []interface{} {
	attrs := make([]interface{}, 0)

	for key, val := range setVal {
		if strings.HasPrefix(key, "@") || strings.HasSuffix(key, "@odata.type") {
			continue
		}

		attrMap := map[string]interface{}{
			"name":          key,
			"value":         "",
			"values":        []interface{}{},
			"boolean_value": false,
		}

		switch v := val.(type) {
		case bool:
			attrMap["boolean_value"] = v

		case []interface{}:
			strVals := make([]interface{}, 0, len(v))
			for _, sv := range v {
				strVals = append(strVals, fmt.Sprintf("%v", sv))
			}
			attrMap["values"] = strVals

		default:
			attrMap["value"] = fmt.Sprintf("%v", v)
		}

		attrs = append(attrs, attrMap)
	}

	return attrs
}

// patchCSA is intentionally unused — PATCH is done inline via raw HTTP in create/update/delete.

// selectOptions implements sdkClient.Options to append a $select query parameter.
type selectOptions struct {
	fields []string
}

func (o *selectOptions) ToHeaders() *sdkClient.Headers { return nil }
func (o *selectOptions) ToOData() *odata.Query {
	return &odata.Query{Select: o.fields}
}
func (o *selectOptions) ToQuery() *sdkClient.QueryParams { return nil }

// parseResourceID splits the composite resource ID "<principalId>/<attributeSet>" into its parts.
func parseResourceID(id string) (principalId string, attributeSet string, err error) {
	parts := strings.SplitN(id, "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("expected resource ID in format <principal_id>/<attribute_set>, got %q", id)
	}
	return parts[0], parts[1], nil
}

// resolveObjectPath does a GET on /directoryObjects/{id} to read the @odata.type,
// then returns the type-specific path (e.g. /servicePrincipals/{id}, /users/{id}, /groups/{id}).
// This is required because /directoryObjects/{id} does not support PATCH of customSecurityAttributes.
func resolveObjectPath(ctx context.Context, c *msgraph.Client, objectId string) (string, pluginsdk.Diagnostics) {
	directoryObjectPath := fmt.Sprintf("/directoryObjects/%s", objectId)

	req, err := c.NewRequest(ctx, sdkClient.RequestOptions{
		ContentType:         "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                directoryObjectPath,
		OptionsObject:       &selectOptions{fields: []string{"id"}},
	})
	if err != nil {
		return "", tf.ErrorDiagF(err, "Building GET request to resolve object type for %q", objectId)
	}

	httpResp, err := req.Execute(ctx)
	if err != nil {
		return "", tf.ErrorDiagF(err, "Resolving object type for principal %q", objectId)
	}
	if httpResp == nil || httpResp.Response == nil {
		return "", tf.ErrorDiagF(errors.New("response was nil"), "Resolving object type for principal %q", objectId)
	}

	var body map[string]interface{}
	if err := httpResp.Unmarshal(&body); err != nil {
		return "", tf.ErrorDiagF(err, "Decoding object type response for principal %q", objectId)
	}

	odataType, _ := body["@odata.type"].(string)
	log.Printf("[DEBUG] Resolved @odata.type for principal %q: %s", objectId, odataType)

	switch {
	case strings.EqualFold(odataType, "#microsoft.graph.servicePrincipal"):
		return fmt.Sprintf("/servicePrincipals/%s", objectId), nil
	case strings.EqualFold(odataType, "#microsoft.graph.user"):
		return fmt.Sprintf("/users/%s", objectId), nil
	case strings.EqualFold(odataType, "#microsoft.graph.group"):
		return fmt.Sprintf("/groups/%s", objectId), nil
	default:
		return "", tf.ErrorDiagF(
			fmt.Errorf("unsupported object type %q — must be servicePrincipal, user, or group", odataType),
			"Resolving object type for principal %q", objectId,
		)
	}
}

func customSecurityAttributeAssignmentResourceCreateUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	spClient := meta.(*clients.Client).CustomSecurityAttributes.ServicePrincipalClientBeta

	principalId := d.Get("principal_id").(string)
	attributeSet := d.Get("attribute_set").(string)
	rawAttrs := d.Get("attribute").(*pluginsdk.Set).List()

	csaPayload, err := expandCustomSecurityAttributes(attributeSet, rawAttrs)
	if err != nil {
		return tf.ErrorDiagF(err, "Building custom security attributes payload")
	}

	payload := map[string]interface{}{
		"customSecurityAttributes": csaPayload,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return tf.ErrorDiagF(err, "Marshalling custom security attributes")
	}

	log.Printf("[DEBUG] Custom security attributes PATCH payload for principal %q: %s", principalId, string(payloadJSON))

	// Resolve the object type so we can PATCH the correct type-specific endpoint.
	// /directoryObjects/{id} does not support customSecurityAttributes in PATCH.
	objectPath, diags := resolveObjectPath(ctx, spClient.Client, principalId)
	if diags != nil {
		return diags
	}

	req, err := spClient.Client.NewRequest(ctx, sdkClient.RequestOptions{
		ContentType:         "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{http.StatusNoContent, http.StatusOK},
		HttpMethod:          http.MethodPatch,
		Path:                objectPath,
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Building PATCH request for principal %q", principalId)
	}

	if err = req.Marshal(payload); err != nil {
		return tf.ErrorDiagF(err, "Marshalling PATCH payload for principal %q", principalId)
	}

	httpResp, err := req.Execute(ctx)
	if err != nil {
		return tf.ErrorDiagF(err, "Setting custom security attributes on principal %q", principalId)
	}
	if httpResp == nil || httpResp.Response == nil {
		return tf.ErrorDiagF(errors.New("response was nil"), "Setting custom security attributes on principal %q", principalId)
	}

	log.Printf("[DEBUG] Successfully PATCHed custom security attributes on principal %q (HTTP %d)", principalId, httpResp.Response.StatusCode)

	// Encode both principal_id and attribute_set into the resource ID so import works
	d.SetId(principalId + "/" + attributeSet)

	return customSecurityAttributeAssignmentResourceRead(ctx, d, meta)
}

func customSecurityAttributeAssignmentResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	spClient := meta.(*clients.Client).CustomSecurityAttributes.ServicePrincipalClientBeta

	principalId, attributeSet, err := parseResourceID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing resource ID")
	}

	objectPath, diags := resolveObjectPath(ctx, spClient.Client, principalId)
	if diags != nil {
		return diags
	}

	req, err := spClient.Client.NewRequest(ctx, sdkClient.RequestOptions{
		ContentType:         "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                objectPath,
		OptionsObject:       &selectOptions{fields: []string{"id", "customSecurityAttributes"}},
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Building GET request for principal %q", principalId)
	}

	httpResp, err := req.Execute(ctx)
	if err != nil {
		if httpResp != nil && response.WasNotFound(httpResp.Response) {
			log.Printf("[DEBUG] Principal %q was not found - removing from state!", principalId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving custom security attributes for principal %q", principalId)
	}

	if httpResp == nil || httpResp.Response == nil {
		return tf.ErrorDiagF(errors.New("response was nil"), "Retrieving principal %q", principalId)
	}

	var body map[string]interface{}
	if err := httpResp.Unmarshal(&body); err != nil {
		return tf.ErrorDiagF(err, "Decoding response for principal %q", principalId)
	}

	rawCSAVal, ok := body["customSecurityAttributes"]
	if !ok || rawCSAVal == nil {
		log.Printf("[DEBUG] customSecurityAttributes not present in response for principal %q (may lack read permission)", principalId)
		tf.Set(d, "principal_id", principalId)
		tf.Set(d, "attribute_set", attributeSet)
		return nil
	}

	rawCSA, ok := rawCSAVal.(map[string]interface{})
	if !ok {
		return tf.ErrorDiagF(fmt.Errorf("unexpected type for customSecurityAttributes"), "Parsing response for principal %q", principalId)
	}

	rawJSON, _ := json.Marshal(rawCSA)
	log.Printf("[DEBUG] Custom security attributes read back for principal %q: %s", principalId, string(rawJSON))

	tf.Set(d, "principal_id", principalId)
	tf.Set(d, "attribute_set", attributeSet)

	if setVal, ok := rawCSA[attributeSet].(map[string]interface{}); ok {
		tf.Set(d, "attribute", flattenAttributes(setVal))
	}

	return nil
}

func customSecurityAttributeAssignmentResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	spClient := meta.(*clients.Client).CustomSecurityAttributes.ServicePrincipalClientBeta

	principalId, attributeSet, err := parseResourceID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing resource ID")
	}

	objectPath, diags := resolveObjectPath(ctx, spClient.Client, principalId)
	if diags != nil {
		return diags
	}
	rawAttrs := d.Get("attribute").(*pluginsdk.Set).List()

	clearPayload := map[string]interface{}{
		"customSecurityAttributes": buildClearPayload(attributeSet, rawAttrs),
	}

	req, err := spClient.Client.NewRequest(ctx, sdkClient.RequestOptions{
		ContentType:         "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{http.StatusNoContent, http.StatusOK},
		HttpMethod:          http.MethodPatch,
		Path:                objectPath,
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Building PATCH request for principal %q", principalId)
	}

	if err = req.Marshal(clearPayload); err != nil {
		return tf.ErrorDiagF(err, "Marshalling clear payload for principal %q", principalId)
	}

	if _, err := req.Execute(ctx); err != nil {
		return tf.ErrorDiagF(err, "Clearing custom security attributes on principal %q", principalId)
	}

	return nil
}

// buildClearPayload constructs a payload that nulls out every attribute tracked by this resource.
func buildClearPayload(attributeSet string, input []interface{}) map[string]interface{} {
	setPayload := map[string]interface{}{
		"@odata.type": "#microsoft.graph.customSecurityAttributeValue",
	}

	for _, rawAttr := range input {
		attrMap, ok := rawAttr.(map[string]interface{})
		if !ok {
			continue
		}
		name := attrMap["name"].(string)
		setPayload[name] = nil
	}

	return map[string]interface{}{
		attributeSet: setPayload,
	}
}

// unmarshalCSA is kept for reference. The resource uses raw HTTP reads instead,
// since the SDK's CustomSecurityAttributeValue type is a stub that cannot hold attribute data.
