// Copyright 2023 CodeMaker AI Inc. All rights reserved.

package client

const (
	ModeDocument       = "DOCUMENT"
	ModeUnitTest       = "UNIT_TEST"
	ModeMigrateSyntax  = "MIGRATE_SYNTAX"
	ModeRefactorNaming = "REFACTOR_NAMING"

	StatusInProgress = "IN_PROGRESS"
	StatusCompleted  = "COMPLETED"
	StatusFailed     = "FAILED"
	StatusTimedOut   = "TIMED_OUT"
)

const (
	LanguageJava = "JAVA"
)

type CreateProcessRequest struct {
	Process Process `json:"process"`
}

type CreateProcessResponse struct {
	Id string `json:"id"`
}

type GetProcessStatusRequest struct {
	Id string `json:"id"`
}

type GetProcessStatusResponse struct {
	Status string `json:"status"`
}

type GetProcessOutputRequest struct {
	Id string `json:"id"`
}

type GetProcessOutputResponse struct {
	Output Output `json:"output"`
}

type Process struct {
	Mode          string         `json:"mode"`
	Language      string         `json:"language"`
	Input         Input          `json:"input"`
	OutputOptions *OutputOptions `json:"outputOptions"`
}

type Input struct {
	Source string `json:"source"`
}

type OutputOptions struct {
	LanguageVersion *string `json:"languageVersion"`
	Framework       *string `json:"framework"`
}

type Output struct {
	Source string `json:"source"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
