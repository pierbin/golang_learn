module apps/learn

go 1.17

require (
	github.com/asmcos/requests v0.0.0-20210319030608-c839e8ae4946
	github.com/gin-gonic/gin v1.7.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

require (
	github.com/ramseyjiang/golang_learn/greetings v0.0.0-20211109004149-270f1ee03f37 // indirect, replace github module to local module
	github.com/ramseyjiang/golang_learn/keyboard v0.0.0-20211109034547-af1efbb27cd5 // indirect
)

replace (
	github.com/ramseyjiang/golang_learn/greetings v0.0.0-20211109004149-270f1ee03f37 => /Users/jiangdawei/go/src/apps/greetings
	github.com/ramseyjiang/golang_learn/keyboard v0.0.0-20211109034547-af1efbb27cd5 => /Users/jiangdawei/go/src/apps/keyboard
)

//below is used for local greetings module work.
//replace greetings => /Users/jiangdawei/go/src/apps/greetings
//replace keyboard => /Users/jiangdawei/go/src/apps/keyboard