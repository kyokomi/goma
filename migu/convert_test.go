package migu

import "testing"

func TestToCamelCase(t *testing.T) {
	type testCase struct {
		in string
		out string
	}

	testCases := []testCase{
		{in: "name", out: "Name"},
		{in: "id", out: "ID"},
		{in: "user_id", out: "UserID"},
		{in: "user_id", out: "UserID"},
		{in: "image_url", out: "ImageURL"},
	}

	for _, testCase := range testCases {
		snake := toCamelCase(testCase.in)
		if testCase.out != snake {
			t.Errorf("error %s != %s", testCase.out, snake)
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	type testCase struct {
		in string
		out string
	}

	testCases := []testCase{
		{in: "Name", out: "name"},
		{in: "ID", out: "id"},
		{in: "UserID", out: "user_id"},
		{in: "userID", out: "user_id"},
		{in: "ImageURL", out: "image_url"},
	}

	for _, testCase := range testCases {
		snake := toSnakeCase(testCase.in)
		if testCase.out != snake {
			t.Errorf("error %s != %s", testCase.out, snake)
		}
	}
}
