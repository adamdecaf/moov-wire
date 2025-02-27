package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIBeneficiary creates a FIBeneficiary
func mockFIBeneficiary() *FIBeneficiary {
	fib := NewFIBeneficiary()
	fib.FIToFI.LineOne = "Line One"
	fib.FIToFI.LineTwo = "Line Two"
	fib.FIToFI.LineThree = "Line Three"
	fib.FIToFI.LineFour = "Line Four"
	fib.FIToFI.LineFive = "Line Five"
	fib.FIToFI.LineSix = "Line Six"
	return fib
}

// TestMockFIBeneficiary validates mockFIBeneficiary
func TestMockFIBeneficiary(t *testing.T) {
	fib := mockFIBeneficiary()

	require.NoError(t, fib.Validate(), "mockFIBeneficiary does not validate and will break other tests")
}

// TestFIBeneficiaryLineOneAlphaNumeric validates FIBeneficiary LineOne is alphanumeric
func TestFIBeneficiaryLineOneAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineOne = "®"

	err := fib.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrNonAlphanumeric, fib.FIToFI.LineOne).Error())
}

// TestFIBeneficiaryLineTwoAlphaNumeric validates FIBeneficiary LineTwo is alphanumeric
func TestFIBeneficiaryLineTwoAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineTwo = "®"

	err := fib.Validate()

	require.EqualError(t, err, fieldError("LineTwo", ErrNonAlphanumeric, fib.FIToFI.LineTwo).Error())
}

// TestFIBeneficiaryLineThreeAlphaNumeric validates FIBeneficiary LineThree is alphanumeric
func TestFIBeneficiaryLineThreeAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineThree = "®"

	err := fib.Validate()

	require.EqualError(t, err, fieldError("LineThree", ErrNonAlphanumeric, fib.FIToFI.LineThree).Error())
}

// TestFIBeneficiaryLineFourAlphaNumeric validates FIBeneficiary LineFour is alphanumeric
func TestFIBeneficiaryLineFourAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineFour = "®"

	err := fib.Validate()

	require.EqualError(t, err, fieldError("LineFour", ErrNonAlphanumeric, fib.FIToFI.LineFour).Error())
}

// TestFIBeneficiaryLineFiveAlphaNumeric validates FIBeneficiary LineFive is alphanumeric
func TestFIBeneficiaryLineFiveAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineFive = "®"

	err := fib.Validate()

	require.EqualError(t, err, fieldError("LineFive", ErrNonAlphanumeric, fib.FIToFI.LineFive).Error())
}

// TestFIBeneficiaryLineSixAlphaNumeric validates FIBeneficiary LineSix is alphanumeric
func TestFIBeneficiaryLineSixAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineSix = "®"

	err := fib.Validate()

	require.EqualError(t, err, fieldError("LineSix", ErrNonAlphanumeric, fib.FIToFI.LineSix).Error())
}

// TestParseFIBeneficiaryWrongLength parses a wrong FIBeneficiary record length
func TestParseFIBeneficiaryWrongLength(t *testing.T) {
	var line = "{6400}Line Six                                                                                                                                                                                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiary()
	require.EqualError(t, err, r.parseError(fieldError("LineOne", ErrRequireDelimiter)).Error())
}

// TestParseFIBeneficiaryReaderParseError parses a wrong FIBeneficiary reader parse error
func TestParseFIBeneficiaryReaderParseError(t *testing.T) {
	var line = "{6400}Line Si®                                                                                                                                                                                          *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiary()

	expected := r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line Si®")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line Si®")).Error()
	require.EqualError(t, err, expected)
}

// TestFIBeneficiaryTagError validates a FIBeneficiary tag
func TestFIBeneficiaryTagError(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.tag = "{9999}"

	err := fib.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, fib.tag).Error())
}

// TestStringFIBeneficiaryVariableLength parses using variable length
func TestStringFIBeneficiaryVariableLength(t *testing.T) {
	var line = "{6400}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiary()
	require.NoError(t, err)

	line = "{6400}                                                                                                                                                                                                                  NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseFIBeneficiary()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{6400}********"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseFIBeneficiary()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{6400}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseFIBeneficiary()
	require.NoError(t, err)
}

// TestStringFIBeneficiaryOptions validates Format() formatted according to the FormatOptions
func TestStringFIBeneficiaryOptions(t *testing.T) {
	var line = "{6400}*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiary()
	require.NoError(t, err)

	record := r.currentFEDWireMessage.FIBeneficiary
	require.Equal(t, "{6400}                              *                                 *                                 *                                 *                                 *                                 *", record.String())
	require.Equal(t, "{6400}*", record.Format(FormatOptions{VariableLengthFields: true}))
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
