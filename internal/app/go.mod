module app

go 1.16

replace domain => ../domain

replace infrastructure => ../infrastructure

require (
	domain v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.3
	infrastructure v0.0.0-00010101000000-000000000000
)
