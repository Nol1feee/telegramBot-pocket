package authServer

import (
	pocket "github.com/Nol1feee/telegramBot-pocket/internal/api/pocketSDK"
	"github.com/Nol1feee/telegramBot-pocket/internal/storage"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type AuthServer struct {
	//httpServer *http.Server

	pocket  *pocket.Client
	storage storage.TokenStorage
}

func NewServer(pocket *pocket.Client, storage storage.TokenStorage) *AuthServer {
	return &AuthServer{
		pocket:  pocket,
		storage: storage,
	}
}

func (s *AuthServer) Start() error {
	r := chi.NewRouter()

	r.Get("/auth", s.handleAuth)

	return http.ListenAndServe(":80", r)
}

func (s *AuthServer) handleAuth(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	idInt, _ := strconv.Atoi(id)

	reqToken, err := s.storage.Get(idInt, storage.RequestToken)
	if err != nil {
		//w.Write([]byte("Are you sure that you complete regestration?"))
		w.WriteHeader(http.StatusInternalServerError)
	}

	userInfo, err := s.pocket.Authetication(reqToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.storage.Save(idInt, userInfo.Access_token, storage.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	http.Redirect(w, r, "https://t.me/apiPocketBot", http.StatusPermanentRedirect)
}
