package usecases

import (
	"context"
	"fiap-tech-challenge-pedidos/client"
	"fiap-tech-challenge-pedidos/internal/adapters/repository"
	"fiap-tech-challenge-pedidos/internal/core/domain"
	"fiap-tech-challenge-pedidos/internal/core/usecases/mapper"
	"fmt"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/errors"
	"strconv"
)

type CadastrarPedido interface {
	Cadastra(ctx context.Context, pedido *domain.PedidoRequest) (*domain.PedidoResponse, error)
}

type cadastraPedido struct {
	repo          repository.PedidoRepo
	mapperPedido  mapper.Pedido
	produtoClient client.Produto
	clienteClient client.Cliente
}

func (uc cadastraPedido) Cadastra(ctx context.Context, req *domain.PedidoRequest) (*domain.PedidoResponse, error) {
	err := uc.clienteClient.PesquisaPorID(ctx, strconv.FormatInt(req.ClienteId, 10))
	if err != nil {
		return nil, errors.BadRequest.New(fmt.Sprintf("cliente id inválido %s", err.Error()))
	}
	err = uc.produtoClient.PesquisaPorIDS(ctx, req.ProdutoIds)
	if err != nil {
		return nil, errors.BadRequest.New(fmt.Sprintf("produto id inválido %s", err.Error()))
	}

	dto, err := uc.repo.Insere(ctx, uc.mapperPedido.MapReqToDTO(req))
	if err != nil {
		return nil, err
	}

	return uc.mapperPedido.MapDTOToResponse(dto), err
}

func NewCadastraPedido(repo repository.PedidoRepo, mapperPedido mapper.Pedido, produtoClient client.Produto, clienteClient client.Cliente) CadastrarPedido {
	return &cadastraPedido{
		repo:          repo,
		mapperPedido:  mapperPedido,
		produtoClient: produtoClient,
		clienteClient: clienteClient,
	}
}
