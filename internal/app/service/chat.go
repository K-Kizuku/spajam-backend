package service

import (
	"context"

	"github.com/K-Kizuku/spajam-backend/internal/domain/entity"
	"github.com/K-Kizuku/spajam-backend/internal/domain/repository"
	"github.com/K-Kizuku/spajam-backend/pkg/uuid"
)

type IChatService interface {
	PostChat(ctx context.Context, chat entity.Chat) (*entity.Chat, error)
	GetChatByUserID(ctx context.Context, userID string) ([]entity.Chat, error)
}

type ChatService struct {
	cr repository.IChatRepository
}

func NewChatService(cr repository.IChatRepository) IChatService {
	return &ChatService{
		cr: cr,
	}
}

func (s *ChatService) PostChat(ctx context.Context, chat entity.Chat) (*entity.Chat, error) {
	chat.ChatID = uuid.New()
	createdChat, err := s.cr.Create(ctx, chat)
	if err != nil {
		return nil, err
	}
	return createdChat, nil
}

func (s *ChatService) GetChatByUserID(ctx context.Context, userID string) ([]entity.Chat, error) {
	chat, err := s.cr.FindChatByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
