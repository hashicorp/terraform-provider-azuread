package azuread

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataDomains() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceActiveDirectoryDomainsRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"tenant_domain_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"domains": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_initial": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_verified": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceActiveDirectoryDomainsRead(d *schema.ResourceData, meta interface{}) error {
	armClient := meta.(*ArmClient)
	client := meta.(*ArmClient).domainsClient
	ctx := meta.(*ArmClient).StopContext

	tenantDomainOnly := false
	if value, ok := d.GetOk("tenant_domain_only"); ok {
		tenantDomainOnly = value.(bool)
	}

	results, err := client.List(ctx, "")
	if err != nil {
		return fmt.Errorf("Error listing Azure AD Domains: %+v", err)
	}

	//iterate across each domain and append it to slice
	domains := make([]map[string]interface{}, 0)
	for _, v := range *results.Value {
		domain := make(map[string]interface{})

		if tenantDomainOnly && !v.AdditionalProperties["isInitial"].(bool) {
			//we only want the tenant root domain, which is always the initial domain
			//if this conditional matches, the current domain result should be skipped
			log.Printf("[DEBUG] Domain %q skipped, as we only want the tenant root domain.", *v.Name)
			continue
		}

		if v.Name != nil {
			domain["domain_name"] = *v.Name
		}
		if v.AdditionalProperties["isInitial"] != nil {
			domain["is_initial"] = v.AdditionalProperties["isInitial"].(bool)
		}
		if v.IsVerified != nil {
			domain["is_verified"] = *v.IsVerified
		}
		if v.IsDefault != nil {
			domain["is_default"] = *v.IsDefault
		}

		domains = append(domains, domain)
	}

	d.SetId("domains-" + armClient.tenantID)
	if err = d.Set("domains", domains); err != nil {
		return fmt.Errorf("Error setting `domains`: %+v", err)
	}

	return nil
}
