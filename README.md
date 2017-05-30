# Chaos Proxy

A controlled means of introducing chaos into your infrastructure for Windows, Mac and Linux.

Features:
- Target some or all endpoints via regular expression
- Ability to target only a proportion of requests (ie: 30% of requests will timeout)
- Force latency on endpoints
- Specify response status code

## How to use it:

### Step 1: Create a config file

```yml
# config.yml

config:
  port: 8082

endpoints:
  - host: (\S+)-consumer-iapi-(\S+)$ # required
    url: \/consumer\/optinflow(\/?)$ # required
    delay: 5000 # optional - default is 0
    range: 50 # required
    responseStatusCode: 504 # optional - default is 200

  - host: (\S+)-public-iapi-(\S+)$
    url: \/authorize\/providers\?tenant=(es|ie|it)$
    delay: 5000
    range: 30

  - host: www.bbc.co.uk$
    url: \/weather(\/?)$
    delay: 5000
    range: 100
    responseStatusCode: 504
```

### Step 2: Run it

```
$ chaos_proxy.exe -logtostderr=true
```

Note: Use `$ chaos_proxy.exe -h` to view all CLI options.

### Step 3: Proxy requests

**IIS:**

Add this to your web.config:

```
<system.net>
    <defaultProxy enabled="true">
        <proxy proxyaddress="http://127.0.0.1:{your port number}" bypassonlocal="False"/>
    </defaultProxy>
</system.net>
```

You may need to restart IIS `iisreset /stop` then `iisreset /start`

**Firefox:** 

In Firefox go to `Options > Advanced > Network > Connection Settings`

## Compatibility:

- Mac
- Windows
- Linux

## Coming soon

- Add ability to target DynamoDB and simulate throttling
- Ability to target requests based on HTTP request methods (GET, POST etc)
- HTTP endpoint allowing you to post behaviour changes to Chaos Proxy
- Dashboard allowing you to visualise traffic flowing through proxy and matching behaviours
- Set behaviour rules to activate/deactivate between certain dates and times
- Watch for config file changes
- Tests!