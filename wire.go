// wire.go
//go:build wireinject

package main

import (
	"fiap-tech-challenge-pedidos/internal/adapters/http"
	"fiap-tech-challenge-pedidos/internal/adapters/http/handlers"
	"fiap-tech-challenge-pedidos/internal/adapters/repository"
	"fiap-tech-challenge-pedidos/internal/core/usecases"
	"fiap-tech-challenge-pedidos/internal/core/usecases/mapper"
	_mongo "github.com/rhuandantas/fiap-tech-challenge-commons/pkg/db/mongo"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/middlewares/auth"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/util"

	"github.com/google/wire"
)

func InitializeWebServer() (*http.Server, error) {
	wire.Build(_mongo.NewMongoConnector,
		util.NewCustomValidator,
		repository.NewPedidoRepo,
		auth.NewJwtToken,
		mapper.NewPedidoMapper,
		usecases.NewListaPedidoPorStatus,
		usecases.NewListaTodosPedidos,
		usecases.NewCadastraPedido,
		usecases.NewAtualizaStatusPedidoUC,
		usecases.NewPegaDetalhePedido,
		handlers.NewHealthCheck,
		handlers.NewPedido,
		http.NewAPIServer)
	return &http.Server{}, nil
}
