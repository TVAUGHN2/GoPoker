# GoPoker
(WORK IN PROGRESS) 
This is a basic deterministic poker engine for best hand. 
It will take input of many n based card hands, determine the best 5 card hand from each, and then compare the hands against each other to determine the best hand. 

TODOs: 
 *  ~~Move logic to  separate files.~~ 
 * 	~~Setup REST api.~~
 *  Add concurrency.
 *  ~~Apply idiomatic Golang test harness.~~
 *	Convert fmt to logging framework. (glog for lightweight use w/ ability to set logging level)
 *  Create logging interface

# Running 
- Navigate to root directory
- Run the following commands to build and run:
```bash
$ docker build -t go-poker-api .
$ docker run -it --rm --name go-poker-api go-poker-api
```