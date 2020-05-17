// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package util

import (
	"fmt"
	"os"
	"testing"
)

func TestGetAllFiles(t *testing.T) {
	type args struct {
		dirPth   string
		spicType string
	}
	tests := []struct {
		name      string
		args      args
		wantFiles []string
		wantErr   bool
	}{
		{
			name: "AllFile",
			args: args{
				dirPth:   "/Users/c/Desktop/glp",
				spicType: "text/plain",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, err := GetAllFiles(tt.args.dirPth, tt.args.spicType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, f := range gotFiles {
				fmt.Println(f)
			}
		})
	}
}

func TestGetFileContentType(t *testing.T) {
	path := "/Users/c/Desktop/glp/归档.zip"
	file, _ := os.Open(path)
	t.Log(GetFileContentType(file))
}
