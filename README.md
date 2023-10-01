# URL Shortner
# Docker Build

    > docker build -t urlservice1.6 .
    > docker run --publish 8080:8080 urlservice1.5

Note : By default this service is using :8080 port
for using custom port pass port in env 

# APIs List
> [GET] localhost:8080/healthCheck 

    RESPONSE
    200 [String]
    Welcome to short url service

> [POST] localhost:8080/createURL

    REQUEST JSON
    {
    "url":"aloksonker.me"
    }
    

    RESPONSE
    200 [String]
    https://url.com/1000000




> [GET] localhost:8080/getURL

    REQUEST JSON
    {
        "url":"https://url.com/1000000"
    }

    RESPONSE
    200 [String]
    aloksonker.me

> [GET] localhost:8080/getAllURL

    RESPONSE
    200 [JSON]
    {
    "1000000": "aloksonker.me"
    }

> [GET] localhost:8080/metrics

    RESPONSE
    200 [JSON]
    {
    "alok.com": 5,
    "r.com": 5,
    "t.com": 7
    }