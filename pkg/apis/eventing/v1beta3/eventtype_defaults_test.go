/*
Copyright 2023 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta3

import (
	"context"
	"testing"

	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/google/go-cmp/cmp"

	"knative.dev/pkg/apis"
)

func TestEventTypeDefaults(t *testing.T) {
	testSource := apis.HTTP("test-source")
	testSchema := apis.HTTP("test-schema")
	testCases := map[string]struct {
		initial  EventType
		expected EventType
	}{
		"nil spec": {
			initial: EventType{},
			expected: EventType{
				Spec: EventTypeSpec{},
			},
		},
		"broker empty": {
			initial: EventType{
				Spec: EventTypeSpec{
					Type:   "test-type",
					Source: testSource,
					Reference: &duckv1.KReference{
						APIVersion: "eventing.knative.dev/v1",
						Kind:       "Broker",
						Name:       "default",
					},
					Schema: testSchema,
				},
			},
			expected: EventType{
				Spec: EventTypeSpec{
					Type:   "test-type",
					Source: testSource,
					Reference: &duckv1.KReference{
						APIVersion: "eventing.knative.dev/v1",
						Kind:       "Broker",
						Name:       "default",
					},
					Schema: testSchema,
				},
			},
		},
		"broker not set": {
			initial: EventType{
				Spec: EventTypeSpec{
					Type:   "test-type",
					Source: testSource,
					Schema: testSchema,
				},
			},
			expected: EventType{
				Spec: EventTypeSpec{
					Type:   "test-type",
					Source: testSource,
					Schema: testSchema,
				},
			},
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			tc.initial.SetDefaults(context.TODO())
			if diff := cmp.Diff(tc.expected, tc.initial); diff != "" {
				t.Fatal("Unexpected defaults (-want, +got):", diff)
			}
		})
	}
}
