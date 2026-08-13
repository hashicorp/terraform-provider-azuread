package administrativeunit

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const defaultApiVersion = "v1.0"

func userAgent() string {
	return "hashicorp/go-azure-sdk/administrativeunit/stable"
}

func AzureAPIVersion() string {
	return defaultApiVersion
}
