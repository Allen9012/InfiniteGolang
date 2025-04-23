package main

import "go-common/library/sync/pipeline/fanout"

var worker *fanout.Fanout = fanout.New("workerSubmit", fanout.Worker(4), fanout.Buffer(102400))
