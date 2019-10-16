# Asset Core

For documentation, see http://docs.imqs.co.za/#Services-AssetCore

# Building
```shell
go build github.com/IMQS/assetcore/core
```

# Testing
```shell
docker run -p 5432:5432 imqs/postgres:unittest-10.5
go test github.com/IMQS/assetcore/core
```