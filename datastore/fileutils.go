package datastore

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

// PermissionType is a type of file permission
type PermissionType int

// Old-school way of defining constants in Go.
// Permission type constants
// const (
// 	Readonly Permission = syscall.O_RDONLY;
// 	Writeonly Permission = syscall.O_WRONLY;
// 	ReadWrite Permission = syscall.O_RDWR;
// )

// Permission of type Permission
// Enum namespacing trick to make it easier to use, find and manage.
var Permission = struct {
    Read   PermissionType
    Write  PermissionType
    ReadWrite PermissionType
	} {
    Read:   syscall.O_RDONLY,
    Write:  syscall.O_WRONLY,
    ReadWrite:  syscall.O_RDWR,
}

func (permission PermissionType) String() string {
	names := [...]string{
        "Readonly", 
        "Writeonly", 
        "ReadWrite",
	}

	if permission < Permission.Read || permission > Permission.ReadWrite {
		return "Unknown";
	}

	// return the string value for the permission
	// constant from the names array above.
	return names[permission];
}


// DoesFileExists checks if a file exists, this does not attempt to open the file.
// we can do better: https://freshman.tech/snippets/go/check-file-exists/
func DoesFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

// HasFileAccess checks if a file exists & have read permission.
func HasFileAccess(filePath string) bool {
	if(!DoesFileExists(filePath)) {
		return false;
	}
	_, error := os.Open(filePath)
	return !errors.Is(error, os.ErrPermission)
}

// HasFileAccessPermission checks if a file exists & have permission specified.
 func HasFileAccessPermission(filePath string, permission PermissionType) bool {
	if(!DoesFileExists(filePath)) {
		return false;
	}
	_, error := os.OpenFile(filePath, int(permission), 0)
	return !errors.Is(error, os.ErrPermission)
}

// LoadFile reads a file and returns the byte array of the file, has extra check for read permission, more detailed error messaging
func LoadFile(filename string)  ([]byte, error)  {
    if( !HasFileAccessPermission(filename, Permission.Read) ) {
        return nil, fmt.Errorf("File=%s is missing read permissions", filename);
    }

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error trying to read file=%s, error=%w", filename, err);
	}    

    return data, nil;
}

