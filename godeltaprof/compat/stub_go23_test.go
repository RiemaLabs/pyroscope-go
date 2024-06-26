//go:build go1.23
// +build go1.23

package compat

import "testing"

func TestRuntimeFrameSymbolName(t *testing.T) {
	checkSignature(t, "runtime/pprof",
		"runtime_FrameSymbolName",
		"func runtime/pprof.runtime_FrameSymbolName(f *runtime.Frame) string")
	checkSignature(t, "github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof",
		"runtime_FrameSymbolName",
		"func github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof.runtime_FrameSymbolName(f *runtime.Frame) string")
}

func TestRuntimeFrameStartLine(t *testing.T) {
	checkSignature(t, "runtime/pprof",
		"runtime_FrameStartLine",
		"func runtime/pprof.runtime_FrameStartLine(f *runtime.Frame) int")
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
	checkSignature(t, "runtime",
		"pprof_cyclesPerSecond",
		"func runtime.pprof_cyclesPerSecond() int64")
	checkSignature(t, "runtime/pprof",
		"pprof_cyclesPerSecond",
		"func runtime/pprof.pprof_cyclesPerSecond() int64")
	checkSignature(t, "github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof",
		"runtime_cyclesPerSecond",
		"func github.com/RiemaLabs/pyroscope-go/godeltaprof/external/pprof.runtime_cyclesPerSecond() int64")
}
