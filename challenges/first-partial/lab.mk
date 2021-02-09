# build & test automation

APP_NAME=remote-sa

serve:
	@echo Starting Remote Shade Analyzer
	go run ./${APP_NAME}.go

test:
	@echo Test 1 - 3 vertices shade
	curl http://localhost:8000\?vertices=\(-3,1\),\(2,3\),\(-1.5,-2.5\)
	@echo Test 2 - 4 vertices shade
	curl http://localhost:8000\?vertices=\(-3,1\),\(2,3\),\(0,0\),\(-1.5,-2.5\)
	@echo Test 3 - 5 vertices shade
	curl http://localhost:8000\?vertices=\(-3,1\),\(2,3\),\(0,0\),\(2,-3\),\(-1.5,-2.5\)
	@echo Test 4 - 2 vertices shade
	curl http://localhost:8000\?vertices=\(-3,1\),\(2,3\)
