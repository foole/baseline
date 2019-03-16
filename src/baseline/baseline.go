package baseline

import (
    //"fmt"
    //"os"
    //"path"
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

// Given the root artifact directory, recurse through it and find all the
// files in that directory. Return an array of absolute paths to the
// artifacts.
func FindArtifacts(testArtifactDir string) []string {
    panic("Not yet implemented)")
}

// Given a list of tests/directories, find all the executable tests and return
// a new list of those tests. If a directory is encountered, we will recurse
// through that directory to find all the tests that match the aforementioned
// criteria
func FindTests(testList []string) []string {
    panic("Not yet implemented")
    //tests := []string
    //for _, path := range testList {
    //    fs, err = os.Stat(path)
    //    if err != nil {
    //        fmt.Println(err)
    //        return
    //    }
    //    switch mode := fs.Mode(); {
    //    case mode.isDir():
    //        // find all files in directory and call FindTests
    //    case mode.IsRegular():
    //        // if executable, add it to list of tests
    //    case mode&os.ModeSymlink != 0:
    //        // attempt to find a file or dir for this symlink and call
    //        // FindTests on that
    //    default:
    //        fmt.Println(path, "is not a file or directory; will not add to tests")
    //    }
    //}
}

// Shell out and run a single test. Collect STDOUT, STDERR, diff, and
// artifacts. Return a struct with the following:
//   - name of test
//   - raw STDOUT
//   - raw STDERR
//   - inline diff
//   - array of artifacts (these will be files on the filesystem)
//   - test status:
//     - pass
//     - fail
//     - no baseline
//     - not executable
func (b *Baseline) RunTest() string {
    panic("Not yet implemented")
}

// Truncate the path to the test to 50 characters. The test is truncated from
// the end of the string (since that's where the filename is)
func TruncateTestName(test string) string {
    if len(test) <= 50 {
        return test
    } else {
        return test[len(test) - 50:len(test)]
    }
}
