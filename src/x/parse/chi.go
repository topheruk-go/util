package parse

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetID(ctx context.Context) uuid.UUID { return uuid.MustParse(chi.URLParamFromCtx(ctx, "id")) }
