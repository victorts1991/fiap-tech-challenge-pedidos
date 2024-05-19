package client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/joomcode/errorx"
	"log"
	"net/http"
	"os"
)

type Cliente interface {
	PesquisaPorID(ctx context.Context, id string) error
}

type clienteClient struct {
	httpClient *http.Client
	url        string
}

func NewCliente() Cliente {
	url := os.Getenv("API_URL")
	if url == "" {
		log.Fatal("API_URL environment variable not set")
	}
	return &clienteClient{
		httpClient: http.DefaultClient,
		url:        url,
	}
}
func (c *clienteClient) PesquisaPorID(ctx context.Context, id string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/internal/clientes/%s", c.url, id), bytes.NewBuffer([]byte{}))
	if err != nil {
		return errorx.InternalError.New(fmt.Sprintf("não foi possível inicializar cliente client %s", err.Error()))
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errorx.InternalError.New(fmt.Sprintf("cliente server retornou error %s", err.Error()))
	}

	if resp == nil {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return errorx.InternalError.New(fmt.Sprintf("cliente server retornou status %s", resp.Status))
	}
	defer resp.Body.Close()

	return nil
}
