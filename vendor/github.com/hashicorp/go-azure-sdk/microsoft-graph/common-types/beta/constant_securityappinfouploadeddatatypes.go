package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAppInfoUploadedDataTypes string

const (
	SecurityAppInfoUploadedDataTypes_CodingFiles   SecurityAppInfoUploadedDataTypes = "codingFiles"
	SecurityAppInfoUploadedDataTypes_CreditCards   SecurityAppInfoUploadedDataTypes = "creditCards"
	SecurityAppInfoUploadedDataTypes_DatabaseFiles SecurityAppInfoUploadedDataTypes = "databaseFiles"
	SecurityAppInfoUploadedDataTypes_Documents     SecurityAppInfoUploadedDataTypes = "documents"
	SecurityAppInfoUploadedDataTypes_MediaFiles    SecurityAppInfoUploadedDataTypes = "mediaFiles"
	SecurityAppInfoUploadedDataTypes_None          SecurityAppInfoUploadedDataTypes = "none"
	SecurityAppInfoUploadedDataTypes_Unknown       SecurityAppInfoUploadedDataTypes = "unknown"
)

func PossibleValuesForSecurityAppInfoUploadedDataTypes() []string {
	return []string{
		string(SecurityAppInfoUploadedDataTypes_CodingFiles),
		string(SecurityAppInfoUploadedDataTypes_CreditCards),
		string(SecurityAppInfoUploadedDataTypes_DatabaseFiles),
		string(SecurityAppInfoUploadedDataTypes_Documents),
		string(SecurityAppInfoUploadedDataTypes_MediaFiles),
		string(SecurityAppInfoUploadedDataTypes_None),
		string(SecurityAppInfoUploadedDataTypes_Unknown),
	}
}

func (s *SecurityAppInfoUploadedDataTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAppInfoUploadedDataTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAppInfoUploadedDataTypes(input string) (*SecurityAppInfoUploadedDataTypes, error) {
	vals := map[string]SecurityAppInfoUploadedDataTypes{
		"codingfiles":   SecurityAppInfoUploadedDataTypes_CodingFiles,
		"creditcards":   SecurityAppInfoUploadedDataTypes_CreditCards,
		"databasefiles": SecurityAppInfoUploadedDataTypes_DatabaseFiles,
		"documents":     SecurityAppInfoUploadedDataTypes_Documents,
		"mediafiles":    SecurityAppInfoUploadedDataTypes_MediaFiles,
		"none":          SecurityAppInfoUploadedDataTypes_None,
		"unknown":       SecurityAppInfoUploadedDataTypes_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAppInfoUploadedDataTypes(input)
	return &out, nil
}
