# Lab 2 GO
## Topic: 
- Continuous integration and test automation

## Purpose: 
- Implementation of the practice of continuous integration, consolidation of the skills of use
dependency injection pattern to simplify testing

## This project create by:

> [Мартинюк Марія](https://github.com/mmarty12) <br>
> [Котенко Ярослав](https://github.com/yarikkot04) <br>
> [Дубчак Аліна](https://github.com/AlinaDubchak) <br>

## Description
This project evaluates postfix expressions. Mathematical operations supported by the project include:
- addition (+)
- subtraction(-)
- multiplication (*)
- division (/)
- exponentiation (^)

## How to run project

1. Clone the repository to your local machine

2. Navigate to the project directory in your terminal or command prompt

3. To run the project
```
$ go run cmd/example/main.go
``` 
4. Reading an input expression from the command line
```
go run main.go -e="4 2 - "
```
5. Reading an input expression from the specified file
```
go run main.go -f="test.txt"
```
6. Write the conversion result to some file
```
go run main.go -e="4 2 - " -o="result.txt
```

## How to run test

```
$ go test
``` 

## References

[GitHub Actions](https://github.com/AlinaDubchak/Lab2-Go/actions)

[Successful build](https://github.com/AlinaDubchak/Lab2-Go/commit/047399cc1c40ae63cfc88b5558d7cb7a765c75f6)

[Failed build](https://github.com/AlinaDubchak/Lab2-Go/commit/49abb4d2ecab0a8223d5439ddeb349e8dda7a3b2)

[Pull request](https://github.com/AlinaDubchak/Lab2-Go/commit/cfde5f678bcc186beea944335e807e13522e3df7)
