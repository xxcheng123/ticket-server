### 1. N/A

1. route definition

- Url: /user/register
- Method: POST
- Request: `RegisterReq`
- Response: `RegisteraResp`

2. request definition



```golang
type RegisterReq struct {
	Username string 
	Password string 
}
```


3. response definition



```golang
type RegisteraResp struct {
	Ok bool 
	Id string 
}
```

### 2. N/A

1. route definition

- Url: /user/info
- Method: POST
- Request: `InfoReq`
- Response: `InfoResp`

2. request definition



```golang
type InfoReq struct {
}
```


3. response definition



```golang
type InfoResp struct {
	Id string 
	Username string 
	Balance float64 
}
```

