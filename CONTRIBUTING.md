# how we start

**init** `go mod init git.elm.sa/devops/seal-cli`
**gin framework** `go get -u github.com/gin-gonic/gin`
**external config with viper** `go get -u github.com/spf13/viper`




# Helpful links from Hewedy
gin and swaggo

https://github.com/codegangsta/gin

https://github.com/swaggo/swag



example sh script to start the main.go with gin and swaggo (auto-reload on any file change)

https://github.com/mhewedy/myApi/blob/master/start



---------

examples on net/http client



struct body

https://github.com/mhewedy/go-gistlog/blob/master/gistlog.go (search for http.NewRequest you will find example get and patch)



binary body

https://github.com/mhewedy/vermin/blob/2ee8d2b7e8fb87d7fafa07ccda52dd44594892b4/images/vagrant/image_url.go#L145



basic auth & NTLM example + log request and response

https://github.com/mhewedy/ews/blob/0268d58d244b247d353c455dff9f04371d7448af/ews.go#L64



converting body to string

https://github.com/mhewedy/onesignal-client/blob/6df0fa0505fc71778832232a485751abd9276c95/main.go#L94