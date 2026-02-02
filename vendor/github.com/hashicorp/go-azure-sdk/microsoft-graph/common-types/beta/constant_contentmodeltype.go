package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentModelType string

const (
	ContentModelType_FreeformSelectionMethod ContentModelType = "freeformSelectionMethod"
	ContentModelType_LayoutMethod            ContentModelType = "layoutMethod"
	ContentModelType_PrebuiltContractModel   ContentModelType = "prebuiltContractModel"
	ContentModelType_PrebuiltInvoiceModel    ContentModelType = "prebuiltInvoiceModel"
	ContentModelType_PrebuiltReceiptModel    ContentModelType = "prebuiltReceiptModel"
	ContentModelType_TeachingMethod          ContentModelType = "teachingMethod"
)

func PossibleValuesForContentModelType() []string {
	return []string{
		string(ContentModelType_FreeformSelectionMethod),
		string(ContentModelType_LayoutMethod),
		string(ContentModelType_PrebuiltContractModel),
		string(ContentModelType_PrebuiltInvoiceModel),
		string(ContentModelType_PrebuiltReceiptModel),
		string(ContentModelType_TeachingMethod),
	}
}

func (s *ContentModelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContentModelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContentModelType(input string) (*ContentModelType, error) {
	vals := map[string]ContentModelType{
		"freeformselectionmethod": ContentModelType_FreeformSelectionMethod,
		"layoutmethod":            ContentModelType_LayoutMethod,
		"prebuiltcontractmodel":   ContentModelType_PrebuiltContractModel,
		"prebuiltinvoicemodel":    ContentModelType_PrebuiltInvoiceModel,
		"prebuiltreceiptmodel":    ContentModelType_PrebuiltReceiptModel,
		"teachingmethod":          ContentModelType_TeachingMethod,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentModelType(input)
	return &out, nil
}
