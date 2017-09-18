package main

import (
	"fmt"
	"os"
	"time"
	"github.com/joshlf13/term"
)

var f *os.File
type printFunc func() string

func makeTimes (now, end time.Time) (accum []string) {
	times := []func(time.Time)(time.Time, string){ How, Now, Brown, Cow }
	var str string
	for {
		for _,v := range times {
			now, str = v(now)
			accum = append(accum, str)
			if !now.Before(end) { return }
		}
	}
}

func advanceTime(now time.Time, duration string, printIt printFunc) (time.Time, string) {
	d,_ := time.ParseDuration(duration)
	tString := fmt.Sprint(now.Format("15:04:05"), "   ")
	fmt.Fprint(f, tString)
	return now.Add(d), tString+printIt()
}

func red(s string) string	{ term.Red(f, s); 		return s}
func yellow(s string) string	{ term.LightYellow(f, s); 	return s}
func green(s string) string	{ term.Green(f, s); 		return s}


func How ( now time.Time ) (time.Time, string) { return 	advanceTime(now, "270s", func () string { return fmt.Sprint( 	green( 	"(N,S) Green  "),	red(	"(E,W) Red\n")) }) }
func Now ( now time.Time ) (time.Time, string) { return 	advanceTime(now, "30s",  func () string { return fmt.Sprint(	yellow(	"(N,S) Yellow "), 	red(	"(E,W) Red\n")) }) }
func Brown ( now time.Time ) (time.Time, string) { return 	advanceTime(now, "270s", func () string { return fmt.Sprint(	red(   	"(N,S) Red    "), 	green(	"(E,W) Green\n")) }) }
func Cow ( now time.Time ) (time.Time, string) { return 	advanceTime(now, "30s",  func () string { return fmt.Sprint(	red(   	"(N,S) Red    "), 	yellow(	"(E,W) Yellow\n")) }) }

func main() {
	f = os.Stdout
	var start, end time.Time
	var err error
	if start, err = time.Parse("15:04", "12:00"); err != nil {
		panic(err)
	}
	if end, err = time.Parse("15:04", "12:30"); err != nil {
                panic(err)
        }
	makeTimes(start, end)
}
