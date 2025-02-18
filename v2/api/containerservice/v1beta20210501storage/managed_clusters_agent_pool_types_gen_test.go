// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501storage

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_ManagedClustersAgentPool_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClustersAgentPool via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClustersAgentPool, ManagedClustersAgentPoolGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClustersAgentPool runs a test to see if a specific instance of ManagedClustersAgentPool round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClustersAgentPool(subject ManagedClustersAgentPool) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClustersAgentPool
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ManagedClustersAgentPool instances for property testing - lazily instantiated by
//ManagedClustersAgentPoolGenerator()
var managedClustersAgentPoolGenerator gopter.Gen

// ManagedClustersAgentPoolGenerator returns a generator of ManagedClustersAgentPool instances for property testing.
func ManagedClustersAgentPoolGenerator() gopter.Gen {
	if managedClustersAgentPoolGenerator != nil {
		return managedClustersAgentPoolGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagedClustersAgentPool(generators)
	managedClustersAgentPoolGenerator = gen.Struct(reflect.TypeOf(ManagedClustersAgentPool{}), generators)

	return managedClustersAgentPoolGenerator
}

// AddRelatedPropertyGeneratorsForManagedClustersAgentPool is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClustersAgentPool(gens map[string]gopter.Gen) {
	gens["Spec"] = ManagedClustersAgentPoolsSpecGenerator()
	gens["Status"] = AgentPoolStatusGenerator()
}

