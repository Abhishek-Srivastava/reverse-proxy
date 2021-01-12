# reverse-proxy
A very simple reverse proxy in go

### Generate self-signed certficates
```
openssl req  -new  -newkey rsa:2048  -nodes  -keyout revpro.key  -out revpro.csr

openssl  x509  -req  -days 365  -in revpro.csr  -signkey revpro.key  -out revpro.crt
```
### Use the revpro.key and revpro.crt in the reverseproxy.go:
```
http.ListenAndServeTLS("0.0.0.0:8080", "revpro.crt", "revpro.key", nil)
```
#### The port might be blocked by your IT. Make sure to check what ports are available.
### For the example IP in the reverseproxy.go, this is a typical curl command to your server would look like
```
curl -XGET -u username:password https://100.67.28.113:8080/yourendpoint  --insecure
```
