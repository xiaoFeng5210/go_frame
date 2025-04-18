package database_test

import (
	"crypto/md5"
	database "dqq/go/frame/post/database/xorm"
	"dqq/go/frame/post/util"
	"encoding/hex"
	"fmt"
	"testing"
)

var (
	uid = 7
)

func hash(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	digest := hasher.Sum(nil)
	return hex.EncodeToString(digest) //md5的输出是128bit，十六进制编码之后长度是32
}

func init() {
	util.InitSlog("../../../log/post.log")
	database.ConnectPostDB("../../conf", "db", util.YAML, "../../../log")
}

func TestRegistUser(t *testing.T) {
	err := database.RegistUser("dlm", hash("123456"))
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Printf("注册成功\n")
	}

	err = database.RegistUser("dlm", hash("123456"))
	if err != nil {
		fmt.Printf("注册失败: %s\n", err)
	} else {
		fmt.Println("重复注册成功！")
		t.Fail()
	}
}

func TestGetUserById(t *testing.T) {
	user := database.GetUserById(uid)
	if user == nil {
		t.Fatalf("could not get user by id %d", uid)
	}

	tmpUid := -1
	user = database.GetUserById(tmpUid)
	if user != nil {
		t.Fatalf("get user by id %d, user %v", tmpUid, *user)
	}
}

func TestGetUserByName(t *testing.T) {
	user := database.GetUserByName("dlm")
	if user == nil {
		t.Fail()
	}

	user = database.GetUserByName("ok")
	if user != nil {
		t.Fail()
	}
}

func TestUpdateUserName(t *testing.T) {
	err := database.UpdateUserName(uid, "zcy")
	if err != nil {
		t.Fatal(err)
	}
	user := database.GetUserById(uid)
	if user == nil {
		t.Fail()
		return
	}
	if user.Name != "zcy" {
		t.Fatalf("user name %s", user.Name)
	}

	tmpUid := -1
	err = database.UpdateUserName(tmpUid, "zcy")
	if err == nil {
		t.Fatal(err)
	}
}

func TestUpdatePassword(t *testing.T) {
	err := database.UpdatePassword(uid, hash("abcdefg"), hash("123456"))
	if err != nil {
		t.Fatal(err)
	}
	user := database.GetUserById(uid)
	if user == nil {
		t.Fail()
		return
	}
	if user.PassWord != hash("abcdefg") {
		t.Fatalf("user password %s", user.PassWord)
	}

	err = database.UpdatePassword(uid, hash("abcdefg"), hash("123456"))
	if err == nil {
		t.Fatal(err)
	}
}

func TestLogOffUser(t *testing.T) {
	err := database.LogOffUser(uid)
	if err != nil {
		t.Fatal(err)
	}

	user := database.GetUserById(uid)
	if user != nil {
		t.Fail()
		return
	}

	err = database.LogOffUser(uid)
	if err == nil {
		t.Fatalf("用户%d第二次删除成功！", uid)
	} else {
		fmt.Printf("用户%d第二次删除失败：%s", uid, err)
	}
}

// go test -v ./post/database/xorm -run=^TestRegistUser$ -count=1
// go test -v ./post/database/xorm -run=^TestGetUserById$ -count=1
// go test -v ./post/database/xorm -run=^TestGetUserByName$ -count=1
// go test -v ./post/database/xorm -run=^TestUpdateUserName$ -count=1
// go test -v ./post/database/xorm -run=^TestUpdatePassword$ -count=1
// go test -v ./post/database/xorm -run=^TestLogOffUser$ -count=1
