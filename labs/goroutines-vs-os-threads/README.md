Lab - Goroutines vs OS Threads
==============================
Read section 9.8 from The [Go Programming Language](https://www.amazon.com/dp/0134190440/ref=cm_sw_em_r_mt_dp_U_Uz0RDbHAH27PH) book.
Then, do the following 2 excercises.

1. Construct a pipeline that connects an arbitrary number of goroutines with channels.
What is the maximum number of pipeline stages you can create without running out of memory? How long does a value take to transite
the entire pipeline?

2. Write a program with 2 goroutines that send messages back and forth over two unbuffered channels in ping-pong fashion.
How many communications per second can the program sustain?

- Update the `README.md` file on instructions about how to run your programs.
- Your 2 programs must automatically generate their proper result report depending on what parameters they received or the response
they're givng to the presented questions.
- Generated reports are free format and written on any file extension (i.e. .txt, .pdf, .md, .svg, .png, .jpg, .tex)

General Requirements and Considerations
---------------------------------------
- Submit only `*.go` files.
- Don't forget to handle errors properly.
- Coding best practices implementation will be also considered.


How to submit your work
=======================
```
GITHUB_USER=<your_github_user>  make submit
```
More details at: [Classify API](../../classify.md)
