fogbugz
=======

Library for automated crash reports on fogbugz for Go applications using the bugscoutz API (http://help.fogcreek.com/7566/bugzscout-for-automatic-crash-reporting)

usage
-----

```go
package main

import (
  "os"

	"github.com/mulander/fogbugz"
)

func main() {
	reporter := fogbugz.New("https://project.fogbugz.com/scoutsubmit.asp",
		"ReporterUserName",
		"MyProject",
		"TelemetryArea",
		"reporter@example.com",
		"The occurrence of this problem has been noted. Thank you for using MyProject!",
		"MyProject v1.0.1",
		"1",
	)
	fogbugz.Init(reporter)

	file, err := os.Open("test.fogbugz")
	if err != nil {
		fogbugz.Fatal(err)
	}
}
```
