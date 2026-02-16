package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MacOSSingleSignOnExtension = MacOSKerberosSingleSignOnExtension{}

type MacOSKerberosSingleSignOnExtension struct {
	// Gets or sets the Active Directory site.
	ActiveDirectorySiteCode nullable.Type[string] `json:"activeDirectorySiteCode,omitempty"`

	// Enables or disables whether the Kerberos extension can automatically determine its site name.
	BlockActiveDirectorySiteAutoDiscovery *bool `json:"blockActiveDirectorySiteAutoDiscovery,omitempty"`

	// Enables or disables Keychain usage.
	BlockAutomaticLogin *bool `json:"blockAutomaticLogin,omitempty"`

	// Gets or sets the Generic Security Services name of the Kerberos cache to use for this profile.
	CacheName nullable.Type[string] `json:"cacheName,omitempty"`

	// Gets or sets a list of app Bundle IDs allowed to access the Kerberos Ticket Granting Ticket.
	CredentialBundleIdAccessControlList *[]string `json:"credentialBundleIdAccessControlList,omitempty"`

	// When set to True, the credential is requested on the next matching Kerberos challenge or network state change. When
	// the credential is expired or missing, a new credential is created. Available for devices running macOS versions 12
	// and later.
	CredentialsCacheMonitored *bool `json:"credentialsCacheMonitored,omitempty"`

	// Gets or sets a list of realms for custom domain-realm mapping. Realms are case sensitive.
	DomainRealms *[]string `json:"domainRealms,omitempty"`

	// Gets or sets a list of hosts or domain names for which the app extension performs SSO.
	Domains *[]string `json:"domains,omitempty"`

	// When true, this profile's realm will be selected as the default. Necessary if multiple Kerberos-type profiles are
	// configured.
	IsDefaultRealm *bool `json:"isDefaultRealm,omitempty"`

	// When set to True, the Kerberos extension allows any apps entered with the app bundle ID, managed apps, and standard
	// Kerberos utilities, such as TicketViewer and klist, to access and use the credential. Available for devices running
	// macOS versions 12 and later.
	KerberosAppsInBundleIdACLIncluded *bool `json:"kerberosAppsInBundleIdACLIncluded,omitempty"`

	// When set to True, the Kerberos extension allows managed apps, and any apps entered with the app bundle ID to access
	// the credential. When set to False, the Kerberos extension allows all apps to access the credential. Available for
	// devices running iOS and iPadOS versions 14 and later.
	ManagedAppsInBundleIdACLIncluded *bool `json:"managedAppsInBundleIdACLIncluded,omitempty"`

	// Select how other processes use the Kerberos Extension credential.
	ModeCredentialUsed nullable.Type[string] `json:"modeCredentialUsed,omitempty"`

	// Enables or disables password changes.
	PasswordBlockModification *bool `json:"passwordBlockModification,omitempty"`

	// Gets or sets the URL that the user will be sent to when they initiate a password change.
	PasswordChangeUrl nullable.Type[string] `json:"passwordChangeUrl,omitempty"`

	// Enables or disables password syncing. This won't affect users logged in with a mobile account on macOS.
	PasswordEnableLocalSync *bool `json:"passwordEnableLocalSync,omitempty"`

	// Overrides the default password expiration in days. For most domains, this value is calculated automatically.
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// Gets or sets the number of days until the user is notified that their password will expire (default is 15).
	PasswordExpirationNotificationDays nullable.Type[int64] `json:"passwordExpirationNotificationDays,omitempty"`

	// Gets or sets the minimum number of days until a user can change their password again.
	PasswordMinimumAgeDays nullable.Type[int64] `json:"passwordMinimumAgeDays,omitempty"`

	// Gets or sets the minimum length of a password.
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Gets or sets the number of previous passwords to block.
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Enables or disables whether passwords must meet Active Directory's complexity requirements.
	PasswordRequireActiveDirectoryComplexity *bool `json:"passwordRequireActiveDirectoryComplexity,omitempty"`

	// Gets or sets a description of the password complexity requirements.
	PasswordRequirementsDescription nullable.Type[string] `json:"passwordRequirementsDescription,omitempty"`

	// Add creates an ordered list of preferred Key Distribution Centers (KDCs) to use for Kerberos traffic. This list is
	// used when the servers are not discoverable using DNS. When the servers are discoverable, the list is used for both
	// connectivity checks, and used first for Kerberos traffic. If the servers don’t respond, then the device uses DNS
	// discovery. Delete removes an existing list, and devices use DNS discovery. Available for devices running macOS
	// versions 12 and later.
	PreferredKDCs *[]string `json:"preferredKDCs,omitempty"`

	// Gets or sets the case-sensitive realm name for this profile.
	Realm *string `json:"realm,omitempty"`

	// Gets or sets whether to require authentication via Touch ID, Face ID, or a passcode to access the keychain entry.
	RequireUserPresence *bool `json:"requireUserPresence,omitempty"`

	// Text displayed to the user at the Kerberos sign in window. Available for devices running iOS and iPadOS versions 14
	// and later.
	SignInHelpText nullable.Type[string] `json:"signInHelpText,omitempty"`

	// When set to True, LDAP connections are required to use Transport Layer Security (TLS). Available for devices running
	// macOS versions 11 and later.
	TlsForLDAPRequired *bool `json:"tlsForLDAPRequired,omitempty"`

	// Gets or sets the principle user name to use for this profile. The realm name does not need to be included.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// When set to True, the user isn’t prompted to set up the Kerberos extension until the extension is enabled by the
	// admin, or a Kerberos challenge is received. Available for devices running macOS versions 11 and later.
	UserSetupDelayed *bool `json:"userSetupDelayed,omitempty"`

	// This label replaces the user name shown in the Kerberos extension. You can enter a name to match the name of your
	// company or organization. Available for devices running macOS versions 11 and later.
	UsernameLabelCustom nullable.Type[string] `json:"usernameLabelCustom,omitempty"`

	// Fields inherited from SingleSignOnExtension

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MacOSKerberosSingleSignOnExtension) MacOSSingleSignOnExtension() BaseMacOSSingleSignOnExtensionImpl {
	return BaseMacOSSingleSignOnExtensionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s MacOSKerberosSingleSignOnExtension) SingleSignOnExtension() BaseSingleSignOnExtensionImpl {
	return BaseSingleSignOnExtensionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOSKerberosSingleSignOnExtension{}

func (s MacOSKerberosSingleSignOnExtension) MarshalJSON() ([]byte, error) {
	type wrapper MacOSKerberosSingleSignOnExtension
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOSKerberosSingleSignOnExtension: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSKerberosSingleSignOnExtension: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSKerberosSingleSignOnExtension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOSKerberosSingleSignOnExtension: %+v", err)
	}

	return encoded, nil
}
