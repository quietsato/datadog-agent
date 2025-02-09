// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build windows

package probe

/*
 * the below represent the REG_DWORD in the registry for the dump type that's
 * currently configured.  Types are not explicitly documented.  These are
 * discovered via combination of helpful web searches and trial & error.
 *
 * the numbers with explicit comments are validated by trial and error.
 * remainder found here under the table "Value of CrashDumpEnabled"
 * https://crashdmp.wordpress.com/crash-mechanism/configuration/
 *
 */
const (
	DumpTypeUnknown      = int(-1)
	DumpTypeNone         = int(0) // none
	DumpTypeFull         = int(1) // complete, active
	DumpTypeSummary      = int(2) // kernel
	DumpTypeHeader       = int(3) // small
	DumpTypeTriage       = int(4)
	DumpTypeBitmapFull   = int(5)
	DumpTypeBitmapKernel = int(6)
	DumpTypeAutomatic    = int(7) // automatic
)

const (
	// WinCrashStatusCodeUnknown indicates an invalid or corrupted code.
	WinCrashStatusCodeUnknown = int(-1)

	// WinCrashStatusCodeSuccess indicates that crash dump processing succeeded
	// or no crash dump was found.
	WinCrashStatusCodeSuccess = int(0)

	// WinCrashStatusCodeBusy indicates that crash dump processing is still busy
	// and no result is yet available.
	WinCrashStatusCodeBusy = int(1)

	// WinCrashStatusCodeFailed indicates that crash dump processing failed or had an error.
	WinCrashStatusCodeFailed = int(2)
)

// WinCrashStatus defines all of the information returned from the system
// probe to the caller
type WinCrashStatus struct {
	StatusCode int    `json:"statuscode"`
	ErrString  string `json:"errstring"`
	FileName   string `json:"filename"`
	Type       int    `json:"dumptype"`
	DateString string `json:"datestring"`
	Offender   string `json:"offender"`
	BugCheck   string `json:"buckcheckcode"`
}
