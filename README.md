# urlShortner

# APIs List
> [GET] localhost:8080/healthCheck
    RESPONSE
    200 [String]
    Welcome to short url service
> [POST] localhost:8080/createURL
    REQUEST JSON
    ```json
    {
    "url":"aloksonker.me"
    }
    ```
    

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

