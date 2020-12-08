package model

type stats struct {
	Freq    float64 `json:"freq"`
	AvgDiff float64 `json:"avg_diff"`
}

type teachingObjective struct {
	Desc        string `json:"desc"`
	Level       string `json:"level"`
	LessonCount int `json:"lesson_count"`
}


type Ktag struct {
	Id                 string `json:"_id"`
	Name               string `json:"name"`
	Type               int `json:"type"`
	Weight             int `json:"weight"`
	Wiki               string `json:"wiki"`
	Desc              string `json:"desc"`
	DescTex           string `json:"desc_tex"`
	ExtraDesc         string `json:"extra_desc"`
	ExtraDescTex      string `json:"extra_desc_tex"`
	Deleted           bool `json:"deleted"`
	AssessDirs        []int `json:"assess_dirs"`
	Path              string `json:"path"`
	ParentId          string `json:"parent_id"`
	Keywords          []string `json:"keywords"`
	Stats             stats `json:"stats"`
	TeachingObjective teachingObjective `json:"teaching_objective"`
}