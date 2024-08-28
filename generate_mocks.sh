mockgen -source=./tools/httpx/client.go -destination=./tools/httpx/client_mocks.go -package=httpx
mockgen -source=./tools/redisx/client.go -destination=./tools/redisx/client_mocks.go -package=redisx
mockgen -source=./log/logrus_logger.go -destination=./log/logrus_logger_mocks.go -package=log