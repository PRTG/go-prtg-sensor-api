# go-prtg-sensor-api
API for writing PRTG custom sensors in Go.

[![Build Status](https://travis-ci.org/PRTG/go-prtg-sensor-api.svg "Travis CI status")](https://travis-ci.org/PRTG/go-prtg-sensor-api)
[![GoDoc](https://godoc.org/github.com/PRTG/go-prtg-sensor-api?status.svg)](https://godoc.org/github.com/PRTG/go-prtg-sensor-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/PRTG/go-prtg-sensor-api)](https://goreportcard.com/report/PRTG/go-prtg-sensor-api/)

## Example EXE Sensor
This simple example sensor sends an HTTP request to http://paessler.com
and returns two channels:
- `Response time` - time it takes to perform the request and read the body
- `Bytes read` - number of bytes the body contained

```golang
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/PRTG/go-prtg-sensor-api"
)

func main() {
	// Initiate Sensor instance
	sensor := prtg.New()

	// Log start time
	start := time.Now()
	// Perform HTTP request
	resp, err := http.Get("http://paessler.com")
	if err != nil {
		sensor.SetError(true)
		sensor.SetSensorText(err.Error())
		json, err := sensor.MarshalToString()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(json)
		return
	}
	// Read the response
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sensor.SetError(true)
		sensor.SetSensorText(err.Error())
		json, err := sensor.MarshalToString()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(json)
		return
	}
	// Evaluate results
	responseTime := time.Since(start)
	responseBytes := len(buffer)

	// Response time channel
	sensor.AddChannel("Response time").SetValue(responseTime.Seconds() * 1000).SetUnit(prtg.TimeResponse)
	// Bytes read channel
	sensor.AddChannel("Bytes read").SetValue(responseBytes).SetUnit(prtg.BytesFile)

	json, err := sensor.MarshalToString()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(json)
}
```

To test this example in your PRTG installation follow these steps:

1. Build the example
2. Copy the `.exe` file to `[PRTG install folder]\Custom Sensors\EXEXML`
3. Go to PRTG web interface
4. Add Sensor to a device of your choice
5. Choose `EXE/SCRIPT ADVANCED` as sensor typee (filter for `Custom Sensors`)
6. Under `EXE/Script` choose the `.exe` you copied in step 2
7. Done


## Example HTTP Data Advanced Sensor

This example sensor uses a http server to serve the sensor data, that can be pulled
by a http data advanced sensor.
- `Some value` - shows sample percentage value
- `Something` - shows if the Something service is up and running

The Sensor returns an error if "Something" is not ok.

```golang
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/PRTG/go-prtg-sensor-api"
)

func main() {
	// Create a webserver listening on "/"
	http.HandleFunc("/", reportStatus)
	http.ListenAndServe(":8080", nil)
}

func reportStatus(w http.ResponseWriter, r *http.Request) {
	// Initiate PRTG instance
	sensor := prtg.New()

	// Set sensor text
	sensor.SetSensorText("This is a test sensor")

	// Add a channel with a random float value in Percent
	sensor.AddChannel("Some value").SetValue(rand.Float64() * 100).
		SetUnit(prtg.Percent).SetMaxWarnLimit(80).SetMaxErrLimit(90)

	// Take a look if Something is working
	isUp, err := isSomethingUp()
	// Create a Sensor that shows the uptime of Something
	sensor.AddChannel("Something").SetValue(isUp).
		SetValueLookup("prtg.standardlookups.connectionstate.stateonlineok")
	// Create error message on sensor if the Something is down
	if err != nil {
		sensor.SetError(true)
		sensor.SetSensorText("Test sensor error: " + err.Error())
	}

	// Create json output
	json, err := sensor.MarshalToString()
	if err != nil {
		log.Fatal(err)
	}

	// Deliver to website
	_, err = fmt.Fprint(w, json)
	if err != nil {
		log.Fatal(err)
	}
}

func isSomethingUp() (bool, error) {
	// Generate a "Thing" to watch, that is working in 90% of the time
	if rand.Intn(10) > 1 {
		return true, nil
	}
	return false, fmt.Errorf("the Something is struggling")
}
```

To test this example in your PRTG installation follow these steps:

1. Build the example
2. Run the binary
3. Go to PRTG web interface
4. Add Sensor to a device of your choice
5. Choose `HTTP Data Advanced` as sensor type
6. Under `URL` choose the ip of your device with the port `8080`
7. Done

## Documentation
- [PRTG Manual](https://www.paessler.com/manuals/prtg/exe_script_advanced_sensor)
