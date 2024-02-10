// functions dealing with timeblock struct

package timeblocks

import (
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
)

// timeblocks dict
// key: timeblock id
// val: the timeblock
type TimeBlocks map[string]*TimeBlock

// main time block struct
type TimeBlock struct {
    Id string
    Title string

    Timerows []TimeRow
}

// time row in a time block
type TimeRow struct {
    Id string

    StartTime time.Time
    // only valid if Ongoing is false
    EndTime time.Time

    Ongoing bool
}

// add a timeblock to timeblocks dict. MUTATES the timeblock
// dict (but also returns same pointer)
func AddTimeBlock(timeblocks TimeBlocks) TimeBlocks {
    var newblock TimeBlock=newTimeBlock()

    timeblocks[newblock.Id]=&newblock

    return timeblocks
}

// toggle running state of a time block in time blocks dict
func ToggleTimeBlock(timeblocks TimeBlocks,timeblockId string) {
    var exists bool
    _,exists=timeblocks[timeblockId]

    if !exists {
        fmt.Printf("could not find timeblock id to toggle: %v\n",timeblockId)
        return
    }

    timeblocks[timeblockId].ToggleTimer()
}

func changeTimeBlockTitle(
    timeblocks TimeBlocks,
    timeblockId string,
    newTitle string,
) {
    var exists bool
    _,exists=timeblocks[timeblockId]

    if !exists {
        fmt.Printf("could not find timeblock id to change title: %v\n",timeblockId)
        return
    }

    timeblocks[timeblockId].Title=newTitle
}

// make new time block with random id
func newTimeBlock() TimeBlock {
    return TimeBlock {
        Id:GenUUid(),
    }
}

// make new timerow with random id, and the time set to now
func newTimeRow() TimeRow {
    return TimeRow {
        Id:GenUUid(),
        StartTime:time.Now(),
        Ongoing:true,
    }
}

// toggle running state of time block. adds a time row if the time block is not
// ongoing, or stops the running timerow
func (timeblock *TimeBlock) ToggleTimer() {
    if !timeblock.running() {
        timeblock.addTimeRow()
    } else {
        timeblock.Timerows[len(timeblock.Timerows)-1].stop()
    }
}

// return if a time block is running or not. it is running if the last time row is ongoing
func (timeblock *TimeBlock) running() bool {
    // if no timerows, timeblock is not running
    if len(timeblock.Timerows)==0 {
        return false
    }

    if timeblock.Timerows[len(timeblock.Timerows)-1].Ongoing {
        return true
    }

    return false
}

// adds a time row to a time block. only works if timeblock is not running
func (timeblock *TimeBlock) addTimeRow() {
    if timeblock.running() {
        fmt.Println("refused to add time row, timeblock already running")
        return
    }

    timeblock.Timerows=append(timeblock.Timerows,newTimeRow())
}

// closes a time row, setting the end time to Now
func (timerow *TimeRow) stop() {
    timerow.EndTime=time.Now()
    timerow.Ongoing=false
}

// compute total time of all timerows
func (timeblock *TimeBlock) totalTime() time.Duration {
    var totalTime time.Duration

    for i := range timeblock.Timerows {
        totalTime=totalTime+timeblock.Timerows[i].duration()
    }

    return totalTime
}

// compute duration of time row. invalid if still ongoing
func (timerow *TimeRow) duration() time.Duration {
    return timerow.EndTime.Sub(timerow.StartTime)
}

// remove timerow from timeblock
func (timeblock *TimeBlock) removeTimeRow(timerowId string) {
    for i := range timeblock.Timerows {
        if timeblock.Timerows[i].Id==timerowId {
            timeblock.Timerows=slices.Delete(timeblock.Timerows,i,i+1)
            return
        }
    }

    fmt.Printf("failed to find timerow id: %v\n",timerowId)
}

// return a short uuid
func GenUUid() string {
    return uuid.New().String()[0:6]
}

// parse custom short date string
func ParseShortDate(datestring string) time.Time {
    var res time.Time
    var e error

    res,e=time.Parse(
        "01/02 15:04",
        datestring,
    )

    res=res.AddDate(2024,0,0)

    if e!=nil {
        fmt.Println("failed to parse time string:",datestring)
    }

    return res
}