package main

import (
	"flag"
	"fmt"
	"time"

	datastore "github.com/secret-server/mock-datastore/datastore"
)

func main() {
    dbPath := "";

	fmt.Println("+-------------------------------------------------------------+");
	fmt.Println("+-------------------------------------------------------------+");
	fmt.Println("\r\nThis is an example of a datastore package");
	t := time.Now()
	fmt.Println("now: ", t.Format("2006-01-02T03:04:05.99Z"));
	fmt.Println("now RFC3339 format: ", t.Format(time.RFC3339));
	
	fmt.Println("+-------------------------------------------------------------+");

	flag.StringVar(&dbPath, "dbPath", "", "Path to datastore files");
	flag.Parse();

	fmt.Println("dbPath value is: ", dbPath)

	storage, err := datastore.New(dbPath);
    if err != nil {
        panic(err)
    }

	var db datastore.Datastore = storage;

	fmt.Println("\r\n+-------------------------------------------------------------+");
	fmt.Println("Users");
	userName := "Kevin Kelche";
	user, err := db.GetUser(userName);
    if err != nil {
		fmt.Println("No user found for userName=", userName);
		return ;
    }

	var haveRoleId bool = db.DoesUserHaveRoleId(user, 0);
	fmt.Println("haveRoleId=", haveRoleId);

	var haveRoleName bool = db.DoesUserHaveRoleName(user, "Admin");
	fmt.Println("haveRoleName=", haveRoleName);
	
	var haveRoleId2 bool = db.DoesUserHaveRoleId(user, 1);
	fmt.Println("haveRoleId2=", haveRoleId2);

	var haveRoleName2 bool = db.DoesUserHaveRoleName(user, "CA Web Admin");
	fmt.Println("haveRoleName2=", haveRoleName2);
}
