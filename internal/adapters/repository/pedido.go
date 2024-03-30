package repository

import (
	"context"
	"fiap-tech-challenge-pedidos/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const tableNamePedido string = "pedidos"

type pedido struct {
	session *mongo.Collection
}

type PedidoRepo interface {
	Insere(ctx context.Context, pedido *domain.PedidoDTO) (*domain.PedidoDTO, error)
	PesquisaPorStatus(ctx context.Context, statuses []string) ([]*domain.PedidoDTO, error)
	PesquisaPorID(ctx context.Context, id int64) (*domain.PedidoDTO, error)
	AtualizaStatus(ctx context.Context, status string, id int64) error
	PesquisaTodos(ctx context.Context) ([]*domain.PedidoDTO, error)
}

func NewPedidoRepo(connector DBConnector) PedidoRepo {
	session := connector.GetDB().Collection(tableNamePedido)
	return &pedido{
		session: session,
	}
}

func (p *pedido) Insere(ctx context.Context, pedido *domain.PedidoDTO) (*domain.PedidoDTO, error) {
	pedido.ID = primitive.NewObjectID()
	now := time.Now()
	pedido.CreatedAt = now
	pedido.UpdatedAt = now
	_, err := p.session.InsertOne(ctx, &pedido)
	if err != nil {
		return nil, err
	}

	return pedido, nil
}

func (p *pedido) PesquisaPorStatus(ctx context.Context, statuses []string) ([]*domain.PedidoDTO, error) {
	pedidos := make([]*domain.PedidoDTO, 0)

	return pedidos, nil
}

func (p *pedido) AtualizaStatus(ctx context.Context, status string, id int64) error {
	return nil
}

func (p *pedido) PesquisaPorID(ctx context.Context, id int64) (*domain.PedidoDTO, error) {
	dto := &domain.PedidoDTO{}

	return dto, nil
}

func (p *pedido) PesquisaTodos(ctx context.Context) ([]*domain.PedidoDTO, error) {
	pedidos := make([]*domain.PedidoDTO, 0)
	filter := bson.M{}
	cur, err := p.session.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &pedidos)
	if err != nil {
		return nil, err
	}

	return pedidos, nil
}
