package http

type TestCase struct {
	Name        string
	PathName    string
	Method      string
	ContentType string
	Content     string
	Status      int
}
