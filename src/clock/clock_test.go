package clock

import (
    "testing"
    "clock"
)


var testInput = []struct {
    balls int
    haltAt int
    expected string
} {
    {30, -1, "30 balls cycle after 15 days."},
    {45, -1, "45 balls cycle after 378 days."},
    {30, 325, "{\"Min\":[],\"FiveMin\":[22,13,25,3,7],\"Hour\":[6,12,17,4,15],\"Main\":[11,5,26,18,2,30,19,8,24,10,29,20,16,21,28,1,23,14,27,9]}"},
    {26, -1, "Invalid number of balls : 26"},
    {27, -1, "27 balls cycle after 23 days."},
    {127, -1, "127 balls cycle after 2415 days."},
    {128, -1, "Invalid number of balls : 128"},
}

func TestRun(t *testing.T) {

    for _, input := range testInput {
        var ballClock clock.BallClock
        clock.NewClock(input.balls, &ballClock)
        output := clock.RunClock(input.haltAt, &ballClock)

        if output != input.expected {
            t.Errorf("Found[%s] Expected[%s] ", output, input.expected)
        }
    }
}


