// Copyright 2017 the gousb Authors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gousb

import (
	"context"
)

type transferIntf interface {
	submit() error
	cancel() error
	wait(context.Context) (int, error)
	free() error
	data() []byte
}

type stream struct {
	// a fifo of USB transfers.
	transfers chan transferIntf
	// err is the first encountered error, returned to the user.
	err error
	// finished is true if transfers has been already closed.
	finished bool
}

func (s *stream) gotError(err error) { _ = "STUB: not implemented"; return }

func (s *stream) noMore() { _ = "STUB: not implemented"; return }

func (s *stream) submitAll() { _ = "STUB: not implemented"; return }

func (s *stream) flushRemaining() { _ = "STUB: not implemented"; return }

func (s *stream) done() { _ = "STUB: not implemented"; return }

// ReadStream is a buffer that tries to prefetch data from the IN endpoint,
// reducing the latency between subsequent Read()s.
// ReadStream keeps prefetching data until Close() is called or until
// an error is encountered. After Close(), the buffer might still have
// data left from transfers that were initiated before Close. Read()ing
// from the ReadStream will keep returning available data. When no more
// data is left, io.EOF is returned.
type ReadStream struct {
	s *stream
	// current holds the last transfer to return.
	current transferIntf
	// total/used are the number of all/used bytes in the current transfer.
	total, used int
}

// Read reads data from the transfer stream.
// The data will come from at most a single transfer, so the returned number
// might be smaller than the length of p.
// After a non-nil error is returned, all subsequent attempts to read will
// return io.ErrClosedPipe.
// Read cannot be called concurrently with other Read, ReadContext
// or Close.
func (r *ReadStream) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadContext reads data from the transfer stream.
// The data will come from at most a single transfer, so the returned number
// might be smaller than the length of p.
// After a non-nil error is returned, all subsequent attempts to read will
// return io.ErrClosedPipe.
// ReadContext cannot be called concurrently with other Read, ReadContext
// or Close.
// The context passed controls the cancellation of this particular read
// operation within the stream. The semantics is identical to
// Endpoint.ReadContext.
func (r *ReadStream) ReadContext(ctx context.Context, p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// no more transfers in flight

// wait error aborts immediately, all remaining data is invalid.

// guaranteed to not block, len(transfers) == number of allocated transfers

// Close signals that the transfer should stop. After Close is called,
// subsequent Read()s will return data from all transfers that were already
// in progress before returning an io.EOF error, unless another error
// was encountered earlier.
// Close cannot be called concurrently with Read.
func (r *ReadStream) Close() error { _ = "STUB: not implemented"; return nil }

// WriteStream is a buffer that will send data asynchronously, reducing
// the latency between subsequent Write()s.
type WriteStream struct {
	s     *stream
	total int
}

// Write sends the data to the endpoint. Write returning a nil error doesn't
// mean that data was written to the device, only that it was written to the
// buffer. Only a call to Close() that returns nil error guarantees that
// all transfers have succeeded.
// If the slice passed to Write does not align exactly with the transfer
// buffer size (as declared in a call to NewStream), the last USB transfer
// of this Write will be sent with less data than the full buffer.
// After a non-nil error is returned, all subsequent attempts to write will
// return io.ErrClosedPipe.
// If Write encounters an error when preparing the transfer, the stream
// will still try to complete any pending transfers. The total number
// of bytes successfully written can be retrieved through a Written()
// call after Close() has returned.
// Write cannot be called concurrently with another Write, Written or Close.
func (w *WriteStream) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteContext sends the data to the endpoint. Write returning a nil error doesn't
// mean that data was written to the device, only that it was written to the
// buffer. Only a call to Close() that returns nil error guarantees that
// all transfers have succeeded.
// If the slice passed to WriteContext does not align exactly with the transfer
// buffer size (as declared in a call to NewStream), the last USB transfer
// of this Write will be sent with less data than the full buffer.
// After a non-nil error is returned, all subsequent attempts to write will
// return io.ErrClosedPipe.
// If WriteContext encounters an error when preparing the transfer, the stream
// will still try to complete any pending transfers. The total number
// of bytes successfully written can be retrieved through a Written()
// call after Close() has returned.
// WriteContext cannot be called concurrently with another Write, WriteContext,
// Written, Close or CloseContext.
func (w *WriteStream) WriteContext(ctx context.Context, p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// unsubmitted transfers will return 0 bytes and no error

// This branch is used only after all the transfers were set in flight.
// That means all transfers left in the queue are in flight.
// They must be ignored, since this wait() failed.

// Even though this submit failed, all the transfers in flight are still valid.
// Don't flush remaining transfers.
// We won't submit any more transfers.

// guaranteed non blocking

// Close signals end of data to write. Close blocks until all transfers
// that were sent are finished. The error returned by Close is the first
// error encountered during writing the entire stream (if any).
// Close returning nil indicates all transfers completed successfully.
// After Close, the total number of bytes successfully written can be
// retrieved using Written().
// Close may not be called concurrently with Write, Close or Written.
func (w *WriteStream) Close() error { _ = "STUB: not implemented"; return nil }

// CloseContext signals end of data to write. CloseContext blocks until all
// transfers that were sent are finished or until the context is canceled. The
// error returned by CloseContext is the first error encountered during writing
// the entire stream (if any).
// CloseContext returning nil indicates all transfers completed successfully.
// After CloseContext, the total number of bytes successfully written can be
// retrieved using Written().
// CloseContext may not be called concurrently with Write, WriteContext, Close,
// CloseContext or Written.
func (w *WriteStream) CloseContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Written returns the number of bytes successfully written by the stream.
// Written may be called only after Close() or CloseContext()
// has been called and returned.
func (w *WriteStream) Written() int { _ = "STUB: not implemented"; return 0 }

func newStream(tt []transferIntf) *stream { _ = "STUB: not implemented"; return nil }
