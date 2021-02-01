package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"context"

	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
)

var (
	cx                 = flag.String("cx", "C023zw3x8", "Gsuites cx number")
	adminEmail         = flag.String("adminEmail", "admin@ddddd.com", "Gsuites Admin to impersonate")
	serviceAccountFile = flag.String("serviceAccountFile", "/path/to/google_apps_svc_dwd.json", "Service Account")
)

func main() {

	serviceAccountJSON, err := ioutil.ReadFile(*serviceAccountFile)
	if err != nil {
		log.Fatal(err)
	}

	config, err := google.JWTConfigFromJSON(serviceAccountJSON, admin.AdminDirectoryUserReadonlyScope,
		admin.AdminDirectoryGroupReadonlyScope)

	config.Subject = *adminEmail

	ctx := context.Background()

	srv, err := admin.New(config.Client(ctx))
	if err != nil {
		log.Fatal(err)
	}

	//https://pkg.go.dev/google.golang.org/api/admin/directory/v1#User
	u, err := srv.Users.Get("user10@esodemoapp2.com").Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User PosixAccounts %s", u.PosixAccounts)

	groupsReport, err := srv.Groups.List().Customer(*cx).Do()
	if err != nil {
		log.Fatal(err)
	}

	if len(groupsReport.Groups) == 0 {
		fmt.Print("No users found.\n")
	} else {
		fmt.Print("Group:\n")
		for _, u := range groupsReport.Groups {
			fmt.Printf("%v\n", u.Id)
		}
	}

}
