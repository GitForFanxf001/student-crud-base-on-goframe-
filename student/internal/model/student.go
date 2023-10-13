package model

type StudentAddInput struct {
	Sname    string
	Sscore   int
	Ssubject string
}

type StudentAddOutput struct {
	Sname string
}

type StudentDeleteInput struct {
	Id int
}

type StudentDeleteOutput struct {
	Sname string
}

type StudentQueryInput struct {
	Id int
}

type StudentQueryOutput struct {
	Id       int    `json:"id"`
	Sname    string `json:"name"`
	Sscore   int    `json:"score"`
	Ssubject string `json:"subject"`
}
