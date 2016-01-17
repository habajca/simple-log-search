# simple-log-search
A simple search of simple pixel server logs. Searches many simple pixel log files for requests within a specific space / time range, aggregates accross domain, and outputs the result to stdout. For details of simple pixel logs see https://github.com/habajca/simple-pixel-server.

## Build

    go get
    go build

## Search

    ./simple-log-search [log_directory] [options]

This searches all the logs in `log_directory` (default: `logs`) with possible options:
    
    -t, --time <unix timestamp>
    
Defining the origin of the search in the time dimension as a unix timestamp (in seconds). Defaults to current time.
    
    -m, --timeframe <seconds>
    
Defining the search space in the time dimension in seconds. Defaults to 3600 seconds (1 hour).
    
    -g, --geo <latitude,longitude>
    
Defining the origin of the search in the geo dimensions as a latitude, longitude tuple. Defaults to `37.7576171,-122.5776844` (San Francisco, CA).
    
    -d, --distance <meters>
    
Defining the search space in the geo dimensions. Defaults to 5000 meters.

Output is in the follow format:

    domain, count

This uses https://github.com/kellydunn/golang-geo to compute great circle distances.

## Generate Test Data

    ./simple-log-search [log_directory] [domains_file] --generate [options]

This generates multiple test log files in `log_directory` (default: `logs`) using the domains defined in `domains_file` with possible options:

    -f, --files <number of files>

Defining the number of test files to create. Defaults to 1000 files.

    -r, --rows <number of rows>
    
Defining the number of rows per test file. Defaults to 1000 rows.
