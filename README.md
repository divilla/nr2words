# nr2words
### 📖 Go int number to English words converter

## 🏃‍♀️ Run
```sh
# List make commands / make help
make

# Run the demo app
make run

# Run tests
make test

# ntw package has 100% test coverage
make test-cover
```

## 📚 How it works?
- package ntw exposes two public functions:
    - **CardinalToWords(i int)** that converts integer to cardinal English language words expression 
    - **OrdinalToWords(i int)** that converts integer to ordinal English language words expression 
- it has 100% test coverage
- it's split into 4 files:
    - nr2words.go contains public functions and main parser logic
    - token.go contains token types used to parse various parts of number
    - metadata.go contains maps to all English words used by parser
    - nr2words_test.go contains tests
- code is, I believe, well commented
- this code is longer version, built with readability, maintainability and configurability instead of performance and minimal code base in mind

## 📌 Usage example
```go
res, err := ntw.CardinalToWords(56945781)
if err != nil {
    log.Fatal(err)
}
// prints "fifty-six million nine hundred and forty-five thousand seven hundred and eighty-one"
fmt.Println(res)

res, err = ntw.OrdinalToWords(1000099)
if err != nil {
    log.Fatal(err)
}
// prints: "one million and ninety-ninth"
fmt.Println(res)
```
