package modeles

type DateOnly struct {
	Year  int
	Month byte
	Day   byte
}

type PdfDoc struct {
	Name          string
	Enterprise    string
	WorkStartDate string
	JobTitle      string
	SubjectArray  []string
}
