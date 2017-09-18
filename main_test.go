package main
import (
	"testing"
	"fmt"
	"time"
)

func TestAdvanceTime(t *testing.T) {
	now, _ := time.Parse("15:04", "12:00")
	newTime, status := advanceTime(now, "270s", func () string { return fmt.Sprint(    green(  "(N,S) Green  "),       red(    "(E,W) Red\n")) })
	correctTime, _ := time.Parse("15:04:05", "12:04:45") 
	if newTime == correctTime {
		fmt.Println("yay")
	}
	if status != "12:00:00   (N,S) Green  (E,W) Red\n" {
		t.Fatalf("Advance time failed!")
	}
}

func TestMakeTimes( t *testing.T) {
	var start, end time.Time
	start, _ = time.Parse("15:04", "12:00")
	end, _ = time.Parse("15:04", "12:30")

	correct := []string{"12:00:00   (N,S) Green  (E,W) Red\n", "12:04:30   (N,S) Yellow (E,W) Red\n", "12:05:00   (N,S) Red    (E,W) Green\n", "12:09:30   (N,S) Red    (E,W) Yellow\n", "12:10:00   (N,S) Green  (E,W) Red\n", "12:14:30   (N,S) Yellow (E,W) Red\n", "12:15:00   (N,S) Red    (E,W) Green\n", "12:19:30   (N,S) Red    (E,W) Yellow\n", "12:20:00   (N,S) Green  (E,W) Red\n", "12:24:30   (N,S) Yellow (E,W) Red\n", "12:25:00   (N,S) Red    (E,W) Green\n", "12:29:30   (N,S) Red    (E,W) Yellow\n" } 
	if fmt.Sprint(makeTimes(start, end)) !=  fmt.Sprint(correct) {
		t.Fatalf("Make times failed!")
	}
}
