package main

import (
    "fmt"
)

type Ball struct {
    number int
}

type Track struct {
    name string
    capacity int
    balls []Ball
}

type BallClock struct {
    queueSize int
    queue []Ball
    minTrack Track
    fiveMinTrack Track
    hourTrack Track
}

func (track *Track) AtCapacity() bool {
    return len(track.balls) >= track.capacity
}

func (track *Track) Spill() (ball Ball, spillage []Ball) {
    ball = track.balls[track.capacity - 1]
    spillage = track.balls[:(track.capacity - 1)]
    track.balls = []Ball{}
    return ball, spillage
}

func (track *Track) Add(ball Ball)  {
    track.balls = append(track.balls, ball)
}

func newClock(queueSize int, ballClock *BallClock) {
    ballClock.queueSize = queueSize
    for i := 1; i <= queueSize; i++ {
        ball := Ball{i}
        ballClock.queue = append(ballClock.queue, ball)
    }
    ballClock.minTrack.capacity = 5
    ballClock.minTrack.name = "Minute"
    ballClock.fiveMinTrack.capacity = 12
    ballClock.fiveMinTrack.name = "Five Minute"
    ballClock.hourTrack.capacity = 12
    ballClock.hourTrack.name = "Hour"
}

func tickMinute(clock *BallClock) {

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

func isOrderReset(clock *BallClock) bool {

    allBallsInQueue := len(clock.queue) == clock.queueSize

    return allBallsInQueue && inOrder(clock.queue)
}

func inOrder(queue []Ball) bool {

    for i := 0; i < len(queue); i++ {
        if queue[i].number != i + 1 {
            return false
        }
    }
    return true
}

//func readInput()  {
//    //b, _ := strconv.ParseInt("123", 0, 64)
//    //fmt.Println(b)
//
//    //reader := bufio.NewReader(os.Stdin)
//    //text, _ := reader.ReadString('\n')
//    ////fmt.Println(text)
//    //i, _ := strconv.Atoi(text)
//    //fmt.Println(i)
//    //fmt.Printf("%d \n", i)
//
//
//    var i, j int
//    fmt.Scan(&i, &j, i, j)
//    fmt.Println(i, j)
//
//
//    //return i, 0
//}

func main() {

    var ballClock BallClock
    newClock(27, &ballClock)

    minutes := 0

    //for i := 0; i < 720; i++ {
    //    tickMinute(&ballClock)
    //    minutes += 1
    //}

    for {
        tickMinute(&ballClock)
        minutes += 1
        if isOrderReset(&ballClock) {
            break
        }
    }


    fmt.Println(ballClock)
    days := float64(minutes) / (60 * 24)
    fmt.Println(days)

}
