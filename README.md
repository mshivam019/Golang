# Golang

## url: https://golang-production-54b0.up.railway.app

## Signup url: https://golang-production-54b0.up.railway.app/signup
### payload json as
{
  "Id":1,
  "Name":"Shivam",
  "username":"mshivam019",
  "password":"password"
}

## Login url: https://golang-production-54b0.up.railway.app/login
### payload json as
{
  "username":"mshivam019",
  "password":"password"
}

## Get all posts https://golang-production-54b0.up.railway.app/post
# also add your jwt token as header

## Delete post by id:-
### url: https://golang-production-54b0.up.railway.app/post/1
### request method: delete


## Add post req method: POST
### url: https://golang-production-54b0.up.railway.app/post
# also add your jwt token as header
### payload json as
 {  
    "Id":1,
	  "Name":"shivam",
    "Body":"hello"
  }



## Get post by id:-
### url: https://golang-production-54b0.up.railway.app/post/3
### request method: get


## Modify data value to anything req method: Put
### url: https://golang-production-54b0.up.railway.app/post
### payload json as
 {  
    "Id":1,
	  "Name":"shivam",
    "Body":"hello"
  }
