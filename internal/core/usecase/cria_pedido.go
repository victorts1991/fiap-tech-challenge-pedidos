package usecase

import (
	"context"
	"fiap-tech-challenge-pedidos/internal/adapters/repository"
	"fiap-tech-challenge-pedidos/internal/core/domain"
	"fiap-tech-challenge-pedidos/internal/core/usecase/mapper"
)

type CadastrarPedido interface {
	Cadastra(ctx context.Context, pedido *domain.PedidoRequest) (*domain.PedidoResponse, error)
}

type cadastraPedido struct {
	repo         repository.PedidoRepo
	mapperPedido mapper.Pedido
}

func (uc cadastraPedido) Cadastra(ctx context.Context, req *domain.PedidoRequest) (*domain.PedidoResponse, error) {
	dto, err := uc.repo.Insere(ctx, uc.mapperPedido.MapReqToDTO(req))
	if err != nil {
		return nil, err
	}

	return uc.mapperPedido.MapDTOToResponse(dto), err
}

func NewCadastraPedido(repo repository.PedidoRepo, mapperPedido mapper.Pedido) CadastrarPedido {
	return &cadastraPedido{
		repo:         repo,
		mapperPedido: mapperPedido,
	}
}
