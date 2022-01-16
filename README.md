# my-crypto-wallets-info

It is a CLI use case that if you provide your Binance API key and secret key
(which will be stored in a .env file in your machine) to use them to make
HTTP requests to the Binance API it will show you (by now)
- Amount of each currency you have with its respective values in USD
- Total USD you have in cryptocurrencies

A key idea is to extend this to more wallets.

## Usage

You need to have Go 1.18 in your machine. Then run at project's root folder

```
go1.18beta1 run main.go
```
or you can build an executable with

```
go1.18beta1 build
```
and then run with
```
./my-crypto-wallets-info
```

## Future versions

Few ideas are
- Add support manipulate information from different wallets and cryptocurrency platforms
- Add functionalities such as historical values that implies a database
- Extend use case to make an interesting CLI application
- As a server-side of a web page


## Programming tools

- Go 1.18 beta 1
- Following CLEAN architecture
- Following TDD for domain package implementation
