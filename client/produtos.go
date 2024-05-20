package client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/joomcode/errorx"
	"log"
	"net/http"
	"os"
	"strings"
)

type Produto interface {
	PesquisaPorIDS(ctx context.Context, ids []string) error
}

type produtoCliente struct {
	httpClient *http.Client
	url        string
}

func NewProduto() Produto {
	url := os.Getenv("API_URL")
	if url == "" {
		log.Fatal("API_URL environment variable not set")
	}
	return &produtoCliente{
		httpClient: http.DefaultClient,
		url:        url,
	}
}
func (c *produtoCliente) PesquisaPorIDS(ctx context.Context, ids []string) error {
	paramIDS := strings.Join(ids, ",")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/internal/produto/%s", c.url, paramIDS), bytes.NewBuffer([]byte{}))
	if err != nil {
		return errorx.InternalError.New(fmt.Sprintf("não foi possível inicializar produto client %s", err.Error()))
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errorx.InternalError.New(fmt.Sprintf("produto server retornou error %s", err.Error()))
	}

	if resp == nil {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return errorx.InternalError.New(fmt.Sprintf("produto server retornou status %s", resp.Status))
	}
	defer resp.Body.Close()

	return nil
}
