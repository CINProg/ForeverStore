package main

import (
	"bytes"
	"io/ioutil"
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

func TestStoreDeleteKey(t *testing.T) {

	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momsspecials"
	data := []byte("some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	//if s.Has("momsspecials")
	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestNewStore(t *testing.T) {

	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momsspecials"
	data := []byte("some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.readStream(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := ioutil.ReadAll(r)
	if string(b) != string(data) {
		t.Errorf("have %s \n want %s", b, data)
	}
}
