package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestAPIGateway(t *testing.T) {
	invokeURL := execute()

	resp, err := http.Get(invokeURL)
	if err != nil {
		t.Fatalf("Error from Invoke URL GET request: %v", err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf ("Error reading Invoke URL body: %v", err)
	}

	// Check the response body is what we expect.
	expected := "\"Hello World!\""
	if string(responseData) != expected {
		t.Errorf("expected: %v got %v ", expected, string(responseData))
	}
}

func execute() string{
	err := os.Chdir("../")
	if err != nil {
		panic(err)
	}

	out, err := exec.Command("terraform", "output", "api_gateway_invoke_url").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	output := strings.Split(string(out), "\n")
	url := strings.TrimSpace(output[0])
	fmt.Printf("The api gateway invoke url is: %v\n", url)

	return url
}