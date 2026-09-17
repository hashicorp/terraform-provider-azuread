/*
 * Copyright IBM Corp. 2014, 2025
 * SPDX-License-Identifier: MPL-2.0
 */

import jetbrains.buildServer.configs.kotlin.ParametrizedWithType

class ClientConfiguration(var clientId: String,
                          var clientSecret: String,
                          val tenantId : String,
                          val vcsRootId : String,
                          val skuId : String,
                          val disabledPlanId : String,
                          val gitHubRepo : String,
                          val gitPat : String,
                          val teamcityToken : String,
                          val labelSuccess : String,
                          val labelFailure : String,
                          val labelOutdated : String,
                          val labelNewFailure : String,
                          val applyTestingLabelsEnabled : Boolean)

class LocationConfiguration(var primary : String, var secondary : String, var ternary : String, var rotate : Boolean)

fun ParametrizedWithType.ConfigureAzureSpecificTestParameters(environment: String, config: ClientConfiguration, locationsForEnv: LocationConfiguration) {
    hiddenPasswordVariable("env.ARM_CLIENT_ID", config.clientId, "The AppID of the Application used for Testing")
    hiddenPasswordVariable("env.ARM_CLIENT_SECRET", config.clientSecret, "The Client Secret of the Application used for Testing")
    hiddenVariable("env.ARM_ENVIRONMENT", environment, "The Azure Environment in which the tests are running")
    hiddenVariable("env.ARM_PROVIDER_DYNAMIC_TEST", "%b".format(locationsForEnv.rotate), "Should tests rotate between the supported regions?")
    hiddenPasswordVariable("env.ARM_TENANT_ID", config.tenantId, "The ID of the Azure Tenant used for Testing")
    hiddenVariable("env.ARM_TEST_LOCATION", locationsForEnv.primary, "The Primary region which should be used for testing")
    hiddenVariable("env.ARM_TEST_LOCATION_ALT", locationsForEnv.secondary, "The Primary region which should be used for testing")
    hiddenVariable("env.ARM_TEST_LOCATION_ALT2", locationsForEnv.ternary, "The Primary region which should be used for testing")
    hiddenVariable("env.AAD_TEST_LICENSE_SKU_ID", config.skuId, "The license SKU ID used for azuread_user_license acceptance tests")
    hiddenVariable("env.AAD_TEST_LICENSE_DISABLED_PLAN_ID", config.disabledPlanId, "A service plan ID within the SKU, used for azuread_user_license disabled_plans acceptance tests")
}

fun ParametrizedWithType.ConfigureGitHubCommentParameters(config: ClientConfiguration) {
    hiddenPasswordVariable("env.GIT_PAT", config.gitPat, "Personal Access Token for GitHub")
    hiddenPasswordVariable("env.TEAMCITY_TOKEN", config.teamcityToken, "Access Token for TeamCity")
    hiddenVariable("env.GITHUB_REPO", config.gitHubRepo, "GitHub Repository")
    hiddenVariable("env.LABEL_SUCCESS", config.labelSuccess, "Label applied when teamcity build passed")
    hiddenVariable("env.LABEL_FAILURE", config.labelFailure, "Label applied when teamcity build failed")
    hiddenVariable("env.LABEL_OUTDATED", config.labelOutdated, "Label applied when teamcity build is outdated")
    hiddenVariable("env.LABEL_NEW_FAILURE", config.labelNewFailure, "Label applied when teamcity build has new failures")
    hiddenVariable("env.APPLY_TESTING_LABELS_ENABLED", config.applyTestingLabelsEnabled.toString(), "Whether to apply testing labels to PRs")
}