package usecases

import (
	"context"
	"errors"
	"fiap-tech-challenge-pedidos/internal/core/domain"
	mock_mapper "fiap-tech-challenge-pedidos/test/mock/mapper"
	mock_repo "fiap-tech-challenge-pedidos/test/mock/repository"
	"time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

var _ = Describe("lista pedido use case testes", func() {
	var (
		ctx                   = context.Background()
		repo                  *mock_repo.MockPedidoRepo
		mapper                *mock_mapper.MockPedido
		listaTodosPedido      ListarTodosPedidos
		listaPedidosPorStatus ListarPedidoPorStatus
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockPedidoRepo(mockCtrl)
		mapper = mock_mapper.NewMockPedido(mockCtrl)
		listaTodosPedido = NewListaTodosPedidos(repo, mapper)
		listaPedidosPorStatus = NewListaPedidoPorStatus(repo, mapper)
	})

	Context("lista todos os pedidos", func() {
		objID := primitive.NewObjectID()
		pedidoDTO := &domain.PedidoDTO{
			ID:        objID,
			ClienteId: 1,
			Status:    "preparando",
			Produtos:  []string{"1", "2"},
		}
		It("pesquisa com sucesso", func() {
			repo.EXPECT().PesquisaTodos(ctx).Return([]*domain.PedidoDTO{pedidoDTO}, nil)
			mapper.EXPECT().MapDTOToModels(gomock.Any()).Return([]*domain.Pedido{
				&domain.Pedido{
					Id:         primitive.ObjectID{},
					ClienteId:  0,
					Status:     "",
					Produtos:   nil,
					Observacao: "",
					CreatedAt:  time.Time{},
					UpdatedAt:  time.Time{},
				},
			})
			pedido, err := listaTodosPedido.ListaTodos(ctx)

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(pedido).ToNot(gomega.BeNil())
		})
		It("pesquisa falha então retorna erro", func() {
			repo.EXPECT().PesquisaTodos(ctx).Return(nil, errors.New("mock error"))
			pedido, err := listaTodosPedido.ListaTodos(ctx)

			gomega.Expect(pedido).To(gomega.BeNil())
			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("mock error"))
		})
	})
	Context("lista pedido por status", func() {
		objID := primitive.NewObjectID()
		pedidoDTO := &domain.PedidoDTO{
			ID:        objID,
			ClienteId: 1,
			Status:    "preparando",
			Produtos:  []string{"1", "2"},
		}
		It("pesquisa com sucesso", func() {
			repo.EXPECT().PesquisaPorStatus(ctx, []string{domain.StatusEmPreparacao}).Return([]*domain.PedidoDTO{pedidoDTO}, nil)
			mapper.EXPECT().MapDTOToModels(gomock.Any()).Return([]*domain.Pedido{
				&domain.Pedido{
					Id:         primitive.ObjectID{},
					ClienteId:  0,
					Status:     "",
					Produtos:   nil,
					Observacao: "",
					CreatedAt:  time.Time{},
					UpdatedAt:  time.Time{},
				},
			})
			pedido, err := listaPedidosPorStatus.ListaPorStatus(ctx, []string{domain.StatusEmPreparacao})

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(pedido).ToNot(gomega.BeNil())
		})
		It("pesquisa falha então retorna erro", func() {
			repo.EXPECT().PesquisaPorStatus(ctx, []string{domain.StatusEmPreparacao}).Return(nil, errors.New("mock error"))
			pedido, err := listaPedidosPorStatus.ListaPorStatus(ctx, []string{domain.StatusEmPreparacao})

			gomega.Expect(pedido).To(gomega.BeNil())
			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("mock error"))
		})
	})
})
