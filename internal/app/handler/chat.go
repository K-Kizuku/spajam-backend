package handler

import (
	"encoding/json"
	"net/http"

	chatschema "github.com/K-Kizuku/spajam-backend/internal/app/handler/schema/chats"
	"github.com/K-Kizuku/spajam-backend/internal/app/service"
	"github.com/K-Kizuku/spajam-backend/internal/domain/entity"
	"github.com/K-Kizuku/spajam-backend/pkg/errors"
	"github.com/K-Kizuku/spajam-backend/pkg/middleware"
)

type IChatHandler interface {
	CreateChat() func(http.ResponseWriter, *http.Request) error
	GetChatByUserID() func(http.ResponseWriter, *http.Request) error
}

type ChatHandler struct {
	cs service.IChatService
}

func NewChatHandler(cs service.IChatService) IChatHandler {
	return &ChatHandler{cs: cs}
}

func (h *ChatHandler) CreateChat() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req chatschema.CreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errors.New(http.StatusBadRequest, err)
		}
		chat := entity.Chat{
			UserID1: req.UserID1,
			UserID2: req.UserID2,
			Content: req.Content,
		}
		h.cs.PostChat(r.Context(), chat)

		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

func (h *ChatHandler) GetChatByUserID() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		email := r.Context().Value(middleware.UserIDKey).(string)
		chats, err := h.cs.GetChatByUserID(r.Context(), email)
		if err != nil {
			return err
		}
		var c []chatschema.Chat
		for _, chat := range chats {
			c = append(c, chatschema.Chat{
				ChatID:  chat.ChatID,
				UserID1: chat.UserID1,
				UserID2: chat.UserID2,
				Content: chat.Content,
			})
		}
		res := chatschema.GetAllByUserIDResponse{
			Chats: c,
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return err
		}
		w.WriteHeader(http.StatusOK)
		return nil
	}
}
