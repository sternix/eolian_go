package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#cgo pkg-config: eolian eina
#include <Eolian.h>
*/
import "C"
import "unsafe"

func init() {
	C.eolian_init()
}

func Shutdown() {
	C.eolian_shutdown()
}

func ParseFile(filePath string) bool {
	cfilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cfilePath))
	return C.eolian_file_parse(cfilePath) == ETrue
}

func AllEoFilePaths() []string {
	return NewIterator(C.eolian_all_eo_file_paths_get()).StringSlice()
}

func AllEotFilePaths() []string {
	return NewIterator(C.eolian_all_eot_file_paths_get()).StringSlice()
}

func AllEoFiles() []string {
	return NewIterator(C.eolian_all_eo_files_get()).StringSlice()
}

func AllEotFiles() []string {
	return NewIterator(C.eolian_all_eot_files_get()).StringSlice()
}

func ParseAllEoFiles() bool {
	return C.eolian_all_eo_files_parse() == ETrue
}

func ParseAllEotFiles() bool {
	return C.eolian_all_eot_files_parse() == ETrue
}

func ScanDirectory(dir string) bool {
	cdir := C.CString(dir)
	defer C.free(unsafe.Pointer(cdir))
	return C.eolian_directory_scan(cdir) == ETrue
}

func ScanSystemDirectory() bool {
	return C.eolian_system_directory_scan() == ETrue
}

func ValidateDatabase(b bool) bool {
	return C.eolian_database_validate(EBool(b)) == ETrue
}
