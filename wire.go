// wire.go
//go:build wireinject

package main

import (
	"fiap-tech-challenge-pedidos/internal/adapters/http"
	"fiap-tech-challenge-pedidos/internal/adapters/http/handlers"
	"fiap-tech-challenge-pedidos/internal/adapters/http/middlewares/auth"
	"fiap-tech-challenge-pedidos/internal/adapters/repository"
	"fiap-tech-challenge-pedidos/internal/core/usecase"
	"fiap-tech-challenge-pedidos/internal/core/usecase/mapper"
	"fiap-tech-challenge-pedidos/internal/util"

	"github.com/google/wire"
)

func InitializeWebServer() (*http.Server, error) {
	wire.Build(repository.NewMySQLConnector,
		util.NewCustomValidator,
		repository.NewPedidoRepo,
		auth.NewJwtToken,
		mapper.NewPedidoMapper,
		usecase.NewListaPedidoPorStatus,
		usecase.NewListaTodosPedidos,
		usecase.NewCadastraPedido,
		usecase.NewAtualizaStatusPedidoUC,
		usecase.NewPegaDetalhePedido,
		handlers.NewHealthCheck,
		handlers.NewPedido,
		http.NewAPIServer)
	return &http.Server{}, nil
}
