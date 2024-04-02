
mock:
	mockgen -source=internal/core/usecases/cria_pedido.go -package=mock_usecases -destination=test/mock/usecases/cria_pedido.go
	mockgen -source=internal/core/usecases/mapper/pedido.go -package=mock_mapper -destination=test/mock/mapper/pedido.go
	mockgen -source=internal/adapters/repository/pedido.go -package=mock_repo -destination=test/mock/repository/pedido.go