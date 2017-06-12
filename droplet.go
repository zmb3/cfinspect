package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"code.cloudfoundry.org/cli/plugin"
)

// droplet downloads the droplet for the app whose name is identified by args[0].
// It saves the droplet to a timestamped tgz file and outputs a summary of its
// content.
func (p *inspectPlugin) droplet(conn plugin.CliConnection, args []string) error {
	appName := args[0]
	app, err := conn.GetApp(appName)
	if err != nil {
		return fmt.Errorf("cannot find app %s: %v", appName, err)
	}

	query := path.Join("v2", "apps", app.Guid, "droplet", "download")
	tstamp := time.Now().Format("20060102_150405")
	fname := fmt.Sprintf("droplet-%s-%s.tgz", appName, tstamp)

	_, err = conn.CliCommandWithoutTerminalOutput("curl", query, "--output", fname)
	if err != nil {
		return err
	}

	contents, err := p.inspectTGZ("droplet-staticapp-20170612-085112.tgz")
	if err != nil {
		return fmt.Errorf("couldn't read %s: %v", fname, err)
	}

	dump(p.out, contents)
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, "-- Droplet saved to", fname)
	return nil
}

func (p *inspectPlugin) inspectTGZ(file string) ([]*tar.Header, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, fmt.Errorf("couldn't extract tgz: %v", err)
	}
	defer gz.Close()

	t := tar.NewReader(gz)
	var result []*tar.Header
	for {
		header, err := t.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return result, err
		}
		result = append(result, header)
	}
	return result, nil
}

func dump(w io.Writer, contents []*tar.Header) {
	fmt.Fprintf(w, "Type   %-40s [Size]\n", "Name")
	fmt.Fprintf(w, "------------------------------------------------------\n")
	for _, header := range contents {
		switch header.Typeflag {
		case tar.TypeDir:
			fmt.Fprintf(w, "d     %-40s\n", header.Name)
		case tar.TypeReg:
			fmt.Fprintf(w, "f       %-40s %d\n", header.Name, header.Size)
		default:
			fmt.Fprintf(w, "?   [Unknown Entry]\n")
		}
	}
}
