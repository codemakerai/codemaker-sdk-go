// Copyright 2023 CodeMaker AI Inc. All rights reserved.

package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func client(endpoint string) Client {
	apiKey := "ABCDE-GHIJK-LMNOP-QRSTU-1"

	return NewClient(Config{
		ApiKey:   apiKey,
		Endpoint: &endpoint,
	})
}

func TestClient(t *testing.T) {

	t.Run("Client is created successfully", func(t *testing.T) {
		got := client("https://127.0.0.1:8080")

		if got == nil {
			t.Fatalf("Client was expect not to be nil")
		}
	})

	t.Run("Empty CreateProcess request is successful", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintln(w, "{\"id\": \"id\"}")
		}))
		defer ts.Close()

		client := client(ts.URL)

		got, err := client.CreateProcess(nil)
		if err != nil {
			t.Fatalf("Request failed with an error %v", err)
		}

		if got == nil {
			t.Fatalf("Response was expect not to be nil")
		}
		if got.Id != "id" {
			t.Fatalf("Response id was incorrect got %s", got.Id)
		}
	})

	t.Run("CreateProcess request is successful", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintln(w, `{"id": "id"}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		got, err := client.CreateProcess(&CreateProcessRequest{
			Process: Process{
				Mode:     ModeDocument,
				Language: LanguageJava,
				Input: Input{
					Source: "",
				},
			},
		})
		if err != nil {
			t.Fatalf("Request failed with an error %v", err)
		}

		if got == nil {
			t.Fatalf("Response was expect not to be nil")
		}
		if got.Id != "id" {
			t.Fatalf("Response id was incorrect got %s", got.Id)
		}
	})

	t.Run("CreateProcess request is successful for all languages", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintln(w, `{"id": "id"}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		languages := []string{LanguageJavaScript, LanguageJava, LanguageKotlin}
		for _, language := range languages {
			got, err := client.CreateProcess(&CreateProcessRequest{
				Process: Process{
					Mode:     ModeDocument,
					Language: LanguageJavaScript,
					Input: Input{
						Source: "",
					},
				},
			})
			if err != nil {
				t.Fatalf("Request failed for language %s with an error %v", language, err)
			}

			if got == nil {
				t.Fatalf("Response was expect not to be nil for language %s", language)
			}
			if got.Id != "id" {
				t.Fatalf("Response id was incorrect got %s for language %s", got.Id, language)
			}
		}
	})

	t.Run("CreateProcess request for mode DOCUMENT is successful", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintln(w, `{"id": "id"}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		got, err := client.CreateProcess(&CreateProcessRequest{
			Process: Process{
				Mode:     ModeDocument,
				Language: LanguageJava,
				Input: Input{
					Source: "",
				},
			},
		})
		if err != nil {
			t.Fatalf("Request failed with an error %v", err)
		}

		if got == nil {
			t.Fatalf("Response was expect not to be nil")
		}
		if got.Id != "id" {
			t.Fatalf("Response id was incorrect got %s", got.Id)
		}
	})

	t.Run("CreateProcess request for mode UNIT_TEST is successful", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintln(w, `{"id": "id"}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		got, err := client.CreateProcess(&CreateProcessRequest{
			Process: Process{
				Mode:     ModeUnitTest,
				Language: LanguageJava,
				Input: Input{
					Source: "",
				},
			},
		})
		if err != nil {
			t.Fatalf("Request failed with an error %v", err)
		}

		if got == nil {
			t.Fatalf("Response was expect not to be nil")
		}
		if got.Id != "id" {
			t.Fatalf("Response id was incorrect got %s", got.Id)
		}
	})

	t.Run("CreateProcess request for mode MIGRATE_SYNTAX is successful", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintln(w, `{"id": "id"}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		got, err := client.CreateProcess(&CreateProcessRequest{
			Process: Process{
				Mode:     ModeMigrateSyntax,
				Language: LanguageJava,
				Input: Input{
					Source: "",
				},
			},
		})
		if err != nil {
			t.Fatalf("Request failed with an error %v", err)
		}

		if got == nil {
			t.Fatalf("Response was expect not to be nil")
		}
		if got.Id != "id" {
			t.Fatalf("Response id was incorrect got %s", got.Id)
		}
	})

	t.Run("CreateProcess request for mode REFACTOR_NAMING is successful", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintln(w, `{"id": "id"}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		got, err := client.CreateProcess(&CreateProcessRequest{
			Process: Process{
				Mode:     ModeRefactorNaming,
				Language: LanguageJava,
				Input: Input{
					Source: "",
				},
			},
		})
		if err != nil {
			t.Fatalf("Request failed with an error %v", err)
		}

		if got == nil {
			t.Fatalf("Response was expect not to be nil")
		}
		if got.Id != "id" {
			t.Fatalf("Response id was incorrect got %s", got.Id)
		}
	})

	t.Run("CreateProcess results in bad request", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(headerRequestId, "123456789")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, `{"code":"BAD_REQUEST","message": "Bad request."}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		_, err := client.CreateProcess(&CreateProcessRequest{
			Process: Process{
				Mode:     ModeDocument,
				Language: LanguageJava,
				Input: Input{
					Source: "",
				},
			},
		})

		if err == nil {
			t.Fatalf("Error was expected")
		}
		if err.Error() != "(123456789) request failed 400 BAD_REQUEST" {
			t.Fatalf("Incorrect error message %v", err)
		}
	})

	t.Run("CreateProcess results in unauthorized request", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(headerRequestId, "123456789")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, `{"code":"UNAUTHORIZED","message":"Unauthorized."}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		_, err := client.CreateProcess(&CreateProcessRequest{
			Process: Process{
				Mode:     ModeDocument,
				Language: LanguageJava,
				Input: Input{
					Source: "",
				},
			},
		})

		if err == nil {
			t.Fatalf("Error was expected")
		}
		if err.Error() != "(123456789) request failed 401 UNAUTHORIZED" {
			t.Fatalf("Incorrect error message %v", err)
		}
	})

	t.Run("CreateProcess results in internal server error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(headerRequestId, "123456789")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"code":"SERVICE_ERROR","message":"Unknown service error"}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		_, err := client.CreateProcess(&CreateProcessRequest{
			Process: Process{
				Mode:     ModeDocument,
				Language: LanguageJava,
				Input: Input{
					Source: "",
				},
			},
		})

		if err == nil {
			t.Fatalf("Error was expected")
		}
		if err.Error() != "(123456789) request failed 500 SERVICE_ERROR" {
			t.Fatalf("Incorrect error message %v", err)
		}
	})

	t.Run("GetProcessStatus request is successful", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(headerRequestId, "123456789")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"status": "IN_PROGRESS"}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		got, err := client.GetProcessStatus(&GetProcessStatusRequest{
			Id: "id",
		})
		if err != nil {
			t.Fatalf("Request failed with an error %v", err)
		}

		if got == nil {
			t.Fatalf("Response was expect not to be nil")
		}
		if got.Status != StatusInProgress {
			t.Fatalf("Response stats was incorrect got %s", got.Status)
		}
	})

	t.Run("GetProcessStatus results in internal server error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(headerRequestId, "123456789")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"code":"SERVICE_ERROR","message":"Unknown service error"}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		_, err := client.GetProcessStatus(&GetProcessStatusRequest{
			Id: "id",
		})

		if err == nil {
			t.Fatalf("Error was expected")
		}
		if err.Error() != "(123456789) request failed 500 SERVICE_ERROR" {
			t.Fatalf("Incorrect error message %v", err)
		}
	})

	t.Run("GetProcessStatus request is successful", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(headerRequestId, "123456789")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"output": {"source": "source"}}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		got, err := client.GetProcessOutput(&GetProcessOutputRequest{
			Id: "id",
		})
		if err != nil {
			t.Fatalf("Request failed with an error %v", err)
		}

		if got == nil {
			t.Fatalf("Response was expect not to be nil")
		}
		if got.Output.Source != "source" {
			t.Fatalf("Response stats was incorrect got %s", got.Output.Source)
		}
	})

	t.Run("GetProcessStatus results in internal server error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(headerRequestId, "123456789")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"code":"SERVICE_ERROR","message":"Unknown service error"}`)
		}))
		defer ts.Close()

		client := client(ts.URL)

		_, err := client.GetProcessOutput(&GetProcessOutputRequest{
			Id: "id",
		})

		if err == nil {
			t.Fatalf("Error was expected")
		}
		if err.Error() != "(123456789) request failed 500 SERVICE_ERROR" {
			t.Fatalf("Incorrect error message %v", err)
		}
	})
}
