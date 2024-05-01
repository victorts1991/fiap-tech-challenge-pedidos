package mapper

import (
	"fiap-tech-challenge-pedidos/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func Test(t *testing.T) {
	t.Parallel()
	gomega.RegisterFailHandler(Fail)
	RunSpecs(t, "mapper suite test")
}

var _ = Describe("mapper testes", func() {
	var (
		mapper Pedido
		objId  = primitive.NewObjectID()
		now    = time.Now()
	)
	dto := &domain.PedidoDTO{
		ID:         objId,
		ClienteId:  1,
		Status:     "em_preparacao",
		Produtos:   []string{},
		Observacao: "",
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	dtos := []*domain.PedidoDTO{dto}
	model := &domain.Pedido{
		Id:         objId,
		ClienteId:  1,
		Status:     "em_preparacao",
		Produtos:   []string{},
		Observacao: "",
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	models := []*domain.Pedido{model}
	req := &domain.PedidoRequest{
		ClienteId:  1,
		ProdutoIds: []string{},
		Observacao: "",
	}

	BeforeEach(func() {
		mapper = NewPedidoMapper()
	})

	It("MapDTOToModels", func() {
		result := mapper.MapDTOToModels(dtos)

		gomega.Expect(models).To(gomega.Equal(result))
	})
	It("MapReqToDTO", func() {
		result := mapper.MapReqToDTO(req)

		gomega.Expect(dto.ClienteId).To(gomega.Equal(result.ClienteId))
		gomega.Expect(dto.Produtos).To(gomega.Equal(result.Produtos))
	})
	It("MapDTOToResponse", func() {
		result := mapper.MapDTOToResponse(dto)

		gomega.Expect(model.ClienteId).To(gomega.Equal(result.ClienteId))
		gomega.Expect(model.Produtos).To(gomega.Equal(result.Produtos))
	})

})
