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
	"fmt"
)

// Bandwidth the model 'Bandwidth'
type Bandwidth string

// List of Bandwidth
const (
	BANDWIDTH_STREAMING Bandwidth = "uav_streaming"
	BANDWIDTH_LOWPOWERMODE Bandwidth = "uav_lowpowermode"
)

// All allowed values of Bandwidth enum
var AllowedBandwidthEnumValues = []Bandwidth{
	"uav_streaming",
	"uav_lowpowermode",
}

func (v *Bandwidth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Bandwidth(value)
	for _, existing := range AllowedBandwidthEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Bandwidth", value)
}

// NewBandwidthFromValue returns a pointer to a valid Bandwidth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBandwidthFromValue(v string) (*Bandwidth, error) {
	ev := Bandwidth(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Bandwidth: valid values are %v", v, AllowedBandwidthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Bandwidth) IsValid() bool {
	for _, existing := range AllowedBandwidthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Bandwidth value
func (v Bandwidth) Ptr() *Bandwidth {
	return &v
}

type NullableBandwidth struct {
	value *Bandwidth
	isSet bool
}

func (v NullableBandwidth) Get() *Bandwidth {
	return v.value
}

func (v *NullableBandwidth) Set(val *Bandwidth) {
	v.value = val
	v.isSet = true
}

func (v NullableBandwidth) IsSet() bool {
	return v.isSet
}

func (v *NullableBandwidth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBandwidth(val *Bandwidth) *NullableBandwidth {
	return &NullableBandwidth{value: val, isSet: true}
}

func (v NullableBandwidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBandwidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

