package mapper

import (
	"fiap-tech-challenge-pedidos/internal/core/domain"
)

type Pedido interface {
	MapReqToDTO(req *domain.PedidoRequest) *domain.PedidoDTO
	MapDTOToModels(req []*domain.PedidoDTO) []*domain.Pedido
	MapDTOToResponse(dto *domain.PedidoDTO) *domain.PedidoResponse
}

type pedido struct {
}

func (p pedido) MapDTOToResponse(dto *domain.PedidoDTO) *domain.PedidoResponse {
	return &domain.PedidoResponse{
		Pedido: &domain.Pedido{
			Id:         dto.ID,
			ClienteId:  dto.ClienteId,
			Produtos:   dto.Produtos,
			Status:     dto.Status,
			Observacao: dto.Observacao,
			CreatedAt:  dto.CreatedAt,
			UpdatedAt:  dto.UpdatedAt,
		},
	}
}

func (p pedido) MapReqToDTO(req *domain.PedidoRequest) *domain.PedidoDTO {
	return &domain.PedidoDTO{
		ClienteId:  req.ClienteId,
		Observacao: req.Observacao,
		Produtos:   req.ProdutoIds,
	}
}

func (p pedido) MapDTOToModels(req []*domain.PedidoDTO) []*domain.Pedido {
	pedidos := make([]*domain.Pedido, len(req))
	for i, dto := range req {
		pedidos[i] = &domain.Pedido{
			ClienteId:  dto.ClienteId,
			Status:     dto.Status,
			Id:         dto.ID,
			Produtos:   dto.Produtos,
			Observacao: dto.Observacao,
			CreatedAt:  dto.CreatedAt,
			UpdatedAt:  dto.UpdatedAt,
		}
	}

	return pedidos
}

func NewPedidoMapper() Pedido {
	return &pedido{}
}
