/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SLOBulkDeleteResponseData An array of service level objective objects.
type SLOBulkDeleteResponseData struct {
	// An array of service level objective object IDs that indicates which objects that were completely deleted.
	Deleted *[]string `json:"deleted,omitempty"`
	// An array of service level objective object IDs that indicates which objects that were modified (objects for which at least one threshold was deleted, but that were not completely deleted).
	Updated *[]string `json:"updated,omitempty"`
}

// NewSLOBulkDeleteResponseData instantiates a new SLOBulkDeleteResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOBulkDeleteResponseData() *SLOBulkDeleteResponseData {
	this := SLOBulkDeleteResponseData{}
	return &this
}

// NewSLOBulkDeleteResponseDataWithDefaults instantiates a new SLOBulkDeleteResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOBulkDeleteResponseDataWithDefaults() *SLOBulkDeleteResponseData {
	this := SLOBulkDeleteResponseData{}
	return &this
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *SLOBulkDeleteResponseData) GetDeleted() []string {
	if o == nil || o.Deleted == nil {
		var ret []string
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOBulkDeleteResponseData) GetDeletedOk() (*[]string, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *SLOBulkDeleteResponseData) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given []string and assigns it to the Deleted field.
func (o *SLOBulkDeleteResponseData) SetDeleted(v []string) {
	o.Deleted = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *SLOBulkDeleteResponseData) GetUpdated() []string {
	if o == nil || o.Updated == nil {
		var ret []string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOBulkDeleteResponseData) GetUpdatedOk() (*[]string, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *SLOBulkDeleteResponseData) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given []string and assigns it to the Updated field.
func (o *SLOBulkDeleteResponseData) SetUpdated(v []string) {
	o.Updated = &v
}

func (o SLOBulkDeleteResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableSLOBulkDeleteResponseData struct {
	value *SLOBulkDeleteResponseData
	isSet bool
}

func (v NullableSLOBulkDeleteResponseData) Get() *SLOBulkDeleteResponseData {
	return v.value
}

func (v *NullableSLOBulkDeleteResponseData) Set(val *SLOBulkDeleteResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOBulkDeleteResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOBulkDeleteResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOBulkDeleteResponseData(val *SLOBulkDeleteResponseData) *NullableSLOBulkDeleteResponseData {
	return &NullableSLOBulkDeleteResponseData{value: val, isSet: true}
}

func (v NullableSLOBulkDeleteResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOBulkDeleteResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}