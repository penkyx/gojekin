/*
goid.gojekapi.com

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goid-gojek

import (
	"encoding/json"
)

// LoginRequestResponseDataNextState struct for LoginRequestResponseDataNextState
type LoginRequestResponseDataNextState struct {
	State *string `json:"state,omitempty"`
	Destination *string `json:"destination,omitempty"`
	TimerInSeconds *float32 `json:"timer_in_seconds,omitempty"`
}

// NewLoginRequestResponseDataNextState instantiates a new LoginRequestResponseDataNextState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginRequestResponseDataNextState() *LoginRequestResponseDataNextState {
	this := LoginRequestResponseDataNextState{}
	return &this
}

// NewLoginRequestResponseDataNextStateWithDefaults instantiates a new LoginRequestResponseDataNextState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginRequestResponseDataNextStateWithDefaults() *LoginRequestResponseDataNextState {
	this := LoginRequestResponseDataNextState{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LoginRequestResponseDataNextState) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginRequestResponseDataNextState) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LoginRequestResponseDataNextState) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *LoginRequestResponseDataNextState) SetState(v string) {
	o.State = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *LoginRequestResponseDataNextState) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginRequestResponseDataNextState) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *LoginRequestResponseDataNextState) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *LoginRequestResponseDataNextState) SetDestination(v string) {
	o.Destination = &v
}

// GetTimerInSeconds returns the TimerInSeconds field value if set, zero value otherwise.
func (o *LoginRequestResponseDataNextState) GetTimerInSeconds() float32 {
	if o == nil || o.TimerInSeconds == nil {
		var ret float32
		return ret
	}
	return *o.TimerInSeconds
}

// GetTimerInSecondsOk returns a tuple with the TimerInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginRequestResponseDataNextState) GetTimerInSecondsOk() (*float32, bool) {
	if o == nil || o.TimerInSeconds == nil {
		return nil, false
	}
	return o.TimerInSeconds, true
}

// HasTimerInSeconds returns a boolean if a field has been set.
func (o *LoginRequestResponseDataNextState) HasTimerInSeconds() bool {
	if o != nil && o.TimerInSeconds != nil {
		return true
	}

	return false
}

// SetTimerInSeconds gets a reference to the given float32 and assigns it to the TimerInSeconds field.
func (o *LoginRequestResponseDataNextState) SetTimerInSeconds(v float32) {
	o.TimerInSeconds = &v
}

func (o LoginRequestResponseDataNextState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.TimerInSeconds != nil {
		toSerialize["timer_in_seconds"] = o.TimerInSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableLoginRequestResponseDataNextState struct {
	value *LoginRequestResponseDataNextState
	isSet bool
}

func (v NullableLoginRequestResponseDataNextState) Get() *LoginRequestResponseDataNextState {
	return v.value
}

func (v *NullableLoginRequestResponseDataNextState) Set(val *LoginRequestResponseDataNextState) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginRequestResponseDataNextState) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginRequestResponseDataNextState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginRequestResponseDataNextState(val *LoginRequestResponseDataNextState) *NullableLoginRequestResponseDataNextState {
	return &NullableLoginRequestResponseDataNextState{value: val, isSet: true}
}

func (v NullableLoginRequestResponseDataNextState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginRequestResponseDataNextState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


