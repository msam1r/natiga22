package result

import (
	"encoding/csv"
	"io"
)

type Result struct {
	Arabic     string
	English    string
	Studies    string
	Algebra    string
	Geometry   string
	Total_math string
	Science    string
	Total      string
	Religion   string
	Art        string
	Computer   string
	Sport      string
}

type Student struct {
	name   string
	number string
	school string
	s_type string
	result *Result
}

func NewStudent(number, name, school, s_type string) *Student {
	return &Student{
		name:   name,
		number: number,
		school: school,
		s_type: s_type,
	}
}

func (s *Student) AttachResult(result *Result) {
	s.result = result
}

func (s *Student) ToCSV(w io.Writer) {
	cw := csv.NewWriter(w)

	cw.Write([]string{
		s.name,
		s.number,
		s.school,
		s.s_type,
		s.result.Arabic,
		s.result.English,
		s.result.Studies,
		s.result.Algebra,
		s.result.Geometry,
		s.result.Total_math,
		s.result.Science,
		s.result.Total,
		s.result.Religion,
		s.result.Art,
		s.result.Computer,
		s.result.Sport,
	})

	cw.Flush()
}
