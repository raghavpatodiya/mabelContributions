# check_install:
# 	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

# swagger: check_install
# 	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

# swagger:
# 	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

doc:
	cd cmd/ && swag init

.PHONY: run
run:
	cd cmd/ && time go build -o ../build/bin && cd .. && ./build/bin