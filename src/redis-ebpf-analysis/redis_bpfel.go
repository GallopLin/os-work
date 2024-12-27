// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type redisL7Event struct {
	Fd                  uint64
	WriteTimeNs         uint64
	Pid                 uint32
	Status              uint32
	Duration            uint64
	Protocol            uint8
	Method              uint8
	Padding             uint16
	Payload             [1024]uint8
	PayloadSize         uint32
	PayloadReadComplete uint8
	Failed              uint8
	IsTls               uint8
	_                   [1]byte
	Seq                 uint32
	Tid                 uint32
	_                   [4]byte
}

type redisL7Request struct {
	WriteTimeNs         uint64
	Protocol            uint8
	Method              uint8
	Payload             [1024]uint8
	_                   [2]byte
	PayloadSize         uint32
	PayloadReadComplete uint8
	RequestType         uint8
	_                   [2]byte
	Seq                 uint32
	Tid                 uint32
	_                   [4]byte
}

type redisSocketKey struct {
	Fd    uint64
	Pid   uint32
	IsTls uint8
	_     [3]byte
}

// loadRedis returns the embedded CollectionSpec for redis.
func loadRedis() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_RedisBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load redis: %w", err)
	}

	return spec, err
}

// loadRedisObjects loads redis and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*redisObjects
//	*redisPrograms
//	*redisMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadRedisObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadRedis()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// redisSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type redisSpecs struct {
	redisProgramSpecs
	redisMapSpecs
}

// redisSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type redisProgramSpecs struct {
	HandleRead      *ebpf.ProgramSpec `ebpf:"handle_read"`
	HandleReadExit  *ebpf.ProgramSpec `ebpf:"handle_read_exit"`
	HandleWrite     *ebpf.ProgramSpec `ebpf:"handle_write"`
	HandleWriteExit *ebpf.ProgramSpec `ebpf:"handle_write_exit"`
}

// redisMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type redisMapSpecs struct {
	ActiveL7Requests *ebpf.MapSpec `ebpf:"active_l7_requests"`
	ActiveReads      *ebpf.MapSpec `ebpf:"active_reads"`
	ActiveWrites     *ebpf.MapSpec `ebpf:"active_writes"`
	L7EventHeap      *ebpf.MapSpec `ebpf:"l7_event_heap"`
	L7Events         *ebpf.MapSpec `ebpf:"l7_events"`
	L7RequestHeap    *ebpf.MapSpec `ebpf:"l7_request_heap"`
}

// redisObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadRedisObjects or ebpf.CollectionSpec.LoadAndAssign.
type redisObjects struct {
	redisPrograms
	redisMaps
}

func (o *redisObjects) Close() error {
	return _RedisClose(
		&o.redisPrograms,
		&o.redisMaps,
	)
}

// redisMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadRedisObjects or ebpf.CollectionSpec.LoadAndAssign.
type redisMaps struct {
	ActiveL7Requests *ebpf.Map `ebpf:"active_l7_requests"`
	ActiveReads      *ebpf.Map `ebpf:"active_reads"`
	ActiveWrites     *ebpf.Map `ebpf:"active_writes"`
	L7EventHeap      *ebpf.Map `ebpf:"l7_event_heap"`
	L7Events         *ebpf.Map `ebpf:"l7_events"`
	L7RequestHeap    *ebpf.Map `ebpf:"l7_request_heap"`
}

func (m *redisMaps) Close() error {
	return _RedisClose(
		m.ActiveL7Requests,
		m.ActiveReads,
		m.ActiveWrites,
		m.L7EventHeap,
		m.L7Events,
		m.L7RequestHeap,
	)
}

// redisPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadRedisObjects or ebpf.CollectionSpec.LoadAndAssign.
type redisPrograms struct {
	HandleRead      *ebpf.Program `ebpf:"handle_read"`
	HandleReadExit  *ebpf.Program `ebpf:"handle_read_exit"`
	HandleWrite     *ebpf.Program `ebpf:"handle_write"`
	HandleWriteExit *ebpf.Program `ebpf:"handle_write_exit"`
}

func (p *redisPrograms) Close() error {
	return _RedisClose(
		p.HandleRead,
		p.HandleReadExit,
		p.HandleWrite,
		p.HandleWriteExit,
	)
}

func _RedisClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed redis_bpfel.o
var _RedisBytes []byte
