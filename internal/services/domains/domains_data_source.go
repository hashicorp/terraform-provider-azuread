// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package domains

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DomainsId string

func (id DomainsId) ID() string {
	return string(id)
}

func (DomainsId) String() string {
	return "Domains"
}

type DomainsDataSourceModel struct {
	AdminManaged      bool     `tfschema:"admin_managed"`
	Domains           []Domain `tfschema:"tags"`
	IncludeUnverified bool     `tfschema:"include_unverified"`
	OnlyDefault       bool     `tfschema:"only_default"`
	OnlyInitial       bool     `tfschema:"only_initial"`
	OnlyRoot          bool     `tfschema:"only_root"`
	SupportsServices  []string `tfschema:"supports_services"`
}

type Domain struct {
	AdminManaged       bool     `tfschema:"admin_managed"`
	AuthenticationType string   `tfschema:"authentication_type"`
	Default            bool     `tfschema:"default"`
	DomainName         string   `tfschema:"domain_name"`
	Initial            bool     `tfschema:"initial"`
	Root               bool     `tfschema:"root"`
	SupportedServices  []string `tfschema:"supported_services"`
	Verified           bool     `tfschema:"verified"`
}

type DomainsDataSource struct{}

var _ sdk.DataSource = DomainsDataSource{}

func (r DomainsDataSource) ResourceType() string {
	return "azurerm_aadb2c_directory"
}

func (r DomainsDataSource) ModelObject() interface{} {
	return &DomainsDataSourceModel{}
}

func (r DomainsDataSource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"admin_managed": {
			Description: "Set to `true` to only return domains whose DNS is managed by Microsoft 365",
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"include_unverified": {
			Description:   "Set to `true` if unverified Azure AD domains should be included",
			Type:          schema.TypeBool,
			Optional:      true,
			ConflictsWith: []string{"only_default", "only_initial"}, // default or initial domains have to be verified
		},

		"only_default": {
			Description:   "Set to `true` to only return the default domain",
			Type:          schema.TypeBool,
			Optional:      true,
			ConflictsWith: []string{"only_initial", "only_root"},
		},

		"only_initial": {
			Description:   "Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain",
			Type:          schema.TypeBool,
			Optional:      true,
			ConflictsWith: []string{"only_default", "only_root"},
		},

		"only_root": {
			Description:   "Set to `true` to only return verified root domains. Excludes subdomains and unverified domains",
			Type:          schema.TypeBool,
			Optional:      true,
			ConflictsWith: []string{"only_default", "only_initial"},
		},

		"supports_services": {
			Description: "A list of supported services that must be supported by a domain",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func (r DomainsDataSource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"domains": {
			Description: "A list of tenant domains",
			Type:        schema.TypeList,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"admin_managed": {
						Description: "Whether the DNS for the domain is managed by Microsoft 365",
						Type:        schema.TypeBool,
						Computed:    true,
					},

					"authentication_type": {
						Description: "The authentication type of the domain. Possible values include `Managed` or `Federated`",
						Type:        schema.TypeString,
						Computed:    true,
					},

					"default": {
						Description: "Whether this is the default domain that is used for user creation",
						Type:        schema.TypeBool,
						Computed:    true,
					},

					"domain_name": {
						Description: "The name of the domain",
						Type:        schema.TypeString,
						Computed:    true,
					},

					"initial": {
						Description: "Whether this is the initial domain created by Azure Active Directory",
						Type:        schema.TypeBool,
						Computed:    true,
					},

					"root": {
						Description: "Whether the domain is a verified root domain (not a subdomain)",
						Type:        schema.TypeBool,
						Computed:    true,
					},

					"supported_services": {
						Description: "A list of capabilities / services supported by the domain. Possible values include `Email`, `Sharepoint`, `EmailInternalRelayOnly`, `OfficeCommunicationsOnline`, `SharePointDefaultDomain`, `FullRedelegation`, `SharePointPublic`, `OrgIdAuthentication`, `Yammer` and `Intune`",
						Type:        schema.TypeList,
						Computed:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},

					"verified": {
						Description: "Whether the domain has completed domain ownership verification",
						Type:        schema.TypeBool,
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r DomainsDataSource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Domains.DomainsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			tenantId := metadata.Client.TenantID

			var state DomainsDataSourceModel
			if err := metadata.Decode(&state); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			// OData filters are not supported for domains
			result, _, err := client.List(ctx, odata.Query{})
			if err != nil {
				return fmt.Errorf("listing domains: %+v", err)
			}

			if result == nil {
				return fmt.Errorf("retrieving domains: result was nil")
			}

			var domains []Domain
			var domainNames []string

			for _, v := range *result {
				if state.AdminManaged && v.IsAdminManaged != nil && !*v.IsAdminManaged {
					continue
				}
				if state.OnlyDefault && v.IsDefault != nil && !*v.IsDefault {
					continue
				}
				if state.OnlyInitial && v.IsInitial != nil && !*v.IsInitial {
					continue
				}
				if state.OnlyRoot && v.IsRoot != nil && !*v.IsRoot {
					continue
				}
				if !state.IncludeUnverified && v.IsVerified != nil && !*v.IsVerified {
					continue
				}
				if len(state.SupportsServices) > 0 && v.SupportedServices != nil {
					supported := 0
					for _, serviceNeeded := range state.SupportsServices {
						for _, serviceSupported := range *v.SupportedServices {
							if serviceNeeded == serviceSupported {
								supported++
								break
							}
						}
					}
					if supported < len(state.SupportsServices) {
						continue
					}
				}

				if v.ID != nil {
					domainNames = append(domainNames, *v.ID)

					var authenticationType string
					if v.AuthenticationType != nil {
						authenticationType = *v.AuthenticationType
					}

					supportedServices := make([]string, 0)
					if v.SupportedServices != nil {
						supportedServices = *v.SupportedServices
					}

					domains = append(domains, Domain{
						AdminManaged:       v.IsAdminManaged != nil && *v.IsAdminManaged,
						AuthenticationType: authenticationType,
						Default:            v.IsDefault != nil && *v.IsDefault,
						DomainName:         *v.ID,
						Initial:            v.IsInitial != nil && *v.IsInitial,
						Root:               v.IsRoot != nil && *v.IsRoot,
						SupportedServices:  supportedServices,
						Verified:           v.IsVerified != nil && *v.IsVerified,
					})
				}
			}

			if len(domains) == 0 {
				return fmt.Errorf("no domains found for the provided filters")
			}

			// Generate a unique ID based on result
			h := sha1.New()
			if _, err := h.Write([]byte(strings.Join(domainNames, "/"))); err != nil {
				return fmt.Errorf("unable to compute hash for domain names: %+v", err)
			}

			metadata.SetID(DomainsId(fmt.Sprintf("domains#%s#%s", tenantId, base64.URLEncoding.EncodeToString(h.Sum(nil)))))

			return nil
		},
	}
}
