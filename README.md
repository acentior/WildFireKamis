## Fetch Random Joke

A production ready web service which combines two existing web services - an api which fetches random first name and last name from https://names.mcquay.me/api/v0 and an api which fetches random Chuck Norris joke from http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy.

## Running the service

```shell
make install
```

## Testing the service

```shell
make app_test
```

## Description

### Environment Values
    app:
        port: 5000
        limit: 5
        count: 50
limit : The maximum number of go routines that will be processed at once <br>
count : The total number of jokes you want to fetch from api