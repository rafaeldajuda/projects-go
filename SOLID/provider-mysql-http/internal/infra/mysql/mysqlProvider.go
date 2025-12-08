package mysql

import (
	"context"
	"provider-mysql-http/internal/domain"
)

type MysqlProvider struct {
}

func NewMysqlProvider() *MysqlProvider {
	return &MysqlProvider{}
}

func (p *MysqlProvider) FetchData(ctx context.Context) (domain.MyData, error) {
	return domain.MyData{}, nil
}
