/*
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

import jetbrains.buildServer.configs.kotlin.ParametrizedWithType

class ClientConfiguration(var clientId: String,
                          var clientSecret: String,
                          val tenantId : String,
                          val vcsRootId : String)

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
}