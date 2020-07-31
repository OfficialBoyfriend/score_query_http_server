package modules

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

var (
	year = []string{"2019"}
	num = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60"}
	//number = []string{"01","02","03","04"}
	//other = []string{"01","02","03"}
	number = []string{"01"}
	other = []string{"01"}
)

func TestId_Get(t *testing.T) {
	Init()
	id := new(Id)
	id.Get(map[string]interface{}{"group":"201802020213"})
	log.Println(id)
}

func TestId_GetFind(t *testing.T) {
	Init()
	id := new(Id)
	ids, err := id.GetFind(map[string]interface{}{"is_sync_ok": false}, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ids[0])
}

func TestStudentNumber_Create(t *testing.T) {
	for _, y := range year {
		for _, o := range other {
			for _, n := range number {
				for _, nMin := range num {
					studentId := fmt.Sprintf("%v02%v%v%v", y, o, n, nMin)
					resp, _ := http.Get(fmt.Sprintf("http://101.132.135.192:9090/send-student-id?id=%v", studentId))
					resp.Body.Close()
					time.Sleep(time.Second*5)
				}
			}
		}
	}
}