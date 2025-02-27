// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"io"
	"os"
	"sync"

	"github.com/nginxinc/nginx-service-mesh/internal/nginx-meshctl/support"
)

type FakeFileWriter struct {
	CloseStub        func(*os.File) error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
		arg1 *os.File
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	MkdirStub        func(string) error
	mkdirMutex       sync.RWMutex
	mkdirArgsForCall []struct {
		arg1 string
	}
	mkdirReturns struct {
		result1 error
	}
	mkdirReturnsOnCall map[int]struct {
		result1 error
	}
	MkdirAllStub        func(string) error
	mkdirAllMutex       sync.RWMutex
	mkdirAllArgsForCall []struct {
		arg1 string
	}
	mkdirAllReturns struct {
		result1 error
	}
	mkdirAllReturnsOnCall map[int]struct {
		result1 error
	}
	OpenFileStub        func(string) (*os.File, error)
	openFileMutex       sync.RWMutex
	openFileArgsForCall []struct {
		arg1 string
	}
	openFileReturns struct {
		result1 *os.File
		result2 error
	}
	openFileReturnsOnCall map[int]struct {
		result1 *os.File
		result2 error
	}
	RemoveAllStub        func(string) error
	removeAllMutex       sync.RWMutex
	removeAllArgsForCall []struct {
		arg1 string
	}
	removeAllReturns struct {
		result1 error
	}
	removeAllReturnsOnCall map[int]struct {
		result1 error
	}
	TempDirStub        func(string) (string, error)
	tempDirMutex       sync.RWMutex
	tempDirArgsForCall []struct {
		arg1 string
	}
	tempDirReturns struct {
		result1 string
		result2 error
	}
	tempDirReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	WriteStub        func(string, string) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 string
		arg2 string
	}
	writeReturns struct {
		result1 error
	}
	writeReturnsOnCall map[int]struct {
		result1 error
	}
	WriteFromReaderStub        func(string, io.ReadCloser) error
	writeFromReaderMutex       sync.RWMutex
	writeFromReaderArgsForCall []struct {
		arg1 string
		arg2 io.ReadCloser
	}
	writeFromReaderReturns struct {
		result1 error
	}
	writeFromReaderReturnsOnCall map[int]struct {
		result1 error
	}
	WriteTarFileStub        func(string, string) error
	writeTarFileMutex       sync.RWMutex
	writeTarFileArgsForCall []struct {
		arg1 string
		arg2 string
	}
	writeTarFileReturns struct {
		result1 error
	}
	writeTarFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFileWriter) Close(arg1 *os.File) error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
		arg1 *os.File
	}{arg1})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{arg1})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileWriter) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeFileWriter) CloseCalls(stub func(*os.File) error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeFileWriter) CloseArgsForCall(i int) *os.File {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	argsForCall := fake.closeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFileWriter) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) Mkdir(arg1 string) error {
	fake.mkdirMutex.Lock()
	ret, specificReturn := fake.mkdirReturnsOnCall[len(fake.mkdirArgsForCall)]
	fake.mkdirArgsForCall = append(fake.mkdirArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.MkdirStub
	fakeReturns := fake.mkdirReturns
	fake.recordInvocation("Mkdir", []interface{}{arg1})
	fake.mkdirMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileWriter) MkdirCallCount() int {
	fake.mkdirMutex.RLock()
	defer fake.mkdirMutex.RUnlock()
	return len(fake.mkdirArgsForCall)
}

func (fake *FakeFileWriter) MkdirCalls(stub func(string) error) {
	fake.mkdirMutex.Lock()
	defer fake.mkdirMutex.Unlock()
	fake.MkdirStub = stub
}

func (fake *FakeFileWriter) MkdirArgsForCall(i int) string {
	fake.mkdirMutex.RLock()
	defer fake.mkdirMutex.RUnlock()
	argsForCall := fake.mkdirArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFileWriter) MkdirReturns(result1 error) {
	fake.mkdirMutex.Lock()
	defer fake.mkdirMutex.Unlock()
	fake.MkdirStub = nil
	fake.mkdirReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) MkdirReturnsOnCall(i int, result1 error) {
	fake.mkdirMutex.Lock()
	defer fake.mkdirMutex.Unlock()
	fake.MkdirStub = nil
	if fake.mkdirReturnsOnCall == nil {
		fake.mkdirReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mkdirReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) MkdirAll(arg1 string) error {
	fake.mkdirAllMutex.Lock()
	ret, specificReturn := fake.mkdirAllReturnsOnCall[len(fake.mkdirAllArgsForCall)]
	fake.mkdirAllArgsForCall = append(fake.mkdirAllArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.MkdirAllStub
	fakeReturns := fake.mkdirAllReturns
	fake.recordInvocation("MkdirAll", []interface{}{arg1})
	fake.mkdirAllMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileWriter) MkdirAllCallCount() int {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	return len(fake.mkdirAllArgsForCall)
}

func (fake *FakeFileWriter) MkdirAllCalls(stub func(string) error) {
	fake.mkdirAllMutex.Lock()
	defer fake.mkdirAllMutex.Unlock()
	fake.MkdirAllStub = stub
}

func (fake *FakeFileWriter) MkdirAllArgsForCall(i int) string {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	argsForCall := fake.mkdirAllArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFileWriter) MkdirAllReturns(result1 error) {
	fake.mkdirAllMutex.Lock()
	defer fake.mkdirAllMutex.Unlock()
	fake.MkdirAllStub = nil
	fake.mkdirAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) MkdirAllReturnsOnCall(i int, result1 error) {
	fake.mkdirAllMutex.Lock()
	defer fake.mkdirAllMutex.Unlock()
	fake.MkdirAllStub = nil
	if fake.mkdirAllReturnsOnCall == nil {
		fake.mkdirAllReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mkdirAllReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) OpenFile(arg1 string) (*os.File, error) {
	fake.openFileMutex.Lock()
	ret, specificReturn := fake.openFileReturnsOnCall[len(fake.openFileArgsForCall)]
	fake.openFileArgsForCall = append(fake.openFileArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.OpenFileStub
	fakeReturns := fake.openFileReturns
	fake.recordInvocation("OpenFile", []interface{}{arg1})
	fake.openFileMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFileWriter) OpenFileCallCount() int {
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	return len(fake.openFileArgsForCall)
}

func (fake *FakeFileWriter) OpenFileCalls(stub func(string) (*os.File, error)) {
	fake.openFileMutex.Lock()
	defer fake.openFileMutex.Unlock()
	fake.OpenFileStub = stub
}

func (fake *FakeFileWriter) OpenFileArgsForCall(i int) string {
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	argsForCall := fake.openFileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFileWriter) OpenFileReturns(result1 *os.File, result2 error) {
	fake.openFileMutex.Lock()
	defer fake.openFileMutex.Unlock()
	fake.OpenFileStub = nil
	fake.openFileReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeFileWriter) OpenFileReturnsOnCall(i int, result1 *os.File, result2 error) {
	fake.openFileMutex.Lock()
	defer fake.openFileMutex.Unlock()
	fake.OpenFileStub = nil
	if fake.openFileReturnsOnCall == nil {
		fake.openFileReturnsOnCall = make(map[int]struct {
			result1 *os.File
			result2 error
		})
	}
	fake.openFileReturnsOnCall[i] = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeFileWriter) RemoveAll(arg1 string) error {
	fake.removeAllMutex.Lock()
	ret, specificReturn := fake.removeAllReturnsOnCall[len(fake.removeAllArgsForCall)]
	fake.removeAllArgsForCall = append(fake.removeAllArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RemoveAllStub
	fakeReturns := fake.removeAllReturns
	fake.recordInvocation("RemoveAll", []interface{}{arg1})
	fake.removeAllMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileWriter) RemoveAllCallCount() int {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	return len(fake.removeAllArgsForCall)
}

func (fake *FakeFileWriter) RemoveAllCalls(stub func(string) error) {
	fake.removeAllMutex.Lock()
	defer fake.removeAllMutex.Unlock()
	fake.RemoveAllStub = stub
}

func (fake *FakeFileWriter) RemoveAllArgsForCall(i int) string {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	argsForCall := fake.removeAllArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFileWriter) RemoveAllReturns(result1 error) {
	fake.removeAllMutex.Lock()
	defer fake.removeAllMutex.Unlock()
	fake.RemoveAllStub = nil
	fake.removeAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) RemoveAllReturnsOnCall(i int, result1 error) {
	fake.removeAllMutex.Lock()
	defer fake.removeAllMutex.Unlock()
	fake.RemoveAllStub = nil
	if fake.removeAllReturnsOnCall == nil {
		fake.removeAllReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeAllReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) TempDir(arg1 string) (string, error) {
	fake.tempDirMutex.Lock()
	ret, specificReturn := fake.tempDirReturnsOnCall[len(fake.tempDirArgsForCall)]
	fake.tempDirArgsForCall = append(fake.tempDirArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.TempDirStub
	fakeReturns := fake.tempDirReturns
	fake.recordInvocation("TempDir", []interface{}{arg1})
	fake.tempDirMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFileWriter) TempDirCallCount() int {
	fake.tempDirMutex.RLock()
	defer fake.tempDirMutex.RUnlock()
	return len(fake.tempDirArgsForCall)
}

func (fake *FakeFileWriter) TempDirCalls(stub func(string) (string, error)) {
	fake.tempDirMutex.Lock()
	defer fake.tempDirMutex.Unlock()
	fake.TempDirStub = stub
}

func (fake *FakeFileWriter) TempDirArgsForCall(i int) string {
	fake.tempDirMutex.RLock()
	defer fake.tempDirMutex.RUnlock()
	argsForCall := fake.tempDirArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFileWriter) TempDirReturns(result1 string, result2 error) {
	fake.tempDirMutex.Lock()
	defer fake.tempDirMutex.Unlock()
	fake.TempDirStub = nil
	fake.tempDirReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFileWriter) TempDirReturnsOnCall(i int, result1 string, result2 error) {
	fake.tempDirMutex.Lock()
	defer fake.tempDirMutex.Unlock()
	fake.TempDirStub = nil
	if fake.tempDirReturnsOnCall == nil {
		fake.tempDirReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.tempDirReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFileWriter) Write(arg1 string, arg2 string) error {
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.WriteStub
	fakeReturns := fake.writeReturns
	fake.recordInvocation("Write", []interface{}{arg1, arg2})
	fake.writeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileWriter) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeFileWriter) WriteCalls(stub func(string, string) error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *FakeFileWriter) WriteArgsForCall(i int) (string, string) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFileWriter) WriteReturns(result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) WriteReturnsOnCall(i int, result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) WriteFromReader(arg1 string, arg2 io.ReadCloser) error {
	fake.writeFromReaderMutex.Lock()
	ret, specificReturn := fake.writeFromReaderReturnsOnCall[len(fake.writeFromReaderArgsForCall)]
	fake.writeFromReaderArgsForCall = append(fake.writeFromReaderArgsForCall, struct {
		arg1 string
		arg2 io.ReadCloser
	}{arg1, arg2})
	stub := fake.WriteFromReaderStub
	fakeReturns := fake.writeFromReaderReturns
	fake.recordInvocation("WriteFromReader", []interface{}{arg1, arg2})
	fake.writeFromReaderMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileWriter) WriteFromReaderCallCount() int {
	fake.writeFromReaderMutex.RLock()
	defer fake.writeFromReaderMutex.RUnlock()
	return len(fake.writeFromReaderArgsForCall)
}

func (fake *FakeFileWriter) WriteFromReaderCalls(stub func(string, io.ReadCloser) error) {
	fake.writeFromReaderMutex.Lock()
	defer fake.writeFromReaderMutex.Unlock()
	fake.WriteFromReaderStub = stub
}

func (fake *FakeFileWriter) WriteFromReaderArgsForCall(i int) (string, io.ReadCloser) {
	fake.writeFromReaderMutex.RLock()
	defer fake.writeFromReaderMutex.RUnlock()
	argsForCall := fake.writeFromReaderArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFileWriter) WriteFromReaderReturns(result1 error) {
	fake.writeFromReaderMutex.Lock()
	defer fake.writeFromReaderMutex.Unlock()
	fake.WriteFromReaderStub = nil
	fake.writeFromReaderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) WriteFromReaderReturnsOnCall(i int, result1 error) {
	fake.writeFromReaderMutex.Lock()
	defer fake.writeFromReaderMutex.Unlock()
	fake.WriteFromReaderStub = nil
	if fake.writeFromReaderReturnsOnCall == nil {
		fake.writeFromReaderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeFromReaderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) WriteTarFile(arg1 string, arg2 string) error {
	fake.writeTarFileMutex.Lock()
	ret, specificReturn := fake.writeTarFileReturnsOnCall[len(fake.writeTarFileArgsForCall)]
	fake.writeTarFileArgsForCall = append(fake.writeTarFileArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.WriteTarFileStub
	fakeReturns := fake.writeTarFileReturns
	fake.recordInvocation("WriteTarFile", []interface{}{arg1, arg2})
	fake.writeTarFileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileWriter) WriteTarFileCallCount() int {
	fake.writeTarFileMutex.RLock()
	defer fake.writeTarFileMutex.RUnlock()
	return len(fake.writeTarFileArgsForCall)
}

func (fake *FakeFileWriter) WriteTarFileCalls(stub func(string, string) error) {
	fake.writeTarFileMutex.Lock()
	defer fake.writeTarFileMutex.Unlock()
	fake.WriteTarFileStub = stub
}

func (fake *FakeFileWriter) WriteTarFileArgsForCall(i int) (string, string) {
	fake.writeTarFileMutex.RLock()
	defer fake.writeTarFileMutex.RUnlock()
	argsForCall := fake.writeTarFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFileWriter) WriteTarFileReturns(result1 error) {
	fake.writeTarFileMutex.Lock()
	defer fake.writeTarFileMutex.Unlock()
	fake.WriteTarFileStub = nil
	fake.writeTarFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) WriteTarFileReturnsOnCall(i int, result1 error) {
	fake.writeTarFileMutex.Lock()
	defer fake.writeTarFileMutex.Unlock()
	fake.WriteTarFileStub = nil
	if fake.writeTarFileReturnsOnCall == nil {
		fake.writeTarFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeTarFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.mkdirMutex.RLock()
	defer fake.mkdirMutex.RUnlock()
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	fake.tempDirMutex.RLock()
	defer fake.tempDirMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.writeFromReaderMutex.RLock()
	defer fake.writeFromReaderMutex.RUnlock()
	fake.writeTarFileMutex.RLock()
	defer fake.writeTarFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFileWriter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ support.FileWriter = new(FakeFileWriter)
