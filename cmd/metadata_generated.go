// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import "github.com/SAP/jenkins-library/pkg/config"

// GetStepMetadata return a map with all the step metadata mapped to their stepName
func GetAllStepMetadata() map[string]config.StepData {
	return map[string]config.StepData{
		"abapAddonAssemblyKitCheckCVs":            abapAddonAssemblyKitCheckCVsMetadata(),
		"abapAddonAssemblyKitCheckPV":             abapAddonAssemblyKitCheckPVMetadata(),
		"abapAddonAssemblyKitCreateTargetVector":  abapAddonAssemblyKitCreateTargetVectorMetadata(),
		"abapAddonAssemblyKitPublishTargetVector": abapAddonAssemblyKitPublishTargetVectorMetadata(),
		"abapAddonAssemblyKitRegisterPackages":    abapAddonAssemblyKitRegisterPackagesMetadata(),
		"abapAddonAssemblyKitReleasePackages":     abapAddonAssemblyKitReleasePackagesMetadata(),
		"abapAddonAssemblyKitReserveNextPackages": abapAddonAssemblyKitReserveNextPackagesMetadata(),
		"abapEnvironmentAssemblePackages":         abapEnvironmentAssemblePackagesMetadata(),
		"abapEnvironmentCheckoutBranch":           abapEnvironmentCheckoutBranchMetadata(),
		"abapEnvironmentCloneGitRepo":             abapEnvironmentCloneGitRepoMetadata(),
		"abapEnvironmentCreateSystem":             abapEnvironmentCreateSystemMetadata(),
		"abapEnvironmentPullGitRepo":              abapEnvironmentPullGitRepoMetadata(),
		"abapEnvironmentRunATCCheck":              abapEnvironmentRunATCCheckMetadata(),
		"checkmarxExecuteScan":                    checkmarxExecuteScanMetadata(),
		"cloudFoundryCreateService":               cloudFoundryCreateServiceMetadata(),
		"cloudFoundryCreateServiceKey":            cloudFoundryCreateServiceKeyMetadata(),
		"cloudFoundryCreateSpace":                 cloudFoundryCreateSpaceMetadata(),
		"cloudFoundryDeleteService":               cloudFoundryDeleteServiceMetadata(),
		"cloudFoundryDeleteSpace":                 cloudFoundryDeleteSpaceMetadata(),
		"cloudFoundryDeploy":                      cloudFoundryDeployMetadata(),
		"detectExecuteScan":                       detectExecuteScanMetadata(),
		"fortifyExecuteScan":                      fortifyExecuteScanMetadata(),
		"gctsCloneRepository":                     gctsCloneRepositoryMetadata(),
		"gctsCreateRepository":                    gctsCreateRepositoryMetadata(),
		"gctsDeploy":                              gctsDeployMetadata(),
		"gctsExecuteABAPUnitTests":                gctsExecuteABAPUnitTestsMetadata(),
		"gctsRollback":                            gctsRollbackMetadata(),
		"githubCheckBranchProtection":             githubCheckBranchProtectionMetadata(),
		"githubCreatePullRequest":                 githubCreatePullRequestMetadata(),
		"githubPublishRelease":                    githubPublishReleaseMetadata(),
		"githubSetCommitStatus":                   githubSetCommitStatusMetadata(),
		"gitopsUpdateDeployment":                  gitopsUpdateDeploymentMetadata(),
		"hadolintExecute":                         hadolintExecuteMetadata(),
		"jsonApplyPatch":                          jsonApplyPatchMetadata(),
		"kanikoExecute":                           kanikoExecuteMetadata(),
		"karmaExecuteTests":                       karmaExecuteTestsMetadata(),
		"kubernetesDeploy":                        kubernetesDeployMetadata(),
		"malwareExecuteScan":                      malwareExecuteScanMetadata(),
		"mavenBuild":                              mavenBuildMetadata(),
		"mavenExecute":                            mavenExecuteMetadata(),
		"mavenExecuteIntegration":                 mavenExecuteIntegrationMetadata(),
		"mavenExecuteStaticCodeChecks":            mavenExecuteStaticCodeChecksMetadata(),
		"mtaBuild":                                mtaBuildMetadata(),
		"nexusUpload":                             nexusUploadMetadata(),
		"npmExecuteLint":                          npmExecuteLintMetadata(),
		"npmExecuteScripts":                       npmExecuteScriptsMetadata(),
		"protecodeExecuteScan":                    protecodeExecuteScanMetadata(),
		"containerSaveImage":                      containerSaveImageMetadata(),
		"sonarExecuteScan":                        sonarExecuteScanMetadata(),
		"vaultRotateSecretId":                     vaultRotateSecretIdMetadata(),
		"artifactPrepareVersion":                  artifactPrepareVersionMetadata(),
		"whitesourceExecuteScan":                  whitesourceExecuteScanMetadata(),
		"xsDeploy":                                xsDeployMetadata(),
	}
}
