package baseline

import (
    "fmt"
    "os"
    //"path"
)


// the basic TestResult struct.
// baseline     path to the baseline
// status       can be one of: passed, failed, skipped, updated, matched
// artifactDir  directory where the artifacts from the test are written
// stdout       path to the stdout from running test (in artifactDir)
// stderr       path to the stderr from running test (in artifactDir)
// diff         path to the diff of baseline and stdout (in artifactDir)
// test         path to the test that was run
// runtime      runtime of the test
type TestResult struct {
    baseline string
    status string
    artifactDir string
    stdout string
    stderr string
    diff string
    test string
    runtime string
}

// Given a path, walk it, find all the files, and return the list of all files
// as a slice of strings
func FindFiles(path string) []string {
    var files []string
    err := filepath.Walk(path func(p string i os.FileInfo, e error) error {
        if e != nil {
            return e
        }
        if !i.IsDir() {
            files = append(files, path)
        }
    }
    return files
}


// Given the root artifact directory, recurse through it and find all the
// files in that directory. Return an array of absolute paths to the
// artifacts.
func FindArtifacts(dir string) []string {
    return FindFiles(dir)
}

// Given a list of tests/directories, find all the executable tests and return
// a new list of those tests. If a directory is encountered, we will recurse
// through that directory to find all the tests that match the aforementioned
// criteria
func FindTests(testList []string) []string {
    var tests []string
    var files []string

    // First, find all the files from all the paths in testList
    for _, path := range testList {
        files = append(files, FindFiles(path)...)
    }

    // Now, select only the executable files (all tests must be executable)
    // and add those to tests
    for _, file := range files {
        fs, err := os.Stat(file)
        if err != nil {
            panic(fmt.Sprintf("%s", err))
        }
        if fs.Mode() & 0555 == 0555 {
            test = append(tests, file)
        }
    }
    return tests
}

// Shell out and run a single test. Collect STDOUT, STDERR, diff, and
// artifacts. Return a struct with the following:
//   - name of test
//   - path to baseline (usually <path to test>.baseline)
//   - path to stdout file
//   - path to stderr file
//   - path to diff file
//   - path to artifactDir
//   - test status:
//     - pass
//     - fail
//     - skipped
//     - updated
//     - matched
// Note: this test will not set pass/updated/matched and may not set fail. It
// will only set a test status if the test is skipped or it fails due to
// STDERR not being empty
func (tr *TestResult) RunTest(test string, artifactDir string) {
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
