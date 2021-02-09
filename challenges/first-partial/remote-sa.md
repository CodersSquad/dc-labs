Remote Shapes Analyzer
======================
The `Remote Shapes Analyzer` will work as a web server that will
calculate the perimeter and area of a given irregular shape.  Given a
number of vertices in the command line, verify that they all together
are forming an irregular shape without collisions (you can see the
following
[link](https://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/)
to verify collisions), then, calculate the area and
perimeter. Finally, reply the web request with your results.


General Requirements
--------------------
- The program must be implemented in Go programming language.
- Use must implement your code in the `remote-sa.go` file.
- You need to follow the output format guidelines.
- [Coding best practices](http://talks.obedmr.com/programming-art) that we learned in class will be reviewed.
- Before submitting you code, make sure it's compiling and running correctly.


Execution example
-----------------

You will need 2 terminals:

- **Terminal 1 - Server**
```
$ make serve
```

- **Terminal 2 - Client**
```
# Working example
$ curl http://localhost:8000\?vertices\=\(-3,\1),\(2,3\),\(0,0\),\(-1.5,-2.5\)
Welcome to the Remote Shapes Analyzer
- Your figure has      : [4] vertices
- Vertices             : (-3,1) (2,3) (0,0) (-1.5, -2.5)
- Calculated Perimeter : 16.7
- Calculated Area      : 10.5

# Error handling example
$ curl http://localhost:8000\?vertices=\(-3,1\)\(2,3\)
Welcome to the Remote Shapes Analyzer
- Your figure has      : [2] vertices
ERROR - Your shape is not compliying with the minimum number of vertices.
```


Test Cases
----------

It can be automatically tested with the following commands. You will be required to open 2 terminals.
This is the way your program will be tested by the Classify API tester.

- **Terminal 1 - Server**
```
$ make serve
```

- **Terminal 2 - Client**
```
$ make test
```

*Extra Note:*

If you want to test your program manually, you open the
[`lab.mk`](lab.mk) file and you the `curl` commands separately.