package smile

//go:generate echo "* using GOPATH" $GOPATH
//go:generate echo "* installing go-swagger..."
//go:generate go get -u github.com/go-openapi/runtime
//go:generate go get -u github.com/go-swagger/go-swagger/cmd/swagger
//go:generate swagger generate server -f swagger.yml
//go:generate swagger generate client -f swagger.yml
//go:generate echo "* installing govendor..."
//go:generate go get -u github.com/kardianos/govendor
//go:generate echo "* fixing dependencies..."
//go:generate govendor fetch +missing
//go:generate govendor add +external
//go:generate govendor sync
//go:generate govendor remove +unused
//go:generate echo "* success"
