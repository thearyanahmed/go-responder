## go-responder 
It is a collection of few (3-4) functions to send generic responses. 

#### response format
```json
{
    "code" : 2001,
    "status": "SUCCESS",
    "message": "User created successfully.",
    "data" : {
        "id": 1213,
        "name": "John Doe",
        "email": "john@doe.com"
    },
}
```

#### sending a response
the `Send` function takes the response write, message, data and a code. The code can be a http status code, can also be some other code like if you want to maintain custom numeric codes for your consumer.
`message`, `code` & `data` all are optional. `data` is an interface, meaining you can have anything (valid) in it.

```go
// the send function
func Send(w http.ResponseWriter,message string, data interface{},code int) {
    sendResponse(w,response{
        Code:    code,
        Status: SUCCESS,
        Message: message,
        Data: data,
    },code)
}
```
#### the error response
The error response is same a the success response but with a 'status' = 'ERROR' and instead of `data` its `errors`.

#### example
```go
func (handler *handler) createNewUser(w http.ResponseWriter,r *http.Request) {

    req := request.Request{}

    if valid := req.Validate(r,w,createNewUserRequest); valid == false {
        return
    }

    validated := req.ValidatedFormData(r,[]string{"first_name","last_name","email"})

    user, errs := createNewUser(validated["first_name"],validated["last_name"],validated["email"])

    if len(errs) > 0 {
        res.SendError(w,"Unprocessable entity.",errs,422)
        return
    }

    res.Send(w,userCreatedSuccessfully,user,200)
}
```
