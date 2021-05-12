/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SyntheticsVariableParser Details of the parser to use for the global variable.
type SyntheticsVariableParser struct {
	Type SyntheticsGlobalVariableParserType `json:"type"`
	// Regex or JSON path used for the parser. Not used with type `raw`.
	Value *string `json:"value,omitempty"`
}

// NewSyntheticsVariableParser instantiates a new SyntheticsVariableParser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsVariableParser(type_ SyntheticsGlobalVariableParserType) *SyntheticsVariableParser {
	this := SyntheticsVariableParser{}
	this.Type = type_
	return &this
}

// NewSyntheticsVariableParserWithDefaults instantiates a new SyntheticsVariableParser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsVariableParserWithDefaults() *SyntheticsVariableParser {
	this := SyntheticsVariableParser{}
	return &this
}

// GetType returns the Type field value
func (o *SyntheticsVariableParser) GetType() SyntheticsGlobalVariableParserType {
	if o == nil {
		var ret SyntheticsGlobalVariableParserType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsVariableParser) GetTypeOk() (*SyntheticsGlobalVariableParserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SyntheticsVariableParser) SetType(v SyntheticsGlobalVariableParserType) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SyntheticsVariableParser) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsVariableParser) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SyntheticsVariableParser) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SyntheticsVariableParser) SetValue(v string) {
	o.Value = &v
}

func (o SyntheticsVariableParser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsVariableParser) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Type *SyntheticsGlobalVariableParserType `json:"type"`
	}{}
	all := struct {
		Type  SyntheticsGlobalVariableParserType `json:"type"}`
		Value *string                            `json:"value,omitempty"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Type = all.Type
	o.Value = all.Value
	return nil
}

type NullableSyntheticsVariableParser struct {
	value *SyntheticsVariableParser
	isSet bool
}

func (v NullableSyntheticsVariableParser) Get() *SyntheticsVariableParser {
	return v.value
}

func (v *NullableSyntheticsVariableParser) Set(val *SyntheticsVariableParser) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsVariableParser) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsVariableParser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsVariableParser(val *SyntheticsVariableParser) *NullableSyntheticsVariableParser {
	return &NullableSyntheticsVariableParser{value: val, isSet: true}
}

func (v NullableSyntheticsVariableParser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsVariableParser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
