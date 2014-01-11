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
	reporter := fogbugz.Scout{
		URL : "https://project.fogbugz.com/scoutsubmit.asp",
		UserName : "ReporterUserName",
		Project: "MyProject",
		Area: "TelemetryArea",
		Email: "reporter@example.com",
		ScoutDefaultMessage: "The occurrence of this problem has been noted. Thank you for using MyProject!",
		Prefix: "MyProject v1.0.1 - ",
		FriendlyResponse: "1",
		Logger: nil, // call Init and use the default logger
	}
	fogbugz.Init(reporter)

	file, err := os.Open("test.fogbugz")
	if err != nil {
		fogbugz.Fatal(err)
	}
}

```
