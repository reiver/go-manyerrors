package manyerrors


import (
	"testing"

	errs "errors"
	"fmt"
	"math/rand"
	"time"
)


func TestManyErrors(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))


	tests := []struct{
		Errors []error
	}{
		{
			Errors: []error{},
		},
		{
			Errors: []error{
				errs.New("one"),
			},
		},
		{
			Errors: []error{
				errs.New("one"),
				errs.New("two"),
			},
		},
		{
			Errors: []error{
				errs.New("one"),
				errs.New("two"),
				errs.New("three"),
			},
		},
		{
			Errors: []error{
				errs.New("apple banana"),
				errs.New("banana cherry"),
				errs.New("kiwi orange"),
				errs.New("cantaloupe honeydew watermelon"),
				errs.New("lemon mango"),
				errs.New("cherry grape"),
			},
		},
	}


	more := 5 + randomness.Intn(50)
	for i:=0; i<more; i++ {

		numErrors := 1 + randomness.Intn(50)
		slice := make([]error, numErrors)
		for ii:=0; ii<numErrors; ii++ {
			slice[ii] = fmt.Errorf("%d %d %d", randomness.Int63(), randomness.Int63(), randomness.Int63())
		}

		datum := struct{Errors []error}{Errors:slice}
		tests = append(tests, datum)
	}



	for testNumber, test := range tests {
		errs := New(test.Errors...)

		if nil == errs {
			t.Errorf("For test #%d, expected non-nil, but instead got: %v", testNumber, errs)
		}

		actual := errs.Errors()

		if actualLen, expectedLen := len(actual), len(test.Errors); expectedLen != actualLen {
			t.Errorf("For test #%d, expected length to be %d, actually was %d.", testNumber, expectedLen, actualLen)
		}

		for i, expectedErr := range test.Errors {
			if actualErr := actual[i]; expectedErr != actualErr {
				t.Errorf("For test #%d and error #%d, expected err (%T) %q, but actually got (%T) %q.", testNumber, i, expectedErr, expectedErr, actualErr, actualErr)
			}
		}
	}
}
