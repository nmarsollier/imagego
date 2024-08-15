mockgen -source=./tools/http_client/client.go -destination=./tools/http_client/client_mocks.go -package=http_client
mockgen -source=./tools/redis_client/client.go -destination=./tools/redis_client/client_mocks.go -package=redis_client
