package synchronizationjob

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateSynchronizationJobCredentialsRequest struct {
	ApplicationIdentifier nullable.Type[string]                             `json:"applicationIdentifier,omitempty"`
	Credentials           *[]stable.SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
	TemplateId            nullable.Type[string]                             `json:"templateId,omitempty"`
	UseSavedCredentials   nullable.Type[bool]                               `json:"useSavedCredentials,omitempty"`
}
