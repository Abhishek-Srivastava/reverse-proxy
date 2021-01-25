# reverse-proxy
A very simple reverse proxy in go

### Generate self-signed certficates
```
openssl req  -new  -newkey rsa:2048  -nodes  -keyout revpro.key  -out revpro.csr

openssl  x509  -req  -days 365  -in revpro.csr  -signkey revpro.key  -out revpro.crt
```
### Use the revpro.key and revpro.crt in the cli:
```
reverse-proxy start --ipaddress 172.200.18.22 --port 443 --protocol https --proxyport 8080 --httptimeout 10 --certfile cert.crt --keyfile key.key
reverse-proxy start --ipaddress 172.200.18.22

Usage:
  reverse-proxy start [flags]

Flags:
      --certfile string    certificate file with path for tls (default "./certs/revpro.crt")
  -h, --help               help for start
      --httptimeout int    http timeout value in seconds (default 10)
      --ipaddress string   ip of the remote server
      --keyfile string     certificate key file with path for tls (default "./certs/revpro.key")
      --port string        port of the remote server (default "443")
      --protocol string    protocol to connect http/https (default "https")
      --proxyport string   proxyport on which the revereproxy would be served (default "8080")

```
#### The port might be blocked by your IT. Make sure to check what ports are available.

### For the example IP in the reverseproxy.go, this is how a typical curl command to your server would look like

```
curl -XGET -u username:password https://100.67.28.113:8080/yourendpoint  --insecure
```
