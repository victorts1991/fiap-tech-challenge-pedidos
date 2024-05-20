package usecases

import (
	"context"
	"errors"
	"fiap-tech-challenge-pedidos/internal/core/domain"
	mock_client "fiap-tech-challenge-pedidos/test/mock/client"
	mock_mapper "fiap-tech-challenge-pedidos/test/mock/mapper"
	mock_repo "fiap-tech-challenge-pedidos/test/mock/repository"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

var _ = ginkgo.Describe("cria pedido use case testes", func() {
	var (
		ctx            = context.Background()
		repo           *mock_repo.MockPedidoRepo
		mapper         *mock_mapper.MockPedido
		produtoCliente *mock_client.MockProduto
		clienteCliente *mock_client.MockCliente
		criaPedido     CadastrarPedido
	)

	ginkgo.BeforeEach(func() {
		mockCtrl := gomock.NewController(ginkgo.GinkgoT())
		repo = mock_repo.NewMockPedidoRepo(mockCtrl)
		mapper = mock_mapper.NewMockPedido(mockCtrl)
		produtoCliente = mock_client.NewMockProduto(mockCtrl)
		clienteCliente = mock_client.NewMockCliente(mockCtrl)
		criaPedido = NewCadastraPedido(repo, mapper, produtoCliente, clienteCliente)
	})

	ginkgo.Context("criar pedido", func() {
		objID := primitive.NewObjectID()
		pedidoDTO := &domain.PedidoDTO{
			ID:        objID,
			ClienteId: 1,
			Status:    "",
			Produtos:  []string{"1", "2"},
		}
		ginkgo.It("cria com sucesso", func() {
			produtoCliente.EXPECT().PesquisaPorIDS(ctx, []string{"1", "2"}).Return(nil)
			clienteCliente.EXPECT().PesquisaPorID(ctx, "1").Return(nil)
			repo.EXPECT().Insere(ctx, gomock.Any()).Return(pedidoDTO, nil)
			mapper.EXPECT().MapReqToDTO(gomock.Any()).Return(pedidoDTO)
			mapper.EXPECT().MapDTOToResponse(gomock.Any()).Return(&domain.PedidoResponse{
				Pedido: &domain.Pedido{
					Id:        objID,
					ClienteId: 1,
					Status:    "",
					Produtos:  []string{"1", "2"},
				},
			})

			resp, err := criaPedido.Cadastra(ctx, &domain.PedidoRequest{
				ClienteId:  1,
				ProdutoIds: []string{"1", "2"},
			})

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(resp.ClienteId).To(gomega.Equal(pedidoDTO.ClienteId))
			gomega.Expect(resp.Id.Hex()).To(gomega.Equal(pedidoDTO.ID.Hex()))
		})
		ginkgo.It("insere falha então retorna erro", func() {
			produtoCliente.EXPECT().PesquisaPorIDS(ctx, []string{"1", "2"}).Return(nil)
			clienteCliente.EXPECT().PesquisaPorID(ctx, "1").Return(nil)
			repo.EXPECT().Insere(ctx, gomock.Any()).Return(nil, errors.New("mock error"))
			mapper.EXPECT().MapReqToDTO(gomock.Any()).Return(pedidoDTO)
			resp, err := criaPedido.Cadastra(ctx, &domain.PedidoRequest{
				ClienteId:  1,
				ProdutoIds: []string{"1", "2"},
			})

			gomega.Expect(resp).To(gomega.BeNil())
			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("mock error"))
		})
	})
})
