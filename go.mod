module github.com/IMQS/assetcore

replace github.com/IMQS/nf => ../nf

replace github.com/IMQS/serviceauth => ../serviceauth

go 1.13

require (
	github.com/IMQS/log v1.0.0
	github.com/IMQS/nf v0.0.0-00010101000000-000000000000
	github.com/IMQS/serviceauth v0.0.0-00010101000000-000000000000
	github.com/IMQS/serviceconfigsgo v1.0.0
	github.com/jinzhu/gorm v1.9.11
	github.com/julienschmidt/httprouter v1.3.0
	github.com/lib/pq v1.1.1
	gotest.tools v2.2.0+incompatible
)
