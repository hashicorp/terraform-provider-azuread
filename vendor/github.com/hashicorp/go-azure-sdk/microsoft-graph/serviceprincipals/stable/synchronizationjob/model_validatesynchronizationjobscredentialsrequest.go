package synchronizationjob

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateSynchronizationJobsCredentialsRequest struct {
	ApplicationIdentifier nullable.Type[string]                             `json:"applicationIdentifier,omitempty"`
	Credentials           *[]stable.SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
	TemplateId            nullable.Type[string]                             `json:"templateId,omitempty"`
	UseSavedCredentials   nullable.Type[bool]                               `json:"useSavedCredentials,omitempty"`
}
