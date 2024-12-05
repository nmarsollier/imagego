mockgen -source=./tools/httpx/client.go -destination=./tools/httpx/client_mocks.go -package=httpx
mockgen -source=./image/image_dao/dao.go -destination=./tools/redisx/client_mocks.go -package=redisx
mockgen -source=./tools/log/logrus_logger.go -destination=./tools/log/logrus_logger_mocks.go -package=log