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

package note

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

type nn struct {
	ID
	VS string
	VT GraphNote
	CS []GraphNote
	TS []GraphNote
}

func (n nn) GetID() ID                            { return n.ID }
func (n nn) GetValue() (string, GraphNote, error) { return n.VS, n.VT, nil }
func (n nn) GetContents() ([]GraphNote, error)    { return n.CS, nil }
func (n nn) GetTypes() ([]GraphNote, error)       { return n.TS, nil }

type brokenValue struct{ nn }

func (n brokenValue) GetValue() (string, GraphNote, error) {
	return "", nil, errors.New("brokenValue")
}

type brokenContents struct{ nn }

func (n brokenContents) GetContents() ([]GraphNote, error) {
	return nil, errors.New("brokenContents")
}

func TestTruncateNote(t *testing.T) {
	actual, err := TruncateNote(nn{
		ID: "id",
		VS: "value",
		VT: EmptyNote("vt"),
		CS: []GraphNote{EmptyNote("c0"), EmptyNote("c1")},
		TS: []GraphNote{EmptyNote("t0"), EmptyNote("t1")},
	})
	if err != nil {
		t.Fatal(err)
	}
	expected := TruncatedNote{
		ID:          "id",
		ValueString: "value",
		ValueType:   "vt",
		Contents:    []ID{"c0", "c1"},
		Types:       []ID{"t0", "t1"},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %#v, expected %#v", actual, expected)
	}
}

func TestTruncateNote_errorIfValueBroken(t *testing.T) {
	_, err := TruncateNote(brokenValue{nn{
		ID: "id",
		VS: "value",
		VT: EmptyNote("vt"),
		CS: []GraphNote{EmptyNote("c0"), EmptyNote("c1")},
	}})
	if err == nil || !strings.HasSuffix(err.Error(), "brokenValue") {
		t.Fatal("got", err, "expected brokenValue")
	}
}

func TestTruncateNote_errorIfContentsBroken(t *testing.T) {
	_, err := TruncateNote(brokenContents{nn{
		ID: "id",
		VS: "value",
		VT: EmptyNote("vt"),
		CS: []GraphNote{EmptyNote("c0"), EmptyNote("c1")},
	}})
	if err == nil || !strings.HasSuffix(err.Error(), "brokenContents") {
		t.Fatal("got", err, "expected brokenContents")
	}
}

func TestTruncateNote_Equals(t *testing.T) {
	for _, test := range []struct {
		A, B  TruncatedNote
		Equal bool
	}{
		{TruncatedNote{}, TruncatedNote{}, true},
		{
			TruncatedNote{},
			TruncatedNote{"", "", "", []ID{}, []ID{}},
			true,
		},
		{TruncatedNote{ID: "0"}, TruncatedNote{}, false},
		{TruncatedNote{ID: "0"}, TruncatedNote{ID: "0"}, true},
		{
			TruncatedNote{ValueString: "x"},
			TruncatedNote{ValueString: "y"},
			false,
		},
		{
			TruncatedNote{ValueString: "x"},
			TruncatedNote{ValueString: "x"},
			true,
		},
		{
			TruncatedNote{ValueType: "x"},
			TruncatedNote{ValueType: "y"},
			false,
		},
		{
			TruncatedNote{ValueType: "x"},
			TruncatedNote{ValueType: "x"},
			true,
		},
		{
			TruncatedNote{Contents: []ID{"x"}},
			TruncatedNote{Contents: []ID{"y"}},
			false,
		},
		{
			TruncatedNote{Contents: []ID{"x"}},
			TruncatedNote{Contents: []ID{"x"}},
			true,
		},
	} {
		if test.A.Equals(test.B) != test.B.Equals(test.A) {
			t.Errorf("A.Equals(B) != B.Equals(A) : %v != %v",
				test.A.Equals(test.B), test.B.Equals(test.A))
		}
		if test.A.Equals(test.B) != test.Equal {
			t.Errorf("%#v==%#v got %v, expected %v",
				test.A, test.B, !test.Equal, test.Equal)
		}
	}
}

func TestPatch_mismatchedID(t *testing.T) {
	tn := TruncatedNote{ID: "id0"}
	ops := OperationSlice(nil).SetValue("id1", "vs", "vt").InsertContent("id1", 0, "c0")
	Patch(&tn, []Operation(ops))
	expect := TruncatedNote{ID: "id0"}
	if !reflect.DeepEqual(tn, expect) {
		t.Error("got", tn, "expected", expect)
	}
}

func TestDiffPatch(t *testing.T) {
	for _, test := range []struct {
		Title string
		A, B  TruncatedNote
	}{
		{Title: "empty notes"},
		{Title: "change value string",
			A: TruncatedNote{ValueString: "a"},
			B: TruncatedNote{ValueString: "b"}},
		{Title: "change value type",
			A: TruncatedNote{ValueType: "vt0"},
			B: TruncatedNote{ValueType: "vt1"}},
		{Title: "add content",
			A: TruncatedNote{Contents: []ID{"a"}},
			B: TruncatedNote{Contents: []ID{"a", "b"}}},
		{Title: "remove content",
			A: TruncatedNote{Contents: []ID{"a", "b"}},
			B: TruncatedNote{Contents: []ID{"a"}}},
		{Title: "insert content",
			A: TruncatedNote{Contents: []ID{"a", "b"}},
			B: TruncatedNote{Contents: []ID{"a", "c", "b"}}},
		{Title: "swap content",
			A: TruncatedNote{Contents: []ID{"a", "b"}},
			B: TruncatedNote{Contents: []ID{"b", "a"}}},
	} {
		t.Run(test.Title, func(t *testing.T) {
			ops := Diff(test.A, test.B)
			newB := test.A
			if err := Patch(&newB, ops); err != nil {
				t.Error(err)
			} else if !newB.Equals(test.B) {
				t.Errorf("got %#v, expected %#v, applying %#v to %#v",
					newB, test.B, ops, test.A)
			}
		})
	}
}

func TestExpandNote(t *testing.T) {
	n := ExpandNote(TruncatedNote{
		ID:          "id0",
		ValueString: "vs",
		ValueType:   "vt",
		Contents:    []ID{"c0"},
	}, EmptyLoader)
	if id := n.GetID(); id != "id0" {
		t.Error("got", id, "expected", ID("id0"))
	}
	if vs, vt, err := n.GetValue(); vs != "vs" || vt.GetID() != "vt" || err != nil {
		t.Error("got", vs, vt, err, "expected", "vs", EmptyNote("vt"), nil)
	}
	if cs, err := n.GetContents(); len(cs) != 1 || cs[0].GetID() != "c0" || err != nil {
		t.Error("got", cs, err, "expected", []GraphNote{EmptyNote("c0")}, nil)
	}
}

func TestExpandNote_noDatatype(t *testing.T) {
	n := ExpandNote(TruncatedNote{
		ID:          "id0",
		ValueString: "vs",
	}, EmptyLoader)
	if vs, vt, err := n.GetValue(); vs != "vs" || !vt.GetID().Empty() || err != nil {
		t.Error("got", vs, vt, err, "expected", "vs", EmptyNote("vt"), nil)
	}
}
