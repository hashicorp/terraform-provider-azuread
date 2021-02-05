package groups

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func groupResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: groupResourceCreate,
		ReadContext:   groupResourceRead,
		UpdateContext: groupResourceUpdate,
		DeleteContext: groupResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:             schema.TypeString,
				Optional:         true, // TODO: v2.0 set Required
				Computed:         true, // TODO: v2.0 remove Computed
				ExactlyOneOf:     []string{"display_name", "name"},
				ForceNew:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Deprecated:       "This property has been renamed to `display_name` and will be removed in v2.0 of this provider.",
				ExactlyOneOf:     []string{"display_name", "name"},
				ForceNew:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"description": {
				Type:     schema.TypeString,
				ForceNew: true, // there is no update method available in the SDK
				Optional: true,
			},

			"mail_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"members": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"owners": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},

			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"prevent_duplicate_names": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"security_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func groupResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.AadClient

	var name string
	if v, ok := d.GetOk("display_name"); ok && v.(string) != "" {
		name = v.(string)
	} else {
		name = d.Get("name").(string)
	}

	if d.Get("prevent_duplicate_names").(bool) {
		existingGroup, err := aadgraph.GroupFindByName(ctx, client, name)
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name", "Could not check for existing group(s)")
		}
		if existingGroup != nil {
			if existingGroup.ObjectID == nil {
				return tf.ImportAsDuplicateDiag("azuread_group", "unknown", name)
			}
			return tf.ImportAsDuplicateDiag("azuread_group", *existingGroup.ObjectID, name)
		}
	}

	mailNickname, err := uuid.GenerateUUID()
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to generate mailNickname")
	}

	properties := graphrbac.GroupCreateParameters{
		DisplayName:          &name,
		MailEnabled:          utils.Bool(false),          // we're defaulting to false, as the API currently only supports the creation of non-mail enabled security groups.
		MailNickname:         utils.String(mailNickname), // this matches the portal behaviour
		SecurityEnabled:      utils.Bool(true),           // we're defaulting to true, as the API currently only supports the creation of non-mail enabled security groups.
		AdditionalProperties: make(map[string]interface{}),
	}

	if v, ok := d.GetOk("description"); ok {
		properties.AdditionalProperties["description"] = v.(string)
	}

	group, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating group %q", name)
	}

	if group.ObjectID == nil || *group.ObjectID == "" {
		return tf.ErrorDiagF(errors.New("API returned group with nil object ID"), "Bad API Response")
	}

	d.SetId(*group.ObjectID)

	_, err = aadgraph.WaitForCreationReplication(ctx, d.Timeout(schema.TimeoutCreate), func() (interface{}, error) {
		return client.Get(ctx, *group.ObjectID)
	})

	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for Group with object ID: %q", *group.ObjectID)
	}

	// Add members if specified
	if v, ok := d.GetOk("members"); ok {
		members := tf.ExpandStringSlicePtr(v.(*schema.Set).List())

		// we could lock here against the group member resource, but they should not be used together (todo conflicts with at a resource level?)
		if err := aadgraph.GroupAddMembers(ctx, client, *group.ObjectID, *members); err != nil {
			return tf.ErrorDiagF(err, "Adding group members")
		}
	}

	// Add owners if specified
	if v, ok := d.GetOk("owners"); ok {
		existingOwners, err := aadgraph.GroupAllOwners(ctx, client, *group.ObjectID)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve group owners")
		}
		members := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		ownersToAdd := utils.Difference(members, existingOwners)

		if err := aadgraph.GroupAddOwners(ctx, client, *group.ObjectID, ownersToAdd); err != nil {
			return tf.ErrorDiagF(err, "Adding group owners")
		}
	}

	return groupResourceRead(ctx, d, meta)
}

func groupResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.AadClient

	resp, err := client.Get(ctx, d.Id())
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] Group with id %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "Retrieving group with object ID: %q", d.Id())
	}

	tf.Set(d, "object_id", resp.ObjectID)
	tf.Set(d, "display_name", resp.DisplayName)
	tf.Set(d, "name", resp.DisplayName)
	tf.Set(d, "mail_enabled", resp.MailEnabled)
	tf.Set(d, "security_enabled", resp.SecurityEnabled)

	description := ""
	if v, ok := resp.AdditionalProperties["description"]; ok {
		description = v.(string)
	}
	tf.Set(d, "description", description)

	members, err := aadgraph.GroupAllMembers(ctx, client, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve members for group with object ID %q", d.Id())
	}
	tf.Set(d, "members", members)

	owners, err := aadgraph.GroupAllOwners(ctx, client, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for group with object ID %q", d.Id())
	}
	tf.Set(d, "owners", owners)

	preventDuplicates := false
	if v := d.Get("prevent_duplicate_names").(bool); v {
		preventDuplicates = v
	}
	tf.Set(d, "prevent_duplicate_names", preventDuplicates)

	return nil
}

func groupResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.AadClient

	if v, ok := d.GetOkExists("members"); ok && d.HasChange("members") { //nolint:SA1019
		existingMembers, err := aadgraph.GroupAllMembers(ctx, client, d.Id())
		if err != nil {
			return tf.ErrorDiagPathF(err, "owners", "Could not retrieve members for group with object ID %q", d.Id())
		}

		desiredMembers := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		membersForRemoval := utils.Difference(existingMembers, desiredMembers)
		membersToAdd := utils.Difference(desiredMembers, existingMembers)

		for _, existingMember := range membersForRemoval {
			log.Printf("[DEBUG] Removing member with id %q from Group with id %q", existingMember, d.Id())
			if err := aadgraph.GroupRemoveMember(ctx, client, d.Timeout(schema.TimeoutDelete), d.Id(), existingMember); err != nil {
				return tf.ErrorDiagF(err, "Removing group members")
			}

			if _, err := aadgraph.WaitForListRemove(ctx, existingMember, func() ([]string, error) {
				return aadgraph.GroupAllMembers(ctx, client, d.Id())
			}); err != nil {
				return tf.ErrorDiagF(err, "Waiting for group membership removal")
			}
		}

		if err := aadgraph.GroupAddMembers(ctx, client, d.Id(), membersToAdd); err != nil {
			return tf.ErrorDiagF(err, "Adding group members")
		}
	}

	if v, ok := d.GetOkExists("owners"); ok && d.HasChange("owners") { //nolint:SA1019
		existingOwners, err := aadgraph.GroupAllOwners(ctx, client, d.Id())
		if err != nil {
			return tf.ErrorDiagPathF(err, "owners", "Could not retrieve owners for group with object ID %q", d.Id())
		}

		desiredOwners := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		ownersForRemoval := utils.Difference(existingOwners, desiredOwners)
		ownersToAdd := utils.Difference(desiredOwners, existingOwners)

		for _, ownerToDelete := range ownersForRemoval {
			log.Printf("[DEBUG] Removing member with ID %q from Group with ID %q", ownerToDelete, d.Id())
			if resp, err := client.RemoveOwner(ctx, d.Id(), ownerToDelete); err != nil {
				if !utils.ResponseWasNotFound(resp) {
					return tf.ErrorDiagF(err, "Removing group owner %q from group with object ID: %q", ownerToDelete, d.Id())
				}
			}
		}

		if err := aadgraph.GroupAddOwners(ctx, client, d.Id(), ownersToAdd); err != nil {
			return tf.ErrorDiagF(err, "Adding group owners")
		}
	}

	return groupResourceRead(ctx, d, meta)
}

func groupResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Groups.AadClient

	if resp, err := client.Delete(ctx, d.Id()); err != nil {
		if !utils.ResponseWasNotFound(resp) {
			return tf.ErrorDiagF(err, "Deleting group with object ID: %q", d.Id())
		}
	}

	return nil
}
