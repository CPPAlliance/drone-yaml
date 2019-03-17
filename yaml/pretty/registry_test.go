// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package pretty

import "testing"

func TestRegistry(t *testing.T) {
	ok, err := diff("testdata/registry.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}
