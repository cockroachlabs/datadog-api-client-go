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

// UserInvitationData Object to create a user invitation.
type UserInvitationData struct {
	Relationships UserInvitationRelationships `json:"relationships"`
	Type          UserInvitationsType         `json:"type"`
}

// NewUserInvitationData instantiates a new UserInvitationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitationData(relationships UserInvitationRelationships, type_ UserInvitationsType) *UserInvitationData {
	this := UserInvitationData{}
	this.Relationships = relationships
	this.Type = type_
	return &this
}

// NewUserInvitationDataWithDefaults instantiates a new UserInvitationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationDataWithDefaults() *UserInvitationData {
	this := UserInvitationData{}
	var type_ UserInvitationsType = USERINVITATIONSTYPE_USER_INVITATIONS
	this.Type = type_
	return &this
}

// GetRelationships returns the Relationships field value
func (o *UserInvitationData) GetRelationships() UserInvitationRelationships {
	if o == nil {
		var ret UserInvitationRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *UserInvitationData) GetRelationshipsOk() (*UserInvitationRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *UserInvitationData) SetRelationships(v UserInvitationRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value
func (o *UserInvitationData) GetType() UserInvitationsType {
	if o == nil {
		var ret UserInvitationsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserInvitationData) GetTypeOk() (*UserInvitationsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserInvitationData) SetType(v UserInvitationsType) {
	o.Type = v
}

func (o UserInvitationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["relationships"] = o.Relationships
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

func (o *UserInvitationData) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Relationships *UserInvitationRelationships `json:"relationships"`
		Type          *UserInvitationsType         `json:"type"`
	}{}
	all := struct {
		Relationships UserInvitationRelationships `json:"relationships"}`
		Type          UserInvitationsType         `json:"type"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Relationships == nil {
		return fmt.Errorf("Required field relationships missing")
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Relationships = all.Relationships
	o.Type = all.Type
	return nil
}

type NullableUserInvitationData struct {
	value *UserInvitationData
	isSet bool
}

func (v NullableUserInvitationData) Get() *UserInvitationData {
	return v.value
}

func (v *NullableUserInvitationData) Set(val *UserInvitationData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitationData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitationData(val *UserInvitationData) *NullableUserInvitationData {
	return &NullableUserInvitationData{value: val, isSet: true}
}

func (v NullableUserInvitationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
