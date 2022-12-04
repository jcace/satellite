package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/tatsushid/go-fastping"
)

var ipList = []string{"shuttle-1.estuary.tech", "shuttle-2.estuary.tech", "shuttle-3.estuary.tech", "shuttle-4.estuary.tech", "shuttle-5.estuary.tech", "shuttle-6.estuary.tech", "shuttle-7.estuary.tech", "shuttle-8.estuary.tech"}

func main() {
	var results = make(map[string]float64)
	p := fastping.NewPinger()

	for _, ip := range ipList {
		var pings []int64
		ra, err := net.ResolveIPAddr("ip4:icmp", ip)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		p.AddIPAddr(ra)
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			pings = append(pings, rtt.Milliseconds())
		}

		p.OnIdle = func() {
			// Disregard nodes that don't respond
			if len(pings) == 0 {
				return
			}
			results[ip] = avg(pings)
		}
		err = p.Run()
		if err != nil {
			fmt.Println(err)
		}

	}

	t := table.NewWriter()
	t.SetStyle(table.StyleLight)

	for shuttle, ping := range results {
		t.AppendRow(table.Row{shuttle, fmt.Sprintf("%.2f", ping)})
	}

	fmt.Printf("%s\n", t.Render())

	fmt.Println("Shuttle precedence order: ")
	for i, s := range sortShuttles(results) {
		fmt.Printf("%v | %v \n", i+1, s)
	}
}

func avg(times []int64) float64 {
	var sum = 0.0

	for _, rtt := range times {
		sum += float64(rtt)
	}

	return sum / float64(len(times))
}

func sortShuttles(m map[string]float64) []string {
	result := make([]string, 0, len(m))

	for k := range m {
		result = append(result, k)
	}

	sort.SliceStable(result, func(i, j int) bool {
		return m[result[i]] < m[result[j]]
	})

	return result
}
