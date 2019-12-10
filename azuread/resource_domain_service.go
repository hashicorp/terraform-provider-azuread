package azuread

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/domainservices/mgmt/2017-06-01/aad"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/hashicorp/go-azure-helpers/response"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

func resourceDomainService() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmDomainServiceCreate,
		Read:   resourceArmDomainServiceRead,
		Update: resourceArmDomainServiceUpdate,
		Delete: resourceArmDomainServiceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"location": azure.SchemaLocation(),

			"resource_group_name": azure.SchemaResourceGroupNameDiffSuppress(),

			"subnet_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: azure.ValidateResourceID,
			},

			"filtered_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"ldaps": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_access": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"ldaps": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"pfx_certificate": {
							Type:         schema.TypeString,
							Optional:     true,
							Sensitive:    true,
							ValidateFunc: validate.Base64String(),
						},
						"pfx_certificate_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"external_access_ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"notifications": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_recipients": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validate.StringIsEmailAddress,
							},
						},
						"notify_dc_admins": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"notify_global_admins": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
			},

			"security": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ntlm_v1": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"sync_ntlm_passwords": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"tls_v1": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
			},

			"domain_controller_ip_address": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceArmDomainServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).domainServicesClient
	vnetClient := meta.(*ArmClient).vnetClient
	ctx := meta.(*ArmClient).StopContext

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)
	location := azure.NormalizeLocation(d.Get("location").(string))
	security := d.Get("security").([]interface{})
	ldaps := d.Get("ldaps").([]interface{})
	notifications := d.Get("notifications").([]interface{})
	subnetId := d.Get("subnet_id").(string)
	filteredSync := aad.FilteredSyncDisabled
	if d.Get("filtered_sync").(bool) {
		filteredSync = aad.FilteredSyncDisabled
	}

	domainService := aad.DomainService{
		Location: &location,
		DomainServiceProperties: &aad.DomainServiceProperties{
			DomainName:             &name,
			DomainSecuritySettings: expandArmDomainServiceSecurity(security),
			FilteredSync:           filteredSync,
			LdapsSettings:          expandArmDomainServiceLdapsSettings(ldaps),
			NotificationSettings:   expandArmDomainServiceNotificationSettings(notifications),
			SubnetID:               &subnetId,
		},
	}

	future, err := client.CreateOrUpdate(ctx, resourceGroup, name, domainService)
	if err != nil {
		return fmt.Errorf("Error creating Domain Service %q (Resource Group %q): %+v", name, resourceGroup, err)
	}
	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for creation of Domain Service %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	// Domain Services is 2 controllers running in the cloud
	// the create api completes Once the first domain controller is up and running, 1 ip address will show up
	// Afterwards, there will be an additional domain controller creating, then the ip address will update and become 2
	// the Azure Portal blocks users from modifying the domain service until both domain controllers are up
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"pending"},
		Target:       []string{"available"},
		Refresh:      domainServiceControllerRefreshFunc(ctx, &client, resourceGroup, name),
		Delay:        30 * time.Second,
		PollInterval: 10 * time.Second,
		Timeout:      1 * time.Hour,
	}

	if domainControllerIPAddress, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("Error waiting for both Domain Service Controller up, err: %+v", err)
	} else {
		// Update DNS server settings of the virtual network
		parsedId, err := azure.ParseAzureResourceID(subnetId)
		if err != nil {
			return err
		}
		resourceGroupName := parsedId.ResourceGroup
		virtualNetworkName := parsedId.Path["virtualNetworks"]

		resp, err := vnetClient.Get(ctx, resourceGroupName, virtualNetworkName, "")
		if err != nil {
			return fmt.Errorf("Error readding Virtual Network %q (Resource Group %q): %+v", virtualNetworkName, resourceGroupName, err)
		}
		dns := domainControllerIPAddress.([]string)
		resp.DhcpOptions = &network.DhcpOptions{
			DNSServers: &dns,
		}
		if _, err := vnetClient.CreateOrUpdate(ctx, resourceGroupName, virtualNetworkName, resp); err != nil {
			return fmt.Errorf("Error updating DNS server to %+v for Virtual Network %q (Resource Group %q): %+v", domainControllerIPAddress, virtualNetworkName, resourceGroupName, err)
		}
	}

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		return fmt.Errorf("Error retrieving Domain Service %q (Resource Group %q): %+v", name, resourceGroup, err)
	}
	if resp.ID == nil {
		return fmt.Errorf("Cannot read Domain Service %q (Resource Group %q) ID", name, resourceGroup)
	}
	d.SetId(*resp.ID)

	return resourceArmDomainServiceRead(d, meta)
}

func resourceArmDomainServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).domainServicesClient
	ctx := meta.(*ArmClient).StopContext

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	name := id.Path["domainServices"]

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		if ar.ResponseWasNotFound(resp.Response) {
			return nil
		}
		return err
	}

	d.Set("name", name)
	d.Set("resource_group_name", resourceGroup)
	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}
	if domainServiceProperties := resp.DomainServiceProperties; domainServiceProperties != nil {
		if err := d.Set("domain_controller_ip_address", tf.FlattenStringSlicePtr(domainServiceProperties.DomainControllerIPAddress)); err != nil {
			return fmt.Errorf("Error setting `domain_controller_ip_address`: %+v", err)
		}
		if err := d.Set("security", flattenArmDomainServiceSecurity(domainServiceProperties.DomainSecuritySettings)); err != nil {
			return fmt.Errorf("Error setting `security`: %+v", err)
		}
		if err := d.Set("ldaps", flattenArmDomainServiceLdaps(domainServiceProperties.LdapsSettings)); err != nil {
			return fmt.Errorf("Error setting `ldaps`: %+v", err)
		}
		if err := d.Set("notifications", flattenArmDomainServiceNotification(domainServiceProperties.NotificationSettings)); err != nil {
			return fmt.Errorf("Error setting `notifications`: %+v", err)
		}

		d.Set("filtered_sync", false)
		if domainServiceProperties.FilteredSync == aad.FilteredSyncEnabled {
			d.Set("filtered_sync", true)
		}
		d.Set("subnet_id", domainServiceProperties.SubnetID)
	}

	return nil
}

func resourceArmDomainServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).domainServicesClient
	ctx := meta.(*ArmClient).StopContext

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)
	security := d.Get("security").([]interface{})
	ldaps := d.Get("ldaps").([]interface{})
	notifications := d.Get("notifications").([]interface{})
	subnetId := d.Get("subnet_id").(string)
	filteredSync := aad.FilteredSyncDisabled
	if d.Get("filtered_sync").(bool) {
		filteredSync = aad.FilteredSyncDisabled
	}

	domainService := aad.DomainService{
		DomainServiceProperties: &aad.DomainServiceProperties{
			DomainName:             &name,
			DomainSecuritySettings: expandArmDomainServiceSecurity(security),
			FilteredSync:           filteredSync,
			LdapsSettings:          expandArmDomainServiceLdapsSettings(ldaps),
			NotificationSettings:   expandArmDomainServiceNotificationSettings(notifications),
			SubnetID:               &subnetId,
		},
	}

	future, err := client.Update(ctx, resourceGroup, name, domainService)
	if err != nil {
		return fmt.Errorf("Error updating Domain Service %q (Resource Group %q): %+v", name, resourceGroup, err)
	}
	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for update of Domain Service %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	return resourceArmDomainServiceRead(d, meta)
}

func resourceArmDomainServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).domainServicesClient
	ctx := meta.(*ArmClient).StopContext

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	name := id.Path["domainServices"]

	future, err := client.Delete(ctx, resourceGroup, name)
	if err != nil {
		if response.WasNotFound(future.Response()) {
			return nil
		}
		return fmt.Errorf("Error deleting Domain Service %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		if !response.WasNotFound(future.Response()) {
			return fmt.Errorf("Error waiting for deleting Domain Service %q (Resource Group %q): %+v", name, resourceGroup, err)
		}
	}

	return nil
}

func domainServiceControllerRefreshFunc(ctx context.Context, client *aad.DomainServicesClient, resourceGroup, name string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Waiting for both Domain Service Controller deploying...")
		resp, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			return nil, "error", err
		}
		if resp.DomainControllerIPAddress == nil || len(*resp.DomainControllerIPAddress) < 2 {
			return *resp.DomainControllerIPAddress, "pending", nil
		}
		return *resp.DomainControllerIPAddress, "available", nil
	}
}

func expandArmDomainServiceSecurity(input []interface{}) *aad.DomainSecuritySettings {
	if len(input) == 0 {
		return nil
	}
	v := input[0].(map[string]interface{})

	ntlmV1 := aad.NtlmV1Disabled
	syncNtlmPasswords := aad.SyncNtlmPasswordsDisabled
	tlsV1 := aad.TLSV1Disabled

	if v["ntlm_v1"].(bool) {
		ntlmV1 = aad.NtlmV1Enabled
	}
	if v["sync_ntlm_passwords"].(bool) {
		syncNtlmPasswords = aad.SyncNtlmPasswordsEnabled
	}
	if v["tls_v1"].(bool) {
		tlsV1 = aad.TLSV1Enabled
	}

	return &aad.DomainSecuritySettings{
		NtlmV1:            ntlmV1,
		SyncNtlmPasswords: syncNtlmPasswords,
		TLSV1:             tlsV1,
	}
}

