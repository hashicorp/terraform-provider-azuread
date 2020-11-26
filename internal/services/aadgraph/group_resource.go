package aadgraph

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
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
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			"description": {
				Type:     schema.TypeString,
				ForceNew: true, // there is no update method available in the SDK
				Optional: true,
			},

			"members": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.UUID,
				},
			},

			"owners": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.UUID,
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
		},
	}
}

func groupResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient

	name := d.Get("name").(string)

	if d.Get("prevent_duplicate_names").(bool) {
		err := graph.GroupCheckNameAvailability(ctx, client, name)
		if err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       err.Error(),
				AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
			}}
		}
	}

	mailNickname, err := uuid.GenerateUUID()
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to generate mailNickname",
			Detail:   err.Error(),
		}}
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
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Creating group %q", name),
			Detail:   err.Error(),
		}}
	}

	if group.ObjectID == nil || *group.ObjectID == "" {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "API returned group with nil object ID",
		}}
	}

	d.SetId(*group.ObjectID)

	_, err = graph.WaitForCreationReplication(d.Timeout(schema.TimeoutCreate), func() (interface{}, error) {
		return client.Get(ctx, *group.ObjectID)
	})

	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Waiting for Group with object ID: %q", *group.ObjectID),
			Detail:   err.Error(),
		}}
	}

	// Add members if specified
	if v, ok := d.GetOk("members"); ok {
		members := tf.ExpandStringSlicePtr(v.(*schema.Set).List())

		// we could lock here against the group member resource, but they should not be used together (todo conflicts with at a resource level?)
		if err := graph.GroupAddMembers(ctx, client, *group.ObjectID, *members); err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Adding group members",
				Detail:   err.Error(),
			}}
		}
	}

	// Add owners if specified
	if v, ok := d.GetOk("owners"); ok {
		existingOwners, err := graph.GroupAllOwners(ctx, client, *group.ObjectID)
		if err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not retrieve group owners",
				Detail:   err.Error(),
			}}
		}
		members := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		ownersToAdd := utils.Difference(members, existingOwners)

		if err := graph.GroupAddOwners(ctx, client, *group.ObjectID, ownersToAdd); err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Adding group owners",
				Detail:   err.Error(),
			}}
		}
	}

	return groupResourceRead(ctx, d, meta)
}

func groupResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient

	resp, err := client.Get(ctx, d.Id())
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] Group with id %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Retrieving group with object ID: %q", d.Id()),
			Detail:   err.Error(),
		}}
	}

	d.Set("name", resp.DisplayName)
	d.Set("object_id", resp.ObjectID)

	if v, ok := resp.AdditionalProperties["description"]; ok {
		d.Set("description", v.(string))
	}

	members, err := graph.GroupAllMembers(ctx, client, d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not retrieve group members",
			Detail:   err.Error(),
		}}
	}
	d.Set("members", members)

	owners, err := graph.GroupAllOwners(ctx, client, d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not retrieve group owners",
			Detail:   err.Error(),
		}}
	}
	d.Set("owners", owners)

	preventDuplicates := false
	if v := d.Get("prevent_duplicate_names").(bool); v {
		preventDuplicates = v
	}
	d.Set("prevent_duplicate_names", preventDuplicates)

	return nil
}

func groupResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient

	if v, ok := d.GetOkExists("members"); ok && d.HasChange("members") {
		existingMembers, err := graph.GroupAllMembers(ctx, client, d.Id())
		if err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not retrieve group members",
				Detail:   err.Error(),
			}}
		}

		desiredMembers := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		membersForRemoval := utils.Difference(existingMembers, desiredMembers)
		membersToAdd := utils.Difference(desiredMembers, existingMembers)

		for _, existingMember := range membersForRemoval {
			log.Printf("[DEBUG] Removing member with id %q from Group with id %q", existingMember, d.Id())
			if err := graph.GroupRemoveMember(ctx, client, d.Timeout(schema.TimeoutDelete), d.Id(), existingMember); err != nil {
				return diag.Diagnostics{diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Removing group members",
					Detail:   err.Error(),
				}}
			}

			if _, err := graph.WaitForListRemove(existingMember, func() ([]string, error) {
				return graph.GroupAllMembers(ctx, client, d.Id())
			}); err != nil {
				return diag.Diagnostics{diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Waiting for group membership removal",
					Detail:   err.Error(),
				}}
			}
		}

		if err := graph.GroupAddMembers(ctx, client, d.Id(), membersToAdd); err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Adding group members",
				Detail:   err.Error(),
			}}
		}
	}

	if v, ok := d.GetOkExists("owners"); ok && d.HasChange("owners") {
		existingOwners, err := graph.GroupAllOwners(ctx, client, d.Id())
		if err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not retrieve group owners",
				Detail:   err.Error(),
			}}
		}

		desiredOwners := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		ownersForRemoval := utils.Difference(existingOwners, desiredOwners)
		ownersToAdd := utils.Difference(desiredOwners, existingOwners)

		for _, ownerToDelete := range ownersForRemoval {
			log.Printf("[DEBUG] Removing member with id %q from Group with id %q", ownerToDelete, d.Id())
			if resp, err := client.RemoveOwner(ctx, d.Id(), ownerToDelete); err != nil {
				if !utils.ResponseWasNotFound(resp) {
					return diag.Diagnostics{diag.Diagnostic{
						Severity: diag.Error,
						Summary:  fmt.Sprintf("Removing group owner %q from group with object ID: %q", ownerToDelete, d.Id()),
						Detail:   err.Error(),
					}}
				}
			}
		}

		if err := graph.GroupAddOwners(ctx, client, d.Id(), ownersToAdd); err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Adding group owners",
				Detail:   err.Error(),
			}}
		}
	}

	return groupResourceRead(ctx, d, meta)
}

func groupResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient

	if resp, err := client.Delete(ctx, d.Id()); err != nil {
		if !utils.ResponseWasNotFound(resp) {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Deleting group with object ID: %q", d.Id()),
				Detail:   err.Error(),
			}}
		}
	}

	return nil
}
