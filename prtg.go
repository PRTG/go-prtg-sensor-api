// Copyright (c) 2019, Paessler AG.
// All Rights Reserved

// Package PRTG implements the API for PRTG custom sensors.
// It provides all structs and constants needed to implement your own advanced exe sensor in Go.
package prtg

import (
	"encoding/json"
)

// SensorStatus has the whole JSON object
type SensorResponse struct {
	SensorResults SensorResults `json:"prtg"`
}

// SensorResults has all the channels
type SensorResults struct {
	SensorChannels []SensorChannel `json:"result"`
	Text           string          `json:"text,omitempty"`
	Error          string          `json:"error,omitempty"`
}

type Sensor interface {
	AddChannel(string) *SensorChannel
	MarshalToString() (string, error)
	SetSensorText(string) *SensorResults
	SetError(bool) *SensorResults
}

// Creates the Sensor instance and returns the interface.
func New() Sensor {
	return &SensorResults{}
}

// Name of the channel as displayed in user interfaces.
func (sr *SensorResults) AddChannel(channelName string) *SensorChannel {
	newChan := SensorChannel{Channel: channelName}
	sr.SensorChannels = append(sr.SensorChannels, newChan)

	return &sr.SensorChannels[len(sr.SensorChannels)-1]
}

// Create a JSON string from the PRTG object
func (sr *SensorResults) MarshalToString() (string, error) {
	bytes, err := json.Marshal(&SensorResponse{*sr})
	return string(bytes), err
}

// Text the sensor returns in the Message field with every scanning interval.
// There can be one message per sensor, regardless of the number of channels.
// Default is OK.
func (sr *SensorResults) SetSensorText(text string) *SensorResults {
	sr.Text = text
	return sr
}

// If enabled, the sensor will return an error status.
// This element can be combined with the SensorText element in order to show an error message.
// Default is 0.
func (sr *SensorResults) SetError(err bool) *SensorResults {
	sr.Error = "0"
	if err {
		sr.Error = "1"
	}

	return sr
}
