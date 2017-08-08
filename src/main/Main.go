package main

import (
    "fmt"
    "strconv"
    "bufio"
    "os"
    "strings"
    "clock"
)

type ClockInput struct {
    numberOfBalls int
    haltAtMinute int
}

// Reads each input line and parses out to two integers.
func ReadInput() ClockInput {
    reader := bufio.NewReader(os.Stdin)
    text, error := reader.ReadString('\n')
    text = strings.TrimSpace(text)
     if error != nil {
        fmt.Errorf("an error: %s", error)
    }

    stringSlice := strings.Split(text, " ")

    var j int = -1

    i, _ := strconv.Atoi(stringSlice[0])
    if len(stringSlice) == 2 {
        j, _ = strconv.Atoi(stringSlice[1])
    }

    return ClockInput{
        numberOfBalls: i,
        haltAtMinute: j,
    }
}

type JsonClock struct {
    Min   []int      `json:"Min"`
    FiveMin   []int      `json:"FiveMin"`
    Hour   []int      `json:"Hour"`
    Main   []int      `json:"Main"`

}

func main() {
    var inputSlice []ClockInput

    for {
        input := ReadInput()
        if input.numberOfBalls == 0 {
          break
        }
        inputSlice = append(inputSlice, input)
    }

    //var ballClock clock.BallClock
    //fmt.Println(ballClock)


    for _, j := range inputSlice {
        var ballClock clock.BallClock
        clock.NewClock(j.numberOfBalls, &ballClock)
        output := clock.RunClock(j.haltAtMinute, &ballClock)
        fmt.Println(output)
    }

}
