package main

import (
    "log"
    "os"
    "os/exec"
    "strconv"
    "time"

    "github.com/warthog618/gpiod"
)

func main() {
    if len(os.Args) != 2 {
        log.Println("Usage: shutdown-monitor <PIN>")
        return
    }

    pinStr := os.Args[1]
    pin, err := strconv.ParseInt(pinStr, 10, 64)
    if err != nil {
        log.Fatal("Invalid argument %s: %w", pinStr, err)
    }

    l, err := gpiod.RequestLine("gpiochip0", int(pin), gpiod.AsInput, gpiod.AsActiveLow, gpiod.WithPullUp)
    if err != nil {
        log.Fatal(err)
    }
    defer l.Close()

    log.Println("Starting run loop waiting for GPIO", pin)
    for {
        time.Sleep(50*time.Millisecond)
        n, err := l.Value()
        if err != nil {
            log.Println(err)
        }

        if n == 1 {
            break
        }
    }
    log.Println("Run loop completed. Shutting down system.")

    cmd := exec.Command("poweroff")
    if err := cmd.Run(); err != nil {
        log.Println("failed to shutdown -", err)
    }
}
