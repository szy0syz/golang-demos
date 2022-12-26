# Best Way To Structuring Golang Code

> golang should be simple, don’t make it worst and complicated
> 
> **Flat Structure**

## Flow

![image](https://user-images.githubusercontent.com/10555820/209550763-7027862a-329a-43e6-b6b1-7abbf6f9299f.png)

## Run

```bash
# Run environment
docker-compose up -d

# Install GRPC
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

#export GO_PATH=~/go
#export PATH=$PATH:/$GO_PATH/bin

# Generate Your Proto file
protoc -I./proto --go_out=./proto --go-grpc_out=./proto ./proto/*.proto

# Run the code
go run cmd/*.go

```
