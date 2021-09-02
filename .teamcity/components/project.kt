import jetbrains.buildServer.configs.kotlin.v2019_2.BuildType
import jetbrains.buildServer.configs.kotlin.v2019_2.Project

const val providerName = "azuread"

var services = mapOf(
        "applications" to "Applications",
        "domains" to "Domains",
        "groups" to "Groups",
        "invitations" to "Invitations",
        "serviceprincipals" to "Service Principals",
        "users" to "Users"
)

fun AzureAD(environment: String, configuration : ClientConfiguration) : Project {
    return Project{
        vcsRoot(providerRepository)

        var pullRequestBuildConfig = pullRequestBuildConfiguration(environment, configuration)
        buildType(pullRequestBuildConfig)

        var buildConfigs = buildConfigurationsForServices(services, providerName, environment, configuration)
        buildConfigs.forEach { buildConfiguration ->
            buildType(buildConfiguration)
        }
    }
}

fun buildConfigurationsForServices(services: Map<String, String>, providerName : String, environment: String, config : ClientConfiguration): List<BuildType> {
    var list = ArrayList<BuildType>()
    var locationsForEnv = locations.get(environment)!!

    services.forEach { (serviceName, displayName) ->
        // TODO: overriding locations
        var testConfig = testConfiguration(defaultParallelism, defaultStartHour)
        var runNightly = runNightly.getOrDefault(environment, false)

        var service = serviceDetails(serviceName, displayName, environment)
        var buildConfig = service.buildConfiguration(providerName, runNightly, testConfig.startHour, testConfig.parallelism)

        buildConfig.params.ConfigureAzureSpecificTestParameters(environment, config, locationsForEnv)

        list.add(buildConfig)
    }

    return list
}

fun pullRequestBuildConfiguration(environment: String, configuration: ClientConfiguration) : BuildType {
    var locationsForEnv = locations.get(environment)!!
    var pullRequest = pullRequest("! Run Pull Request", environment)
    var buildConfiguration = pullRequest.buildConfiguration(providerName)
    buildConfiguration.params.ConfigureAzureSpecificTestParameters(environment, configuration, locationsForEnv)
    return buildConfiguration
}

class testConfiguration(parallelism: Int, startHour: Int) {
    var parallelism = parallelism
    var startHour = startHour
}
