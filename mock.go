package gitstorage

type Mock struct {
	ReturnString  string
	ReturnError   error
	ReturnStrings []string
}

func NewMock(ReturnString string, ReturnError error) Abstractor {
	return &Mock{}
}

func (m *Mock) Create(string, string) error {
	return m.ReturnError
}

func (m *Mock) List() ([]string, error) {
	return m.ReturnStrings, m.ReturnError
}

func (m *Mock) Read(string) (string, error) {
	return m.ReturnString, m.ReturnError
}

func (m *Mock) Update(string, string) error {
	return m.ReturnError
}

func (m *Mock) Delete(string) error {
	return m.ReturnError
}

func (m *Mock) getHead() (string, error) {
	return m.ReturnString, m.ReturnError
}

func (m *Mock) push(string) error {
	return m.ReturnError
}

func (m *Mock) createFile(string, string) error {
	return m.ReturnError
}

func (m *Mock) add(string) (string, error) {
	return m.ReturnString, m.ReturnError
}

func (m *Mock) get(string) (string, error) {
	return m.ReturnString, m.ReturnError
}

func (m *Mock) commit() (string, error) {
	return m.ReturnString, m.ReturnError
}

func (m *Mock) deleteFile(string) error {
	return m.ReturnError
}

func (m *Mock) objects() ([]string, error) {
	return m.ReturnStrings, m.ReturnError
}
