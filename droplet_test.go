package main

import (
	"archive/tar"
	"io/ioutil"
	"testing"
)

func TestInspectTGZ(t *testing.T) {
	p := &inspectPlugin{out: ioutil.Discard}
	headers, err := p.inspectTGZ("testdata/droplet-staticapp-20170612-085112.tgz")
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		index int
		typ   byte
		name  string
	}{
		{1, tar.TypeReg, "./staging_info.yml"},
		{7, tar.TypeDir, "./app/nginx/logs/"},
		{18, tar.TypeReg, "./app/.profile.d/staticfile.sh"},
	}

	for _, test := range tests {
		h := headers[test.index]
		if h.Typeflag != test.typ {
			t.Errorf("invalid type; want %v, got %v", test.typ, h.Typeflag)
		}
		if h.Name != test.name {
			t.Errorf("invalid name; want %v, got %v", test.name, h.Name)
		}
	}
}
