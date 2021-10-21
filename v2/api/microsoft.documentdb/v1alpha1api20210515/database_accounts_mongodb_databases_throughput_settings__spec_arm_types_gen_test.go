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

func Test_DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesThroughputSettingsSpecARM, DatabaseAccountsMongodbDatabasesThroughputSettingsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesThroughputSettingsSpecARM runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesThroughputSettingsSpecARM(subject DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM
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

// Generator of DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM instances for property testing - lazily
//instantiated by DatabaseAccountsMongodbDatabasesThroughputSettingsSpecARMGenerator()
var databaseAccountsMongodbDatabasesThroughputSettingsSpecARMGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesThroughputSettingsSpecARMGenerator returns a generator of DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesThroughputSettingsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesThroughputSettingsSpecARMGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesThroughputSettingsSpecARMGenerator != nil {
		return databaseAccountsMongodbDatabasesThroughputSettingsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpecARM(generators)
	databaseAccountsMongodbDatabasesThroughputSettingsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpecARM(generators)
	databaseAccountsMongodbDatabasesThroughputSettingsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM{}), generators)

	return databaseAccountsMongodbDatabasesThroughputSettingsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSettingsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = ThroughputSettingsUpdatePropertiesARMGenerator()
}
