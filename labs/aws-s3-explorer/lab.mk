# build & test automation

PROXY_APP=awsS3Explorer
CLIENT_APP=awsS3Client

proxy:
	@echo Start AWS S3 Explorer Proxy
	go run ./${PROXY_APP}.go -port 8000
test:
	@echo Test 1 - only bucket case
	go run ./${CLIENT_APP}.go -bucket ryft-public-sample-data
	@echo Test 2 - bucket and directory case
	go run ./${CLIENT_APP}.go -bucket ryft-public-sample-data -directory DNS
	@echo Test 3 - missing parameters
	go run ./${CLIENT_APP}.go
	@echo Test 4 - permission denied bucket
	go run ./${CLIENT_APP}.go -bucket prodnh
