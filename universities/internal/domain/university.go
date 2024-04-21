package domain

type University struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	LongName     string  `json:"longName"`
	Logo         string  `json:"logo"`
	Rating       float64 `json:"rating"`
	Region       string  `json:"region"`
	Type         string  `json:"type"`
	StudyFields  int     `json:"studyFields"` // Направлений подготовки
	BudgetPlaces int     `json:"budgetPlaces"`
}
