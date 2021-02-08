# build & test automation

test:
	@echo Test slices.go
	go run ./slices.go
	@echo Test maps.go
	go run ./maps.go
