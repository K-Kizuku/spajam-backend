package repository

import (
	"context"

	"github.com/K-Kizuku/spajam-backend/db/sql/query"
	"github.com/K-Kizuku/spajam-backend/internal/domain/entity"
	"github.com/K-Kizuku/spajam-backend/internal/domain/repository"
	"github.com/K-Kizuku/spajam-backend/pkg/errors"
)

type ChatRepository struct {
	queries *query.Queries
}

func NewChatRepository(queries *query.Queries) repository.IChatRepository {
	return &ChatRepository{queries: queries}
}

func (r *ChatRepository) Create(ctx context.Context, chat entity.Chat) (*entity.Chat, error) {
	c, err := r.queries.CreateChat(ctx, query.CreateChatParams{
		ChatID:  chat.ChatID,
		UserID1: chat.UserID1,
		UserID2: chat.UserID2,
		Content: chat.Content,
	})
	if err != nil {
		return nil, errors.HandleDBError(err)
	}
	createdChat := &entity.Chat{
		ChatID:  c.ChatID,
		UserID1: c.UserID1,
		UserID2: c.UserID2,
		Content: c.Content,
	}
	return createdChat, nil
}

func (r *ChatRepository) FindChatByUserID(ctx context.Context, userID string) ([]entity.Chat, error) {
	user, err := r.queries.GetUserByEmail(ctx, userID)
	if err != nil {
		return nil, errors.HandleDBError(err)
	}
	id := user.UserID
	chat, err := r.queries.GetChatByUserID(ctx, id)
	if err != nil {
		return nil, errors.HandleDBError(err)
	}
	var e []entity.Chat
	for _, c := range chat {
		e = append(e, entity.Chat{
			ChatID:  c.ChatID,
			UserID1: c.UserID1,
			UserID2: c.UserID2,
			Content: c.Content,
		})
	}
	return e, nil
}
