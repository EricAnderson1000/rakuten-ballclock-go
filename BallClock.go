package main

import "fmt"
import "encoding/json"

type Track struct {
    name string
    capacity int
    balls []int
}

type BallClock struct {
    queueSize int
    queue []int `json:"main"`
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

func (ballClock *BallClock) Json() string {
    fmt.Println(ballClock)
    result, _ := json.Marshal(ballClock)
    fmt.Println(string(result))
    return string(result)
}

func newClock(queueSize int, ballClock *BallClock) {
    ballClock.queueSize = queueSize
    for i := 1; i <= queueSize; i++ {
        ball := i
        ballClock.queue = append(ballClock.queue, ball)
    }
    ballClock.minTrack.capacity = 5
    ballClock.minTrack.name = "Min"
    ballClock.fiveMinTrack.capacity = 12
    ballClock.fiveMinTrack.name = "FiveMin"
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

func inOrder(queue []int) bool {

    for i := 0; i < len(queue); i++ {
        if queue[i] != i + 1 {
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



type JsonClock struct {
    Min   []int      `json:"Min"`
    FiveMin   []int      `json:"FiveMin"`
    Hour   []int      `json:"Hour"`
    Main   []int      `json:"Main"`

}

func runClock(ballClock *BallClock) {
    minutes := 0

    for {
        tickMinute(ballClock)
        minutes += 1
        if isOrderReset(ballClock) {
            break
        }
    }

    days := minutes / (60 * 24)
    fmt.Printf("%d balls cycle after %d days.\n", ballClock.queueSize, days)

    jsonClock := &JsonClock{
        Min:     ballClock.minTrack.balls,
        FiveMin: ballClock.fiveMinTrack.balls,
        Hour:        ballClock.hourTrack.balls,
        Main:        ballClock.queue,
    }
    result, _ := json.Marshal(jsonClock)
    fmt.Println(string(result))
}

func main() {

    var ballClock BallClock
    
    newClock(127, &ballClock)
    runClock(&ballClock)

}
