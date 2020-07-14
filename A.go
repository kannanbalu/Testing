
package main

import ( 
  "fmt"
  "os"
  "flag"
//  "runtime/trace"
)

func main() {
   //trace.Start(os.Stdout);
   //defer trace.Stop();
   fmt.Println("Args length: ", len(os.Args))
   fmt.Println("command line args: ", os.Args);
   parseCmdLineArgs();
}

func parseCmdLineArgs() {
    namePtr := flag.String("Name", "Kannan", "Please enter your name");
    countPtr := flag.Int("Count", 100, "Please enter a count");
    flag.Parse();
    fmt.Println("name entered is ", *namePtr);
    fmt.Println("count entered is ", *countPtr);
}

