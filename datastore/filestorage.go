package datastore

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
)

type FileStorage struct {
	path         string
	userFilepath string
	userData     map[string]User

	roleFilepath string
	roleData     map[string]Role

	secretsFilepath string
	secretsData     map[string]Secret
	secretsDataById map[int]Secret
}

// DoesUserHaveRoleId implements Datastore.
func (fs *FileStorage) DoesUserHaveRoleId(user User, roleId int) bool {
	if user.Roles != nil {
		for _, role := range user.Roles {
			if role == roleId {
				return true;
			}
		}
	}
	return false
}

// DoesUserHaveRoleName implements Datastore.
func (fs *FileStorage) DoesUserHaveRoleName(user User, roleName string) bool {
	if user.Roles != nil {
		for _, userRoleId := range user.Roles {
			for _, role := range fs.roleData {
				if role.ID == userRoleId {
					if roleName == role.Name {
						return true;
					}
					
				}
			}
		}
	}
	return false;
}

func New(filePath string) (*FileStorage, error) {
	log.Printf("FileStorage filePath=%s", filePath)

	if Empty(filePath) {
		log.Printf("Empty filePath")
		return nil, fmt.Errorf("Empty filePath string")
	}

	if !HasFileAccessPermission(filePath, Permission.Read) {
		return nil, fmt.Errorf("Invalid filePath=%s", filePath)
	}

	fileStorage := &FileStorage{path: filePath}
	fileStorage.userData = loadUser(fileStorage)
	fileStorage.roleData = loadRole(fileStorage)
	fileStorage.secretsData = loadSecret(fileStorage)

	fileStorage.secretsDataById = make(map[int]Secret);
	for _, value := range fileStorage.secretsData {
		fileStorage.secretsDataById[value.ID] = value;
    }

	return fileStorage, nil
}

// AddSecret implements Datastore.
func (fs *FileStorage) AddSecret(string, Secret) error {
	panic("unimplemented")
}

// AddUser implements Datastore.
func (fs *FileStorage) AddUser(string, User) error {
	panic("unimplemented")
}

// DeleteSecret implements Datastore.
func (fs *FileStorage) DeleteSecret(string) error {
	panic("unimplemented")
}

// DeleteUser implements Datastore.
func (fs *FileStorage) DeleteUser(string) error {
	panic("unimplemented")
}

func (fs FileStorage) GetSecret(roleId string) (Secret, error) {
	return fs.secretsData[roleId], nil
}

func (fs FileStorage) GetSecretById(id int) (Secret, error) {
	if(id <= 0) { 
		return Secret{}, fmt.Errorf("Invalid id=%d", id)
	}

	return fs.secretsDataById[id], nil
}

func (fs FileStorage) GetSecrets() ([]Secret, error) {
	results := make([]Secret, 0, len(fs.secretsData))
	for _, key := range fs.secretsData {
		results = append(results, key)
	}
	return results, nil
}

