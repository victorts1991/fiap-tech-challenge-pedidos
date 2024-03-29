package usecase

import (
	"context"
	"fiap-tech-challenge-pedidos/internal/adapters/repository"
	"fiap-tech-challenge-pedidos/internal/core/domain"
	"fiap-tech-challenge-pedidos/internal/core/usecase/mapper"
)

type PegarDetalhePedido interface {
	Pesquisa(ctx context.Context, id int64) (*domain.PedidoResponse, error)
}

type pegaDetalhePedido struct {
	repo         repository.PedidoRepo
	mapperPedido mapper.Pedido
}

func (uc *pegaDetalhePedido) Pesquisa(ctx context.Context, id int64) (*domain.PedidoResponse, error) {
	dto, err := uc.repo.PesquisaPorID(ctx, id)
	if err != nil {
		return nil, err
	}

	return uc.mapperPedido.MapDTOToResponse(dto), nil
}

func NewPegaDetalhePedido(repo repository.PedidoRepo,
	mapperPedido mapper.Pedido,
) PegarDetalhePedido {
	return &pegaDetalhePedido{
		repo:         repo,
		mapperPedido: mapperPedido,
	}
}
