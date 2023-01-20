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

// GoidTokenData struct for GoidTokenData
type GoidTokenData struct {
	// OTP from whatsapp or SMS
	Otp *string `json:"otp,omitempty"`
	// OTP token from login request
	OtpToken *string `json:"otp_token,omitempty"`
}

// NewGoidTokenData instantiates a new GoidTokenData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoidTokenData() *GoidTokenData {
	this := GoidTokenData{}
	return &this
}

// NewGoidTokenDataWithDefaults instantiates a new GoidTokenData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoidTokenDataWithDefaults() *GoidTokenData {
	this := GoidTokenData{}
	return &this
}

// GetOtp returns the Otp field value if set, zero value otherwise.
func (o *GoidTokenData) GetOtp() string {
	if o == nil || o.Otp == nil {
		var ret string
		return ret
	}
	return *o.Otp
}

// GetOtpOk returns a tuple with the Otp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoidTokenData) GetOtpOk() (*string, bool) {
	if o == nil || o.Otp == nil {
		return nil, false
	}
	return o.Otp, true
}

// HasOtp returns a boolean if a field has been set.
func (o *GoidTokenData) HasOtp() bool {
	if o != nil && o.Otp != nil {
		return true
	}

	return false
}

// SetOtp gets a reference to the given string and assigns it to the Otp field.
func (o *GoidTokenData) SetOtp(v string) {
	o.Otp = &v
}

// GetOtpToken returns the OtpToken field value if set, zero value otherwise.
func (o *GoidTokenData) GetOtpToken() string {
	if o == nil || o.OtpToken == nil {
		var ret string
		return ret
	}
	return *o.OtpToken
}

// GetOtpTokenOk returns a tuple with the OtpToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoidTokenData) GetOtpTokenOk() (*string, bool) {
	if o == nil || o.OtpToken == nil {
		return nil, false
	}
	return o.OtpToken, true
}

// HasOtpToken returns a boolean if a field has been set.
func (o *GoidTokenData) HasOtpToken() bool {
	if o != nil && o.OtpToken != nil {
		return true
	}

	return false
}

// SetOtpToken gets a reference to the given string and assigns it to the OtpToken field.
func (o *GoidTokenData) SetOtpToken(v string) {
	o.OtpToken = &v
}

func (o GoidTokenData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Otp != nil {
		toSerialize["otp"] = o.Otp
	}
	if o.OtpToken != nil {
		toSerialize["otp_token"] = o.OtpToken
	}
	return json.Marshal(toSerialize)
}

type NullableGoidTokenData struct {
	value *GoidTokenData
	isSet bool
}

func (v NullableGoidTokenData) Get() *GoidTokenData {
	return v.value
}

func (v *NullableGoidTokenData) Set(val *GoidTokenData) {
	v.value = val
	v.isSet = true
}

func (v NullableGoidTokenData) IsSet() bool {
	return v.isSet
}

func (v *NullableGoidTokenData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoidTokenData(val *GoidTokenData) *NullableGoidTokenData {
	return &NullableGoidTokenData{value: val, isSet: true}
}

func (v NullableGoidTokenData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoidTokenData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


