package roster

import (
	"fmt"
	"strings"
)

type Student struct {
	ID    string
	Name  string
	Score int
}

type Roster struct {
	students    []Student
	comparisons int
}

func New() *Roster {
	return &Roster{
		students: make([]Student, 0),
	}
}

func (r *Roster) Add(s Student) {
	s.Name = strings.TrimSpace(s.Name)
	r.students = append(r.students, s)
}

func (r *Roster) Len() int {
	return len(r.students)
}

func (r *Roster) FindByID(id string) (Student, bool) {
	for _, s := range r.students {
		r.comparisons++
		if s.ID == id {
			return s, true
		}
	}
	return Student{}, false
}

func (r *Roster) TotalScoreByName(name string) int {
	total := 0
	for _, s := range r.students {
		r.comparisons++
		if s.Name == name {
			total += s.Score
		}
	}
	return total
}

func (r *Roster) Comparisons() int {
	return r.comparisons
}

func FormatText(r *Roster) string {
	if r.Len() == 0 {
		return ""
	}
	res := ""
	for _, s := range r.students {
		res += fmt.Sprintf("%s %s %d\n", s.ID, s.Name, s.Score)
	}
	return res
}

func FormatCSV(r *Roster) string {
	res := "id,name,score\n"
	for _, s := range r.students {
		res += fmt.Sprintf("%s,%s,%d\n", s.ID, s.Name, s.Score)
	}
	return res
}