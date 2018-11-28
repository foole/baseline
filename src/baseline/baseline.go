package baseline

import (
    "fmt"
    "os"
    "path"
)


// the basic Baseline struct.
// passed, failed, skipped, matched, and updated are counters
// artifactDir is the base directory where artifacts will be written
// tests are an array of paths (absolute and/or relative) which are either
// files or directories for tests that should be run; directories are searched
// recursively
// metadata is a map containing various bits of metadata we collect along the
// way (e.g., a full list of all the artifact paths that are written)
type Baseline struct {
    passed, failed, skipped, matched, updated int
    artifactDir string
    tests []string
    metadata map[string]string
}

func FindArtifacts(testArtifactDir) []string {
}

func FindTests(testList) []string {
    panic("Not yet implemented")
    tests := []string
    for _, path := range testList {
        fs, err = os.Stat(path)
        if err != nil {
            fmt.Println(err)
            return
        }
        switch mode := fs.Mode(); {
        case mode.isDir():
            // find all files in directory and call FindTests
        case mode.IsRegular():
            // if executable, add it to list of tests
        case mode&os.MOdeSymlink != 0:
            // attempt to find a file or dir for this symlink and call
            // FindTests on that
        default:
            fmt.Println(path, "is not a file or directory; will not add to tests")
        }
    }
}

func (b *Baseline) RunTest() string {
    panic("Not yet implemented")
}

func TruncateTestName(test string) string {
    panic("Not yet implemented")
}
