Lab - Amazon S3 Explorer
========================

Create a Simple Amazon S3 objects explorer. Your program will receive the name of a public bucket and then, it will explore details about it.

The main idea of the exercise is to interact with the RESTFul API that AWS is providing in order to get its objects and buckets metadata.

General Requirements and Considerations
---------------------------------------
- Use the `aws_s3_explorer.go` file for your implementation.
- Don't forget to handle errors properly.
- Coding best practices implementation will be also considered.

Handy resources
---------------
- https://golang.org/pkg/net/http/
- https://golang.org/pkg/encoding/xml/
- https://docs.aws.amazon.com/sdk-for-go/api/service/s3/ (optional)

Test Cases
==========
Your program will be tested with the following URLs:

**Note:** Below output data was invented. Please don't take those numbers as your goal.

- Test 1: https://ryft-public-sample-data.s3.amazonaws.com
```
$ go run aws_s3_explorer.go  --bucket ryft-public-sample-data
AWS S3 Explorer
Bucket Name            : ryft-public-sample-data
Number of objects      : 719
Number of directories  : 10
Extensions             : pdf(250), gif(100), txt(300), png(69)
```
- Test 2: http://prodnh.s3.amazonaws.com
```

$ go run aws_s3_explorer.go --bucket prodnh
AWS S3 Explorer
Bucket Name            : prodnh
Number of objects      : 543
Number of directories  : 12
Extensions             : docx(200), gif(100), tgz(200), jpeg(43)
```

How to submit your work
=======================
```
GITHUB_USER=<your_github_user>  make submit
```
More details at: [Classify API](../../classify.md)
