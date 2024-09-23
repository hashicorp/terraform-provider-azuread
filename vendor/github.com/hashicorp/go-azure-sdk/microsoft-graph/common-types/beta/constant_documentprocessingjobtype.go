package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentProcessingJobType string

const (
	DocumentProcessingJobType_File   DocumentProcessingJobType = "file"
	DocumentProcessingJobType_Folder DocumentProcessingJobType = "folder"
)

func PossibleValuesForDocumentProcessingJobType() []string {
	return []string{
		string(DocumentProcessingJobType_File),
		string(DocumentProcessingJobType_Folder),
	}
}

func (s *DocumentProcessingJobType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDocumentProcessingJobType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDocumentProcessingJobType(input string) (*DocumentProcessingJobType, error) {
	vals := map[string]DocumentProcessingJobType{
		"file":   DocumentProcessingJobType_File,
		"folder": DocumentProcessingJobType_Folder,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DocumentProcessingJobType(input)
	return &out, nil
}
