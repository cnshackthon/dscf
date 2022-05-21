/*
Network as Code

Manipulate network conditions via simplified REST calls

API version: 2
Contact: todd.levi@nokia.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datarepository

import (
	"encoding/json"
)

// BandwidthResponse struct for BandwidthResponse
type BandwidthResponse struct {
	Id string `json:"id" yaml:"id" bson:"id" mapstructure:"Id"`
}

// NewBandwidthResponse instantiates a new BandwidthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBandwidthResponse(id string) *BandwidthResponse {
	this := BandwidthResponse{}
	this.Id = id
	return &this
}

// NewBandwidthResponseWithDefaults instantiates a new BandwidthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBandwidthResponseWithDefaults() *BandwidthResponse {
	this := BandwidthResponse{}
	return &this
}

// GetId returns the Id field value
func (o *BandwidthResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BandwidthResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BandwidthResponse) SetId(v string) {
	o.Id = v
}

func (o BandwidthResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBandwidthResponse struct {
	value *BandwidthResponse
	isSet bool
}

func (v NullableBandwidthResponse) Get() *BandwidthResponse {
	return v.value
}

func (v *NullableBandwidthResponse) Set(val *BandwidthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBandwidthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBandwidthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBandwidthResponse(val *BandwidthResponse) *NullableBandwidthResponse {
	return &NullableBandwidthResponse{value: val, isSet: true}
}

func (v NullableBandwidthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBandwidthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

