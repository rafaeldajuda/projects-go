package domain

import "context"

type Provider interface {
	FetchData(ctx context.Context) (MyData, error)
}
