// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package main

import (
	"log"

	"github.com/google/note-maps/otgen"
)

func main() {
	err := otgen.Generate(otgen.Options{
		PackageName: "note",
		ElementType: "ID",
		SliceType:   "IDSlice",
		BaseName:    "id_ot",
		OpStringer:  true,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
