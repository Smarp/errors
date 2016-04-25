package errors

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error Test:", func() {
	It("create new error", func() {
		err := New(Placebo, nil)
		Expect(err.Type()).To(Equal(Placebo))
		Expect(err.Error()).To(Equal("Placebo"))
	})
	It("create new error with meta", func() {
		err := New(Placebo, map[string]interface{}{"Key": "Val"})
		Expect(err.Type()).To(Equal(Placebo))
		Expect(err.Error()).To(Equal("Placebo{Key: Val}"))
	})
	It("should add error to stack", func() {
		err := New(Placebo, map[string]interface{}{"Key": "Val"}).
			Add(errors.New("Some error")).
			Add(errors.New("Other error"))
		Expect(err.Type()).To(Equal(Placebo))
		Expect(err.Error()).To(Equal("Placebo{Key: Val};Some error;Other error"))
	})
	It("should add our error to stack", func() {
		err := New(Placebo, map[string]interface{}{"Key": "Val"}).
			Add(New(Aspirin, nil)).
			Add(New(Ibuprofen, nil))
		Expect(err.Type()).To(Equal(Placebo))
		Expect(err.Error()).To(Equal("Placebo{Key: Val};Aspirin;Ibuprofen"))
	})
	It("should wrap error to Errors", func() {
		e := errors.New("Some error")
		err := Wrap(e)
		Expect(err.Type()).To(Equal(e))
		Expect(err.Error()).To(Equal("Some error"))
	})
	It("should wrap error to Errors", func() {
		eFunc := func() error {
			return New(Placebo, map[string]interface{}{"Key": "Val"})
		}
		err := Wrap(eFunc())
		Expect(err.Type()).To(Equal(Placebo))
	})
})
