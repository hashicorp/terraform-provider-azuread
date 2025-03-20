/*
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

// specifies the default hour (UTC) at which tests should be triggered, if enabled
var defaultStartHour = 0

// specifies the default level of parallelism per-service-package
var defaultParallelism = 20

// specifies the default version of Terraform Core which should be used for testing
var defaultTerraformCoreVersion = "1.11.2"

var locations = mapOf(
        "public" to LocationConfiguration("westeurope", "eastus2", "francecentral", false)
)

// specifies the list of Azure Environments where tests should be run nightly
var runNightly = mapOf(
        "public" to true
)
