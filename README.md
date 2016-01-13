# simple-log-search
A simple search of simple pixel server logs. Searches many simple pixel log files for requests within a specific space / time range, aggregates accross domain, and outputs the result to stdout. For details of simple pixel logs see https://github.com/habajca/simple-pixel-server.

## Build

    go get
    go build

## Search

    ./simple-log-search [log_directory] --time <unix timestamp> --timeframe <seconds> --lat <latitude> --lon <longitude> --distance <meters>

This searches all the logs in `log_directory` (default: `logs`) for requests within `timeframe` seconds of `time` and `distance` meters of the coordinte `lat, lon` and aggreagates by domain. Output is in the follow format:

    domain, count

This uses https://github.com/kellydunn/golang-geo to compute great circle distances.

## Generate Test Data

    ./simple-log-search [log_directory] --generate

This generates multiple test log files in `log_directory` (default: `logs`).
