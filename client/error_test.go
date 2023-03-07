// Copyright 2023 CodeMaker AI Inc. All rights reserved.

package client

import (
	"fmt"
	"testing"
)

func TestClientError(t *testing.T) {
	t.Run("NewClientError is created", func(t *testing.T) {
		got := NewClientError("Error")

		if got == nil {
			t.Fatalf("Error was expected not to be nil")
		}
		if got.Error() != "Error" {
			t.Fatalf("Error message is incorrt got %s", got.Error())
		}
	})

	t.Run("NewClientErrorWithCause is created", func(t *testing.T) {
		cause := fmt.Errorf("cause")
		got := NewClientErrorWithCause("Error", cause)

		if got == nil {
			t.Fatalf("Error was expected not to be nil")
		}
		if got.Error() != "Error" {
			t.Fatalf("Error message is incorrt got %s", got.Error())
		}
		if got.Unwrap() != cause {
			t.Fatalf("Error cause is incorrt got %v", got.Unwrap())
		}
	})
}
