package main

import (
    "fmt"
    "encoding/json"
    "strconv"
    "bufio"
    "os"
    "strings"
)

type Track struct {
    capacity int
    balls []int
}

type BallClock struct {
    queueSize int
    queue []int
    minTrack Track
    fiveMinTrack Track
    hourTrack Track
}

func Reverse(balls []int) []int {
    for i, j := 0, len(balls)-1; i < j; i, j = i+1, j-1 {
        balls[i], balls[j] = balls[j], balls[i]
    }
    return balls
}

func (track *Track) AtCapacity() bool {
    return len(track.balls) >= track.capacity
}

func (track *Track) Spill() (ball int, spillage []int) {
    ball = track.balls[track.capacity - 1]
    spillage = Reverse(track.balls[:(track.capacity - 1)])
    track.balls = []int{}
    return ball, spillage
}

func (track *Track) Add(ball int) {
    track.balls = append(track.balls, ball)
}

func NewClock(queueSize int, ballClock *BallClock) {
    ballClock.queueSize = queueSize
    for i := 1; i <= queueSize; i++ {
        ball := i
        ballClock.queue = append(ballClock.queue, ball)
    }
    ballClock.minTrack.capacity = 5
    ballClock.fiveMinTrack.capacity = 12
    ballClock.hourTrack.capacity = 12
}

func TickMinute(clock *BallClock) {

    ball := clock.queue[0]
    clock.queue = append(clock.queue[:0], clock.queue[1:]...)
    clock.minTrack.balls = append(clock.minTrack.balls, ball)

    if clock.minTrack.AtCapacity() {
        ball, spillage := clock.minTrack.Spill()
        clock.fiveMinTrack.Add(ball)
        clock.queue = append(clock.queue, spillage...)
    }

    if clock.fiveMinTrack.AtCapacity() {
        ball, spillage := clock.fiveMinTrack.Spill()
        clock.hourTrack.Add(ball)
        clock.queue = append(clock.queue, spillage...)
    }

    if clock.hourTrack.AtCapacity() {
        ball, spillage := clock.hourTrack.Spill()
        clock.queue = append(clock.queue, spillage...)
        clock.queue = append(clock.queue, ball)
    }
}

func OrderReset(clock *BallClock) bool {

    allBallsInQueue := len(clock.queue) == clock.queueSize

    return allBallsInQueue && InOrder(clock.queue)
}

func InOrder(queue []int) bool {

    for i := 0; i < len(queue); i++ {
        if queue[i] != i + 1 {
            return false
        }
    }
    return true
}

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

func RunClock(haltAt int, ballClock *BallClock) string {
    minutes := 0

    for {
        TickMinute(ballClock)
        minutes += 1
        if haltAt == minutes || OrderReset(ballClock) {
            break
        }
    }

    if haltAt == minutes {
        jsonClock := &JsonClock{
            Min:     ballClock.minTrack.balls,
            FiveMin: ballClock.fiveMinTrack.balls,
            Hour:        ballClock.hourTrack.balls,
            Main:        ballClock.queue,
        }
        result, _ := json.Marshal(jsonClock)
        return string(result)
    } else {
        days := minutes / (60 * 24)
        return fmt.Sprintf("%d balls cycle after %d days.", ballClock.queueSize, days)
    }
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

    for _, j := range inputSlice {
        var ballClock BallClock
        NewClock(j.numberOfBalls, &ballClock)
        output := RunClock(j.haltAtMinute, &ballClock)
        fmt.Println(output)
    }

}
