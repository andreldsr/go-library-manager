package dtos

type BookListDto struct {
	Id                int    `json:"id"`
	Title             string `json:"title"`
	RegisterNumber    int    `json:"registerNumber"`
	CDD               string `json:"cdd"`
	Index             string `json:"index"`
	PublisherName     string `json:"publisherName" gorm:"references:publisher.name"`
	AuthorsNames      string `json:"authorsNames"`
	LendingId         int    `json:"lendingId,omitempty"`
	LendingReturnDate string `json:"lendingReturnDate,omitempty"`
}

type BookStats struct {
	Total   int `json:"total,omitempty"`
	Lent    int `json:"lent"`
	Today   int `json:"today"`
	Delayed int `json:"delayed"`
}

type CreateBookDto struct {
	RegisterNumber   string `json:"registerNumber"`
	RegistrationDate string `json:"registrationDate"`
	Authors          string `json:"authors"`
	Title            string `json:"title"`
	Volume           string `json:"volume"`
	Copy             string `json:"copy"`
	Location         string `json:"location"`
	Publisher        string `json:"publisher"`
	PublicationYear  int    `json:"publicationYear"`
	AcquisitionForm  string `json:"acquisitionForm"`
	Index            string `json:"index"`
	Cdd              string `json:"cdd"`
	Observation      string `json:"observation"`
}

type UpdateBookDto struct {
	RegisterNumber  *string `json:"registerNumber,omitempty"`
	Authors         *string `json:"authors,omitempty"`
	Title           *string `json:"title,omitempty"`
	Volume          *string `json:"volume,omitempty"`
	Copy            *string `json:"copy,omitempty"`
	Location        *string `json:"location,omitempty"`
	Publisher       *string `json:"publisher,omitempty"`
	PublicationYear *int    `json:"publicationYear,omitempty"`
	AcquisitionForm *string `json:"acquisitionForm,omitempty"`
	Index           *string `json:"index,omitempty"`
	Cdd             *string `json:"cdd,omitempty"`
	Observation     *string `json:"observation,omitempty"`
}
