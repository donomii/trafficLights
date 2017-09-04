#lang racket

[require srfi/1]
[require srfi/19]

(require rackunit)

;List the timings, in seconds, follow by the state
[define time-table [circular-list 
                    '[270  "(N,S) Green,  (E,W) Red"]
                    '[30   "(N,S) Yellow, (E,W) Red"]
                    '[270  "(N,S) Red,    (E,W) Green"]
                    '[30   "(N,S) Red,    (E,W) Yellow"]
                    ]]

;Add a some seconds to a time (which must be in YYYY-MM-DD HH:MM:SS format)
[define [addTime aTime seconds]
  [date->string
   [time-utc->date
    [add-duration [date->time-utc
                   [string->date aTime "~Y-~m-~d ~H:~M:~S"]] [make-time time-duration 0 seconds]]] "~Y-~m-~d ~H:~M:~S"]]

;True if time a < time b
[define [compareTime a b]
  [time<? [date->time-utc [string->date a "~Y-~m-~d ~H:~M:~S"]] [date->time-utc [string->date b "~Y-~m-~d ~H:~M:~S"]]]]

;Now print the times, until nowTime >= endTime
[define[makeTimes nowTime endTime times]
  [if [compareTime nowTime endTime]
      [cons [format "~a: ~a" nowTime [cadar times]]
            [makeTimes [addTime nowTime [caar times]] endTime [cdr times]]]
      '[]]]

;Do some tests
[test-case "add time"
           (check-equal?  [addTime "2017-09-01 12:30:00" 0] "2017-09-01 12:30:00" "add zero time")
           (check-equal?  [addTime "2017-09-01 12:30:00" 30] "2017-09-01 12:30:30" "add 30 seconds")
           (check-equal?  [addTime "2017-09-01 12:30:00" 270] "2017-09-01 12:34:30" "add 4:30")
           (check-equal?  [addTime "2017-12-31 23:59:59" 30] "2018-01-01 00:00:29" "check rollover new year")
           (check-equal?  [addTime "2020-02-29 23:59:59" 30] "2020-03-01 00:00:29" "leap year")
           (check-equal?  [addTime "2019-02-28 23:59:59" 30] "2019-03-01 00:00:29" "non leap year")
           ]

[test-case "compare time"
           (check-equal?  [compareTime "2019-02-28 23:59:59" "2019-03-01 00:00:29"] #t "less than")
           (check-equal?  [compareTime "2019-03-01 00:00:29" "2019-02-28 23:59:59"] #f "greater than")
           ]

[test-case "sample output"
           (check-equal?  [makeTimes "2017-09-01 12:30:00" "2017-09-01 13:00:00" time-table]
                          '("2017-09-01 12:30:00: (N,S) Green,  (E,W) Red"
                            "2017-09-01 12:34:30: (N,S) Yellow, (E,W) Red"
                            "2017-09-01 12:35:00: (N,S) Red,    (E,W) Green"
                            "2017-09-01 12:39:30: (N,S) Red,    (E,W) Yellow"
                            "2017-09-01 12:40:00: (N,S) Green,  (E,W) Red"
                            "2017-09-01 12:44:30: (N,S) Yellow, (E,W) Red"
                            "2017-09-01 12:45:00: (N,S) Red,    (E,W) Green"
                            "2017-09-01 12:49:30: (N,S) Red,    (E,W) Yellow"
                            "2017-09-01 12:50:00: (N,S) Green,  (E,W) Red"
                            "2017-09-01 12:54:30: (N,S) Yellow, (E,W) Red"
                            "2017-09-01 12:55:00: (N,S) Red,    (E,W) Green"
                            "2017-09-01 12:59:30: (N,S) Red,    (E,W) Yellow")
                          "sample output")]

;Print the times
[let [[startTime "2017-09-01 12:30:00"][endTime "2017-09-01 13:00:00"]]
  [map displayln [makeTimes startTime endTime time-table]]
  "Demo complete"]