func expandArmDomainServiceLdapsSettings(input []interface{}) *aad.LdapsSettings {
	if len(input) == 0 {
		return nil
	}
	v := input[0].(map[string]interface{})

	externalAccess := aad.Disabled
	ldaps := aad.LdapsDisabled
	pfxCertificate := v["pfx_certificate"].(string)
	pfxCertificatePassword := v["pfx_certificate_password"].(string)

	if v["external_access"].(bool) {
		externalAccess = aad.Enabled
	}
	if v["ldaps"].(bool) {
		ldaps = aad.LdapsEnabled
	}
	return &aad.LdapsSettings{
		ExternalAccess:         externalAccess,
		Ldaps:                  ldaps,
		PfxCertificate:         &pfxCertificate,
		PfxCertificatePassword: &pfxCertificatePassword,
	}
}

func expandArmDomainServiceNotificationSettings(input []interface{}) *aad.NotificationSettings {
	if len(input) == 0 {
		return nil
	}
	v := input[0].(map[string]interface{})

	additionalRecipients := v["additional_recipients"].([]interface{})
	notifyDcAdmins := aad.NotifyDcAdminsDisabled
	notifyGlobalAdmins := aad.NotifyGlobalAdminsDisabled
	if v["notify_dc_admins"].(bool) {
		notifyDcAdmins = aad.NotifyDcAdminsEnabled
	}
	if v["notify_global_admins"].(bool) {
		notifyGlobalAdmins = aad.NotifyGlobalAdminsEnabled
	}

	return &aad.NotificationSettings{
		AdditionalRecipients: tf.ExpandStringSlicePtr(additionalRecipients),
		NotifyDcAdmins:       notifyDcAdmins,
		NotifyGlobalAdmins:   notifyGlobalAdmins,
	}
}

func flattenArmDomainServiceSecurity(input *aad.DomainSecuritySettings) []interface{} {
	if input == nil {
		return make([]interface{}, 0)
	}

	result := map[string]bool{
		"ntlm_v1":             false,
		"sync_ntlm_passwords": false,
		"tls_v1":              false,
	}
	if input.NtlmV1 == aad.NtlmV1Enabled {
		result["ntlm_v1"] = true
	}
	if input.SyncNtlmPasswords == aad.SyncNtlmPasswordsEnabled {
		result["sync_ntlm_passwords"] = true
	}
	if input.TLSV1 == aad.TLSV1Enabled {
		result["tls_v1"] = true
	}

	return []interface{}{result}
}

func flattenArmDomainServiceLdaps(input *aad.LdapsSettings) []interface{} {
	if input == nil {
		return make([]interface{}, 0)
	}

	result := map[string]interface{}{
		"external_access": false,
		"ldaps":           false,
	}

	if input.ExternalAccess == aad.Enabled {
		result["external_access"] = true
	}
	if input.Ldaps == aad.LdapsEnabled {
		result["ldaps"] = true
	}
	if pfxCertificate := input.PfxCertificate; pfxCertificate != nil {
		result["pfx_certificate"] = *pfxCertificate
	}
	if pfxCertificatePassword := input.PfxCertificatePassword; pfxCertificatePassword != nil {
		result["pfx_certificate_password"] = *pfxCertificatePassword
	}
	if externalAccessIPAddress := input.ExternalAccessIPAddress; externalAccessIPAddress != nil {
		result["external_access_ip_address"] = *externalAccessIPAddress
	}

	return []interface{}{result}
}

func flattenArmDomainServiceNotification(input *aad.NotificationSettings) []interface{} {
	if input == nil {
		return make([]interface{}, 0)
	}

	result := map[string]interface{}{
		"notify_dc_admins":     false,
		"notify_global_admins": false,
	}

	result["additional_recipients"] = tf.FlattenStringSlicePtr(input.AdditionalRecipients)
	if input.NotifyDcAdmins == aad.NotifyDcAdminsEnabled {
		result["notify_dc_admins"] = true
	}
	if input.NotifyGlobalAdmins == aad.NotifyGlobalAdminsEnabled {
		result["notify_global_admins"] = true
	}

	return []interface{}{result}
}
