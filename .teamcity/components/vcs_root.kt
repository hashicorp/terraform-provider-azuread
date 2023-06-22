/*
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot

object providerRepository : GitVcsRoot({
    name = "terraform-provider-azuread"
    url = "https://github.com/hashicorp/terraform-provider-azuread.git"
    agentCleanPolicy = AgentCleanPolicy.ALWAYS
    agentCleanFilesPolicy = AgentCleanFilesPolicy.ALL_UNTRACKED
    branchSpec = "+:*"
    branch = "refs/heads/main"
    authMethod = anonymous()
})
