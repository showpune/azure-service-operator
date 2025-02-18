// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/eventhub/v1alpha1api20211101storage"
	"github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101storage"
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

func Test_NamespacesEventhubsAuthorizationRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsAuthorizationRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesEventhubsAuthorizationRule, NamespacesEventhubsAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesEventhubsAuthorizationRule tests if a specific instance of NamespacesEventhubsAuthorizationRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesEventhubsAuthorizationRule(subject NamespacesEventhubsAuthorizationRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1beta20211101storage.NamespacesEventhubsAuthorizationRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesEventhubsAuthorizationRule
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesEventhubsAuthorizationRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsAuthorizationRule to NamespacesEventhubsAuthorizationRule via AssignPropertiesToNamespacesEventhubsAuthorizationRule & AssignPropertiesFromNamespacesEventhubsAuthorizationRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsAuthorizationRule, NamespacesEventhubsAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsAuthorizationRule tests if a specific instance of NamespacesEventhubsAuthorizationRule can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsAuthorizationRule(subject NamespacesEventhubsAuthorizationRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRule
	err := copied.AssignPropertiesToNamespacesEventhubsAuthorizationRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubsAuthorizationRule
	err = actual.AssignPropertiesFromNamespacesEventhubsAuthorizationRule(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesEventhubsAuthorizationRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsAuthorizationRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule, NamespacesEventhubsAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule runs a test to see if a specific instance of NamespacesEventhubsAuthorizationRule round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule(subject NamespacesEventhubsAuthorizationRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsAuthorizationRule
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

// Generator of NamespacesEventhubsAuthorizationRule instances for property testing - lazily instantiated by
//NamespacesEventhubsAuthorizationRuleGenerator()
var namespacesEventhubsAuthorizationRuleGenerator gopter.Gen

// NamespacesEventhubsAuthorizationRuleGenerator returns a generator of NamespacesEventhubsAuthorizationRule instances for property testing.
func NamespacesEventhubsAuthorizationRuleGenerator() gopter.Gen {
	if namespacesEventhubsAuthorizationRuleGenerator != nil {
		return namespacesEventhubsAuthorizationRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule(generators)
	namespacesEventhubsAuthorizationRuleGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsAuthorizationRule{}), generators)

	return namespacesEventhubsAuthorizationRuleGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesEventhubsAuthorizationRulesSpecGenerator()
	gens["Status"] = AuthorizationRuleStatusGenerator()
}

func Test_NamespacesEventhubsAuthorizationRules_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsAuthorizationRules_Spec to NamespacesEventhubsAuthorizationRules_Spec via AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec & AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsAuthorizationRulesSpec, NamespacesEventhubsAuthorizationRulesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsAuthorizationRulesSpec tests if a specific instance of NamespacesEventhubsAuthorizationRules_Spec can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsAuthorizationRulesSpec(subject NamespacesEventhubsAuthorizationRules_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRules_Spec
	err := copied.AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubsAuthorizationRules_Spec
	err = actual.AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesEventhubsAuthorizationRules_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsAuthorizationRules_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsAuthorizationRulesSpec, NamespacesEventhubsAuthorizationRulesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsAuthorizationRulesSpec runs a test to see if a specific instance of NamespacesEventhubsAuthorizationRules_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsAuthorizationRulesSpec(subject NamespacesEventhubsAuthorizationRules_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsAuthorizationRules_Spec
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

// Generator of NamespacesEventhubsAuthorizationRules_Spec instances for property testing - lazily instantiated by
//NamespacesEventhubsAuthorizationRulesSpecGenerator()
var namespacesEventhubsAuthorizationRulesSpecGenerator gopter.Gen

// NamespacesEventhubsAuthorizationRulesSpecGenerator returns a generator of NamespacesEventhubsAuthorizationRules_Spec instances for property testing.
func NamespacesEventhubsAuthorizationRulesSpecGenerator() gopter.Gen {
	if namespacesEventhubsAuthorizationRulesSpecGenerator != nil {
		return namespacesEventhubsAuthorizationRulesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsAuthorizationRulesSpec(generators)
	namespacesEventhubsAuthorizationRulesSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsAuthorizationRules_Spec{}), generators)

	return namespacesEventhubsAuthorizationRulesSpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsAuthorizationRulesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsAuthorizationRulesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(AuthorizationRulePropertiesRightsListen, AuthorizationRulePropertiesRightsManage, AuthorizationRulePropertiesRightsSend))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
