// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

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

func Test_DatabaseAccountsMongodbDatabases_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabases_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesSpecARM, DatabaseAccountsMongodbDatabasesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesSpecARM runs a test to see if a specific instance of DatabaseAccountsMongodbDatabases_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesSpecARM(subject DatabaseAccountsMongodbDatabases_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabases_SpecARM
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

// Generator of DatabaseAccountsMongodbDatabases_SpecARM instances for property testing - lazily instantiated by
//DatabaseAccountsMongodbDatabasesSpecARMGenerator()
var databaseAccountsMongodbDatabasesSpecARMGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesSpecARMGenerator returns a generator of DatabaseAccountsMongodbDatabases_SpecARM instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesSpecARMGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesSpecARMGenerator != nil {
		return databaseAccountsMongodbDatabasesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpecARM(generators)
	databaseAccountsMongodbDatabasesSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabases_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpecARM(generators)
	databaseAccountsMongodbDatabasesSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabases_SpecARM{}), generators)

	return databaseAccountsMongodbDatabasesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBDatabaseCreateUpdatePropertiesARMGenerator())
}

func Test_MongoDBDatabaseCreateUpdatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseCreateUpdatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseCreateUpdatePropertiesARM, MongoDBDatabaseCreateUpdatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseCreateUpdatePropertiesARM runs a test to see if a specific instance of MongoDBDatabaseCreateUpdatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseCreateUpdatePropertiesARM(subject MongoDBDatabaseCreateUpdatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseCreateUpdatePropertiesARM
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

// Generator of MongoDBDatabaseCreateUpdatePropertiesARM instances for property testing - lazily instantiated by
//MongoDBDatabaseCreateUpdatePropertiesARMGenerator()
var mongoDBDatabaseCreateUpdatePropertiesARMGenerator gopter.Gen

// MongoDBDatabaseCreateUpdatePropertiesARMGenerator returns a generator of MongoDBDatabaseCreateUpdatePropertiesARM instances for property testing.
func MongoDBDatabaseCreateUpdatePropertiesARMGenerator() gopter.Gen {
	if mongoDBDatabaseCreateUpdatePropertiesARMGenerator != nil {
		return mongoDBDatabaseCreateUpdatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdatePropertiesARM(generators)
	mongoDBDatabaseCreateUpdatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseCreateUpdatePropertiesARM{}), generators)

	return mongoDBDatabaseCreateUpdatePropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdatePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdatePropertiesARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsARMGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseResourceARMGenerator())
}

func Test_MongoDBDatabaseResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseResourceARM, MongoDBDatabaseResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseResourceARM runs a test to see if a specific instance of MongoDBDatabaseResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseResourceARM(subject MongoDBDatabaseResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseResourceARM
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

// Generator of MongoDBDatabaseResourceARM instances for property testing - lazily instantiated by
//MongoDBDatabaseResourceARMGenerator()
var mongoDBDatabaseResourceARMGenerator gopter.Gen

// MongoDBDatabaseResourceARMGenerator returns a generator of MongoDBDatabaseResourceARM instances for property testing.
func MongoDBDatabaseResourceARMGenerator() gopter.Gen {
	if mongoDBDatabaseResourceARMGenerator != nil {
		return mongoDBDatabaseResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseResourceARM(generators)
	mongoDBDatabaseResourceARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseResourceARM{}), generators)

	return mongoDBDatabaseResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseResourceARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
