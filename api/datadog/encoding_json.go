// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

//go:build !goccy_gojson

package datadog

import (
	"encoding/json"
	"io"

	"github.com/bytedance/sonic"
)

func Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return sonic.Unmarshal(data, v)
}

func NewEncoder(w io.Writer) *json.Encoder {
	enc := sonic.ConfigDefault.NewEncoder(w)
	return enc.(*json.Encoder)
}

func NewDecoder(r io.Reader) *json.Decoder {
	dec := sonic.ConfigDefault.NewDecoder(r)
	return dec.(*json.Decoder)
}
