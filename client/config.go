// Copyright 2023 CodeMaker AI Inc. All rights reserved.

package client

import "time"

type Config struct {
	ApiKey            string
	Endpoint          *string
	ConnectionTimeout *time.Duration
	RequestTimeout    *time.Duration
}
