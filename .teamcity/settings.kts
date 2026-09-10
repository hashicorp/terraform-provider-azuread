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
var vcsRootId = DslContext.getParameter("vcsRootId", "TF_HashiCorp_AzureAD_Repository")
var skuId = DslContext.getParameter("skuId", "")
var disabledPlanId = DslContext.getParameter("disabledPlanId", "")
var gitHubRepo = DslContext.getParameter("gitHubRepo", "hashicorp/terraform-provider-azuread")
var gitPat = DslContext.getParameter("gitPat", "")
var teamcityToken = DslContext.getParameter("teamcityToken", "")
var labelSuccess = DslContext.getParameter("labelSuccess", "teamcity-passed")
var labelFailure = DslContext.getParameter("labelFailure", "teamcity-failed")
var labelOutdated = DslContext.getParameter("labelOutdated", "teamcity-outdated")
var labelNewFailure = DslContext.getParameter("labelNewFailure", "teamcity-new-failure")
var applyTestingLabelsEnabled = DslContext.getParameter("applyTestingLabelsEnabled", "false").equals("true", ignoreCase = true)

var clientConfig = ClientConfiguration(
    clientId,
    clientSecret,
    tenantId,
    vcsRootId,
    skuId,
    disabledPlanId,
    gitHubRepo,
    gitPat,
    teamcityToken,
    labelSuccess,
    labelFailure,
    labelOutdated,
    labelNewFailure,
    applyTestingLabelsEnabled
)

project(AzureAD(environment, clientConfig))
