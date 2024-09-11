// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

//go:build !goccy_gojson

package datadog

import (
	"io"

	"github.com/bytedance/sonic"
)

func Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return sonic.Unmarshal(data, v)
}

func NewEncoder(w io.Writer) sonic.Encoder {
	return sonic.ConfigDefault.NewEncoder(w)
}

func NewDecoder(r io.Reader) sonic.Decoder {
	return sonic.ConfigDefault.NewDecoder(r)
}
