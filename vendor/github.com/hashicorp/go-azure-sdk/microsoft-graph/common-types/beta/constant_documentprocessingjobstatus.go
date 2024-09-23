package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentProcessingJobStatus string

const (
	DocumentProcessingJobStatus_Completed  DocumentProcessingJobStatus = "completed"
	DocumentProcessingJobStatus_Failed     DocumentProcessingJobStatus = "failed"
	DocumentProcessingJobStatus_InProgress DocumentProcessingJobStatus = "inProgress"
)

func PossibleValuesForDocumentProcessingJobStatus() []string {
	return []string{
		string(DocumentProcessingJobStatus_Completed),
		string(DocumentProcessingJobStatus_Failed),
		string(DocumentProcessingJobStatus_InProgress),
	}
}

func (s *DocumentProcessingJobStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDocumentProcessingJobStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDocumentProcessingJobStatus(input string) (*DocumentProcessingJobStatus, error) {
	vals := map[string]DocumentProcessingJobStatus{
		"completed":  DocumentProcessingJobStatus_Completed,
		"failed":     DocumentProcessingJobStatus_Failed,
		"inprogress": DocumentProcessingJobStatus_InProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DocumentProcessingJobStatus(input)
	return &out, nil
}
