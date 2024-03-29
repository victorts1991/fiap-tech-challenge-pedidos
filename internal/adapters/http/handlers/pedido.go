package handlers

import (
	serverErr "fiap-tech-challenge-pedidos/internal/adapters/http/error"
	"fiap-tech-challenge-pedidos/internal/adapters/http/middlewares/auth"
	"fiap-tech-challenge-pedidos/internal/core/commons"
	"fiap-tech-challenge-pedidos/internal/core/domain"
	"fiap-tech-challenge-pedidos/internal/core/usecase"
	"fiap-tech-challenge-pedidos/internal/util"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/joomcode/errorx"
)

type Pedido struct {
	validator           util.Validator
	listaPorStatusUC    usecase.ListarPedidoPorStatus
	listaTodosUC        usecase.ListarTodosPedidos
	cadastraPedidoUC    usecase.CadastrarPedido
	atualizaStatusUC    usecase.AtualizaStatusPedidoUC
	pegaDetalhePedidoUC usecase.PegarDetalhePedido
	tokenJwt            auth.Token
}

func NewPedido(validator util.Validator,
	listaPorStatusUC usecase.ListarPedidoPorStatus,
	listaTodosUC usecase.ListarTodosPedidos,
	cadastraPedidoUC usecase.CadastrarPedido,
	atualizaStatusUC usecase.AtualizaStatusPedidoUC,
	pegaDetalhePedidoUC usecase.PegarDetalhePedido,
	tokenJwt auth.Token,
) *Pedido {
	return &Pedido{
		validator:           validator,
		listaPorStatusUC:    listaPorStatusUC,
		listaTodosUC:        listaTodosUC,
		cadastraPedidoUC:    cadastraPedidoUC,
		atualizaStatusUC:    atualizaStatusUC,
		pegaDetalhePedidoUC: pegaDetalhePedidoUC,
		tokenJwt:            tokenJwt,
	}
}

func (h *Pedido) RegistraRotasPedido(server *echo.Echo) {
	server.POST("/pedido", h.cadastra, h.tokenJwt.VerifyToken)
	server.GET("/pedidos/:statuses", h.listaPorStatus)
	server.GET("/pedidos", h.listaTodos)
	server.GET("/pedido/detail/:id", h.listaDetail, h.tokenJwt.VerifyToken)
	server.PATCH("/pedido/:id", h.atualizaStatus)
}

// cadastra godoc
// @Summary cadastra um novo pedido
// @Tags Pedido
// @Param			pedido	body		domain.PedidoRequest	true	"cria pedido"
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce json
// @Router /pedido [post]
func (h *Pedido) cadastra(ctx echo.Context) error {
	var (
		req domain.PedidoRequest
		err error
	)

	if err = ctx.Bind(&req); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	if err = h.validatePedidoBody(&req); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	response, err := h.cadastraPedidoUC.Cadastra(ctx.Request().Context(), &req)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"id": response.Id, "status": response.Status})
}

// listaPorStatus godoc
// @Summary lista pedido por status
// @Tags Pedido
// @Produce json
// @Param        statuses   path      string  true  "status dos pedidos a ser pesquisado:(recebido, em_preparacao, pronto, finalizado)"
// @Success 200 {array} domain.Pedido
// @Router /pedidos/{statuses} [get]
func (h *Pedido) listaPorStatus(ctx echo.Context) error {
	statuses := ctx.Param("statuses")
	filter := strings.Split(statuses, ",")

	pedidos, err := h.listaPorStatusUC.ListaPorStatus(ctx.Request().Context(), filter)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	return ctx.JSON(http.StatusOK, pedidos)
}

func (h *Pedido) validatePedidoBody(p *domain.PedidoRequest) error {
	if err := h.validator.ValidateStruct(p); err != nil {
		return err
	}

	// check out something more?

	return nil
}

// atualizaStatus godoc
// @Summary atualiza o status do pedido
// @Tags Pedido
// @Accept json
// @Param        id   path      integer  true  "id do pedido"
// @Param        id   body      domain.StatusRequest  true  "status permitido: recebido, em_preparacao, pronto, finalizado"
// @Produce json
// @Router /pedido/{id} [patch]
func (h *Pedido) atualizaStatus(ctx echo.Context) error {
	var (
		status struct {
			Status string `json:"status"`
		}
		//pedidoID int
		err error
	)

	if err = ctx.Bind(&status); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	id := ctx.Param("id")
	if _, err = strconv.Atoi(id); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(fmt.Sprintf("%s não é um id válido", id)))
	}

	//err = h.atualizaStatusUC.Atualiza(ctx.Request().Context(), status.Status, int64(pedidoID))
	//if err != nil {
	//	return serverErr.HandleError(ctx, errorx.Cast(err))
	//}

	return ctx.JSON(http.StatusOK, status)
}

// listaDetail godoc
// @Summary lista detalhes do pedido
// @Tags Pedido
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      integer  true  "id do pedido a ser lista"
// @Success 200 {object} domain.Pedido
// @Router /pedido/detail/{id} [get]
func (h *Pedido) listaDetail(ctx echo.Context) error {
	var (
		pedidoID int
		err      error
	)

	id := ctx.Param("id")
	if pedidoID, err = strconv.Atoi(id); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(fmt.Sprintf("%s não é um id válido", id)))
	}

	pedido, err := h.pegaDetalhePedidoUC.Pesquisa(ctx.Request().Context(), int64(pedidoID))
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	return ctx.JSON(http.StatusOK, pedido)
}

// listaTodos godoc
// @Summary lista todos os pedidos
// @Tags Pedido
// @Produce json
// @Success 200 {array} domain.Pedido
// @Router /pedidos [get]
func (h *Pedido) listaTodos(ctx echo.Context) error {

	pedidos, err := h.listaTodosUC.ListaTodos(ctx.Request().Context())
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	return ctx.JSON(http.StatusOK, pedidos)
}
