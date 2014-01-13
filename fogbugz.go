package gofogbugz

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime"
)

type Scout struct {
	URL                 string
	UserName            string
	Project             string
	Area                string
	Email               string
	ScoutDefaultMessage string
	FriendlyResponse    string
	Prefix              string
	Logger              *log.Logger
}

func Init(s *Scout) {
	std = s
}

func checkInit() {
	if std == nil {
		panic("gofogbugz: Call gofogbugz.Init before using reporting")
	}
}

// Standard reporter
var std *Scout = nil

// Report obtains a stack trace and reports
// it to fogbugz.
// The err string is used as the bug title.
// Remaining fogbugz data is obtained from
// the Scout struct.
func (s *Scout) Report(title string) error {
	// Heavily based on http://play.golang.org/p/65P0tZOZe0
	var n int
	buf := make([]byte, 1<<20) // 1 MB buffer
	for i := 0; ; i++ {
		n = runtime.Stack(buf, true)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		if len(buf) >= 64<<20 {
			// Filled 64 MB - stop here
			break
		}
		buf = make([]byte, 2*len(buf))
	}
	values := url.Values{
		"ScoutUserName":       {s.UserName},
		"ScoutProject":        {s.Project},
		"ScoutArea":           {s.Area},
		"Description":         {s.Prefix + title},
		"ForceNewBug":         {"0"},
		"Extra":               {string(buf[:n])}, // stack trace
		"Email":               {s.Email},
		"ScoutDefaultMessage": {s.ScoutDefaultMessage},
		"FriendlyResponse":    {s.FriendlyResponse},
	}
	resp, err := http.PostForm(s.URL, values)
	if err != nil {
		log.Print(err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Sets the prefix for fogbugz bug reports.
// Suggested value is the application version.
func (s *Scout) SetPrefix(prefix string) {
	s.Prefix = prefix
}

// Fatal is equivalent to l.Print() followed by
// a call to os.Exit(1) reporting the bug to
// fogbugz beforehand.
func (s *Scout) Fatal(v ...interface{}) {
	str := fmt.Sprint(v...)
	s.Report(str)
	s.Logger.Fatal(str)
}

// Fatalf is equivalent to l.Printf() followed by
// a call to os.Exit(1) reporting the bug to
// fogbugz beforehand.
func (s *Scout) Fatalf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	s.Report(str)
	s.Logger.Fatal(str)
}

// Fatalln is equivalent to l.Println() followed by
// a call to os.Exit(1) reporting the bug to
// fogbugz beforehand.
func (s *Scout) Fatalln(v ...interface{}) {
	str := fmt.Sprintln(v...)
	s.Report(str)
	s.Logger.Fatal(str)
}

// Panic is equivalent to l.Print() followed by
// a call to panic() reporting the bug to
// fogbugz beforehand.
func (s *Scout) Panic(v ...interface{}) {
	str := fmt.Sprint(v...)
	s.Report(str)
	s.Logger.Panic(str)
}

// Panicf is equivalent to l.Printf() followed by
// a call to panic() reporting the bug to
// fogbugz beforehand.
func (s *Scout) Panicf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	s.Report(str)
	s.Logger.Panic(str)
}

// Panicln is equivalent to l.Println() followed by
// a call to panic() reporting the bug to
// fogbugz beforehand.
func (s *Scout) Panicln(v ...interface{}) {
	str := fmt.Sprintln(v...)
	s.Report(str)
	s.Logger.Panic(str)
}

// Print calls l.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// The bug is reported to fogbugz beforehand.
func (s *Scout) Print(v ...interface{}) {
	str := fmt.Sprint(v...)
	s.Report(str)
	s.Logger.Print(str)
}

// Printf calls l.Printf to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// The bug is reported to fogbugz beforehand.
func (s *Scout) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	s.Report(str)
	s.Logger.Print(str)
}

// Println calls l.Println to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// The bug is reported to fogbugz beforehand.
func (s *Scout) Println(format string, v ...interface{}) {
	str := fmt.Sprintln(v...)
	s.Report(str)
	s.Logger.Print(str)
}

// Standard reporter

// Sets the prefix for the standard fogbugz bug reports.
// Suggested value is the application version.
func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

// Fatal is equivalent to l.Print() followed by
// a call to os.Exit(1) reporting the bug to
// fogbugz beforehand.
func Fatal(v ...interface{}) {
	checkInit()
	s := fmt.Sprint(v...)
	std.Report(s)
	log.Fatal(s)
}

// Fatalf is equivalent to l.Printf() followed by
// a call to os.Exit(1) reporting the bug to
// fogbugz beforehand.
func Fatalf(format string, v ...interface{}) {
	checkInit()
	s := fmt.Sprintf(format, v...)
	std.Report(s)
	log.Fatal(s)
}

// Fatalln is equivalent to l.Println() followed by
// a call to os.Exit(1) reporting the bug to
// fogbugz beforehand.
func Fatalln(v ...interface{}) {
	checkInit()
	s := fmt.Sprintln(v...)
	std.Report(s)
	log.Fatal(s)
}

// Panic is equivalent to l.Print() followed by
// a call to panic() reporting the bug to
// fogbugz beforehand.
func Panic(v ...interface{}) {
	checkInit()
	s := fmt.Sprint(v...)
	std.Report(s)
	log.Panic(s)
}

// Panicf is equivalent to l.Printf() followed by
// a call to panic() reporting the bug to
// fogbugz beforehand.
func Panicf(format string, v ...interface{}) {
	checkInit()
	s := fmt.Sprintf(format, v...)
	std.Report(s)
	log.Panic(s)
}

// Panicln is equivalent to l.Println() followed by
// a call to panic() reporting the bug to
// fogbugz beforehand.
func Panicln(v ...interface{}) {
	checkInit()
	s := fmt.Sprintln(v...)
	std.Report(s)
	log.Panic(s)
}

// Print calls l.Print to print to the logger.
// Arguments are handled in the manner of fmt.Print.
// The bug is reported to fogbugz beforehand.
func Print(v ...interface{}) {
	checkInit()
	s := fmt.Sprint(v...)
	std.Report(s)
	log.Print(s)
}

// Printf calls l.Print to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
// The bug is reported to fogbugz beforehand.
func Printf(format string, v ...interface{}) {
	checkInit()
	s := fmt.Sprintf(format, v...)
	std.Report(s)
	log.Print(s)
}

// Println calls l.Print to print to the logger.
// Arguments are handled in the manner of fmt.Println.
// The bug is reported to fogbugz beforehand.
func Println(format string, v ...interface{}) {
	checkInit()
	s := fmt.Sprintln(v...)
	std.Report(s)
	log.Print(s)
}
