/*
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

import jetbrains.buildServer.configs.kotlin.*

class serviceDetails(name: String, displayName: String, environment: String, vcsRootId : String) {
    val packageName = name
    val displayName = displayName
    val environment = environment
    val vcsRootId = vcsRootId

    fun buildConfiguration(providerName : String, nightlyTestsEnabled: Boolean, startHour: Int, parallelism: Int) : BuildType {
        return BuildType {
            // TC needs a consistent ID for dynamically generated packages
            id(uniqueID(providerName))

            name = "%s - Acceptance Tests".format(displayName)

            vcs {
                root(rootId = AbsoluteId(vcsRootId))
                cleanCheckout = true
            }

            steps {
                ConfigureGoEnv()
                DownloadTerraformBinary()
                RunAcceptanceTests(packageName)
            }

            failureConditions {
                errorMessage = true
            }

            features {
                Golang()
            }

            params {
                TerraformAcceptanceTestParameters(parallelism, "TestAcc", "12")
                TerraformAcceptanceTestsFlag()
                TerraformCoreBinaryTesting()
                TerraformShouldPanicForSchemaErrors()
                ReadOnlySettings()
                WorkingDirectory(packageName)
            }

            triggers {
                RunNightly(nightlyTestsEnabled, startHour)
            }
        }
    }

    fun uniqueID(provider : String) : String {
        return "%s_SERVICE_%s_%s".format(provider.toUpperCase(), environment.toUpperCase(), packageName.toUpperCase())
    }
}
