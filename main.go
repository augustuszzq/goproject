// package main
//
// import (
//   "fmt"
//   "time"
// )
//
// func main() {
//     for{
//       fmt.Println("hello world")#System D
//       time.Sleep(10*time.Second)
//     }
// }
// package main
//
// import (
//     "log"
//
//     "github.com/fsnotify/fsnotify"
// )
//
// func main() {
//     // Create new watcher.
//     watcher, err := fsnotify.NewWatcher()
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer watcher.Close()
//
//     // Start listening for events.
//     go func() {
//         for {
//             select {
//             case event, ok := <-watcher.Events:
//                 if !ok {
//                     return
//                 }
//                 log.Println("event:", event)
//                 if event.Has(fsnotify.Write) {
//                     log.Println("modified file:", event.Name)
//                 }
//             case err, ok := <-watcher.Errors:
//                 if !ok {
//                     return
//                 }
//                 log.Println("error:", err)
//             }
//         }
//     }()
//
//     // Add a path.
//     err = watcher.Add("/etc/supervisor/conf.d/supervisord.conf")
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     // Block main goroutine forever.
//     <-make(chan struct{})
// }
package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "time"

)
func main() {
    fmt.Println("start")
    // Create a new command
    cmd := exec.Command("supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    // Start the command
    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Waiting for supervisord to start...")
    time.Sleep(500 * time.Second) // Wait for supervisord to start

    // // Reread the configuration
    // cmd = exec.Command("supervisorctl", "reread")
    // cmd.Stdout = os.Stdout
    // cmd.Stderr = os.Stderr
    // err = cmd.Run()
    // if err != nil {
    //     log.Fatalf("Failed to reread: %s", err)
    // }
    //
    // // Update the processes
    // cmd = exec.Command("supervisorctl", "update")
    // cmd.Stdout = os.Stdout
    // cmd.Stderr = os.Stderr
    // err = cmd.Run()
    // if err != nil {
    //     log.Fatalf("Failed to update: %s", err)
    // }

    log.Printf("Waiting for comsmand to finish...")
    err = cmd.Wait()
    if err != nil {
        log.Printf("Command finished with error: %v", err)
    }
}
