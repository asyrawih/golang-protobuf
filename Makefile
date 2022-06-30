gen:
	 @echo "Remove Current Proto Generation"
	 @rm -rf "model/*.go";
	 @echo "Running Progo Generator .... 🤘🏼"
	 @protoc --proto_path=proto_model  --go_out=model   --go-grpc_out=model  --go-grpc_opt=paths=source_relative  --go_opt=paths=source_relative proto_model/*.proto 
	 @echo "Success"
	 @tree model 
	 @make run

run:
	@echo "Running Server"
	@go run "server/main.go"

consumer:
	@echo "Running Client"
	@go run "client/main.go"
	
