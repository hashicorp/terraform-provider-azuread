package azuread

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/p"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

func resourceApplicationPassword() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationPasswordCreate,
		Read:   resourceApplicationPasswordRead,
		Delete: resourceApplicationPasswordDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"application_object_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
			},

			"key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
			},

			"value": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Sensitive:    true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"start_date": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.ValidateRFC3339TimeString,
			},

			"end_date": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"end_date_relative"},
				ValidateFunc:  validation.ValidateRFC3339TimeString,
			},

			"end_date_relative": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"end_date"},
				ValidateFunc:  validate.NoEmptyStrings,
			},
		},
	}
}

func resourceApplicationPasswordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	applicationObjId := d.Get("application_object_id").(string)
	value := d.Get("value").(string)
	// errors will be handled by the validation

	var keyId string
	if v, ok := d.GetOk("key_id"); ok {
		keyId = v.(string)
	} else {
		kid, err := uuid.GenerateUUID()
		if err != nil {
			return err
		}

		keyId = kid
	}

	var endDate time.Time
	if v := d.Get("end_date").(string); v != "" {
		endDate, _ = time.Parse(time.RFC3339, v)
	} else if v := d.Get("end_date_relative").(string); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			return fmt.Errorf("unable to parse `end_date_relative` (%s) as a duration", v)
		}
		endDate = time.Now().Add(d)
	} else {
		return fmt.Errorf("one of `end_date` or `end_date_relative` must be specified")
	}

	credential := graphrbac.PasswordCredential{
		KeyID:   p.String(keyId),
		Value:   p.String(value),
		EndDate: &date.Time{Time: endDate},
	}

	if v, ok := d.GetOk("start_date"); ok {
		// errors will be handled by the validation
		startDate, _ := time.Parse(time.RFC3339, v.(string))
		credential.StartDate = &date.Time{Time: startDate}
	}

	azureADLockByName(applicationObjId, applicationResourceName)
	defer azureADUnlockByName(applicationObjId, applicationResourceName)

	existingCredentials, err := client.ListPasswordCredentials(ctx, applicationObjId)
	if err != nil {
		return fmt.Errorf("Error Listing Password Credentials for Application Object ID %q: %+v", applicationObjId, err)
	}

	id := fmt.Sprintf("%s/%s", applicationObjId, keyId)
	updatedCredentials := make([]graphrbac.PasswordCredential, 0)
	if existingCredentials.Value != nil {
		if requireResourcesToBeImported {
			for _, v := range *existingCredentials.Value {
				if v.KeyID == nil {
					continue
				}

				if *v.KeyID == keyId {
					return tf.ImportAsExistsError("azuread_application_password", id)
				}
			}
		}

		updatedCredentials = *existingCredentials.Value
	}
	updatedCredentials = append(updatedCredentials, credential)

	parameters := graphrbac.PasswordCredentialsUpdateParameters{
		Value: &updatedCredentials,
	}

	if _, err = client.UpdatePasswordCredentials(ctx, applicationObjId, parameters); err != nil {
		return fmt.Errorf("Error creating Password Credential %q for Application Object ID %q: %+v", keyId, applicationObjId, err)
	}

	d.SetId(id)

	return resourceApplicationPasswordRead(d, meta)
}

func resourceApplicationPasswordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	id := strings.Split(d.Id(), "/")
	if len(id) != 2 {
		return fmt.Errorf("ID should be in the format {applicationObjId}/{keyId} - but got %q", d.Id())
	}

	applicationObjId := id[0]
	keyId := id[1]

	// ensure the parent Application exists
	application, err := client.Get(ctx, applicationObjId)
	if err != nil {
		// the parent Application has been removed - skip it
		if ar.ResponseWasNotFound(application.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", applicationObjId)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error retrieving Application Object ID %q: %+v", applicationObjId, err)
	}

	credentials, err := client.ListPasswordCredentials(ctx, applicationObjId)
	if err != nil {
		return fmt.Errorf("Error Listing Password Credentials for Application with Object ID %q: %+v", applicationObjId, err)
	}

	var credential *graphrbac.PasswordCredential
	for _, c := range *credentials.Value {
		if c.KeyID == nil {
			continue
		}

		if *c.KeyID == keyId {
			credential = &c
			break
		}
	}

	if credential == nil {
		log.Printf("[DEBUG] Applcation Password %q (ID %q) was not found - removing from state!", keyId, applicationObjId)
		d.SetId("")
		return nil
	}

	// value is available in the SDK but isn't returned from the API
	d.Set("key_id", credential.KeyID)
	d.Set("application_object_id", applicationObjId)

	if endDate := credential.EndDate; endDate != nil {
		d.Set("end_date", endDate.Format(time.RFC3339))
	}

	if startDate := credential.StartDate; startDate != nil {
		d.Set("start_date", startDate.Format(time.RFC3339))
	}

	return nil
}

func resourceApplicationPasswordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	id := strings.Split(d.Id(), "/")
	if len(id) != 2 {
		return fmt.Errorf("ID should be in the format {applicationObjId}/{keyId} - but got %q", d.Id())
	}

	applicationObjId := id[0]
	keyId := id[1]

	azureADLockByName(applicationObjId, applicationResourceName)
	defer azureADUnlockByName(applicationObjId, applicationResourceName)

	// ensure the parent Application exists
	application, err := client.Get(ctx, applicationObjId)
	if err != nil {
		// the parent Application was removed - skip it
		if ar.ResponseWasNotFound(application.Response) {
			return nil
		}

		return fmt.Errorf("Error retrieving Application Object ID %q: %+v", applicationObjId, err)
	}

	existing, err := client.ListPasswordCredentials(ctx, applicationObjId)
	if err != nil {
		return fmt.Errorf("Error Listing Password Credentials for Application with Object ID %q: %+v", applicationObjId, err)
	}

	updatedCredentials := make([]graphrbac.PasswordCredential, 0)
	for _, credential := range *existing.Value {
		if credential.KeyID == nil {
			continue
		}

		if *credential.KeyID != keyId {
			updatedCredentials = append(updatedCredentials, credential)
		}
	}

	parameters := graphrbac.PasswordCredentialsUpdateParameters{
		Value: &updatedCredentials,
	}
	_, err = client.UpdatePasswordCredentials(ctx, applicationObjId, parameters)
	if err != nil {
		return fmt.Errorf("Error removing Password %q from Application Object ID %q: %+v", keyId, applicationObjId, err)
	}

	return nil
}
