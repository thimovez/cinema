package api

type FilmRepo interface {
	StoreMessage(msg ChatMessage) error
	LastMessage() []string
}

type FilmService struct {
	chat FilmRepo
}

func NewFilmService(chat FilmRepo) *FilmService {
	return &FilmService{
		chat: chat,
	}
}
