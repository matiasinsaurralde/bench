package main

import(
	"testing"
	"os/exec"
  "runtime"
  "fmt"
)

func execute( command string, args string) {
	_, err := exec.Command( command, args ).Output()
	if err != nil {
		panic(err)
	}
}

func benchmarkGo( b *testing.B ) {
  switch runtime.GOOS {
  case "windows":
    execute( "go/main.exe", "" )
  default:
    execute( "go/main", "" )
  }
}

func benchmarkNode( b *testing.B ) {
  execute( "node", "node/index.js" )
}

func benchmarkOtto( b *testing.B ) {
  switch runtime.GOOS {
  case "windows":
    execute( "otto/main.exe", "" )
  default:
    execute( "otto/main", "" )
  }
}

func main() {

  fmt.Println("GOOS =", runtime.GOOS, "\n")
  GoResult := testing.Benchmark(benchmarkGo)
  fmt.Println("Go:")
  nanoseconds := float64(GoResult.T.Nanoseconds()) / float64(GoResult.N)
  milliseconds := nanoseconds / 1000000.0
  fmt.Printf("%13.2f ns/op | %13.10f ms/op | %d Iterations\n", nanoseconds, milliseconds, GoResult.N)

  NodeResult := testing.Benchmark(benchmarkNode)
  fmt.Println("Node:")
  nanoseconds = float64(NodeResult.T.Nanoseconds()) / float64(NodeResult.N)
  milliseconds = nanoseconds / 1000000.0

  fmt.Printf("%13.2f ns/op | %13.10f ms/op | %d Iterations\n", nanoseconds, milliseconds, NodeResult.N)

  OttoResult := testing.Benchmark(benchmarkOtto)
  fmt.Println("Otto:")
  nanoseconds = float64(OttoResult.T.Nanoseconds()) / float64(OttoResult.N)
  milliseconds = nanoseconds / 1000000.0

  fmt.Printf("%13.2f ns/op | %13.10f ms/op | %d Iterations\n", nanoseconds, milliseconds, OttoResult.N)
}
