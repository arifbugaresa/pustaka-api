package book

type Service interface {
	FindAll() ([]Book, error)
	Create(book BookRequest) (Book, error)
	FindById(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(ID int) (Book, error) {
	return s.repository.FindById(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()

	book := Book{
		Title: bookRequest.Title,
		Price: int(price),
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
}




