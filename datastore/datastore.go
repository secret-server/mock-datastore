package datastore

import (
	"strings"
	"time"
)

// Datastore is an In-Memory Database.  This is the interface that will be used to interact with the datastore.
type Datastore interface {
    // user based operations
    GetUser(string) (User, error)
    GetUsers() ([]User, error)
    AddUser(string, User) error
    UpdateUser(string, User) error
    DeleteUser(string) error
    UserLookup(searchText string) ([]User, error)

    // role based operations
    GetRole(int) (Role, error)   
    GetRoles() ([]Role, error)
    CreateRole(string, bool) (Role, error) 
    UpdateRole(roleId int, name string, enabled bool)(Role, error) 

    // secret based operations
    GetSecret(string) (Secret, error)
    GetSecrets() ([]Secret, error)
    AddSecret(string, Secret) error
    UpdateSecret(string, Secret) error
    DeleteSecret(string) error

    DoesUserHaveRoleId(User, int) (bool)
	DoesUserHaveRoleName(User, string) (bool)
}

// User comment
type User struct {
    Name        string `yaml:"name"`
    DisplayName string `yaml:"display_name"`
    Email       string `yaml:"email"`
    ID    	    int    `yaml:"id"`
    Password    string `yaml:"password"`
    Roles       []int  `yaml:"roles"`
}

type Date struct {
    Time time.Time
}
  
// Role comment
type Role struct {
    Name     string `yaml:"name"`
    ID    	 int 	`yaml:"id"`
    Created  string `yaml:"created"`
    Enabled  bool 	`yaml:"enabled"`
    IsSystem bool 	`yaml:"isSystem"`
}

// Secret comment
type Secret struct {
    Name     string `yaml:"name"`
    ID    	 int 	`yaml:"id"`
	Slug     map[string]string `yaml:"fields"`
}



// NonEmpty checks if a string value is not empty.
func NonEmpty(value string) bool {
    return len(Trim(value)) > 0;
}

// Empty checks if a string value is empty.
func Empty(value string) bool {
    return len(Trim(value)) == 0;
}

// Trim the blank ' ' spaces from both sides of a sting, return the string.
func Trim(s string) string {
    return strings.TrimSpace(s)
}

