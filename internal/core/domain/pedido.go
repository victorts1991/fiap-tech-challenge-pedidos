package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	StatusRecebido            string = "recebido"
	StatusEmPreparacao        string = "em_preparacao"
	StatusPronto              string = "pronto"
	StatusFinalizado          string = "finalizado"
	StatusAguardandoPagamento string = "aguardando_pagamento"
	StatusPagamentoAprovado   string = "pagamento_aprovado"
	StatusPagamentoRecusado   string = "pagamento_recusado"
)

type PedidoRequest struct {
	ClienteId  int64    `json:"cliente_id" validate:"required"`
	ProdutoIds []string `json:"produtos" validate:"required"`
	Observacao string   `json:"observacao"`
}

type PedidoDTO struct {
	ID         primitive.ObjectID `bson:"_id"`
	ClienteId  int64              `bson:"cliente_id"`
	Produtos   []string           `bson:"produtos"`
	Status     string             `bson:"status"`
	Observacao string             `bson:"observacao"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

type PedidosDTO []*PedidoDTO

func (a PedidosDTO) Len() int      { return len(a) }
func (a PedidosDTO) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PedidosDTO) Less(i, j int) bool {
	if a[i].Status != a[j].Status {
		return a[i].Status == StatusPronto || (a[i].Status == StatusEmPreparacao && a[j].Status != StatusPronto)
	}

	return a[i].CreatedAt.Before(a[j].CreatedAt)

}

func (dto *PedidoDTO) TableName() string {
	return "pedidos"
}

type PedidoResponse struct {
	*Pedido
}

type Pedido struct {
	Id         primitive.ObjectID `json:"id"`
	ClienteId  int64              `json:"cliente_id"`
	Status     string             `json:"status"`
	Produtos   []string           `json:"produtos"`
	Observacao string             `json:"observacao"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

type Fila struct {
	Id         int64
	PedidoId   int64  `xorm:"index unique"`
	Status     string `xorm:"status"`
	Observacao string
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

type StatusRequest struct {
	Status string `json:"status" validate:"required"`
}
