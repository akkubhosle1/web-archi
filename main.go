package main

import (
    "bufio"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "sort"
    "strings"
    "sync"
)

type result struct {
    url           string
    contentLength int
    statusCode    int
}

func main() {
    var numConcurrent int
    flag.IntVar(&numConcurrent, "c", 1, "Number of concurrent requests to make")
    flag.Parse()

    scanner := bufio.NewScanner(os.Stdin)

    var results []result
    var wg sync.WaitGroup
    requests := make(chan string)

    for i := 0; i < numConcurrent; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for url := range requests {
                resp, err := http.Get(url)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "error fetching %s: %v\n", url, err)
                    continue
                }
                defer resp.Body.Close()

                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "error reading body for %s: %v\n", url, err)
                    continue
                }

                results = append(results, result{url, len(body), resp.StatusCode})
            }
        }()
    }

    for scanner.Scan() {
        url := scanner.Text()
        requests <- url
    }

    close(requests)

    wg.Wait()

    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
        os.Exit(1)
    }

    sort.Slice(results, func(i, j int) bool {
        if results[i].statusCode == results[j].statusCode {
            return results[i].contentLength < results[j].contentLength
        }
        return results[i].statusCode < results[j].statusCode
    })

    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error creating file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    defer writer.Flush()

    for _, r := range results {
        var statusColor string

        switch {
        case strings.HasPrefix(fmt.Sprintf("%d", r.statusCode), "2"):
            statusColor = "\033[32m" // green
        case strings.HasPrefix(fmt.Sprintf("%d", r.statusCode), "4"):
            statusColor = "\033[31m" // red
        case strings.HasPrefix(fmt.Sprintf("%d", r.statusCode), "5"):
            statusColor = "\033[34m" // blue
        }

        fmt.Fprintf(writer, "%s: %d (%s%d\033[0m)\n", r.url, r.contentLength, statusColor, r.statusCode)
        fmt.Printf("%s: %d (%s%d\033[0m)\n", r.url, r.contentLength, statusColor, r.statusCode)
    }
}