func Test_AgentPool_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPool_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolStatus, AgentPoolStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolStatus runs a test to see if a specific instance of AgentPool_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolStatus(subject AgentPool_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPool_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AgentPool_Status instances for property testing - lazily instantiated by AgentPoolStatusGenerator()
var agentPoolStatusGenerator gopter.Gen

// AgentPoolStatusGenerator returns a generator of AgentPool_Status instances for property testing.
// We first initialize agentPoolStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AgentPoolStatusGenerator() gopter.Gen {
	if agentPoolStatusGenerator != nil {
		return agentPoolStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolStatus(generators)
	agentPoolStatusGenerator = gen.Struct(reflect.TypeOf(AgentPool_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolStatus(generators)
	AddRelatedPropertyGeneratorsForAgentPoolStatus(generators)
	agentPoolStatusGenerator = gen.Struct(reflect.TypeOf(AgentPool_Status{}), generators)

	return agentPoolStatusGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolStatus(gens map[string]gopter.Gen) {
	gens["AvailabilityZones"] = gen.SliceOf(gen.AlphaString())
	gens["Count"] = gen.PtrOf(gen.Int())
	gens["EnableAutoScaling"] = gen.PtrOf(gen.Bool())
	gens["EnableEncryptionAtHost"] = gen.PtrOf(gen.Bool())
	gens["EnableFIPS"] = gen.PtrOf(gen.Bool())
	gens["EnableNodePublicIP"] = gen.PtrOf(gen.Bool())
	gens["EnableUltraSSD"] = gen.PtrOf(gen.Bool())
	gens["GpuInstanceProfile"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["KubeletDiskType"] = gen.PtrOf(gen.AlphaString())
	gens["MaxCount"] = gen.PtrOf(gen.Int())
	gens["MaxPods"] = gen.PtrOf(gen.Int())
	gens["MinCount"] = gen.PtrOf(gen.Int())
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["NodeImageVersion"] = gen.PtrOf(gen.AlphaString())
	gens["NodeLabels"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["NodePublicIPPrefixID"] = gen.PtrOf(gen.AlphaString())
	gens["NodeTaints"] = gen.SliceOf(gen.AlphaString())
	gens["OrchestratorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OsDiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsDiskType"] = gen.PtrOf(gen.AlphaString())
	gens["OsSKU"] = gen.PtrOf(gen.AlphaString())
	gens["OsType"] = gen.PtrOf(gen.AlphaString())
	gens["PodSubnetID"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesType"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ProximityPlacementGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetEvictionPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetPriority"] = gen.PtrOf(gen.AlphaString())
	gens["SpotMaxPrice"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["VmSize"] = gen.PtrOf(gen.AlphaString())
	gens["VnetSubnetID"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAgentPoolStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAgentPoolStatus(gens map[string]gopter.Gen) {
	gens["KubeletConfig"] = gen.PtrOf(KubeletConfigStatusGenerator())
	gens["LinuxOSConfig"] = gen.PtrOf(LinuxOSConfigStatusGenerator())
	gens["PowerState"] = gen.PtrOf(PowerStateStatusGenerator())
	gens["UpgradeSettings"] = gen.PtrOf(AgentPoolUpgradeSettingsStatusGenerator())
}

func Test_ManagedClustersAgentPools_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClustersAgentPools_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClustersAgentPoolsSpec, ManagedClustersAgentPoolsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClustersAgentPoolsSpec runs a test to see if a specific instance of ManagedClustersAgentPools_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClustersAgentPoolsSpec(subject ManagedClustersAgentPools_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClustersAgentPools_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ManagedClustersAgentPools_Spec instances for property testing - lazily instantiated by
//ManagedClustersAgentPoolsSpecGenerator()
var managedClustersAgentPoolsSpecGenerator gopter.Gen

// ManagedClustersAgentPoolsSpecGenerator returns a generator of ManagedClustersAgentPools_Spec instances for property testing.
// We first initialize managedClustersAgentPoolsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClustersAgentPoolsSpecGenerator() gopter.Gen {
	if managedClustersAgentPoolsSpecGenerator != nil {
		return managedClustersAgentPoolsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClustersAgentPoolsSpec(generators)
	managedClustersAgentPoolsSpecGenerator = gen.Struct(reflect.TypeOf(ManagedClustersAgentPools_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClustersAgentPoolsSpec(generators)
	AddRelatedPropertyGeneratorsForManagedClustersAgentPoolsSpec(generators)
	managedClustersAgentPoolsSpecGenerator = gen.Struct(reflect.TypeOf(ManagedClustersAgentPools_Spec{}), generators)

	return managedClustersAgentPoolsSpecGenerator
}

// AddIndependentPropertyGeneratorsForManagedClustersAgentPoolsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClustersAgentPoolsSpec(gens map[string]gopter.Gen) {
	gens["AvailabilityZones"] = gen.SliceOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["Count"] = gen.PtrOf(gen.Int())
	gens["EnableAutoScaling"] = gen.PtrOf(gen.Bool())
	gens["EnableEncryptionAtHost"] = gen.PtrOf(gen.Bool())
	gens["EnableFIPS"] = gen.PtrOf(gen.Bool())
	gens["EnableNodePublicIP"] = gen.PtrOf(gen.Bool())
	gens["EnableUltraSSD"] = gen.PtrOf(gen.Bool())
	gens["GpuInstanceProfile"] = gen.PtrOf(gen.AlphaString())
	gens["KubeletDiskType"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MaxCount"] = gen.PtrOf(gen.Int())
	gens["MaxPods"] = gen.PtrOf(gen.Int())
	gens["MinCount"] = gen.PtrOf(gen.Int())
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
	gens["NodeLabels"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["NodeTaints"] = gen.SliceOf(gen.AlphaString())
	gens["OrchestratorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["OsDiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsDiskType"] = gen.PtrOf(gen.AlphaString())
	gens["OsSKU"] = gen.PtrOf(gen.AlphaString())
	gens["OsType"] = gen.PtrOf(gen.AlphaString())
	gens["ProximityPlacementGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetEvictionPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetPriority"] = gen.PtrOf(gen.AlphaString())
	gens["SpotMaxPrice"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["VmSize"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedClustersAgentPoolsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClustersAgentPoolsSpec(gens map[string]gopter.Gen) {
	gens["KubeletConfig"] = gen.PtrOf(KubeletConfigGenerator())
	gens["LinuxOSConfig"] = gen.PtrOf(LinuxOSConfigGenerator())
	gens["UpgradeSettings"] = gen.PtrOf(AgentPoolUpgradeSettingsGenerator())
}

func Test_AgentPoolUpgradeSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolUpgradeSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolUpgradeSettings, AgentPoolUpgradeSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolUpgradeSettings runs a test to see if a specific instance of AgentPoolUpgradeSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolUpgradeSettings(subject AgentPoolUpgradeSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolUpgradeSettings
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AgentPoolUpgradeSettings instances for property testing - lazily instantiated by
//AgentPoolUpgradeSettingsGenerator()
var agentPoolUpgradeSettingsGenerator gopter.Gen

// AgentPoolUpgradeSettingsGenerator returns a generator of AgentPoolUpgradeSettings instances for property testing.
func AgentPoolUpgradeSettingsGenerator() gopter.Gen {
	if agentPoolUpgradeSettingsGenerator != nil {
		return agentPoolUpgradeSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings(generators)
	agentPoolUpgradeSettingsGenerator = gen.Struct(reflect.TypeOf(AgentPoolUpgradeSettings{}), generators)

	return agentPoolUpgradeSettingsGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings(gens map[string]gopter.Gen) {
	gens["MaxSurge"] = gen.PtrOf(gen.AlphaString())
}

func Test_AgentPoolUpgradeSettings_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolUpgradeSettings_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolUpgradeSettingsStatus, AgentPoolUpgradeSettingsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolUpgradeSettingsStatus runs a test to see if a specific instance of AgentPoolUpgradeSettings_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolUpgradeSettingsStatus(subject AgentPoolUpgradeSettings_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolUpgradeSettings_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AgentPoolUpgradeSettings_Status instances for property testing - lazily instantiated by
//AgentPoolUpgradeSettingsStatusGenerator()
var agentPoolUpgradeSettingsStatusGenerator gopter.Gen

// AgentPoolUpgradeSettingsStatusGenerator returns a generator of AgentPoolUpgradeSettings_Status instances for property testing.
func AgentPoolUpgradeSettingsStatusGenerator() gopter.Gen {
	if agentPoolUpgradeSettingsStatusGenerator != nil {
		return agentPoolUpgradeSettingsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettingsStatus(generators)
	agentPoolUpgradeSettingsStatusGenerator = gen.Struct(reflect.TypeOf(AgentPoolUpgradeSettings_Status{}), generators)

	return agentPoolUpgradeSettingsStatusGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettingsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettingsStatus(gens map[string]gopter.Gen) {
	gens["MaxSurge"] = gen.PtrOf(gen.AlphaString())
}

func Test_KubeletConfig_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KubeletConfig via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKubeletConfig, KubeletConfigGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKubeletConfig runs a test to see if a specific instance of KubeletConfig round trips to JSON and back losslessly
func RunJSONSerializationTestForKubeletConfig(subject KubeletConfig) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KubeletConfig
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of KubeletConfig instances for property testing - lazily instantiated by KubeletConfigGenerator()
var kubeletConfigGenerator gopter.Gen

// KubeletConfigGenerator returns a generator of KubeletConfig instances for property testing.
func KubeletConfigGenerator() gopter.Gen {
	if kubeletConfigGenerator != nil {
		return kubeletConfigGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKubeletConfig(generators)
	kubeletConfigGenerator = gen.Struct(reflect.TypeOf(KubeletConfig{}), generators)

	return kubeletConfigGenerator
}

// AddIndependentPropertyGeneratorsForKubeletConfig is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKubeletConfig(gens map[string]gopter.Gen) {
	gens["AllowedUnsafeSysctls"] = gen.SliceOf(gen.AlphaString())
	gens["ContainerLogMaxFiles"] = gen.PtrOf(gen.Int())
	gens["ContainerLogMaxSizeMB"] = gen.PtrOf(gen.Int())
	gens["CpuCfsQuota"] = gen.PtrOf(gen.Bool())
	gens["CpuCfsQuotaPeriod"] = gen.PtrOf(gen.AlphaString())
	gens["CpuManagerPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["FailSwapOn"] = gen.PtrOf(gen.Bool())
	gens["ImageGcHighThreshold"] = gen.PtrOf(gen.Int())
	gens["ImageGcLowThreshold"] = gen.PtrOf(gen.Int())
	gens["PodMaxPids"] = gen.PtrOf(gen.Int())
	gens["TopologyManagerPolicy"] = gen.PtrOf(gen.AlphaString())
}

func Test_KubeletConfig_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KubeletConfig_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKubeletConfigStatus, KubeletConfigStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKubeletConfigStatus runs a test to see if a specific instance of KubeletConfig_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForKubeletConfigStatus(subject KubeletConfig_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KubeletConfig_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of KubeletConfig_Status instances for property testing - lazily instantiated by
//KubeletConfigStatusGenerator()
var kubeletConfigStatusGenerator gopter.Gen

// KubeletConfigStatusGenerator returns a generator of KubeletConfig_Status instances for property testing.
func KubeletConfigStatusGenerator() gopter.Gen {
	if kubeletConfigStatusGenerator != nil {
		return kubeletConfigStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKubeletConfigStatus(generators)
	kubeletConfigStatusGenerator = gen.Struct(reflect.TypeOf(KubeletConfig_Status{}), generators)

	return kubeletConfigStatusGenerator
}

// AddIndependentPropertyGeneratorsForKubeletConfigStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKubeletConfigStatus(gens map[string]gopter.Gen) {
	gens["AllowedUnsafeSysctls"] = gen.SliceOf(gen.AlphaString())
	gens["ContainerLogMaxFiles"] = gen.PtrOf(gen.Int())
	gens["ContainerLogMaxSizeMB"] = gen.PtrOf(gen.Int())
	gens["CpuCfsQuota"] = gen.PtrOf(gen.Bool())
	gens["CpuCfsQuotaPeriod"] = gen.PtrOf(gen.AlphaString())
	gens["CpuManagerPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["FailSwapOn"] = gen.PtrOf(gen.Bool())
	gens["ImageGcHighThreshold"] = gen.PtrOf(gen.Int())
	gens["ImageGcLowThreshold"] = gen.PtrOf(gen.Int())
	gens["PodMaxPids"] = gen.PtrOf(gen.Int())
	gens["TopologyManagerPolicy"] = gen.PtrOf(gen.AlphaString())
}

func Test_LinuxOSConfig_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinuxOSConfig via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinuxOSConfig, LinuxOSConfigGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinuxOSConfig runs a test to see if a specific instance of LinuxOSConfig round trips to JSON and back losslessly
func RunJSONSerializationTestForLinuxOSConfig(subject LinuxOSConfig) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinuxOSConfig
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of LinuxOSConfig instances for property testing - lazily instantiated by LinuxOSConfigGenerator()
var linuxOSConfigGenerator gopter.Gen

// LinuxOSConfigGenerator returns a generator of LinuxOSConfig instances for property testing.
// We first initialize linuxOSConfigGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LinuxOSConfigGenerator() gopter.Gen {
	if linuxOSConfigGenerator != nil {
		return linuxOSConfigGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig(generators)
	linuxOSConfigGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig(generators)
	AddRelatedPropertyGeneratorsForLinuxOSConfig(generators)
	linuxOSConfigGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig{}), generators)

	return linuxOSConfigGenerator
}

// AddIndependentPropertyGeneratorsForLinuxOSConfig is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinuxOSConfig(gens map[string]gopter.Gen) {
	gens["SwapFileSizeMB"] = gen.PtrOf(gen.Int())
	gens["TransparentHugePageDefrag"] = gen.PtrOf(gen.AlphaString())
	gens["TransparentHugePageEnabled"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLinuxOSConfig is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLinuxOSConfig(gens map[string]gopter.Gen) {
	gens["Sysctls"] = gen.PtrOf(SysctlConfigGenerator())
}

func Test_LinuxOSConfig_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinuxOSConfig_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinuxOSConfigStatus, LinuxOSConfigStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinuxOSConfigStatus runs a test to see if a specific instance of LinuxOSConfig_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForLinuxOSConfigStatus(subject LinuxOSConfig_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinuxOSConfig_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of LinuxOSConfig_Status instances for property testing - lazily instantiated by
//LinuxOSConfigStatusGenerator()
var linuxOSConfigStatusGenerator gopter.Gen

// LinuxOSConfigStatusGenerator returns a generator of LinuxOSConfig_Status instances for property testing.
// We first initialize linuxOSConfigStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LinuxOSConfigStatusGenerator() gopter.Gen {
	if linuxOSConfigStatusGenerator != nil {
		return linuxOSConfigStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfigStatus(generators)
	linuxOSConfigStatusGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfigStatus(generators)
	AddRelatedPropertyGeneratorsForLinuxOSConfigStatus(generators)
	linuxOSConfigStatusGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_Status{}), generators)

	return linuxOSConfigStatusGenerator
}

// AddIndependentPropertyGeneratorsForLinuxOSConfigStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinuxOSConfigStatus(gens map[string]gopter.Gen) {
	gens["SwapFileSizeMB"] = gen.PtrOf(gen.Int())
	gens["TransparentHugePageDefrag"] = gen.PtrOf(gen.AlphaString())
	gens["TransparentHugePageEnabled"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLinuxOSConfigStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLinuxOSConfigStatus(gens map[string]gopter.Gen) {
	gens["Sysctls"] = gen.PtrOf(SysctlConfigStatusGenerator())
}

func Test_SysctlConfig_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SysctlConfig via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSysctlConfig, SysctlConfigGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSysctlConfig runs a test to see if a specific instance of SysctlConfig round trips to JSON and back losslessly
func RunJSONSerializationTestForSysctlConfig(subject SysctlConfig) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SysctlConfig
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SysctlConfig instances for property testing - lazily instantiated by SysctlConfigGenerator()
var sysctlConfigGenerator gopter.Gen

// SysctlConfigGenerator returns a generator of SysctlConfig instances for property testing.
func SysctlConfigGenerator() gopter.Gen {
	if sysctlConfigGenerator != nil {
		return sysctlConfigGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSysctlConfig(generators)
	sysctlConfigGenerator = gen.Struct(reflect.TypeOf(SysctlConfig{}), generators)

	return sysctlConfigGenerator
}

// AddIndependentPropertyGeneratorsForSysctlConfig is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSysctlConfig(gens map[string]gopter.Gen) {
	gens["FsAioMaxNr"] = gen.PtrOf(gen.Int())
	gens["FsFileMax"] = gen.PtrOf(gen.Int())
	gens["FsInotifyMaxUserWatches"] = gen.PtrOf(gen.Int())
	gens["FsNrOpen"] = gen.PtrOf(gen.Int())
	gens["KernelThreadsMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreNetdevMaxBacklog"] = gen.PtrOf(gen.Int())
	gens["NetCoreOptmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreSomaxconn"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemMax"] = gen.PtrOf(gen.Int())
	gens["NetIpv4IpLocalPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["NetIpv4NeighDefaultGcThresh1"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh2"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh3"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpFinTimeout"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveProbes"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveTime"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxSynBacklog"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxTwBuckets"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpTwReuse"] = gen.PtrOf(gen.Bool())
	gens["NetIpv4TcpkeepaliveIntvl"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackBuckets"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackMax"] = gen.PtrOf(gen.Int())
	gens["VmMaxMapCount"] = gen.PtrOf(gen.Int())
	gens["VmSwappiness"] = gen.PtrOf(gen.Int())
	gens["VmVfsCachePressure"] = gen.PtrOf(gen.Int())
}

func Test_SysctlConfig_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SysctlConfig_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSysctlConfigStatus, SysctlConfigStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSysctlConfigStatus runs a test to see if a specific instance of SysctlConfig_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSysctlConfigStatus(subject SysctlConfig_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SysctlConfig_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SysctlConfig_Status instances for property testing - lazily instantiated by SysctlConfigStatusGenerator()
var sysctlConfigStatusGenerator gopter.Gen

// SysctlConfigStatusGenerator returns a generator of SysctlConfig_Status instances for property testing.
func SysctlConfigStatusGenerator() gopter.Gen {
	if sysctlConfigStatusGenerator != nil {
		return sysctlConfigStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSysctlConfigStatus(generators)
	sysctlConfigStatusGenerator = gen.Struct(reflect.TypeOf(SysctlConfig_Status{}), generators)

	return sysctlConfigStatusGenerator
}

// AddIndependentPropertyGeneratorsForSysctlConfigStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSysctlConfigStatus(gens map[string]gopter.Gen) {
	gens["FsAioMaxNr"] = gen.PtrOf(gen.Int())
	gens["FsFileMax"] = gen.PtrOf(gen.Int())
	gens["FsInotifyMaxUserWatches"] = gen.PtrOf(gen.Int())
	gens["FsNrOpen"] = gen.PtrOf(gen.Int())
	gens["KernelThreadsMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreNetdevMaxBacklog"] = gen.PtrOf(gen.Int())
	gens["NetCoreOptmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreSomaxconn"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemMax"] = gen.PtrOf(gen.Int())
	gens["NetIpv4IpLocalPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["NetIpv4NeighDefaultGcThresh1"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh2"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh3"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpFinTimeout"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveProbes"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveTime"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxSynBacklog"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxTwBuckets"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpTwReuse"] = gen.PtrOf(gen.Bool())
	gens["NetIpv4TcpkeepaliveIntvl"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackBuckets"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackMax"] = gen.PtrOf(gen.Int())
	gens["VmMaxMapCount"] = gen.PtrOf(gen.Int())
	gens["VmSwappiness"] = gen.PtrOf(gen.Int())
	gens["VmVfsCachePressure"] = gen.PtrOf(gen.Int())
}
