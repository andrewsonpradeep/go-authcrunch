// Copyright 2022 Paul Greenberg greenpau@outlook.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validate

import (
	"github.com/andrewsonpradeep/go-authcrunch/internal/tests"
	"github.com/andrewsonpradeep/go-authcrunch/pkg/errors"
	"testing"
)

func TestLoginHint(t *testing.T) {
	var testcases = []struct {
		name       string
		loginHint  string
		validators []string
		shouldErr  bool
		err        error
	}{
		{
			name:       "doesn't return an error if the provided login_hint is a valid email",
			loginHint:  "foo@bar.com",
			validators: []string{"email"},
			shouldErr:  false,
			err:        nil,
		},
		{
			name:       "returns an error if the provided login_hint has an invalid domain",
			loginHint:  "foo@",
			validators: []string{"email"},
			shouldErr:  true,
			err:        errors.ErrInvalidLoginHint,
		},
		{
			name:       "returns an error if the provided login_hint has an invalid domain",
			loginHint:  "foo@().com",
			validators: []string{"email"},
			shouldErr:  true,
			err:        errors.ErrInvalidLoginHint,
		},
		{
			name:       "doesn't return an error if the provided login_hint is a valid phone number",
			loginHint:  "+1-555-55 55",
			validators: []string{"phone"},
			shouldErr:  false,
			err:        nil,
		},
		{
			name:       "returns an error if the provided login_hint is not a valid phone number",
			loginHint:  ".5dsa-dasdas55",
			validators: []string{"phone"},
			shouldErr:  true,
			err:        errors.ErrInvalidLoginHint,
		},
		{
			name:       "doesn't return an error if the provided login_hint is a valid alphanumeric string",
			loginHint:  "foobar",
			validators: []string{"alphanumeric"},
			shouldErr:  false,
			err:        nil,
		},
		{
			name:       "returns an error if the provided login_hint is not a valid alphanumeric string",
			loginHint:  "!^$^#&$&abc",
			validators: []string{"alphanumeric"},
			shouldErr:  true,
			err:        errors.ErrInvalidLoginHint,
		},
		{
			name:       "doesn't return an error if the provided login_hint matches one of the validators",
			loginHint:  "+1-555-55 55",
			validators: []string{"email", "phone"},
			shouldErr:  false,
			err:        nil,
		},
		{
			name:       "returns an error if the provided login_hint doesn't match any validator",
			loginHint:  "ma!e^",
			validators: []string{"email", "phone", "alphanumeric"},
			shouldErr:  true,
			err:        errors.ErrInvalidLoginHint,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := LoginHint(tc.loginHint, tc.validators)

			if tests.EvalErrWithLog(t, err, nil, tc.shouldErr, tc.err, []string{}) {
				return
			}
		})
	}
}
