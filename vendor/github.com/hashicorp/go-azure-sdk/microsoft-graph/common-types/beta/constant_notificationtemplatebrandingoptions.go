package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationTemplateBrandingOptions string

const (
	NotificationTemplateBrandingOptions_IncludeCompanyLogo        NotificationTemplateBrandingOptions = "includeCompanyLogo"
	NotificationTemplateBrandingOptions_IncludeCompanyName        NotificationTemplateBrandingOptions = "includeCompanyName"
	NotificationTemplateBrandingOptions_IncludeCompanyPortalLink  NotificationTemplateBrandingOptions = "includeCompanyPortalLink"
	NotificationTemplateBrandingOptions_IncludeContactInformation NotificationTemplateBrandingOptions = "includeContactInformation"
	NotificationTemplateBrandingOptions_IncludeDeviceDetails      NotificationTemplateBrandingOptions = "includeDeviceDetails"
	NotificationTemplateBrandingOptions_None                      NotificationTemplateBrandingOptions = "none"
)

func PossibleValuesForNotificationTemplateBrandingOptions() []string {
	return []string{
		string(NotificationTemplateBrandingOptions_IncludeCompanyLogo),
		string(NotificationTemplateBrandingOptions_IncludeCompanyName),
		string(NotificationTemplateBrandingOptions_IncludeCompanyPortalLink),
		string(NotificationTemplateBrandingOptions_IncludeContactInformation),
		string(NotificationTemplateBrandingOptions_IncludeDeviceDetails),
		string(NotificationTemplateBrandingOptions_None),
	}
}

func (s *NotificationTemplateBrandingOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNotificationTemplateBrandingOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNotificationTemplateBrandingOptions(input string) (*NotificationTemplateBrandingOptions, error) {
	vals := map[string]NotificationTemplateBrandingOptions{
		"includecompanylogo":        NotificationTemplateBrandingOptions_IncludeCompanyLogo,
		"includecompanyname":        NotificationTemplateBrandingOptions_IncludeCompanyName,
		"includecompanyportallink":  NotificationTemplateBrandingOptions_IncludeCompanyPortalLink,
		"includecontactinformation": NotificationTemplateBrandingOptions_IncludeContactInformation,
		"includedevicedetails":      NotificationTemplateBrandingOptions_IncludeDeviceDetails,
		"none":                      NotificationTemplateBrandingOptions_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotificationTemplateBrandingOptions(input)
	return &out, nil
}
