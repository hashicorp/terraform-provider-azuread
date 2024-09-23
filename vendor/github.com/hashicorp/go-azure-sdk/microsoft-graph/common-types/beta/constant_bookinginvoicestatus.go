package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingInvoiceStatus string

const (
	BookingInvoiceStatus_Canceled   BookingInvoiceStatus = "canceled"
	BookingInvoiceStatus_Corrective BookingInvoiceStatus = "corrective"
	BookingInvoiceStatus_Draft      BookingInvoiceStatus = "draft"
	BookingInvoiceStatus_Open       BookingInvoiceStatus = "open"
	BookingInvoiceStatus_Paid       BookingInvoiceStatus = "paid"
	BookingInvoiceStatus_Reviewing  BookingInvoiceStatus = "reviewing"
)

func PossibleValuesForBookingInvoiceStatus() []string {
	return []string{
		string(BookingInvoiceStatus_Canceled),
		string(BookingInvoiceStatus_Corrective),
		string(BookingInvoiceStatus_Draft),
		string(BookingInvoiceStatus_Open),
		string(BookingInvoiceStatus_Paid),
		string(BookingInvoiceStatus_Reviewing),
	}
}

func (s *BookingInvoiceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingInvoiceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingInvoiceStatus(input string) (*BookingInvoiceStatus, error) {
	vals := map[string]BookingInvoiceStatus{
		"canceled":   BookingInvoiceStatus_Canceled,
		"corrective": BookingInvoiceStatus_Corrective,
		"draft":      BookingInvoiceStatus_Draft,
		"open":       BookingInvoiceStatus_Open,
		"paid":       BookingInvoiceStatus_Paid,
		"reviewing":  BookingInvoiceStatus_Reviewing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingInvoiceStatus(input)
	return &out, nil
}
