package aadgraph

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func GroupResource() *schema.Resource {
	return &schema.Resource{
		Create: groupResourceCreate,
		Read:   groupResourceRead,
		Update: groupResourceUpdate,
		Delete: groupResourceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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

func groupResourceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient
	ctx := meta.(*clients.AadClient).StopContext

	name := d.Get("name").(string)

	if d.Get("prevent_duplicate_names").(bool) {
		err := graph.GroupCheckNameAvailability(ctx, client, name)
		if err != nil {
			return err
		}
	}

	properties := graphrbac.GroupCreateParameters{
		DisplayName:          &name,
		MailEnabled:          utils.Bool(false),                 // we're defaulting to false, as the API currently only supports the creation of non-mail enabled security groups.
		MailNickname:         utils.String(uuid.New().String()), // this matches the portal behaviour
		SecurityEnabled:      utils.Bool(true),                  // we're defaulting to true, as the API currently only supports the creation of non-mail enabled security groups.
		AdditionalProperties: make(map[string]interface{}),
	}

	if v, ok := d.GetOk("description"); ok {
		properties.AdditionalProperties["description"] = v.(string)
	}

	group, err := client.Create(ctx, properties)
	if err != nil {
		return fmt.Errorf("creating Group (%q): %+v", name, err)
	}
	if group.ObjectID == nil {
		return fmt.Errorf("nil Group ID for %q: %+v", name, err)
	}

	d.SetId(*group.ObjectID)

	_, err = graph.WaitForCreationReplication(d.Timeout(schema.TimeoutCreate), func() (interface{}, error) {
		return client.Get(ctx, *group.ObjectID)
	})

	// Add members if specified
	if v, ok := d.GetOk("members"); ok {
		members := tf.ExpandStringSlicePtr(v.(*schema.Set).List())

		// we could lock here against the group member resource, but they should not be used together (todo conflicts with at a resource level?)
		if err := graph.GroupAddMembers(ctx, client, *group.ObjectID, *members); err != nil {
			return err
		}
	}

	// Add owners if specified
	if v, ok := d.GetOk("owners"); ok {
		existingOwners, err := graph.GroupAllOwners(ctx, client, *group.ObjectID)
		if err != nil {
			return err
		}
		members := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		ownersToAdd := utils.Difference(members, existingOwners)

		if err := graph.GroupAddOwners(ctx, client, *group.ObjectID, ownersToAdd); err != nil {
			return err
		}
	}

	if err != nil {
		return fmt.Errorf("waiting for Group (%s) with ObjectId %q: %+v", name, *group.ObjectID, err)
	}

	return groupResourceRead(d, meta)
}

func groupResourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient
	ctx := meta.(*clients.AadClient).StopContext

	resp, err := client.Get(ctx, d.Id())
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] Group with id %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("retrieving Group with ID %q: %+v", d.Id(), err)
	}

	d.Set("name", resp.DisplayName)
	d.Set("object_id", resp.ObjectID)

	if v, ok := resp.AdditionalProperties["description"]; ok {
		d.Set("description", v.(string))
	}

	members, err := graph.GroupAllMembers(ctx, client, d.Id())
	if err != nil {
		return err
	}
	d.Set("members", members)

	owners, err := graph.GroupAllOwners(ctx, client, d.Id())
	if err != nil {
		return err
	}
	d.Set("owners", owners)

	if preventDuplicates := d.Get("prevent_duplicate_names").(bool); !preventDuplicates {
		d.Set("prevent_duplicate_names", false)
	}

	return nil
}

func groupResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient
	ctx := meta.(*clients.AadClient).StopContext

	if v, ok := d.GetOkExists("members"); ok && d.HasChange("members") {
		existingMembers, err := graph.GroupAllMembers(ctx, client, d.Id())
		if err != nil {
			return err
		}

		desiredMembers := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		membersForRemoval := utils.Difference(existingMembers, desiredMembers)
		membersToAdd := utils.Difference(desiredMembers, existingMembers)

		for _, existingMember := range membersForRemoval {
			var err error
			var resp autorest.Response
			log.Printf("[DEBUG] Removing member with id %q from Group with id %q", existingMember, d.Id())
			attempts := 10
			for i := 0; i <= attempts; i++ {
				if resp, err = client.RemoveMember(ctx, d.Id(), existingMember); err != nil {
					if utils.ResponseWasNotFound(resp) {
						break
					}
				}
				if i == attempts {
					return fmt.Errorf("deleting group member %q from Group with ID %q: %+v", existingMember, d.Id(), err)
				}
				time.Sleep(time.Second * 2)
			}

			if _, err := graph.WaitForListRemove(existingMember, func() ([]string, error) {
				return graph.GroupAllMembers(ctx, client, d.Id())
			}); err != nil {
				return fmt.Errorf("waiting for group membership removal: %+v", err)
			}
		}

		if err := graph.GroupAddMembers(ctx, client, d.Id(), membersToAdd); err != nil {
			return err
		}
	}

	if v, ok := d.GetOkExists("owners"); ok && d.HasChange("owners") {
		existingOwners, err := graph.GroupAllOwners(ctx, client, d.Id())
		if err != nil {
			return err
		}

		desiredOwners := *tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		ownersForRemoval := utils.Difference(existingOwners, desiredOwners)
		ownersToAdd := utils.Difference(desiredOwners, existingOwners)

		for _, ownerToDelete := range ownersForRemoval {
			log.Printf("[DEBUG] Removing member with id %q from Group with id %q", ownerToDelete, d.Id())
			if resp, err := client.RemoveOwner(ctx, d.Id(), ownerToDelete); err != nil {
				if !utils.ResponseWasNotFound(resp) {
					return fmt.Errorf("deleting group member %q from Group with ID %q: %+v", ownerToDelete, d.Id(), err)
				}
			}
		}

		if err := graph.GroupAddOwners(ctx, client, d.Id(), ownersToAdd); err != nil {
			return err
		}
	}

	return groupResourceRead(d, meta)
}

func groupResourceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient
	ctx := meta.(*clients.AadClient).StopContext

	if resp, err := client.Delete(ctx, d.Id()); err != nil {
		if !utils.ResponseWasNotFound(resp) {
			return fmt.Errorf("deleting Group with ID %q: %+v", d.Id(), err)
		}
	}

	return nil
}
