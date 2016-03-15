## Encoding - Standard Library

Encoding is the process or marshaling or unmarshaling data into different forms. Taking JSON string documents and convert them to values of our user defined types is a very common practice in many go programs today. Go's support for encoding is amazing and improves and gets faster with every release.

## Notes

* Support for Decoding and Encoding JSON and XML are provided by the standard libary.
* This package gets better and better with every release.

## Links

http://www.goinggo.net/2014/01/decode-json-documents-in-go.html

## Code Review

[Unmarshal JSON documents](example1/example1.go) ([Go Playground](https://play.golang.org/p/AYZyeNHDoD))

[Unmarshal JSON files](example2/example2.go) ([Go Playground](https://play.golang.org/p/pjLnUZOwgj))

[Unmarshal JSON from Postman ](example3/example3.go) ([Go Playground](https://play.golang.org/p/907b4IH3NE))

[Custom Marshaler and Unmarshler](example4/example4.go) ([Go Playground](http://play.golang.org/p/IIz0WiuP4Q))

[Unmarshal JSON files](example5/example5.go) ([Go Playground](https://play.golang.org/p/Dmplcr-Z53))

## Exercises

### Exercise 1

**Part A** Create a an array of JSON documents that contain a user name and email address. Declare a struct type that maps to the JSON document. Using the json package and create a slice of this struct type. Display the unmarshaled contents.

([Go Playground](https://play.golang.org/p/c7PUd1vNEI)) |
 Answer ([Go Playground](https://play.golang.org/p/JXn1yQK5U-))

 **Part B** Send a POST request from Postman and sent success data as response.

 ([Go Playground](https://play.golang.org/p/rA2rdCJSVz) |
	Answer ([Go Playground](https://play.golang.org/p/HQiImWUw86))
