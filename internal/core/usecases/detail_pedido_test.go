package usecases

import (
	"context"
	"errors"
	"fiap-tech-challenge-pedidos/internal/core/domain"
	mock_mapper "fiap-tech-challenge-pedidos/test/mock/mapper"
	mock_repo "fiap-tech-challenge-pedidos/test/mock/repository"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

var _ = Describe("detalha pedido use case testes", func() {
	var (
		ctx           = context.Background()
		repo          *mock_repo.MockPedidoRepo
		mapper        *mock_mapper.MockPedido
		detalhePedido PegarDetalhePedido
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockPedidoRepo(mockCtrl)
		mapper = mock_mapper.NewMockPedido(mockCtrl)
		detalhePedido = NewPegaDetalhePedido(repo, mapper)
	})

	Context("atualiza pedido", func() {
		objID := primitive.NewObjectID()
		pedidoDTO := &domain.PedidoDTO{
			ID:        objID,
			ClienteId: 1,
			Status:    "preparando",
			Produtos:  []string{"1", "2"},
		}
		It("pesquisa com sucesso", func() {
			repo.EXPECT().PesquisaPorID(ctx, gomock.Any()).Return(pedidoDTO, nil)
			mapper.EXPECT().MapDTOToResponse(gomock.Any()).Return(&domain.PedidoResponse{
				Pedido: &domain.Pedido{
					Id:        objID,
					ClienteId: 1,
					Status:    "",
					Produtos:  []string{"1", "2"},
				},
			})
			pedido, err := detalhePedido.Pesquisa(ctx, objID)

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(pedido).ToNot(gomega.BeNil())
		})
		It("update falha então retorna erro", func() {
			repo.EXPECT().PesquisaPorID(ctx, gomock.Any()).Return(nil, errors.New("mock error"))
			pedido, err := detalhePedido.Pesquisa(ctx, objID)

			gomega.Expect(pedido).To(gomega.BeNil())
			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("mock error"))
		})
	})
})
