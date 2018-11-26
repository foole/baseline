# Design doc for baseline utility

## Command line opts

 -b 		creates or rebaselines a file
 -c <n> 	runs baseline continuously n times (0 is forever) or
until failure
 -d <path>	use path as the artifact directory (overrides env
variable); created automatically if none given (and env variable unset)
 -j 		write results to a json file (results.json, overwritten
		each iteration)
 -v		verbose mode; write out diff below test pass/fail (if
non-empty)

## Environment variables

BL_ARTIFACT_DIR		directory to write artifacts to
BL_TEST_TIMEOUT		per test timeout for each test (defaults to
600s)

## Output

As each test is run, the test path is written out along with it's result
(the result can be different depending on which mode we are running in).

Once all the tests are written out, we then write out a summary of the
results (number passed, failed, skipped, etc)

## Modes

Baseline runs in two modes: checking and writing. The checking mode
means that we are checking to see if the test runs according to a
previous baseline. If the baseline does not exist, then the test is
skipped.

The second mode is writing mode. This writes out the baseline (and will
overwrite the previous baseline) so long as STDERR is empty.

## Operation

Baseline accepts as its final parameters an arbitrary list of files and
directories. It will recurse through directories to find tests. Tests
which are non-executable will be pruned. The final list is sorted in
lexigraphically.

## Test Statuses

### Checking Mode

 - PASS			test's output matched baseline
 - FAIL			test's output did not match baseline or STDERR
   was non-empty
 - SKIP		test was skipped

### Writing Mode
 - UPDATED		baseline was created/updated
 - MATCHED		baseline matched existing baseline (equivalent
   to PASS in check mode)
 - FAILED		baseline could not be updated
