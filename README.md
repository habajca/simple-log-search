# simple-log-search
A simple search of simple pixel server logs. Searches many simple pixel log files for requests within a specific space / time range, aggregates accross domain, and outputs the result to stdout. For details of simple pixel logs see https://github.com/habajca/simple-pixel-server.

## Build

    go get
    go build

## Usage

        > ./simple-log-search -h
        Usage:
        ./simple-log-search log_directory [options]
        or
        ./simple-log-search log_directory domains_file --generate [options]
        with options:
        -d, --distance=5000: The search space in the geo dimensions (in meters).
            --generate=false: Indicates that test files should be generated.
        -f, --generation-files=1000: The number of test files to create. (generation only)
        -r, --generation-rows=1000: The number of rows per test file. (generation only)
        -u, --generation-uids=1000: The number of uids in all test files. (generation only)
        -g, --geo=37.7576171,-122.5776844: The origin of the search in the geo dimensions as a latitude, longitude tuple.
        -t, --time=1453286759: The origin of the search in the time dimension as a unix timestamp (in seconds). (defaults to now)
        -m, --timeframe=3600: The search space in the time dimension in seconds.
