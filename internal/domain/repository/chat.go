package repository

import (
	"context"

	"github.com/K-Kizuku/spajam-backend/internal/domain/entity"
)

type IChatRepository interface {
	Create(ctx context.Context, chat entity.Chat) (*entity.Chat, error)
	FindChatByUserID(ctx context.Context, userID string) ([]entity.Chat, error)
}
