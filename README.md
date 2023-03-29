# Golang

## url: https://restless-snowflake-5074.fly.dev

## Signup url: https://restless-snowflake-5074.fly.dev/signup
### payload json as
{
  "Id":1,
  "Name":"Shivam",
  "username":"mshivam019",
  "password":"password"
}

## Login url: https://restless-snowflake-5074.fly.dev/login
### payload json as
{
  "username":"mshivam019",
  "password":"password"
}

## Get all posts https://restless-snowflake-5074.fly.dev/post
## also add your jwt token as header

## Delete post by id:-
### url: https://restless-snowflake-5074.fly.dev/post/1
### request method: delete


## Add post req method: POST
### url: https://restless-snowflake-5074.fly.dev/post
## also add your jwt token as header
### payload json as
 {  "Id":1,
    "Name":"shivam",
    "Body":"hello" }



## Get post by id:-
### url: https://restless-snowflake-5074.fly.dev/post/3
### request method: get


## Modify data value to anything req method: Put
### url: https://restless-snowflake-5074.fly.dev/post
### payload json as
 {  "Id":1,
    "Name":"shivam",
    "Body":"hello" }
