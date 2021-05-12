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

// MetricTagConfigurationUpdateRequest Request object that includes the metric that you would like to edit the tag configuration on.
type MetricTagConfigurationUpdateRequest struct {
	Data MetricTagConfigurationUpdateData `json:"data"`
}

// NewMetricTagConfigurationUpdateRequest instantiates a new MetricTagConfigurationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricTagConfigurationUpdateRequest(data MetricTagConfigurationUpdateData) *MetricTagConfigurationUpdateRequest {
	this := MetricTagConfigurationUpdateRequest{}
	this.Data = data
	return &this
}

// NewMetricTagConfigurationUpdateRequestWithDefaults instantiates a new MetricTagConfigurationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricTagConfigurationUpdateRequestWithDefaults() *MetricTagConfigurationUpdateRequest {
	this := MetricTagConfigurationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *MetricTagConfigurationUpdateRequest) GetData() MetricTagConfigurationUpdateData {
	if o == nil {
		var ret MetricTagConfigurationUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationUpdateRequest) GetDataOk() (*MetricTagConfigurationUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MetricTagConfigurationUpdateRequest) SetData(v MetricTagConfigurationUpdateData) {
	o.Data = v
}

func (o MetricTagConfigurationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *MetricTagConfigurationUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Data *MetricTagConfigurationUpdateData `json:"data"`
	}{}
	all := struct {
		Data MetricTagConfigurationUpdateData `json:"data"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Data == nil {
		return fmt.Errorf("Required field data missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Data = all.Data
	return nil
}

type NullableMetricTagConfigurationUpdateRequest struct {
	value *MetricTagConfigurationUpdateRequest
	isSet bool
}

func (v NullableMetricTagConfigurationUpdateRequest) Get() *MetricTagConfigurationUpdateRequest {
	return v.value
}

func (v *NullableMetricTagConfigurationUpdateRequest) Set(val *MetricTagConfigurationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTagConfigurationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTagConfigurationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTagConfigurationUpdateRequest(val *MetricTagConfigurationUpdateRequest) *NullableMetricTagConfigurationUpdateRequest {
	return &NullableMetricTagConfigurationUpdateRequest{value: val, isSet: true}
}

func (v NullableMetricTagConfigurationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTagConfigurationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
