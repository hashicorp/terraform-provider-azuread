package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferenceAttachmentProvider string

const (
	ReferenceAttachmentProvider_Dropbox          ReferenceAttachmentProvider = "dropbox"
	ReferenceAttachmentProvider_OneDriveBusiness ReferenceAttachmentProvider = "oneDriveBusiness"
	ReferenceAttachmentProvider_OneDriveConsumer ReferenceAttachmentProvider = "oneDriveConsumer"
	ReferenceAttachmentProvider_Other            ReferenceAttachmentProvider = "other"
)

func PossibleValuesForReferenceAttachmentProvider() []string {
	return []string{
		string(ReferenceAttachmentProvider_Dropbox),
		string(ReferenceAttachmentProvider_OneDriveBusiness),
		string(ReferenceAttachmentProvider_OneDriveConsumer),
		string(ReferenceAttachmentProvider_Other),
	}
}

func (s *ReferenceAttachmentProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReferenceAttachmentProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReferenceAttachmentProvider(input string) (*ReferenceAttachmentProvider, error) {
	vals := map[string]ReferenceAttachmentProvider{
		"dropbox":          ReferenceAttachmentProvider_Dropbox,
		"onedrivebusiness": ReferenceAttachmentProvider_OneDriveBusiness,
		"onedriveconsumer": ReferenceAttachmentProvider_OneDriveConsumer,
		"other":            ReferenceAttachmentProvider_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReferenceAttachmentProvider(input)
	return &out, nil
}
