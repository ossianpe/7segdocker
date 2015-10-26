# Code template taken from http://gobot.io/documentation/platforms/raspi/

package main

import (
    "time"

    "github.com/hybridgroup/gobot"
    "github.com/hybridgroup/gobot/platforms/raspi"
    "github.com/hybridgroup/gobot/platforms/gpio"

    "github.com/samalba/dockerclient"
)

func runningContainers() {
    // Init the client
    docker, _ := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
    // Get only running containers
    containers, err := docker.ListContainers(false, false, "")
    if err != nil {
        log.Fatal(err)
    }
    return len(containers)
}

func main() {
    gbot := gobot.NewGobot()

    raspiAdaptor := raspi.NewRaspiAdaptor("raspi")
    a := gpio.NewLedDriver(raspiAdaptor, "a", "16")
    b := gpio.NewLedDriver(raspiAdaptor, "b", "12")
    c := gpio.NewLedDriver(raspiAdaptor, "c", "16") # changed from 5 to 16
    d := gpio.NewLedDriver(raspiAdaptor, "d", "6")
    e := gpio.NewLedDriver(raspiAdaptor, "e", "0")
    f := gpio.NewLedDriver(raspiAdaptor, "f", "1")
    g := gpio.NewLedDriver(raspiAdaptor, "g", "24")
    dp := gpio.NewLedDriver(raspiAdaptor, "dp", "26")

    zero := func() {
        a.On()
        b.On()
        c.On()
        d.On()
        e.On()
        f.On()
    }

    one := func() {
        b.On()
        c.On()
    }

    two := func() {
        a.On()
        b.On()
        d.On()
        e.On()
        g.On()
    }

    three := func() {
        a.On()
        b.On()
        c.On()
        d.On()
        g.On()
    }

    four := func() {
        b.On()
        c.On()
    }

    one := func() {
        b.On()
        c.On()
    }

    one := func() {
        b.On()
        c.On()
    }

    one := func() {
        b.On()
        c.On()
    }

    one := func() {
        b.On()
        c.On()
    }

    one := func() {
        b.On()
        c.On()
    }

    robot := gobot.NewRobot("bot",
        []gobot.Connection{raspiAdaptor},
        []gobot.Device{a},
        work,
    )

    gbot.AddRobot(robot)

    gbot.Start()
}
