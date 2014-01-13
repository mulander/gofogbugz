gofogbugz
=======

Library for automated crash reports on fogbugz for Go applications using the bugscoutz API (http://help.fogcreek.com/7566/bugzscout-for-automatic-crash-reporting)

usage
-----

```go
package main

import (
	"os"

	"github.com/mulander/gofogbugz"
)

func main() {
	reporter := &gofogbugz.Scout{
		URL:                 "https://project.fogbugz.com/scoutsubmit.asp",
		UserName:            "ReporterUserName",
		Project:             "MyProject",
		Area:                "TelemetryArea",
		Email:               "reporter@example.com",
		ScoutDefaultMessage: "The occurrence of this problem has been noted. Thank you for using MyProject!",
		Prefix:              "MyProject v1.0.1 - ",
		FriendlyResponse:    "1",
		Logger:              nil, // call Init and use the default logger
	}
	gofogbugz.Init(reporter)

	file, err := os.Open("test.gofogbugz")
	if err != nil {
		gofogbugz.Fatal(err)
	}
	defer file.Close()
}
```

output
------

Running the above code on a configured fogbugz instance will result in the following bug report

![fogbugz bug report](http://i.imgur.com/BcfZoMq.png)
