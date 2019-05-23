package azuread

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/p"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

const objectNotFound = "Object not found"

// valid types are `application` and `service_pricipal`
func resourceObjectPasswordSchema(object_type string) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		object_type + "_id": {
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
	}
}

func resourceApplicationPassword() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectPasswordCreate,
		Read:   resourceObjectPasswordRead,
		Delete: resourceObjectPasswordDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourceObjectPasswordSchema("application"),
	}
}

func resourceObjectPasswordGetPasswordCredential(d *schema.ResourceData, objectType string) (*graphrbac.PasswordCredential, error) {
	objectId := d.Get(objectType+"_id").(string)
	value := d.Get("value").(string)
	// errors will be handled by the validation

	var keyId string
	if v, ok := d.GetOk("key_id"); ok {
		keyId = v.(string)
	} else {
		kid, err := uuid.GenerateUUID()
		if err != nil {
			return nil, err
		}

		keyId = kid
	}

	var endDate time.Time
	if v := d.Get("end_date").(string); v != "" {
		endDate, _ = time.Parse(time.RFC3339, v)
	} else if v := d.Get("end_date_relative").(string); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			return nil, fmt.Errorf("unable to parse `end_date_relative` (%s) as a duration", v)
		}
		endDate = time.Now().Add(d)
	} else {
		return nil, fmt.Errorf("one of `end_date` or `end_date_relative` must be specified")
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

	return &credential, nil
}

func resourceObjectPasswordCreate(d *schema.ResourceData, meta interface{}) error {
	client := passwordsClient{
		appClient: meta.(*ArmClient).applicationsClient,
		spClient:  meta.(*ArmClient).servicePrincipalsClient,
		ctx:       meta.(*ArmClient).StopContext,
	}



	resourceName, err := client.GetObjectType(objectId)
	if err != nil {
		return fmt.Errorf("Error getting resource name for Object ID %q: %+v", objectId, err)
	}

	azureADLockByName(objectId, resourceName)
	defer azureADUnlockByName(objectId, resourceName)

	existingCredentials, err := client.ListPasswordCredentials(objectId)
	if err != nil {
		return fmt.Errorf("Error Listing Password Credentials for Object ID %q: %+v", objectId, err)
	}
	id := fmt.Sprintf("%s/%s", objectId, keyId)

	updatedCredentials := make([]graphrbac.PasswordCredential, 0)
	if *existingCredentials.Value != nil {
		if requireResourcesToBeImported {
			for _, v := range *existingCredentials.Value {
				if v.KeyID == nil {
					continue
				}

				if *v.KeyID == keyId {
					return tf.ImportAsExistsError(passwordResourceName, id)
				}
			}
		}

		updatedCredentials = *existingCredentials.Value
	}
	updatedCredentials = append(updatedCredentials, credential)

	parameters := graphrbac.PasswordCredentialsUpdateParameters{
		Value: &updatedCredentials,
	}

	if _, err = client.UpdatePasswordCredentials(objectId, parameters); err != nil {
		return fmt.Errorf("Error creating Password Credential %q for Object ID %q: %+v", keyId, objectId, err)
	}

	d.SetId(id)

	return resourceObjectPasswordRead(d, meta)
}

func resourceObjectPasswordRead(d *schema.ResourceData, meta interface{}) error {
	client := passwordsClient{
		appClient: meta.(*ArmClient).applicationsClient,
		spClient:  meta.(*ArmClient).servicePrincipalsClient,
		ctx:       meta.(*ArmClient).StopContext,
	}

	id := strings.Split(d.Id(), "/")
	if len(id) != 2 {
		return fmt.Errorf("ID should be in the format {objectId}/{keyId} - but got %q", d.Id())
	}

	objectId := id[0]
	keyId := id[1]

	// ensure the parent Object exists
	exists, err := client.ObjectExists(objectId)
	if err != nil {
		return fmt.Errorf("Error retrieving Object with ID %q: %+v", id, err)
	}
	if !exists {
		log.Printf("[DEBUG] Object with ID %q was not found - removing from state!", id)
		d.SetId("")
		return nil
	}

	credentials, err := client.ListPasswordCredentials(objectId)
	if err != nil {
		return fmt.Errorf("Error Listing Password Credentials for Application with Object ID %q: %+v", objectId, err)
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
		log.Printf("[DEBUG] Password %q (ID %q) was not found - removing from state!", keyId, objectId)
		d.SetId("")
		return nil
	}

	// value is available in the SDK but isn't returned from the API
	d.Set("key_id", credential.KeyID)
	d.Set("object_id", objectId)

	if endDate := credential.EndDate; endDate != nil {
		d.Set("end_date", endDate.Format(time.RFC3339))
	}

	if startDate := credential.StartDate; startDate != nil {
		d.Set("start_date", startDate.Format(time.RFC3339))
	}

	return nil
}

