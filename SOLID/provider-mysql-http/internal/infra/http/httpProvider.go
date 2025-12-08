package http

import (
	"context"
	"fmt"
	"net/http"
	"provider-mysql-http/internal/domain"
)

type HttpProvider struct {
	Method string
	URL    string
}

func NewHttpProvider(method, url string) *HttpProvider {
	return &HttpProvider{
		Method: method,
		URL:    url,
	}
}

func (p *HttpProvider) FetchData(ctx context.Context) (domain.MyData, error) {
	myData := domain.MyData{}

	req, err := http.NewRequest(p.Method, p.URL, nil)
	if err != nil {
		err = fmt.Errorf("Erro ao criar requisição:", err)
		return myData, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("Erro ao enviar requisição: %s", err)
		return myData, err
	}
	defer resp.Body.Close()

	body := resp.Body

	return domain.MyData{}, nil
}
