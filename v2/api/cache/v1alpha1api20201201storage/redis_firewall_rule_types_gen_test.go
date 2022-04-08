// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201storage

import (
	"encoding/json"
	v20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201storage"
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

func Test_RedisFirewallRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisFirewallRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForRedisFirewallRule, RedisFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRedisFirewallRule tests if a specific instance of RedisFirewallRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRedisFirewallRule(subject RedisFirewallRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20201201s.RedisFirewallRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RedisFirewallRule
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

func Test_RedisFirewallRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisFirewallRule to RedisFirewallRule via AssignPropertiesToRedisFirewallRule & AssignPropertiesFromRedisFirewallRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisFirewallRule, RedisFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisFirewallRule tests if a specific instance of RedisFirewallRule can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisFirewallRule(subject RedisFirewallRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.RedisFirewallRule
	err := copied.AssignPropertiesToRedisFirewallRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisFirewallRule
	err = actual.AssignPropertiesFromRedisFirewallRule(&other)
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

func Test_RedisFirewallRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRule, RedisFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRule runs a test to see if a specific instance of RedisFirewallRule round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRule(subject RedisFirewallRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRule
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

// Generator of RedisFirewallRule instances for property testing - lazily instantiated by RedisFirewallRuleGenerator()
var redisFirewallRuleGenerator gopter.Gen

// RedisFirewallRuleGenerator returns a generator of RedisFirewallRule instances for property testing.
func RedisFirewallRuleGenerator() gopter.Gen {
	if redisFirewallRuleGenerator != nil {
		return redisFirewallRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisFirewallRule(generators)
	redisFirewallRuleGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule{}), generators)

	return redisFirewallRuleGenerator
}

// AddRelatedPropertyGeneratorsForRedisFirewallRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisFirewallRule(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisFirewallRulesSpecGenerator()
	gens["Status"] = RedisFirewallRuleStatusGenerator()
}

func Test_RedisFirewallRule_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisFirewallRule_Status to RedisFirewallRule_Status via AssignPropertiesToRedisFirewallRuleStatus & AssignPropertiesFromRedisFirewallRuleStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisFirewallRuleStatus, RedisFirewallRuleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisFirewallRuleStatus tests if a specific instance of RedisFirewallRule_Status can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisFirewallRuleStatus(subject RedisFirewallRule_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.RedisFirewallRule_Status
	err := copied.AssignPropertiesToRedisFirewallRuleStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisFirewallRule_Status
	err = actual.AssignPropertiesFromRedisFirewallRuleStatus(&other)
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

func Test_RedisFirewallRule_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRule_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRuleStatus, RedisFirewallRuleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRuleStatus runs a test to see if a specific instance of RedisFirewallRule_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRuleStatus(subject RedisFirewallRule_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRule_Status
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

// Generator of RedisFirewallRule_Status instances for property testing - lazily instantiated by
//RedisFirewallRuleStatusGenerator()
var redisFirewallRuleStatusGenerator gopter.Gen

// RedisFirewallRuleStatusGenerator returns a generator of RedisFirewallRule_Status instances for property testing.
func RedisFirewallRuleStatusGenerator() gopter.Gen {
	if redisFirewallRuleStatusGenerator != nil {
		return redisFirewallRuleStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRuleStatus(generators)
	redisFirewallRuleStatusGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule_Status{}), generators)

	return redisFirewallRuleStatusGenerator
}

// AddIndependentPropertyGeneratorsForRedisFirewallRuleStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisFirewallRuleStatus(gens map[string]gopter.Gen) {
	gens["EndIP"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StartIP"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisFirewallRules_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisFirewallRules_Spec to RedisFirewallRules_Spec via AssignPropertiesToRedisFirewallRulesSpec & AssignPropertiesFromRedisFirewallRulesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisFirewallRulesSpec, RedisFirewallRulesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisFirewallRulesSpec tests if a specific instance of RedisFirewallRules_Spec can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisFirewallRulesSpec(subject RedisFirewallRules_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.RedisFirewallRules_Spec
	err := copied.AssignPropertiesToRedisFirewallRulesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisFirewallRules_Spec
	err = actual.AssignPropertiesFromRedisFirewallRulesSpec(&other)
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

func Test_RedisFirewallRules_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRules_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRulesSpec, RedisFirewallRulesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRulesSpec runs a test to see if a specific instance of RedisFirewallRules_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRulesSpec(subject RedisFirewallRules_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRules_Spec
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

// Generator of RedisFirewallRules_Spec instances for property testing - lazily instantiated by
//RedisFirewallRulesSpecGenerator()
var redisFirewallRulesSpecGenerator gopter.Gen

// RedisFirewallRulesSpecGenerator returns a generator of RedisFirewallRules_Spec instances for property testing.
func RedisFirewallRulesSpecGenerator() gopter.Gen {
	if redisFirewallRulesSpecGenerator != nil {
		return redisFirewallRulesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRulesSpec(generators)
	redisFirewallRulesSpecGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRules_Spec{}), generators)

	return redisFirewallRulesSpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisFirewallRulesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisFirewallRulesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EndIP"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["StartIP"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
