package fastly

import "testing"

func TestClient_HealthChecks(t *testing.T) {
	t.Parallel()

	tv := testVersion(t)

	// Create
	hc, err := testClient.CreateHealthCheck(&CreateHealthCheckInput{
		Service:          testServiceID,
		Version:          tv.Number,
		Name:             "test-healthcheck",
		Method:           "HEAD",
		Host:             "example.com",
		Path:             "/foo",
		HTTPVersion:      "1.1",
		Timeout:          1500,
		CheckInterval:    2500,
		ExpectedResponse: 200,
		Window:           5000,
		Threshold:        10,
		Initial:          10,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Ensure deleted
	defer func() {
		testClient.DeleteHealthCheck(&DeleteHealthCheckInput{
			Service: testServiceID,
			Version: tv.Number,
			Name:    "test-healthcheck",
		})

		testClient.DeleteHealthCheck(&DeleteHealthCheckInput{
			Service: testServiceID,
			Version: tv.Number,
			Name:    "new-test-healthcheck",
		})
	}()

	if hc.Name != "test-healthcheck" {
		t.Errorf("bad name: %q", hc.Name)
	}
	if hc.Method != "HEAD" {
		t.Errorf("bad address: %q", hc.Method)
	}
	if hc.Host != "example.com" {
		t.Errorf("bad host: %q", hc.Host)
	}
	if hc.Path != "/foo" {
		t.Errorf("bad path: %q", hc.Path)
	}
	if hc.HTTPVersion != "1.1" {
		t.Errorf("bad http_version: %q", hc.HTTPVersion)
	}
	if hc.Timeout != 1500 {
		t.Errorf("bad timeout: %q", hc.Timeout)
	}
	if hc.CheckInterval != 2500 {
		t.Errorf("bad check_interval: %q", hc.CheckInterval)
	}
	if hc.ExpectedResponse != 200 {
		t.Errorf("bad timeout: %q", hc.ExpectedResponse)
	}
	if hc.Window != 5000 {
		t.Errorf("bad window: %q", hc.Window)
	}
	if hc.Threshold != 10 {
		t.Errorf("bad threshold: %q", hc.Threshold)
	}
	if hc.Initial != 10 {
		t.Errorf("bad initial: %q", hc.Initial)
	}

	// List
	hcs, err := testClient.ListHealthChecks(&ListHealthChecksInput{
		Service: testServiceID,
		Version: tv.Number,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(hcs) < 1 {
		t.Errorf("bad health checks: %v", hcs)
	}

	// Get
	nhc, err := testClient.GetHealthCheck(&GetHealthCheckInput{
		Service: testServiceID,
		Version: tv.Number,
		Name:    "test-healthcheck",
	})
	if err != nil {
		t.Fatal(err)
	}
	if hc.Name != nhc.Name {
		t.Errorf("bad name: %q (%q)", hc.Name, nhc.Name)
	}
	if hc.Method != nhc.Method {
		t.Errorf("bad address: %q", hc.Method)
	}
	if hc.Host != nhc.Host {
		t.Errorf("bad host: %q", hc.Host)
	}
	if hc.Path != nhc.Path {
		t.Errorf("bad path: %q", hc.Path)
	}
	if hc.HTTPVersion != nhc.HTTPVersion {
		t.Errorf("bad http_version: %q", hc.HTTPVersion)
	}
	if hc.Timeout != nhc.Timeout {
		t.Errorf("bad timeout: %q", hc.Timeout)
	}
	if hc.CheckInterval != nhc.CheckInterval {
		t.Errorf("bad check_interval: %q", hc.CheckInterval)
	}
	if hc.ExpectedResponse != nhc.ExpectedResponse {
		t.Errorf("bad timeout: %q", hc.ExpectedResponse)
	}
	if hc.Window != nhc.Window {
		t.Errorf("bad window: %q", hc.Window)
	}
	if hc.Threshold != nhc.Threshold {
		t.Errorf("bad threshold: %q", hc.Threshold)
	}
	if hc.Initial != nhc.Initial {
		t.Errorf("bad initial: %q", hc.Initial)
	}

	// Update
	ub, err := testClient.UpdateHealthCheck(&UpdateHealthCheckInput{
		Service: testServiceID,
		Version: tv.Number,
		Name:    "test-healthcheck",
		NewName: "new-test-healthcheck",
	})
	if err != nil {
		t.Fatal(err)
	}
	if ub.Name != "new-test-healthcheck" {
		t.Errorf("bad name: %q", ub.Name)
	}

	// Delete
	if err := testClient.DeleteHealthCheck(&DeleteHealthCheckInput{
		Service: testServiceID,
		Version: tv.Number,
		Name:    "new-test-healthcheck",
	}); err != nil {
		t.Fatal(err)
	}
}

func TestClient_ListHealthChecks_validation(t *testing.T) {
	var err error
	_, err = testClient.ListHealthChecks(&ListHealthChecksInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.ListHealthChecks(&ListHealthChecksInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_CreateHealthCheck_validation(t *testing.T) {
	var err error
	_, err = testClient.CreateHealthCheck(&CreateHealthCheckInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.CreateHealthCheck(&CreateHealthCheckInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_GetHealthCheck_validation(t *testing.T) {
	var err error
	_, err = testClient.GetHealthCheck(&GetHealthCheckInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.GetHealthCheck(&GetHealthCheckInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.GetHealthCheck(&GetHealthCheckInput{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_UpdateHealthCheck_validation(t *testing.T) {
	var err error
	_, err = testClient.UpdateHealthCheck(&UpdateHealthCheckInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.UpdateHealthCheck(&UpdateHealthCheckInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.UpdateHealthCheck(&UpdateHealthCheckInput{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_DeleteHealthCheck_validation(t *testing.T) {
	var err error
	err = testClient.DeleteHealthCheck(&DeleteHealthCheckInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	err = testClient.DeleteHealthCheck(&DeleteHealthCheckInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	err = testClient.DeleteHealthCheck(&DeleteHealthCheckInput{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}
