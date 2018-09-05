package main

import (
	//"sort"
	"fmt"
	//"encoding/json"
	"time"
)

// Voting is ...
type Voting struct {
	VotingName      string      `json="votingname"`
	UserID          []string    `json="userid"`
	Candidate       map[string]int
	VotingNumber    int         `json="VotingNumber"`
	StartTime       int64       `json="starttime"`
	EndTime         int64       `json="endtime"`
	CurrentState    int         `json="currentstate"`  // state 0 : 투표 시작 전,   1 : 투표 가능,   2 : 투표 종료
}

// votingSlice is ...
var votingSlice []Voting // 투표 목록

// createVote creates Voting structure
func createVote(name string, startTime string, endTime string) { // Voting 구조체 생성
	layout := "01/02/2006 3:04:05 PM" // string으로 받은 시간을 Unix 시간으로 바꿔준다
	tStart, _ := time.Parse(layout, startTime) // ...
	tEnd, _ := time.Parse(layout, endTime) // ...
	tStartUTC := tStart.Unix() - 32400  // 받은 시간은 KST, Unix() 시간은 UTC
	tEndUTC := tEnd.Unix() - 32400
	v := Voting{Candidate: make(map[string]int)}
	votingSlice = append(votingSlice, v)
	votingInit(name, tStartUTC, tEndUTC)
}

// votingInit is ...
func votingInit(name string, startTime int64, endTime int64) {
	num := len(votingSlice) - 1
	votingSlice[num].VotingName = name
	votingSlice[num].VotingNumber = num + 1
	votingSlice[num].StartTime = startTime
	votingSlice[num].EndTime = endTime
	votingSlice[num].CurrentState = 0
}

// registerCandidate register candidate in Voting structure
func (v *Voting) registerCandidate(cd string) { // 후보 등록, cd는 후보 이름
	v.Candidate[cd] = 0
}

// getCandidate gets candidate in Voting structure
func (v *Voting) getCandidate() { // 후보 및 표 확인
	for key, val := range v.Candidate {
		fmt.Print(key, " ", val, " ")
	}
	fmt.Println()
}

// deleteCandidate deletes candidate in Voting structure
func (v *Voting) deleteCandidate(cd string) { // cd는 후보
	if v.StartTime < time.Now().Unix() { // 시작전에만 삭제 가능
		fmt.Println("후보를 삭제할 수 없습니다")
		return
	}
	delete(v.Candidate, cd)
}


func (v *Voting) check(id string) bool { // 투표를 이미 한 id인지 체크
	b := true
	for _, i := range v.UserID {
		if i == id {
			b = false
			break
		}
	}
	return b
}


// vote increases Poll belong to selected candidate
func (v *Voting) vote(cd string, userID string) { // 투표, cd는 후보
	id := userID
	if v.CurrentState == 1 {
		if v.check(id) {
			v.Candidate[cd]++
			v.saveCompleteID(id)
		} else {
			fmt.Println("중복")
		} 
	} else if v.CurrentState == 0 {
		fmt.Println("아직 투표할 수 없습니다")		
	} else {
		fmt.Println("투표가 끝났습니다")
	}
}

// changeState change Voting structure's CurrentState
func (v *Voting) changeState() { // Voting 상태 변화
	// TODO : 어떻게 할지 생각해봐야함
	//        시간이 되면 자동으로 Currentstate가 변해야 하는데 이걸 어떻게 구현 할지 생각
	//        무한 루프 go changeState()
	if v.StartTime < time.Now().Unix() && v.EndTime > time.Now().Unix() { // 투표 시작
		v.CurrentState = 1
	} else if v.EndTime < time.Now().Unix() { // 투표가 끝난 상태
		v.CurrentState = 2
	}
}

// viewCompleteVoting views completed Voting
func viewCompleteVoting() { // 전체 투표 목록 중 완료된 투표 조회
	for i := 0; i < len(votingSlice); i++ {
		if votingSlice[i].CurrentState == 2 { // 상태가 2인 투표들은 투표가 완료된 것들
			fmt.Println(votingSlice[i])
		}
	}
}

// saveCompleteID saves ID
func (v *Voting) saveCompleteID(id string) {
	v.UserID = append(v.UserID, id)
}

func main() { // Test
		createVote("First", "09/05/2018 2:06:05 PM", "09/05/2018 2:22:05 PM")
		votingSlice[0].registerCandidate("이상현")
		votingSlice[0].registerCandidate("김도정")
		votingSlice[0].registerCandidate("김현우")
		votingSlice[0].registerCandidate("유상욱")
		votingSlice[0].registerCandidate("최현빈")
		votingSlice[0].vote("이상현", "a")
		votingSlice[0].vote("이상현", "b")
		votingSlice[0].changeState()
		votingSlice[0].vote("이상현", "c")
		votingSlice[0].vote("김현우", "d")
		votingSlice[0].vote("김현우", "a")
		votingSlice[0].vote("김도정", "b")
		votingSlice[0].getCandidate()
		fmt.Println(votingSlice[0])
		fmt.Println(votingSlice[0])
		votingSlice[0].vote("이상현", "e")
		fmt.Println(votingSlice[0])
}