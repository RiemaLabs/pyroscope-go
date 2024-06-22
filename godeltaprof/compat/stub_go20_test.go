//go:build go1.16 && !go1.21
// +build go1.16,!go1.21

package compat

import "testing"

func TestRuntimeFrameSymbolName(t *testing.T) {
	checkSignature(t, "github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof",
		"runtime_FrameSymbolName",
		"func github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof.runtime_FrameSymbolName(f *runtime.Frame) string")
}

func TestRuntimeFrameStartLine(t *testing.T) {
	checkSignature(t, "github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof",
		"runtime_FrameStartLine",
		"func github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof.runtime_FrameStartLine(f *runtime.Frame) int")
}

func TestRuntimeExpandFinalInlineFrame(t *testing.T) {
	checkSignature(t, "runtime/pprof",
		"runtime_expandFinalInlineFrame",
		"func runtime/pprof.runtime_expandFinalInlineFrame(stk []uintptr) []uintptr")
	checkSignature(t, "github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof",
		"runtime_expandFinalInlineFrame",
		"func github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof.runtime_expandFinalInlineFrame(stk []uintptr) []uintptr")
}

func TestRuntimeCyclesPerSecond(t *testing.T) {
	checkSignature(t, "runtime/pprof",
		"runtime_cyclesPerSecond",
		"func runtime/pprof.runtime_cyclesPerSecond() int64")
	checkSignature(t, "github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof",
		"runtime_cyclesPerSecond",
		"func github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof.runtime_cyclesPerSecond() int64")
}
