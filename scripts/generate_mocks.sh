mockgen -source=./tools/httpx/client.go -destination=./tools/httpx/client_mocks.go -package=httpx
mockgen -source=./db/dao.go -destination=./db/client_mocks.go -package=db
mockgen -source=./tools/log/logrus_logger.go -destination=./tools/log/logrus_logger_mocks.go -package=log