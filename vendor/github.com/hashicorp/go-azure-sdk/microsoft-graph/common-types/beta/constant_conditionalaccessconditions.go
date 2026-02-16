package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessConditions string

const (
	ConditionalAccessConditions_Application                     ConditionalAccessConditions = "application"
	ConditionalAccessConditions_AuthenticationFlows             ConditionalAccessConditions = "authenticationFlows"
	ConditionalAccessConditions_Client                          ConditionalAccessConditions = "client"
	ConditionalAccessConditions_ClientType                      ConditionalAccessConditions = "clientType"
	ConditionalAccessConditions_DevicePlatform                  ConditionalAccessConditions = "devicePlatform"
	ConditionalAccessConditions_DeviceState                     ConditionalAccessConditions = "deviceState"
	ConditionalAccessConditions_IPAddressSeenByAzureAD          ConditionalAccessConditions = "ipAddressSeenByAzureAD"
	ConditionalAccessConditions_IPAddressSeenByResourceProvider ConditionalAccessConditions = "ipAddressSeenByResourceProvider"
	ConditionalAccessConditions_InsiderRisk                     ConditionalAccessConditions = "insiderRisk"
	ConditionalAccessConditions_Location                        ConditionalAccessConditions = "location"
	ConditionalAccessConditions_None                            ConditionalAccessConditions = "none"
	ConditionalAccessConditions_ServicePrincipalRisk            ConditionalAccessConditions = "servicePrincipalRisk"
	ConditionalAccessConditions_ServicePrincipals               ConditionalAccessConditions = "servicePrincipals"
	ConditionalAccessConditions_SignInRisk                      ConditionalAccessConditions = "signInRisk"
	ConditionalAccessConditions_Time                            ConditionalAccessConditions = "time"
	ConditionalAccessConditions_UserRisk                        ConditionalAccessConditions = "userRisk"
	ConditionalAccessConditions_Users                           ConditionalAccessConditions = "users"
)

func PossibleValuesForConditionalAccessConditions() []string {
	return []string{
		string(ConditionalAccessConditions_Application),
		string(ConditionalAccessConditions_AuthenticationFlows),
		string(ConditionalAccessConditions_Client),
		string(ConditionalAccessConditions_ClientType),
		string(ConditionalAccessConditions_DevicePlatform),
		string(ConditionalAccessConditions_DeviceState),
		string(ConditionalAccessConditions_IPAddressSeenByAzureAD),
		string(ConditionalAccessConditions_IPAddressSeenByResourceProvider),
		string(ConditionalAccessConditions_InsiderRisk),
		string(ConditionalAccessConditions_Location),
		string(ConditionalAccessConditions_None),
		string(ConditionalAccessConditions_ServicePrincipalRisk),
		string(ConditionalAccessConditions_ServicePrincipals),
		string(ConditionalAccessConditions_SignInRisk),
		string(ConditionalAccessConditions_Time),
		string(ConditionalAccessConditions_UserRisk),
		string(ConditionalAccessConditions_Users),
	}
}

func (s *ConditionalAccessConditions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessConditions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessConditions(input string) (*ConditionalAccessConditions, error) {
	vals := map[string]ConditionalAccessConditions{
		"application":                     ConditionalAccessConditions_Application,
		"authenticationflows":             ConditionalAccessConditions_AuthenticationFlows,
		"client":                          ConditionalAccessConditions_Client,
		"clienttype":                      ConditionalAccessConditions_ClientType,
		"deviceplatform":                  ConditionalAccessConditions_DevicePlatform,
		"devicestate":                     ConditionalAccessConditions_DeviceState,
		"ipaddressseenbyazuread":          ConditionalAccessConditions_IPAddressSeenByAzureAD,
		"ipaddressseenbyresourceprovider": ConditionalAccessConditions_IPAddressSeenByResourceProvider,
		"insiderrisk":                     ConditionalAccessConditions_InsiderRisk,
		"location":                        ConditionalAccessConditions_Location,
		"none":                            ConditionalAccessConditions_None,
		"serviceprincipalrisk":            ConditionalAccessConditions_ServicePrincipalRisk,
		"serviceprincipals":               ConditionalAccessConditions_ServicePrincipals,
		"signinrisk":                      ConditionalAccessConditions_SignInRisk,
		"time":                            ConditionalAccessConditions_Time,
		"userrisk":                        ConditionalAccessConditions_UserRisk,
		"users":                           ConditionalAccessConditions_Users,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessConditions(input)
	return &out, nil
}
