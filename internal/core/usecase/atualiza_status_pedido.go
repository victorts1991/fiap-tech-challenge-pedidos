package usecase

import (
	"fiap-tech-challenge-pedidos/internal/adapters/repository"
	"fiap-tech-challenge-pedidos/internal/core/domain"
)

type AtualizaStatusPedidoUC interface {
}

type atualizaStatusPedido struct {
	repo repository.PedidoRepo
}

func couldNotUpdateStatus(status string) bool {
	return status == domain.StatusAguardandoPagamento ||
		status == domain.StatusPagamentoRecusado
}

func NewAtualizaStatusPedidoUC(repo repository.PedidoRepo) AtualizaStatusPedidoUC {
	return &atualizaStatusPedido{
		repo: repo,
	}
}
