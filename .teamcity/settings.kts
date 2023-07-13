/*
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

import AzureAD
import ClientConfiguration
import jetbrains.buildServer.configs.kotlin.*

version = "2023.05"

var clientId = DslContext.getParameter("clientId", "")
var clientSecret = DslContext.getParameter("clientSecret", "")
var tenantId = DslContext.getParameter("tenantId", "")
var environment = DslContext.getParameter("environment", "public")
var vcsRootId = DslContext.getParameter("vcsRootId", "")

var clientConfig = ClientConfiguration(clientId, clientSecret, tenantId, vcsRootId)

project(AzureAD(environment, clientConfig))
