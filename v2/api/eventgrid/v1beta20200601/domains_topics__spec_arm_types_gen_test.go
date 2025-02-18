// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

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

func Test_DomainsTopics_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopics_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopicsSpecARM, DomainsTopicsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopicsSpecARM runs a test to see if a specific instance of DomainsTopics_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopicsSpecARM(subject DomainsTopics_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopics_SpecARM
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

// Generator of DomainsTopics_SpecARM instances for property testing - lazily instantiated by
//DomainsTopicsSpecARMGenerator()
var domainsTopicsSpecARMGenerator gopter.Gen

// DomainsTopicsSpecARMGenerator returns a generator of DomainsTopics_SpecARM instances for property testing.
func DomainsTopicsSpecARMGenerator() gopter.Gen {
	if domainsTopicsSpecARMGenerator != nil {
		return domainsTopicsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopicsSpecARM(generators)
	domainsTopicsSpecARMGenerator = gen.Struct(reflect.TypeOf(DomainsTopics_SpecARM{}), generators)

	return domainsTopicsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainsTopicsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainsTopicsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
