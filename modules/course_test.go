package modules

import (
	"testing"
)

func TestCourse_GetAutoCreate(t *testing.T) {
	Init()
	course := new(Course)
	t.Log(course.GetAutoCreate(map[string]interface{}{"name": "900909"}))
	t.Log(course)
}