// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamCreateAttributes Team creation attributes
type TeamCreateAttributes struct {
	// Free-form markdown description/content for the team's homepage
	Description *string `json:"description,omitempty"`
	// The team's identifier
	Handle string `json:"handle"`
	// The number of links belonging to the team
	LinkCount *int32 `json:"link_count,omitempty"`
	// The name of the team
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewTeamCreateAttributes instantiates a new TeamCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamCreateAttributes(handle string, name string) *TeamCreateAttributes {
	this := TeamCreateAttributes{}
	this.Handle = handle
	this.Name = name
	return &this
}

// NewTeamCreateAttributesWithDefaults instantiates a new TeamCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamCreateAttributesWithDefaults() *TeamCreateAttributes {
	this := TeamCreateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TeamCreateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamCreateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TeamCreateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetHandle returns the Handle field value.
func (o *TeamCreateAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *TeamCreateAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *TeamCreateAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetLinkCount returns the LinkCount field value if set, zero value otherwise.
func (o *TeamCreateAttributes) GetLinkCount() int32 {
	if o == nil || o.LinkCount == nil {
		var ret int32
		return ret
	}
	return *o.LinkCount
}

// GetLinkCountOk returns a tuple with the LinkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamCreateAttributes) GetLinkCountOk() (*int32, bool) {
	if o == nil || o.LinkCount == nil {
		return nil, false
	}
	return o.LinkCount, true
}

// HasLinkCount returns a boolean if a field has been set.
func (o *TeamCreateAttributes) HasLinkCount() bool {
	return o != nil && o.LinkCount != nil
}

// SetLinkCount gets a reference to the given int32 and assigns it to the LinkCount field.
func (o *TeamCreateAttributes) SetLinkCount(v int32) {
	o.LinkCount = &v
}

// GetName returns the Name field value.
func (o *TeamCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TeamCreateAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["handle"] = o.Handle
	if o.LinkCount != nil {
		toSerialize["link_count"] = o.LinkCount
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Handle *string `json:"handle"`
		Name   *string `json:"name"`
	}{}
	all := struct {
		Description *string `json:"description,omitempty"`
		Handle      string  `json:"handle"`
		LinkCount   *int32  `json:"link_count,omitempty"`
		Name        string  `json:"name"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if required.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "handle", "link_count", "name"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Handle = all.Handle
	o.LinkCount = all.LinkCount
	o.Name = all.Name
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}