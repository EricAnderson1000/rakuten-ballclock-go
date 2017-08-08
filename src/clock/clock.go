package clock

import (
    "fmt"
    "encoding/json"
)

type track struct {
    capacity int
    balls []int
}

type BallClock struct {
    queueSize int
    queue []int
    minTrack track
    fiveMinTrack track
    hourTrack track
}

func reverse(balls []int) []int {
    for i, j := 0, len(balls)-1; i < j; i, j = i+1, j-1 {
        balls[i], balls[j] = balls[j], balls[i]
    }
    return balls
}

func (track *track) AtCapacity() bool {
    return len(track.balls) >= track.capacity
}

func (track *track) spill() (ball int, spillage []int) {
    ball = track.balls[track.capacity - 1]
    spillage = reverse(track.balls[:(track.capacity - 1)])
    track.balls = []int{}
    return ball, spillage
}

func (track *track) add(ball int) {
    track.balls = append(track.balls, ball)
}

func tickMinute(clock *BallClock) {

    ball := clock.queue[0]
    clock.queue = append(clock.queue[:0], clock.queue[1:]...)
    clock.minTrack.balls = append(clock.minTrack.balls, ball)

    if clock.minTrack.AtCapacity() {
        ball, spillage := clock.minTrack.spill()
        clock.fiveMinTrack.add(ball)
        clock.queue = append(clock.queue, spillage...)
    }

    if clock.fiveMinTrack.AtCapacity() {
        ball, spillage := clock.fiveMinTrack.spill()
        clock.hourTrack.add(ball)
        clock.queue = append(clock.queue, spillage...)
    }

    if clock.hourTrack.AtCapacity() {
        ball, spillage := clock.hourTrack.spill()
        clock.queue = append(clock.queue, spillage...)
        clock.queue = append(clock.queue, ball)
    }
}

func orderReset(clock *BallClock) bool {

    allBallsInQueue := len(clock.queue) == clock.queueSize

    return allBallsInQueue && inOrder(clock.queue)
}

func inOrder(queue []int) bool {

    for i := 0; i < len(queue); i++ {
        if queue[i] != i + 1 {
            return false
        }
    }
    return true
}

type jsonClock struct {
    Min   []int      `json:"Min"`
    FiveMin   []int      `json:"FiveMin"`
    Hour   []int      `json:"Hour"`
    Main   []int      `json:"Main"`

}

func RunClock(haltAt int, ballClock *BallClock) string {
    if ballClock.queueSize < 27 || ballClock.queueSize > 127 {
        return fmt.Sprintf("Invalid number of balls : %d", ballClock.queueSize)
    }

    minutes := 0

    for {
        tickMinute(ballClock)
        minutes += 1
        if haltAt == minutes || orderReset(ballClock) {
            break
        }
    }

    if haltAt == minutes {
        jsonClock := &jsonClock{
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