package util

import (
	"context"
	"log"
    "path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func SetupFirebaseClient(path string)*auth.Client{
    servicAccountfilePath,err := filepath.Abs(path);

    if err!=nil{
        log.Fatal("unable to load firebase config file: ",err)
    }

    opt:= option.WithCredentialsFile(servicAccountfilePath)

    //Firebase sdk initialization
    app,err:=firebase.NewApp(context.Background(),nil,opt)
    if err!=nil{
        log.Fatal("Firebase load error")
    }

    //Firebase auth
    auth,err:=app.Auth(context.Background())
    if err!= nil{
        log.Fatal("Firebase load error: ",err);
    }

    return auth
}
