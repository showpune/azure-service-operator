// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
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

func Test_SqlDatabaseContainerThroughputSetting_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerThroughputSetting to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseContainerThroughputSetting, SqlDatabaseContainerThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseContainerThroughputSetting tests if a specific instance of SqlDatabaseContainerThroughputSetting round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseContainerThroughputSetting(subject SqlDatabaseContainerThroughputSetting) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1beta20210515storage.SqlDatabaseContainerThroughputSetting
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseContainerThroughputSetting
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

func Test_SqlDatabaseContainerThroughputSetting_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerThroughputSetting to SqlDatabaseContainerThroughputSetting via AssignPropertiesToSqlDatabaseContainerThroughputSetting & AssignPropertiesFromSqlDatabaseContainerThroughputSetting returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting, SqlDatabaseContainerThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting tests if a specific instance of SqlDatabaseContainerThroughputSetting can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting(subject SqlDatabaseContainerThroughputSetting) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.SqlDatabaseContainerThroughputSetting
	err := copied.AssignPropertiesToSqlDatabaseContainerThroughputSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerThroughputSetting
	err = actual.AssignPropertiesFromSqlDatabaseContainerThroughputSetting(&other)
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

func Test_SqlDatabaseContainerThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting, SqlDatabaseContainerThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting runs a test to see if a specific instance of SqlDatabaseContainerThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting(subject SqlDatabaseContainerThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerThroughputSetting
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

// Generator of SqlDatabaseContainerThroughputSetting instances for property testing - lazily instantiated by
//SqlDatabaseContainerThroughputSettingGenerator()
var sqlDatabaseContainerThroughputSettingGenerator gopter.Gen

// SqlDatabaseContainerThroughputSettingGenerator returns a generator of SqlDatabaseContainerThroughputSetting instances for property testing.
func SqlDatabaseContainerThroughputSettingGenerator() gopter.Gen {
	if sqlDatabaseContainerThroughputSettingGenerator != nil {
		return sqlDatabaseContainerThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting(generators)
	sqlDatabaseContainerThroughputSettingGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerThroughputSetting{}), generators)

	return sqlDatabaseContainerThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator()
	gens["Status"] = ThroughputSettingsGetResultsStatusGenerator()
}

func Test_DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec to DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec via AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec & AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec, DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec tests if a specific instance of DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(subject DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
	err := copied.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(&other)
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

func Test_DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec, DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(subject DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
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

// Generator of DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec instances for property testing - lazily
//instantiated by DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator()
var databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator != nil {
		return databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(generators)
	databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(generators)
	databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}), generators)

	return databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
}
