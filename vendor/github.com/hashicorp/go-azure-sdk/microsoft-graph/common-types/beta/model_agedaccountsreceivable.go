package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgedAccountsReceivable struct {
	AgedAsOfDate   nullable.Type[string]  `json:"agedAsOfDate,omitempty"`
	BalanceDue     nullable.Type[float64] `json:"balanceDue,omitempty"`
	CurrencyCode   nullable.Type[string]  `json:"currencyCode,omitempty"`
	CurrentAmount  nullable.Type[float64] `json:"currentAmount,omitempty"`
	CustomerId     *string                `json:"customerId,omitempty"`
	CustomerNumber nullable.Type[string]  `json:"customerNumber,omitempty"`
	Id             *string                `json:"id,omitempty"`
	Name           nullable.Type[string]  `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Period1Amount      nullable.Type[float64] `json:"period1Amount,omitempty"`
	Period2Amount      nullable.Type[float64] `json:"period2Amount,omitempty"`
	Period3Amount      nullable.Type[float64] `json:"period3Amount,omitempty"`
	PeriodLengthFilter nullable.Type[string]  `json:"periodLengthFilter,omitempty"`
}
