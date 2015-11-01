// Code template taken from http://gobot.io/documentation/platforms/raspi/

package main

import (
    "fmt"
    "log"

    "github.com/hybridgroup/gobot"
    "github.com/hybridgroup/gobot/platforms/raspi"
    "github.com/hybridgroup/gobot/platforms/gpio"

    "github.com/samalba/dockerclient"
)

func runningContainers() int {
    // Init the client
    docker, _ := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
    // Get only running containers
    containers, err := docker.ListContainers(false, false, "")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(len(containers))
    return len(containers)
}

func main() {
    raspiAdaptor := raspi.NewRaspiAdaptor("raspi")
    a := gpio.NewLedDriver(raspiAdaptor, "a", "16")
    b := gpio.NewLedDriver(raspiAdaptor, "b", "12")
    c := gpio.NewLedDriver(raspiAdaptor, "c", "16") // changed from 5 to 16
    d := gpio.NewLedDriver(raspiAdaptor, "d", "6")
    e := gpio.NewLedDriver(raspiAdaptor, "e", "0")
    f := gpio.NewLedDriver(raspiAdaptor, "f", "1")
    g := gpio.NewLedDriver(raspiAdaptor, "g", "24")
    dp := gpio.NewLedDriver(raspiAdaptor, "dp", "26")

    work := func() {
        containerCount := runningContainers()
//        fmt.Println(containerCount)
        switch containerCount {
          case 0:
            a.On()
            b.On()
            c.On()
            d.On()
            e.On()
            f.On()
            g.Off()
            dp.On()
            return
          case 1:
            a.Off()
            b.On()
            c.On()
            d.Off()
            e.Off()
            f.Off()
            g.Off()
            dp.On()
            return
          case '2':
            a.On()
            b.On()
            c.Off()
            d.On()
            e.On()
            f.Off()
            g.On()
            dp.On()
            return
          case '3':
            a.On()
            b.On()
            c.On()
            d.On()
            e.Off()
            f.Off()
            g.On()
            dp.On()
            return
          case '4':
            a.Off()
            b.On()
            c.On()
            d.Off()
            e.Off()
            f.Off()
            g.Off()
            dp.On()
            return
          case '5':
            a.Off()
            b.On()
            c.On()
            d.Off()
            e.Off()
            f.Off()
            g.Off()
            dp.On()
            return
          case '6':
            a.Off()
            b.On()
            c.On()
            d.Off()
            e.Off()
            f.Off()
            g.Off()
            dp.On()
            return
          case '7':
            a.Off()
            b.On()
            c.On()
            d.Off()
            e.Off()
            f.Off()
            g.Off()
            dp.On()
            return
          case '8':
            a.Off()
            b.On()
            c.On()
            d.Off()
            e.Off()
            f.Off()
            g.Off()
            dp.On()
            return
          case '9':
            a.Off()
            b.On()
            c.On()
            d.Off()
            e.Off()
            f.Off()
            g.Off()
            dp.On()
            return
          default:
            a.On()
            b.On()
            c.On()
            d.On()
            e.On()
            f.On()
            g.On()
            dp.On()
        }
    }

    robot := gobot.NewRobot("bot",
        []gobot.Connection{raspiAdaptor},
        []gobot.Device{a},
        []gobot.Device{b},
        []gobot.Device{c},
        []gobot.Device{d},
        []gobot.Device{e},
        []gobot.Device{f},
        []gobot.Device{g},
        work,
    )

    gbot.AddRobot(robot)

    gbot.Start()
}
