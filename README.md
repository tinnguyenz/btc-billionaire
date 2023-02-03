# BTC Billionaire

## Local Setup

Make sure you have Golang installed in your PC
Please refer to this how-to for setting up: https://go.dev/doc/install


### Clone the btc-billionaire repository

```bash
git clone https://github.com/tinnguyenz/btc-billionaire.git

git clone https://github.com/tinnguyenz/btc-billionaire.git

go run
```


### Run test cases

```
cd tests
go test
```

### Play around

#### Get some crypto
```bash
curl -H "Content-Type: application/json" -X POST -d '{"Amount":5000,"datetime":"2022-02-03T06:48:02Z"}' localhost:8080/records

curl -H "Content-Type: application/json" -X POST -d '{"Amount":5000,"datetime":"2023-02-03T08:48:02Z"}' localhost:8080/records
```

#### Check transaction history
```bash
curl -H "Content-Type: application/json" -X GET -d '{"startDatetime":"2022-01-05T10:48:01+00:00","endDatetime":"2022-12-05T18:48:02+00:00"}' localhost:8080/showHistory
```

## Improvement Points
1. Enhance validation of JSON input
2. Support multiple databases
3. Implement Goroutines for better resource management

## Note
Any comments are welcome as I'll take them for learning
