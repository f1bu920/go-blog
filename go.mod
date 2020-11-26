module github.com/EDDYCJY/go-gin-example

go 1.15

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/ugorji/go v1.2.0 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9 // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => ../go-application/go-gin-example/pkg/conf
	github.com/EDDYCJY/go-gin-example/middleware => ../go-application/go-gin-example/middleware
	github.com/EDDYCJY/go-gin-example/models => ../go-application/go-gin-example/models
	github.com/EDDYCJY/go-gin-example/pkg => ../go-application/go-gin-example/pkg/setting
	github.com/EDDYCJY/go-gin-example/pkg/e => ../go-application/go-gin-example/pkg/e
	github.com/EDDYCJY/go-gin-example/pkg/setting => ../go-application/go-gin-example/pkg/setting
	github.com/EDDYCJY/go-gin-example/pkg/util => ../go-application/go-gin-example/pkg/util
	github.com/EDDYCJY/go-gin-example/routers => ../go-application/go-gin-example/routers
//github.com/EDDYCJY/go-gin-example/routers/api/v1 => ../go-application/go-gin-example/routers/api/v1
)
