rm -rf ./test/mockgen

set -e

mockgen -source=./internal/image/repository.go -destination=./test/mockgen/image_repository_mocks.go -package=mockgen
mockgen -source=./internal/image/service.go -destination=./test/mockgen/image_service_mocks.go -package=mockgen
