rm -rf ./test/mockgen

mockgen -source=./internal/image/repository.go -destination=./test/mockgen/image_repository_mocks.go -package=mockgen
mockgen -source=./internal/image/service.go -destination=./test/mockgen/image_service_mocks.go -package=mockgen
mockgen -source=./internal/security/repository.go -destination=./test/mockgen/security_repository_mocks.go -package=mockgen
mockgen -source=./internal/security/service.go -destination=./test/mockgen/security_service_mocks.go -package=mockgen
mockgen -source=./internal/rabbit/consume_logout.go -destination=./test/mockgen/consume_logout_mocks.go -package=mockgen
