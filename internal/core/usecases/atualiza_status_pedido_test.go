package usecases

import (
	"context"
	"errors"
	"fiap-tech-challenge-pedidos/internal/core/domain"
	mock_repo "fiap-tech-challenge-pedidos/test/mock/repository"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

var _ = Describe("atualiza pedido use case testes", func() {
	var (
		ctx            = context.Background()
		repo           *mock_repo.MockPedidoRepo
		atualizaPedido AtualizaStatusPedidoUC
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockPedidoRepo(mockCtrl)
		atualizaPedido = NewAtualizaStatusPedidoUC(repo)
	})

	Context("atualiza pedido", func() {
		objID := primitive.NewObjectID()
		pedidoDTO := &domain.PedidoDTO{
			ID:        objID,
			ClienteId: 1,
			Status:    "preparando",
			Produtos:  []string{"1", "2"},
		}
		It("atualiza com sucesso", func() {
			repo.EXPECT().PesquisaPorID(ctx, gomock.Any()).Return(pedidoDTO, nil)
			repo.EXPECT().AtualizaStatus(ctx, gomock.Any(), gomock.Any()).Return(nil)
			err := atualizaPedido.Atualiza(ctx, "preparando", objID)

			gomega.Expect(err).To(gomega.BeNil())
		})
		It("update falha então retorna erro", func() {
			repo.EXPECT().PesquisaPorID(ctx, gomock.Any()).Return(pedidoDTO, nil)
			repo.EXPECT().AtualizaStatus(ctx, gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			err := atualizaPedido.Atualiza(ctx, "praparando", objID)

			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("mock error"))
		})
		It("pega por id falha então retorna erro", func() {
			repo.EXPECT().PesquisaPorID(ctx, gomock.Any()).Return(nil, errors.New("mock error"))
			err := atualizaPedido.Atualiza(ctx, "praparando", objID)

			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("mock error"))
		})
		It("falha ao tentar atualizar com status aguardando pagamento", func() {
			pedidoDTO.Status = domain.StatusAguardandoPagamento
			repo.EXPECT().PesquisaPorID(ctx, gomock.Any()).Return(pedidoDTO, nil)
			err := atualizaPedido.Atualiza(ctx, "finalizado", objID)

			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("common.bad_request: não é possível atualizar status de aguardando_pagamento para finalizado"))
		})
		It("falha ao tentar atualizar com status aguardando pagamento", func() {
			pedidoDTO.Status = domain.StatusPagamentoRecusado
			repo.EXPECT().PesquisaPorID(ctx, gomock.Any()).Return(pedidoDTO, nil)
			err := atualizaPedido.Atualiza(ctx, "finalizado", objID)

			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("common.bad_request: não é possível atualizar status de pagamento_recusado para finalizado"))
		})
	})
})
