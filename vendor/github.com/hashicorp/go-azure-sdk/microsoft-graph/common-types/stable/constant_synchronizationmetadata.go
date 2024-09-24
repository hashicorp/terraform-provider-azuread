package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationMetadata string

const (
	SynchronizationMetadata_ConfigurationFields                      SynchronizationMetadata = "ConfigurationFields"
	SynchronizationMetadata_GalleryApplicationIdentifier             SynchronizationMetadata = "GalleryApplicationIdentifier"
	SynchronizationMetadata_GalleryApplicationKey                    SynchronizationMetadata = "GalleryApplicationKey"
	SynchronizationMetadata_IsOAuthEnabled                           SynchronizationMetadata = "IsOAuthEnabled"
	SynchronizationMetadata_IsSynchronizationAgentAssignmentRequired SynchronizationMetadata = "IsSynchronizationAgentAssignmentRequired"
	SynchronizationMetadata_IsSynchronizationAgentRequired           SynchronizationMetadata = "IsSynchronizationAgentRequired"
	SynchronizationMetadata_IsSynchronizationInPreview               SynchronizationMetadata = "IsSynchronizationInPreview"
	SynchronizationMetadata_OAuthSettings                            SynchronizationMetadata = "OAuthSettings"
	SynchronizationMetadata_SynchronizationLearnMoreIbizaFwLink      SynchronizationMetadata = "SynchronizationLearnMoreIbizaFwLink"
)

func PossibleValuesForSynchronizationMetadata() []string {
	return []string{
		string(SynchronizationMetadata_ConfigurationFields),
		string(SynchronizationMetadata_GalleryApplicationIdentifier),
		string(SynchronizationMetadata_GalleryApplicationKey),
		string(SynchronizationMetadata_IsOAuthEnabled),
		string(SynchronizationMetadata_IsSynchronizationAgentAssignmentRequired),
		string(SynchronizationMetadata_IsSynchronizationAgentRequired),
		string(SynchronizationMetadata_IsSynchronizationInPreview),
		string(SynchronizationMetadata_OAuthSettings),
		string(SynchronizationMetadata_SynchronizationLearnMoreIbizaFwLink),
	}
}

func (s *SynchronizationMetadata) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationMetadata(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationMetadata(input string) (*SynchronizationMetadata, error) {
	vals := map[string]SynchronizationMetadata{
		"configurationfields":                      SynchronizationMetadata_ConfigurationFields,
		"galleryapplicationidentifier":             SynchronizationMetadata_GalleryApplicationIdentifier,
		"galleryapplicationkey":                    SynchronizationMetadata_GalleryApplicationKey,
		"isoauthenabled":                           SynchronizationMetadata_IsOAuthEnabled,
		"issynchronizationagentassignmentrequired": SynchronizationMetadata_IsSynchronizationAgentAssignmentRequired,
		"issynchronizationagentrequired":           SynchronizationMetadata_IsSynchronizationAgentRequired,
		"issynchronizationinpreview":               SynchronizationMetadata_IsSynchronizationInPreview,
		"oauthsettings":                            SynchronizationMetadata_OAuthSettings,
		"synchronizationlearnmoreibizafwlink":      SynchronizationMetadata_SynchronizationLearnMoreIbizaFwLink,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationMetadata(input)
	return &out, nil
}
