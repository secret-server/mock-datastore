package main

import (
	"flag"
	"fmt"
	"time"

	datastore "github.com/secret-server/mock-datastore/datastore"
)

var(
    dbPath string
)

func main() {

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
	fmt.Println("User.Name=", user.Name);
	fmt.Println("User.DisplayName=", user.DisplayName);
	fmt.Println("User.Email=", user.Email);
	fmt.Println("User.Password=", user.Password);

	fmt.Println("");
	userEmail := "kevin.kelche@example.com";
	fmt.Println("Look up vai email=", userEmail);
	user2, err := db.GetUser(userEmail);
    if err != nil {
		fmt.Println("No user found for userName=", userEmail);
		return ;
    }
	fmt.Println("User2.Name=", user2.Name);
	fmt.Println("User2.DisplayName=", user2.DisplayName);
	fmt.Println("User2.Email=", user2.Email);


	fmt.Println("");
	users, err := db.GetUsers();
    if err != nil {
		fmt.Println("No users found");
		return ;
    }
	fmt.Println("User count=", len(users));
	for  _, value := range users {
		fmt.Println("");
		fmt.Println("\tUser=", value.Name);
		fmt.Println("\tUser=", value.Password);
		fmt.Println("\tUser=", value.Roles);
		for  _, roleId := range value.Roles {
			role, err := db.GetRole(roleId);
			if err == nil {
				fmt.Println("\t\troleId=", roleId);
				fmt.Println("\t\trole.Name=", role.Name);
				fmt.Println("\t\trole.ID=", role.ID);
				fmt.Println("\t\trole.Enabled=", role.Enabled);
				fmt.Println("");
			}
		}
	 }

	 fmt.Println("+-------------------------------------------------------------+");
	 fmt.Println("Roles");
	 roles, err := db.GetRoles();
	 if err != nil {
		 fmt.Printf("No roles found");
		 return ;
	 }
	 fmt.Printf("Roles count=%d", len(roles));
	 for  _, role := range roles {
		fmt.Println("");
		fmt.Println("\trole.name=", role.Name);
		fmt.Println("\trole.id=", role.ID);
		fmt.Println("\trole.enabled=", role.Enabled);
	 }
 

	//  // WORKS:
	//  fmt.Println("");
	//  fmt.Println("Update role");
	//  adminRole,_ := db.UpdateRole(0, "AdminU", false);
	//  fmt.Println("\tadminRole.name=", adminRole.Name);
	//  fmt.Println("\tadminRole.id=", adminRole.ID);
	//  fmt.Println("\tadminRole.enabled=", adminRole.Enabled);

	//  fmt.Println("");
	//  adminRoleU, err := db.GetRole(0);	 
	//  if(err != nil) {
	// 	 fmt.Println("No role found for roleId. ", err.Error());
	//  } else {
	// 	fmt.Println("\tread admin changes.name=", adminRoleU.Name);
	// 	fmt.Println("\tread admin changes.id=", adminRoleU.ID);
	// 	fmt.Println("\tread admin changes.enabled=", adminRoleU.Enabled);
	// }

	//  // WORKS:
	//  fmt.Println("");
	//  fmt.Println("Add role");
	//  adminNewRole, _ := db.CreateRole("AdminNew", true);
	//  fmt.Println("\tadminNewRole.name=", adminNewRole.Name);
	//  fmt.Println("\tadminNewRole.id=", adminNewRole.ID);
	//  fmt.Println("\tadminNewRole.enabled=", adminNewRole.Enabled);

	 // func (fs *FileStorage) UpdateRole(roleId int, name string, enabled bool) (Role, error) {
	// fmt.Println("");
	// roles2, err := db.GetRoles();
	// if err != nil {
	// 	fmt.Printf("No roles found");
	// 	return ;
	// }
	// fmt.Printf("Roles2 count=%d", len(roles2));
	// for  _, role := range roles2 {
	// 	fmt.Println("");
	// 	fmt.Println("\trole.name=", role.Name);
	// 	fmt.Println("\trole.id=", role.ID);
	// 	fmt.Println("\trole.enabled=", role.Enabled);
	// }
 


	fmt.Println("\r\n+-------------------------------------------------------------+");
	fmt.Println("Secrets");
	secrets, err := db.GetSecrets();
	if err != nil {
		fmt.Printf("No secrets found");
		return ;
	}
	fmt.Println("Secrets count=", len(secrets));
	for  _, secret := range secrets {
		fmt.Println("\tsecret.name=", secret.Name);
		fmt.Println("\tsecret.id=", secret.ID);
		fmt.Println("");
	}

	fmt.Println("\r\n+-------------------------------------------------------------+");
	secret, err := db.GetSecret("caDatabase");
    if err != nil {
		fmt.Println("No secret found for key=", "caDatabase");
		return ;
    }
	fmt.Println("secret.ID=", secret.ID);
	fmt.Println("secret.Name=", secret.Name);
	fmt.Println("\tsecret.server=", secret.Slug["server"]);
	fmt.Println("\tsecret.user=", secret.Slug["user"]);
	fmt.Println("\tsecret.password=", secret.Slug["password"]);

	fmt.Println("\r\n+-------------------------------------------------------------+");
	privateKeySecret, err := db.GetSecret("CA Private Key Passphrase");
    if err != nil {
		fmt.Println("No secret found for key=", "caDatabase");
		return ;
    }
	fmt.Println("secret=", privateKeySecret.Name);
	fmt.Println("\tsecret.passphrase=", privateKeySecret.Slug["private-key-passphrase"]);

}