// key to a specific user
func (fs FileStorage) GetUser(searchText string) (User, error) {
	for key, user := range fs.userData {
		if user.Name == searchText || user.DisplayName == searchText || user.Email == searchText || key == searchText {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("No user found for searchText=%s", searchText)
}

// search for all users
func (fs FileStorage) UserLookup(searchText string) ([]User, error) {
	ret := []User{};
	for key, user := range fs.userData {
		if strings.Contains(key, searchText) || 
			strings.Contains(user.Name, searchText) || 
			strings.Contains(user.DisplayName, searchText) || 
			strings.Contains(user.Email, searchText) {
            ret = append(ret, user)
		}
    }
    return ret, nil;
}

func (fs FileStorage) RoleLookup(searchText string) ([]Role, error) {
	ret := []Role{};
	for key, role := range fs.roleData {
		if strings.Contains(key, searchText) || 
			strings.Contains(role.Name, searchText) || 
			strings.Contains(strconv.Itoa(role.ID), searchText) {
            ret = append(ret, role)
		}
    }
    return ret, nil;
}

func (fs FileStorage) SecretLookup(searchText string) ([]Secret, error) {
	ret := []Secret{};
	for key, secret := range fs.secretsData {
		// check basic fields
		if strings.Contains(key, searchText) || 
			strings.Contains(secret.Name, searchText) || 
			strings.Contains(strconv.Itoa(secret.ID), searchText) {
            ret = append(ret, secret)
			continue;
		}
		// check slugs
		for k, v := range secret.Slug {
			if strings.Contains(k, searchText) || strings.Contains(v, searchText)  {
            	ret = append(ret, secret);
				break;
			}
		}
			
    }
    return ret, nil;
}


func (fs FileStorage) GetUsers() ([]User, error) {
	userArray := make([]User, 0, len(fs.userData))
	for _, value := range fs.userData {
		userArray = append(userArray, value)
	}
	return userArray, nil
}

// UpdateSecret implements Datastore.
func (fs *FileStorage) UpdateSecret(string, Secret) error {
	panic("unimplemented")
}

// UpdateUser implements Datastore.
func (fs *FileStorage) UpdateUser(string, User) error {
	panic("unimplemented")
}

func (fs *FileStorage) CreateRole(name string, enable bool) (Role, error) {
	if Empty(name) {
		return Role{}, fmt.Errorf("Invalid name, name is an empty string")
	}

	for _, role := range fs.roleData {
		if role.Name == name {
			return Role{}, fmt.Errorf("Role `%s` already exists", name)
		}
	}

	// len(fs.roles)+1 could fail if roles are deleted or someone screws with the data manually
	// the following is used instead
	newId := 0
	for _, role := range fs.roleData {
		if role.ID > newId {
			newId = role.ID
		}
	}
	newId++

	// "created": "2024-09-14T20:26:36.903Z",
	// https://go.dev/src/time/format.go
	// 20060102150405
	t := time.Now()
	role := Role{
		Name:     name,
		ID:       newId,
		Created:  t.Format("2006-01-02T03:04:05.99Z"),
		Enabled:  enable,
		IsSystem: false,
	}

	fs.roleData[name] = role
	fs.savedRoles()
	return role, nil
}

func (fs FileStorage) GetRole(roleId int) (Role, error) {
	if roleId < 0 {
		return Role{}, fmt.Errorf("Invalid roleId=%d", roleId)
	}
	for _, role := range fs.roleData {
		if role.ID == roleId {
			return role, nil
		}
	}
	return Role{}, fmt.Errorf("No Role found for roleId=%d", roleId)
}

func (fs *FileStorage) GetRoleByName(name string) (Role, error) {
	if Empty(name) {
		return Role{}, fmt.Errorf("Invalid name, name is an empty string")
	}
	for _, role := range fs.roleData {
		if strings.EqualFold(role.Name, name) {
			return role, nil
		}
	}
	return Role{}, fmt.Errorf("No Role found for name=%s", name)
}


func (fs FileStorage) GetRoles() ([]Role, error) {
	results := make([]Role, 0, len(fs.roleData))
	for _, role := range fs.roleData {
		results = append(results, role)
	}
	return results, nil
}

func (fs *FileStorage) UpdateRole(roleId int, name string, enabled bool) (Role, error) {
	if roleId < 0 {
		return Role{}, fmt.Errorf("Invalid roleId=%d", roleId)
	}

	if Empty(name) {
		return Role{}, fmt.Errorf("Invalid name, name is an empty string")
	}

	keyName := ""
	for _, role := range fs.roleData {
		if role.ID == roleId {
			keyName = role.Name
			break
		}
	}

	if Empty(keyName) {
		return Role{}, fmt.Errorf("Invalid role id.  No role for roleId=%d", roleId)
	}

	role := fs.roleData[keyName]
	roleUpdate := Role{
		Name:     name,
		ID:       role.ID,
		Created:  role.Created,
		Enabled:  enabled,
		IsSystem: role.IsSystem,
	}

	fmt.Println("")
	fmt.Println("Update role=", roleUpdate)
	fmt.Println("")

	// delete the orginal role
	delete(fs.roleData, role.Name)

	// add the updated role
	fs.roleData[name] = roleUpdate
	fmt.Println("")
	fmt.Println("Roles=", fs.roleData)
	fmt.Println("")
	fs.savedRoles()
	return roleUpdate, nil
}

func (fs *FileStorage) savedRoles() {
	data, err := yaml.Marshal(&fs.roleData)
	if err != nil {
		log.Fatal(err)
	}

	err2 := os.WriteFile(fs.roleFilepath, data, 0)
	if err2 != nil {
		log.Fatalf("Error saving roles to role file=%s, error=%v", fs.roleFilepath, err2)
		return
	}
}

func loadUser(fs *FileStorage) map[string]User {
	fs.userFilepath = filepath.Join(fs.path, "users.yaml")
	yamlFile, err := os.ReadFile(fs.userFilepath)
	if err != nil {
		panic(err)
	}

	var users map[string]User
	err = yaml.Unmarshal(yamlFile, &users)
	if err != nil {
		panic(err)
	}

	return users
}

func loadRole(fs *FileStorage) map[string]Role {
	fs.roleFilepath = filepath.Join(fs.path, "roles.yaml")
	yamlFile, err := os.ReadFile(fs.roleFilepath)
	if err != nil {
		panic(err)
	}

	var roles map[string]Role
	err = yaml.Unmarshal(yamlFile, &roles)
	if err != nil {
		panic(err)
	}

	return roles
}

func loadSecret(fs *FileStorage) map[string]Secret {
	fs.secretsFilepath = filepath.Join(fs.path, "secrets.yaml")
	yamlFile, err := os.ReadFile(fs.secretsFilepath)
	if err != nil {
		panic(err)
	}

	var secrets map[string]Secret
	err = yaml.Unmarshal(yamlFile, &secrets)
	if err != nil {
		panic(err)
	}

	return secrets
}

func (fs *FileStorage) addSecret(filename string, secrets map[string]Secret, secret Secret) (map[string]Secret, error) {
	if Empty(filename) {
		return nil, fmt.Errorf("Filename is empty")
	}

	key := strconv.Itoa(secret.ID)
	name := Trim(secret.Name)

	if Empty(key) {
		return nil, fmt.Errorf("key is empty")
	}

	if secret.ID == 0 {
		return nil, fmt.Errorf("Missing value for secret.ID")
	}

	if Empty(name) {
		return nil, fmt.Errorf("Missing name for secret.Name")
	}

	value := secrets[key]
	if value.ID != 0 && value.Name != "" {
		return nil, fmt.Errorf("Secret `%s` already exists", key)
	}

	value = secrets[name]
	if value.ID != 0 && value.Name != "" {
		return nil, fmt.Errorf("Secret `%s` already exists", name)
	}

	secrets[key] = secret
	secrets[name] = secret

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, _ := yaml.MarshalWithOptions(secrets, yaml.UseLiteralStyleIfMultiline(true))
	_, err = io.WriteString(file, string(data))
	if err != nil {
		return nil, err
	}

	// field := reflect.ValueOf(value).Field(0)
	// fmt.Println("field.Interface()", field.Interface()==nil);
	// fmt.Println("field.Interface()", reflect.TypeOf(field.Interface()));
	// fmt.Println("reflect.Zero", reflect.Zero(field.Type()));
	// fmt.Println("reflect.TypeOf(reflect.Zero)", reflect.TypeOf(reflect.Zero(field.Type())))
	return secrets, nil
}
