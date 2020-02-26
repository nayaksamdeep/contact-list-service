module github.com/nayaksamdeep/contact-list-service

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/jinzhu/gorm v1.9.12
	github.com/lib/pq v1.3.0
	github.com/nayaksamdeep/contact-list-service/Config v0.0.0
)

replace github.com/nayaksamdeep/contact-list-service/Config v0.0.0 => ./Config
