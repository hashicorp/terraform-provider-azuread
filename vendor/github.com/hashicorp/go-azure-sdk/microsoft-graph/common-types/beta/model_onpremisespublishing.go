package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesPublishing struct {
	// If you're configuring a traffic manager in front of multiple app proxy applications, this user-friendly URL points to
	// the traffic manager.
	AlternateUrl nullable.Type[string] `json:"alternateUrl,omitempty"`

	// The duration the connector waits for a response from the backend application before closing the connection. Possible
	// values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set
	// to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to
	// respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default
	// value is default.
	ApplicationServerTimeout nullable.Type[string] `json:"applicationServerTimeout,omitempty"`

	// System-defined value that indicates whether this application is an application proxy configured application. The
	// possible values are quickaccessapp and nonwebapp. Read-only.
	ApplicationType nullable.Type[string] `json:"applicationType,omitempty"`

	// Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate
	// before accessing the app. Pass through doesn't require authentication. Possible values are: passthru,
	// aadPreAuthentication.
	ExternalAuthenticationType *ExternalAuthenticationType `json:"externalAuthenticationType,omitempty"`

	// The published external URL for the application. For example, https://intranet-contoso.msappproxy.net/.
	ExternalUrl nullable.Type[string] `json:"externalUrl,omitempty"`

	// The internal url of the application. For example, https://intranet/.
	InternalUrl nullable.Type[string] `json:"internalUrl,omitempty"`

	// Indicates whether the application is accessible via a Global Secure Access client on a managed device.
	IsAccessibleViaZTNAClient nullable.Type[bool] `json:"isAccessibleViaZTNAClient,omitempty"`

	// Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy
	// apps, the property is set to true by default. For all existing apps, the property is set to false.
	IsBackendCertificateValidationEnabled nullable.Type[bool] `json:"isBackendCertificateValidationEnabled,omitempty"`

	// Indicates Microsoft Entra Private Access should handle DNS resolution. false by default.
	IsDnsResolutionEnabled nullable.Type[bool] `json:"isDnsResolutionEnabled,omitempty"`

	// Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have
	// Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services,
	// set this value to False. Default value is false.
	IsHttpOnlyCookieEnabled nullable.Type[bool] `json:"isHttpOnlyCookieEnabled,omitempty"`

	// Indicates if the application is currently being published via Application Proxy or not. This is preset by the system.
	// Read-only.
	IsOnPremPublishingEnabled nullable.Type[bool] `json:"isOnPremPublishingEnabled,omitempty"`

	// Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false.
	// Only use this setting for applications that can't share cookies between processes. For more information about cookie
	// settings, see Cookie settings for accessing on-premises applications in Microsoft Entra ID. Default value is false.
	IsPersistentCookieEnabled nullable.Type[bool] `json:"isPersistentCookieEnabled,omitempty"`

	// Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit
	// cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
	IsSecureCookieEnabled nullable.Type[bool] `json:"isSecureCookieEnabled,omitempty"`

	// Indicates whether validation of the state parameter when the client uses the OAuth 2.0 authorization code grant flow
	// is enabled. This setting allows admins to specify whether they want to enable CSRF protection for their apps.
	IsStateSessionEnabled nullable.Type[bool] `json:"isStateSessionEnabled,omitempty"`

	// Indicates if the application should translate urls in the response headers. Keep this value as true unless your
	// application required the original host header in the authentication request. Default value is true.
	IsTranslateHostHeaderEnabled nullable.Type[bool] `json:"isTranslateHostHeaderEnabled,omitempty"`

	// Indicates if the application should translate urls in the application body. Keep this value as false unless you have
	// hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link
	// translation with Application Proxy. Default value is false.
	IsTranslateLinksInBodyEnabled nullable.Type[bool] `json:"isTranslateLinksInBodyEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the application segment collection for an on-premises wildcard application. This property is deprecated
	// and will stop returning data on June 1, 2023. Use segmentsConfiguration instead.
	OnPremisesApplicationSegments *[]OnPremisesApplicationSegment `json:"onPremisesApplicationSegments,omitempty"`

	// Represents the collection of application segments for an on-premises wildcard application that's published through
	// Microsoft Entra application proxy.
	SegmentsConfiguration SegmentConfiguration `json:"segmentsConfiguration"`

	// Represents the single sign-on configuration for the on-premises application.
	SingleSignOnSettings *OnPremisesPublishingSingleSignOn `json:"singleSignOnSettings,omitempty"`

	// Indicates whether the application should use alternateUrl instead of externalUrl.
	UseAlternateUrlForTranslationAndRedirect nullable.Type[bool] `json:"useAlternateUrlForTranslationAndRedirect,omitempty"`

	// Details of the certificate associated with the application when a custom domain is in use. null when using the
	// default domain. Read-only.
	VerifiedCustomDomainCertificatesMetadata *VerifiedCustomDomainCertificatesMetadata `json:"verifiedCustomDomainCertificatesMetadata,omitempty"`

	// The associated key credential for the custom domain used.
	VerifiedCustomDomainKeyCredential *KeyCredential `json:"verifiedCustomDomainKeyCredential,omitempty"`

	// The associated password credential for the custom domain used.
	VerifiedCustomDomainPasswordCredential *PasswordCredential `json:"verifiedCustomDomainPasswordCredential,omitempty"`

	WafAllowedHeaders *WafAllowedHeadersDictionary `json:"wafAllowedHeaders,omitempty"`
	WafIPRanges       *[]IPRange                   `json:"wafIpRanges,omitempty"`
	WafProvider       nullable.Type[string]        `json:"wafProvider,omitempty"`
}

var _ json.Marshaler = OnPremisesPublishing{}

func (s OnPremisesPublishing) MarshalJSON() ([]byte, error) {
	type wrapper OnPremisesPublishing
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnPremisesPublishing: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnPremisesPublishing: %+v", err)
	}

	delete(decoded, "applicationType")
	delete(decoded, "isOnPremPublishingEnabled")
	delete(decoded, "verifiedCustomDomainCertificatesMetadata")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnPremisesPublishing: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OnPremisesPublishing{}

func (s *OnPremisesPublishing) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AlternateUrl                             nullable.Type[string]                     `json:"alternateUrl,omitempty"`
		ApplicationServerTimeout                 nullable.Type[string]                     `json:"applicationServerTimeout,omitempty"`
		ApplicationType                          nullable.Type[string]                     `json:"applicationType,omitempty"`
		ExternalAuthenticationType               *ExternalAuthenticationType               `json:"externalAuthenticationType,omitempty"`
		ExternalUrl                              nullable.Type[string]                     `json:"externalUrl,omitempty"`
		InternalUrl                              nullable.Type[string]                     `json:"internalUrl,omitempty"`
		IsAccessibleViaZTNAClient                nullable.Type[bool]                       `json:"isAccessibleViaZTNAClient,omitempty"`
		IsBackendCertificateValidationEnabled    nullable.Type[bool]                       `json:"isBackendCertificateValidationEnabled,omitempty"`
		IsDnsResolutionEnabled                   nullable.Type[bool]                       `json:"isDnsResolutionEnabled,omitempty"`
		IsHttpOnlyCookieEnabled                  nullable.Type[bool]                       `json:"isHttpOnlyCookieEnabled,omitempty"`
		IsOnPremPublishingEnabled                nullable.Type[bool]                       `json:"isOnPremPublishingEnabled,omitempty"`
		IsPersistentCookieEnabled                nullable.Type[bool]                       `json:"isPersistentCookieEnabled,omitempty"`
		IsSecureCookieEnabled                    nullable.Type[bool]                       `json:"isSecureCookieEnabled,omitempty"`
		IsStateSessionEnabled                    nullable.Type[bool]                       `json:"isStateSessionEnabled,omitempty"`
		IsTranslateHostHeaderEnabled             nullable.Type[bool]                       `json:"isTranslateHostHeaderEnabled,omitempty"`
		IsTranslateLinksInBodyEnabled            nullable.Type[bool]                       `json:"isTranslateLinksInBodyEnabled,omitempty"`
		ODataId                                  *string                                   `json:"@odata.id,omitempty"`
		ODataType                                *string                                   `json:"@odata.type,omitempty"`
		OnPremisesApplicationSegments            *[]OnPremisesApplicationSegment           `json:"onPremisesApplicationSegments,omitempty"`
		SingleSignOnSettings                     *OnPremisesPublishingSingleSignOn         `json:"singleSignOnSettings,omitempty"`
		UseAlternateUrlForTranslationAndRedirect nullable.Type[bool]                       `json:"useAlternateUrlForTranslationAndRedirect,omitempty"`
		VerifiedCustomDomainCertificatesMetadata *VerifiedCustomDomainCertificatesMetadata `json:"verifiedCustomDomainCertificatesMetadata,omitempty"`
		VerifiedCustomDomainKeyCredential        *KeyCredential                            `json:"verifiedCustomDomainKeyCredential,omitempty"`
		VerifiedCustomDomainPasswordCredential   *PasswordCredential                       `json:"verifiedCustomDomainPasswordCredential,omitempty"`
		WafAllowedHeaders                        *WafAllowedHeadersDictionary              `json:"wafAllowedHeaders,omitempty"`
		WafProvider                              nullable.Type[string]                     `json:"wafProvider,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AlternateUrl = decoded.AlternateUrl
	s.ApplicationServerTimeout = decoded.ApplicationServerTimeout
	s.ApplicationType = decoded.ApplicationType
	s.ExternalAuthenticationType = decoded.ExternalAuthenticationType
	s.ExternalUrl = decoded.ExternalUrl
	s.InternalUrl = decoded.InternalUrl
	s.IsAccessibleViaZTNAClient = decoded.IsAccessibleViaZTNAClient
	s.IsBackendCertificateValidationEnabled = decoded.IsBackendCertificateValidationEnabled
	s.IsDnsResolutionEnabled = decoded.IsDnsResolutionEnabled
	s.IsHttpOnlyCookieEnabled = decoded.IsHttpOnlyCookieEnabled
	s.IsOnPremPublishingEnabled = decoded.IsOnPremPublishingEnabled
	s.IsPersistentCookieEnabled = decoded.IsPersistentCookieEnabled
	s.IsSecureCookieEnabled = decoded.IsSecureCookieEnabled
	s.IsStateSessionEnabled = decoded.IsStateSessionEnabled
	s.IsTranslateHostHeaderEnabled = decoded.IsTranslateHostHeaderEnabled
	s.IsTranslateLinksInBodyEnabled = decoded.IsTranslateLinksInBodyEnabled
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.OnPremisesApplicationSegments = decoded.OnPremisesApplicationSegments
	s.SingleSignOnSettings = decoded.SingleSignOnSettings
	s.UseAlternateUrlForTranslationAndRedirect = decoded.UseAlternateUrlForTranslationAndRedirect
	s.VerifiedCustomDomainCertificatesMetadata = decoded.VerifiedCustomDomainCertificatesMetadata
	s.VerifiedCustomDomainKeyCredential = decoded.VerifiedCustomDomainKeyCredential
	s.VerifiedCustomDomainPasswordCredential = decoded.VerifiedCustomDomainPasswordCredential
	s.WafAllowedHeaders = decoded.WafAllowedHeaders
	s.WafProvider = decoded.WafProvider

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OnPremisesPublishing into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["segmentsConfiguration"]; ok {
		impl, err := UnmarshalSegmentConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SegmentsConfiguration' for 'OnPremisesPublishing': %+v", err)
		}
		s.SegmentsConfiguration = impl
	}

	if v, ok := temp["wafIpRanges"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling WafIPRanges into list []json.RawMessage: %+v", err)
		}

		output := make([]IPRange, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIPRangeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'WafIPRanges' for 'OnPremisesPublishing': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.WafIPRanges = &output
	}

	return nil
}
