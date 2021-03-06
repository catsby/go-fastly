package fastly

import "testing"

func TestClient_Wordpresses(t *testing.T) {
	t.Parallel()

	tv := testVersion(t)

	// Create
	wp, err := testClient.CreateWordpress(&CreateWordpressInput{
		Service: testServiceID,
		Version: tv.Number,
		Name:    "test-wordpress",
		Path:    "/foo",
		Comment: "comment",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Ensure deleted
	defer func() {
		testClient.DeleteWordpress(&DeleteWordpressInput{
			Service: testServiceID,
			Version: tv.Number,
			Name:    "test-wordpress",
		})

		testClient.DeleteWordpress(&DeleteWordpressInput{
			Service: testServiceID,
			Version: tv.Number,
			Name:    "new-test-wordpress",
		})
	}()

	if wp.Name != "test-wordpress" {
		t.Errorf("bad name: %q", wp.Name)
	}
	if wp.Path != "/foo" {
		t.Errorf("bad path: %q", wp.Path)
	}
	if wp.Comment != "comment" {
		t.Errorf("bad port: %q", wp.Comment)
	}

	// List
	wps, err := testClient.ListWordpresses(&ListWordpressesInput{
		Service: testServiceID,
		Version: tv.Number,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(wps) < 1 {
		t.Errorf("bad wordpresss: %v", wps)
	}

	// Get
	nwp, err := testClient.GetWordpress(&GetWordpressInput{
		Service: testServiceID,
		Version: tv.Number,
		Name:    "test-wordpress",
	})
	if err != nil {
		t.Fatal(err)
	}
	if wp.Name != nwp.Name {
		t.Errorf("bad name: %q", wp.Name)
	}
	if wp.Path != nwp.Path {
		t.Errorf("bad path: %q", wp.Path)
	}
	if wp.Comment != nwp.Comment {
		t.Errorf("bad port: %q", wp.Comment)
	}

	// Update
	uwp, err := testClient.UpdateWordpress(&UpdateWordpressInput{
		Service: testServiceID,
		Version: tv.Number,
		Name:    "test-wordpress",
		NewName: "new-test-wordpress",
	})
	if err != nil {
		t.Fatal(err)
	}
	if uwp.Name != "new-test-wordpress" {
		t.Errorf("bad name: %q", uwp.Name)
	}

	// Delete
	if err := testClient.DeleteWordpress(&DeleteWordpressInput{
		Service: testServiceID,
		Version: tv.Number,
		Name:    "new-test-wordpress",
	}); err != nil {
		t.Fatal(err)
	}
}

func TestClient_ListWordpresses_validation(t *testing.T) {
	var err error
	_, err = testClient.ListWordpresses(&ListWordpressesInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.ListWordpresses(&ListWordpressesInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_CreateWordpress_validation(t *testing.T) {
	var err error
	_, err = testClient.CreateWordpress(&CreateWordpressInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.CreateWordpress(&CreateWordpressInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_GetWordpress_validation(t *testing.T) {
	var err error
	_, err = testClient.GetWordpress(&GetWordpressInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.GetWordpress(&GetWordpressInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.GetWordpress(&GetWordpressInput{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_UpdateWordpress_validation(t *testing.T) {
	var err error
	_, err = testClient.UpdateWordpress(&UpdateWordpressInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.UpdateWordpress(&UpdateWordpressInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.UpdateWordpress(&UpdateWordpressInput{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_DeleteWordpress_validation(t *testing.T) {
	var err error
	err = testClient.DeleteWordpress(&DeleteWordpressInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	err = testClient.DeleteWordpress(&DeleteWordpressInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	err = testClient.DeleteWordpress(&DeleteWordpressInput{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}
