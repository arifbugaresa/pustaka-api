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
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title: bookRequest.Title,
		Price: int(price),
		Description: bookRequest.Description,
		Rating: int(rating),
		Discount: int(discount),
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
}




