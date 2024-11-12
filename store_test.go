package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathname := CASPathTransformFunc(key)
	expectedOriginalKey := "d9e06924cbe4f7c5f59269e6267f971d02774564"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathname.PathName != expectedPathName {
		t.Errorf("have %s \n want %s", pathname.PathName, expectedPathName)
	}
	if pathname.Filename != expectedPathName {
		t.Errorf("have %s \n want %s", pathname.Filename, expectedOriginalKey)
	}
}

func TestNewStore(t *testing.T) {

	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}
}