func resourceObjectPasswordDelete(d *schema.ResourceData, meta interface{}) error {
	client := passwordsClient{
		appClient: meta.(*ArmClient).applicationsClient,
		spClient:  meta.(*ArmClient).servicePrincipalsClient,
		ctx:       meta.(*ArmClient).StopContext,
	}

	id := strings.Split(d.Id(), "/")
	if len(id) != 2 {
		return fmt.Errorf("ID should be in the format {objectId}/{keyId} - but got %q", d.Id())
	}

	objectId := id[0]
	keyId := id[1]

	resourceName, err := client.GetObjectType(objectId)
	if err != nil {
		return fmt.Errorf("Error getting resource name for Object ID %q: %+v", objectId, err)
	}

	azureADLockByName(objectId, resourceName)
	defer azureADUnlockByName(objectId, resourceName)

	// ensure the parent Application exists
	exists, err := client.ObjectExists(objectId)
	if err != nil {
		return fmt.Errorf("Error retrieving Object with ID %q: %+v", id, err)
	}
	if !exists {
		return nil
	}

	existing, err := client.ListPasswordCredentials(objectId)
	if err != nil {
		return fmt.Errorf("Error Listing Password Credentials for Object with ID %q: %+v", objectId, err)
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
	_, err = client.UpdatePasswordCredentials(objectId, parameters)
	if err != nil {
		return fmt.Errorf("Error removing Password %q from Application Object ID %q: %+v", keyId, objectId, err)
	}

	return nil
}

type passwordsClient struct {
	appClient graphrbac.ApplicationsClient
	spClient  graphrbac.ServicePrincipalsClient
	ctx       context.Context
}

func (c passwordsClient) GetObjectType(id string) (string, error) {
	// try if id is an application
	application, err := c.appClient.Get(c.ctx, id)
	if err == nil {
		return applicationResourceName, nil
	}
	if !ar.ResponseWasNotFound(application.Response) {
		return "", fmt.Errorf("Error retrieving Application Object ID %q: %+v", id, err)
	}

	// try if id is a service principal
	sp, err := c.spClient.Get(c.ctx, id)
	if err == nil {
		return servicePrincipalResourceName, nil
	}
	if !ar.ResponseWasNotFound(sp.Response) {
		return "", fmt.Errorf("Error retrieving Service Principal ID %q: %+v", id, err)
	}

	return "", errors.New(objectNotFound)
}

func (c passwordsClient) ObjectExists(id string) (bool, error) {
	_, err := c.GetObjectType(id)
	if err != nil {
		if err.Error() == objectNotFound {
			return false, nil
		}
		return false, fmt.Errorf("Error retrieving Object with ID %q: %+v", id, err)
	}
	return true, nil

}

func (c passwordsClient) ListPasswordCredentials(id string) (*graphrbac.PasswordCredentialListResult, error) {
	objectType, err := c.GetObjectType(id)
	if err != nil {
		return nil, err
	}
	switch objectType {
	case applicationResourceName:
		credentials, err := c.appClient.ListPasswordCredentials(c.ctx, id)
		if err != nil {
			return nil, err
		}
		return &credentials, nil
	case servicePrincipalResourceName:
		credentials, err := c.spClient.ListPasswordCredentials(c.ctx, id)
		if err != nil {
			return nil, err
		}
		return &credentials, nil
	}
	return nil, fmt.Errorf("Object type not supported: %s", objectType)
}

func (c passwordsClient) UpdatePasswordCredentials(id string, parameters graphrbac.PasswordCredentialsUpdateParameters) (*autorest.Response, error) {
	objectType, err := c.GetObjectType(id)
	if err != nil {
		return nil, err
	}
	switch objectType {
	case applicationResourceName:
		resp, err := c.appClient.UpdatePasswordCredentials(c.ctx, id, parameters)
		if err != nil {
			return nil, err
		}
		return &resp, nil
	case servicePrincipalResourceName:
		resp, err := c.spClient.UpdatePasswordCredentials(c.ctx, id, parameters)
		if err != nil {
			return nil, err
		}
		return &resp, nil
	}
	return nil, fmt.Errorf("Object type not supported: %s", objectType)
}
