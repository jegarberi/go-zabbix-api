package zabbix_test

import (
	zapi "github.com/atypon/go-zabbix-api"
	"testing"
)

func TestRole(t *testing.T) {
	testRoleName := "TestRole"
	role := &zapi.Role{Name: testRoleName, Type: zapi.UserRole}
	testCRUDAPIObjectOperations(t, role)
}
