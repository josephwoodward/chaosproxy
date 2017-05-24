# Chaos Kitten (better name coming soon?)

A controlled means of introducing chaos into your infrastructure for Windows, Mac and Linux.

Available behaviours (more coming soon):
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
    methods: [GET] # coming soon
    delay: 5000 # optional - default is 0
    responseStatusCode: 504 # optional - default is 200

  - host: (\S+)-public-iapi-(\S+)$
    url: \/authorize\/providers\?tenant=(es|ie|it)$
    methods: [GET]
    delay: 5000

  - host: www.bbc.co.uk$
    url: \/weather(\/?)$
    methods: [GET]
    delay: 5000
    responseStatusCode: 504
```

### Step 2: Run it  
```
$ chaos_kitten.exe -logtostderr=true
```

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

- HTTP endpoint allowing you to post behaviour changes to your Chaos Kitten
- Watch for config file changes
- Tests!