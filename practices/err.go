package main

import (
	"fmt"
	"learnGo/customizePackages/goerr"
)

func main() {
	fmt.Println(ResourceNotFound("1234", "User", nil).Error())
}

// ResourceNotFound error abstraction
func ResourceNotFound(id, kind string, cause error) goerr.Error {
	data := map[string]interface{}{"kind": kind, "id": id}
	return goerr.NewGoError("ResourceNotFound", data, cause).
		SetComponent(goerr.ErrService).SetResponseType(goerr.NotFound)
}
