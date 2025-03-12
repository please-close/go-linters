build:
	go build -buildmode=plugin -o ./detectany.so plugin/detectany.go

lint-any:
	./custom-gcl run --config ./.golangci.yml