package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#cgo pkg-config: eolian eina
#include <Eolian.h>
*/
import "C"
import "unsafe"

var (
	allEoFilePaths  []string
	allEotFilePaths []string
	allEoFiles      []string
	allEotFiles     []string
)

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
	if allEoFilePaths != nil {
		return allEoFilePaths
	}
	allEoFilePaths = NewIterator(C.eolian_all_eo_file_paths_get()).StringSlice()
	return allEoFilePaths
}

func AllEotFilePaths() []string {
	if allEotFilePaths != nil {
		return allEotFilePaths
	}
	allEotFilePaths = NewIterator(C.eolian_all_eot_file_paths_get()).StringSlice()
	return allEotFilePaths
}

func AllEoFiles() []string {
	if allEoFiles != nil {
		return allEoFiles
	}
	allEoFiles = NewIterator(C.eolian_all_eo_files_get()).StringSlice()
	return allEoFiles
}

func AllEotFiles() []string {
	if allEotFiles != nil {
		return allEotFiles
	}
	allEotFiles = NewIterator(C.eolian_all_eot_files_get()).StringSlice()
	return allEotFiles
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
