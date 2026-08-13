package entitlementmanagementaccesspackage

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const defaultApiVersion = "beta"

func userAgent() string {
	return "hashicorp/go-azure-sdk/entitlementmanagementaccesspackage/beta"
}

func AzureAPIVersion() string {
	return defaultApiVersion
}
