// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
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

func Test_MongodbDatabaseThroughputSetting_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from MongodbDatabaseThroughputSetting to hub returns original",
		prop.ForAll(RunResourceConversionTestForMongodbDatabaseThroughputSetting, MongodbDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForMongodbDatabaseThroughputSetting tests if a specific instance of MongodbDatabaseThroughputSetting round trips to the hub storage version and back losslessly
func RunResourceConversionTestForMongodbDatabaseThroughputSetting(subject MongodbDatabaseThroughputSetting) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1beta20210515storage.MongodbDatabaseThroughputSetting
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual MongodbDatabaseThroughputSetting
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

func Test_MongodbDatabaseThroughputSetting_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from MongodbDatabaseThroughputSetting to MongodbDatabaseThroughputSetting via AssignPropertiesToMongodbDatabaseThroughputSetting & AssignPropertiesFromMongodbDatabaseThroughputSetting returns original",
		prop.ForAll(RunPropertyAssignmentTestForMongodbDatabaseThroughputSetting, MongodbDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForMongodbDatabaseThroughputSetting tests if a specific instance of MongodbDatabaseThroughputSetting can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForMongodbDatabaseThroughputSetting(subject MongodbDatabaseThroughputSetting) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20210515storage.MongodbDatabaseThroughputSetting
	err := copied.AssignPropertiesToMongodbDatabaseThroughputSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual MongodbDatabaseThroughputSetting
	err = actual.AssignPropertiesFromMongodbDatabaseThroughputSetting(&other)
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

func Test_MongodbDatabaseThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabaseThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabaseThroughputSetting, MongodbDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabaseThroughputSetting runs a test to see if a specific instance of MongodbDatabaseThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabaseThroughputSetting(subject MongodbDatabaseThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabaseThroughputSetting
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

// Generator of MongodbDatabaseThroughputSetting instances for property testing - lazily instantiated by
//MongodbDatabaseThroughputSettingGenerator()
var mongodbDatabaseThroughputSettingGenerator gopter.Gen

// MongodbDatabaseThroughputSettingGenerator returns a generator of MongodbDatabaseThroughputSetting instances for property testing.
func MongodbDatabaseThroughputSettingGenerator() gopter.Gen {
	if mongodbDatabaseThroughputSettingGenerator != nil {
		return mongodbDatabaseThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongodbDatabaseThroughputSetting(generators)
	mongodbDatabaseThroughputSettingGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseThroughputSetting{}), generators)

	return mongodbDatabaseThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator()
	gens["Status"] = ThroughputSettingsGetResultsStatusGenerator()
}

func Test_DatabaseAccountsMongodbDatabasesThroughputSettings_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsMongodbDatabasesThroughputSettings_Spec to DatabaseAccountsMongodbDatabasesThroughputSettings_Spec via AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec & AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec, DatabaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec tests if a specific instance of DatabaseAccountsMongodbDatabasesThroughputSettings_Spec can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(subject DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20210515storage.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
	err := copied.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
	err = actual.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(&other)
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

func Test_DatabaseAccountsMongodbDatabasesThroughputSettings_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesThroughputSettings_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec, DatabaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesThroughputSettings_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(subject DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
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

// Generator of DatabaseAccountsMongodbDatabasesThroughputSettings_Spec instances for property testing - lazily
//instantiated by DatabaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator()
var databaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator returns a generator of DatabaseAccountsMongodbDatabasesThroughputSettings_Spec instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator != nil {
		return databaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(generators)
	databaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(generators)
	databaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{}), generators)

	return databaseAccountsMongodbDatabasesThroughputSettingsSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
}
