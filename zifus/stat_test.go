package zifus

import "testing"

func Test_Stat64String(t *testing.T) {
	t.Log(Stat64String(10003,""))
	t.Log(Stat64String(9999,""))
	t.Log(Stat64String(19941999,""))
	t.Log(Stat64String(19943999,""))
	t.Log(Stat64String(19984001,""))
	t.Log(Stat64String(1998000099,""))
	t.Log(Stat64String(1998003399,""))
	t.Log(Stat64String(1940003399,""))
}