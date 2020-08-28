package aadgraph

import (
	"fmt"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func UserResource() *schema.Resource {
	return &schema.Resource{
		Create: userResourceCreate,
		Read:   userResourceRead,
		Update: userResourceUpdate,
		Delete: userResourceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"user_principal_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.StringIsEmailAddress,
			},

			"display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"mail_nickname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"account_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"password": {
				Type:         schema.TypeString,
				Required:     true,
				Sensitive:    true,
				ValidateFunc: validation.StringLenBetween(1, 256), //currently the max length for AAD passwords is 256
			},

			"force_password_change": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"mail": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"onpremises_sam_account_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"onpremises_user_principal_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"immutable_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "This must be specified if you are using a federated domain for the user's userPrincipalName (UPN) property when creating a new user account. " +
					"It is used to associate an on-premises Active Directory user account with their Azure AD user object.",
			},

			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"usage_location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "A two letter country code (ISO standard 3166). " +
					"Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. " +
					"Examples include: `NO`, `JP`, and `GB`. Not nullable.",
			},
		},
	}
}

func userResourceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.UsersClient
	ctx := meta.(*clients.AadClient).StopContext

	upn := d.Get("user_principal_name").(string)
	mailNickName := d.Get("mail_nickname").(string)

	//default mail nickname to the first part of the UPN (matches the portal)
	if mailNickName == "" {
		mailNickName = strings.Split(upn, "@")[0]
	}

	userCreateParameters := graphrbac.UserCreateParameters{
		AccountEnabled: utils.Bool(d.Get("account_enabled").(bool)),
		DisplayName:    utils.String(d.Get("display_name").(string)),
		MailNickname:   &mailNickName,
		PasswordProfile: &graphrbac.PasswordProfile{
			ForceChangePasswordNextLogin: utils.Bool(d.Get("force_password_change").(bool)),
			Password:                     utils.String(d.Get("password").(string)),
		},
		UserPrincipalName: &upn,
	}

	if v, ok := d.GetOk("usage_location"); ok {
		userCreateParameters.UsageLocation = utils.String(v.(string))
	}

	if v, ok := d.GetOk("immutable_id"); ok {
		userCreateParameters.ImmutableID = utils.String(v.(string))
	}

	user, err := client.Create(ctx, userCreateParameters)
	if err != nil {
		return fmt.Errorf("creating User (%q): %+v", upn, err)
	}
	if user.ObjectID == nil {
		return fmt.Errorf("nil User ID for %q: %+v", upn, err)
	}
	d.SetId(*user.ObjectID)

	_, err = graph.WaitForCreationReplication(d.Timeout(schema.TimeoutCreate), func() (interface{}, error) {
		return client.Get(ctx, *user.ObjectID)
	})
	if err != nil {
		return fmt.Errorf("waiting for User (%s) with ObjectId %q: %+v", upn, *user.ObjectID, err)
	}

	return userResourceRead(d, meta)
}

func userResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.UsersClient
	ctx := meta.(*clients.AadClient).StopContext

	var userUpdateParameters graphrbac.UserUpdateParameters

	if d.HasChange("display_name") {
		userUpdateParameters.DisplayName = utils.String(d.Get("display_name").(string))
	}

	if d.HasChange("mail_nickname") {
		userUpdateParameters.MailNickname = utils.String(d.Get("mail_nickname").(string))
	}

	if d.HasChange("account_enabled") {
		userUpdateParameters.AccountEnabled = utils.Bool(d.Get("account_enabled").(bool))
	}

	if d.HasChange("password") {
		userUpdateParameters.PasswordProfile = &graphrbac.PasswordProfile{
			ForceChangePasswordNextLogin: utils.Bool(d.Get("force_password_change").(bool)),
			Password:                     utils.String(d.Get("password").(string)),
		}
	}

	if d.HasChange("usage_location") {
		userUpdateParameters.UsageLocation = utils.String(d.Get("usage_location").(string))
	}

	if d.HasChange("immutable_id") {
		userUpdateParameters.ImmutableID = utils.String(d.Get("immutable_id").(string))
	}

	if _, err := client.Update(ctx, d.Id(), userUpdateParameters); err != nil {
		return fmt.Errorf("updating User with ID %q: %+v", d.Id(), err)
	}

	return userResourceRead(d, meta)
}

func userResourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.UsersClient
	ctx := meta.(*clients.AadClient).StopContext

	objectId := d.Id()

	user, err := client.Get(ctx, objectId)
	if err != nil {
		if utils.ResponseWasNotFound(user.Response) {
			log.Printf("[DEBUG] User with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving User with ID %q: %+v", objectId, err)
	}

	d.Set("user_principal_name", user.UserPrincipalName)
	d.Set("display_name", user.DisplayName)
	d.Set("mail", user.Mail)
	d.Set("mail_nickname", user.MailNickname)
	d.Set("account_enabled", user.AccountEnabled)
	d.Set("object_id", user.ObjectID)
	d.Set("usage_location", user.UsageLocation)
	d.Set("immutable_id", user.ImmutableID)

	d.Set("onpremises_sam_account_name", user.AdditionalProperties["onPremisesSamAccountName"])
	d.Set("onpremises_user_principal_name", user.AdditionalProperties["onPremisesUserPrincipalName"])

	return nil
}

func userResourceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.UsersClient
	ctx := meta.(*clients.AadClient).StopContext

	resp, err := client.Delete(ctx, d.Id())
	if err != nil {
		if !utils.ResponseWasNotFound(resp) {
			return fmt.Errorf("deleting User with ID %q: %+v", d.Id(), err)
		}
	}

	return nil
}
