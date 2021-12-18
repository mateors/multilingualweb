# MultiLingualWeb
golang multi lingual web application development


## Command executed
```
go install golang.org/x/text/cmd/gotext@latest
which gotext

mkdir bookstore
cd bookstore
go mod init bookstore.example.com
mkdir -p cmd/www
touch cmd/www/main.go  cmd/www/handlers.go
go mod tidy
go run ./cmd/www/

curl localhost:8080/en-gb
curl localhost:8080/bn-bd
curl localhost:4018/fr-ch
curl localhost:4018/da-DK

mkdir -p internal/translations
touch internal/translations/translations.go


go generate ./internal/translations/translations.go

cp internal/translations/locales/bn-BD/out.gotext.json internal/translations/locales/bn-BD/messages.gotext.json

cp internal/translations/locales/de-DE/out.gotext.json internal/translations/locales/de-DE/messages.gotext.json

go generate ./internal/translations/translations.go

go mod tidy
go run ./cmd/www/

go build ./cmd/www/

```


## Resources
* [Tutorial-1](https://www.alexedwards.net/blog/i18n-managing-translations)
* [Tutorial-2](https://phrase.com/blog/posts/internationalization-i18n-go)
* [Language-code](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes)
* [Region-code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
