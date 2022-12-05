package directoryroles

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/suppress"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func directoryRoleResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: directoryRoleResourceCreate,
		ReadContext:   directoryRoleResourceRead,
		DeleteContext: directoryRoleResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Description:      "The display name of the directory role",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ExactlyOneOf:     []string{"display_name", "template_id"},
				DiffSuppressFunc: suppress.CaseDifference,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"template_id": {
				Description:      "The object ID of the template associated with the directory role",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ExactlyOneOf:     []string{"display_name", "template_id"},
				ValidateDiagFunc: validate.UUID,
			},

			"description": {
				Description: "The description of the directory role",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"object_id": {
				Description: "The object ID of the directory role",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func directoryRoleResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient
	directoryRoleTemplatesClient := meta.(*clients.Client).DirectoryRoles.DirectoryRoleTemplatesClient
	displayName := d.Get("display_name").(string)
	templateId := d.Get("template_id").(string)

	// First we find the directory role template
	var template *msgraph.DirectoryRoleTemplate
	if displayName != "" {
		templates, _, err := directoryRoleTemplatesClient.List(ctx)
		if err != nil {
			return tf.ErrorDiagF(err, "Retrieving directory role templates: %q", err)
		}
		if templates == nil {
			return tf.ErrorDiagF(errors.New("API error: nil result returned"), "Retrieving directory role templates")
		}

		for _, t := range *templates {
			if t.DisplayName != nil && strings.EqualFold(displayName, *t.DisplayName) {
				template = &t
				break
			}
		}

		if template == nil {
			return tf.ErrorDiagPathF(errors.New("template not found"), "Directory role template not found with display name %q", displayName)
		}
	} else {
		var status int
		var err error
		template, status, err = directoryRoleTemplatesClient.Get(ctx, templateId)
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagPathF(nil, "template_id", "Directory role template with object ID %q was not found", templateId)
			}
			return tf.ErrorDiagPathF(err, "template_id", "Retrieving directory role template with object ID %q: %+v", templateId, err)
		}

		if template == nil {
			return tf.ErrorDiagPathF(errors.New("template not found"), "Directory role template not found with object ID %q", templateId)
		}
	}

	if template == nil {
		return tf.ErrorDiagF(errors.New("template was nil"), "No template found")
	}

	if template.ID == nil {
		return tf.ErrorDiagF(errors.New("API error: template returned with nil ID"), "Retrieving directory role template")
	}

	templateId = *template.ID

	// Now look for the directory role created from that template
	directoryRole, status, err := client.GetByTemplateId(ctx, templateId)
	if err != nil {
		if status == http.StatusNotFound {
			// Directory role was not found, so activate it
			directoryRole, _, err = client.Activate(ctx, templateId)
			if err != nil {
				return tf.ErrorDiagPathF(err, "template_id", "Activating directory role for template ID %q: %+v", templateId, err)
			}
		} else {
			return tf.ErrorDiagPathF(err, "template_id", "Retrieving directory role with template ID %q: %+v", templateId, err)
		}
	}

	if directoryRole == nil {
		return tf.ErrorDiagF(errors.New("unexpected: directoryRole was nil"), "Retrieving directory role for template ID %q", templateId)
	}
	if directoryRole.ID() == nil {
		return tf.ErrorDiagF(errors.New("API error: directoryRole returned with nil ID"), "Retrieving directory role for template ID %q", templateId)
	}

	d.SetId(*directoryRole.ID())

	return directoryRoleResourceRead(ctx, d, meta)
}

func directoryRoleResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient

	directoryRole, status, err := client.Get(ctx, d.Id())
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Directory Role with ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "template_id", "Retrieving directory role with object ID %q: %+v", d.Id(), err)
	}
	if directoryRole == nil {
		return tf.ErrorDiagF(errors.New("API error: nil directoryRole was returned"), "Retrieving directory role with object ID %q", d.Id())
	}

	tf.Set(d, "description", directoryRole.Description)
	tf.Set(d, "display_name", directoryRole.DisplayName)
	tf.Set(d, "object_id", directoryRole.ID())
	tf.Set(d, "template_id", directoryRole.RoleTemplateId)

	return nil
}

func directoryRoleResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Directory roles cannot be deactivated or deleted, so this is a no-op
	return nil
}
