package usecase

import (
	"context"
	"fiap-tech-challenge-pedidos/internal/adapters/repository"
	"fiap-tech-challenge-pedidos/internal/core/commons"
	"fiap-tech-challenge-pedidos/internal/core/domain"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AtualizaStatusPedidoUC interface {
	Atualiza(ctx context.Context, status string, id primitive.ObjectID) error
}

type atualizaStatusPedido struct {
	repo repository.PedidoRepo
}

func (p atualizaStatusPedido) Atualiza(ctx context.Context, status string, id primitive.ObjectID) error {
	pedidoDTO, err := p.repo.PesquisaPorID(ctx, id)
	if err != nil {
		return err
	}

	if couldNotUpdateStatus(pedidoDTO.Status) {
		return commons.BadRequest.New(fmt.Sprintf("não é possível atualizar status de %s para %s", pedidoDTO.Status, status))
	}

	err = p.repo.AtualizaStatus(ctx, status, id)
	if err != nil {
		return err
	}

	return nil
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
