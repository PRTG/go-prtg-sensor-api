// Copyright Paessler AG.
// All Rights Reserved

package prtg

import (
	"encoding/json"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	os.Exit(m.Run())
}

func TestAddChannel(t *testing.T) {
	// Create an empty SensorResponse
	response := &SensorResponse{}
	// Create channel with random data
	channel := createRandomChannel()
	// Add channel to response
	response.AddChannel(channel)
	// Check if it safely arrived
	if channel != response.PRTG.Result[0] {
		t.Error("response.AddChannel(channel) => failed to add channel")
	}
}

func TestString(t *testing.T) {
	// Create response with some random channels
	response := &SensorResponse{}
	for i := 0; i < 3; i++ {
		response.AddChannel(createRandomChannel())
	}
	s := response.String()
	// Unmarshal string and check if we get back what we started with
	var checkResponse *SensorResponse
	if json.Unmarshal([]byte(s), &checkResponse) != nil {
		t.Errorf("Failed to unmarshal stringified response")
		t.FailNow()
	}
	if !reflect.DeepEqual(response, checkResponse) {
		// Errors little helper function
		toString := func(v interface{}) string {
			b, _ := json.MarshalIndent(v, "", " ")
			return string(b)
		}
		t.Errorf("Response marshaled from stringified response does not match original response.\nExpected:\n%s\nGot:\n%s", toString(response), toString(checkResponse))
	}
}

func TestStringError(t *testing.T) {
	response := &SensorResponse{}
	var result SensorResult
	response.PRTG = result
	response.String()
}

// createRandomChannel generates SensorChannels with random settings for testing
func createRandomChannel() SensorChannel {
	// Helper to create semi-random string
	randomString := func() string {
		return strconv.FormatInt(rand.Int63n(245285), 10)
	}
	// Create channel and fill it with random data
	channel := SensorChannel{}
	channel.Channel = randomString()
	channel.Value = randomString()
	channel.CustomUnit = randomString()
	channel.DecimalMode = randomString()
	channel.Float = rand.Intn(2)
	channel.LimitErrorMsg = randomString()
	channel.LimitMaxError = rand.Int()
	channel.LimitMaxWarning = rand.Int()
	channel.LimitMinError = rand.Int()
	channel.LimitMode = rand.Intn(2)
	channel.LimitWarningMsg = randomString()
	channel.Mode = randomString()
	channel.NotifyChanged = rand.Intn(2) != 0
	channel.ShowChart = rand.Intn(2)
	channel.ShowTable = rand.Intn(2)
	channel.SpeedSize = randomString()
	channel.SpeedTime = randomString()
	channel.Unit = randomString()
	channel.ValueLookup = randomString()
	channel.Warning = rand.Int()
	return channel
}
