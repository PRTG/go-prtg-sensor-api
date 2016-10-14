# go-prtg-api
API for writing PRTG custom sensors in Go.

## Example
This simple exampe sensor sends an HTTP request to http://paessler.com
and returns two channels:
- `Response time` - time it takes to perform the request and read the body
- `Bytes read` - number of bytes the body contained

```golang
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PaesslerAG/go-prtg-sensor-api"
)

func main() {
	// Create empty response and log start time
	r := &prtg.SensorResponse{}
	start := time.Now()
	// Perform HTTP request
	resp, err := http.Get("http://paessler.com")
	if err != nil {
		r.PRTG.Error = "1"
		r.PRTG.Text = err.Error()
		fmt.Println(r.String())
		return
	}
	// Read the response
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.PRTG.Error = "1"
		r.PRTG.Text = err.Error()
		fmt.Println(r.String())
		return
	}
	// Evaluate results
	responseTime := time.Since(start)
	responseBytes := len(buffer)
	// Response time channel
	r.AddChannel(prtg.SensorChannel{
		Channel:   "Response time",
		Value:     fmt.Sprintf("%f", responseTime.Seconds()*1000),
		Float:     1,
		ShowChart: 1,
		ShowTable: 1,
		Unit:      prtg.UnitTimeResponse,
	})
	// Bytes read channel
	r.AddChannel(prtg.SensorChannel{
		Channel:   "Bytes read",
		Value:     fmt.Sprintf("%d", responseBytes),
		ShowChart: 1,
		ShowTable: 1,
		Unit:      prtg.UnitBytesFile,
	})

	fmt.Println(r.String())
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

## Documentation
- [godoc.org/github.com/PaesslerAG/go-prtg-sensor-api](https://godoc.org/github.com/PaesslerAG/go-prtg-sensor-api)
- [PRTG Manual](https://www.paessler.com/manuals/prtg/exe_script_advanced_sensor)
