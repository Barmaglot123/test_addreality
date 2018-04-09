package dockertest

import (
    "github.com/fsouza/go-dockerclient"
    "testing"
    "untitled3/datasource"
    "time"
)

func CreateOptions(imageName string) docker.CreateContainerOptions {
    opts := docker.CreateContainerOptions{
        Config: &docker.Config{
            Image:        imageName,
            ExposedPorts: map[docker.Port]struct{}{
                "5432/tcp": {},
                "6379/tcp": {},
            },
        },
    }

    hostConf := docker.HostConfig{
        PortBindings: map[docker.Port][]docker.PortBinding{
            "5432/tcp": {
                {
                    HostIP: "0.0.0.0",
                    HostPort: "9668",
                },
            },
            "6379/tcp": {
                {
                    HostIP: "0.0.0.0",
                    HostPort: "9667",
                },
            },
        },
    }
    opts.HostConfig = &hostConf
    return opts
}

func StartContainer(imageName string, t *testing.T) (*docker.Client, *docker.Container) {
    client, err := docker.NewClientFromEnv()

    if err != nil {
        t.Fatalf("Cannot connect to Docker daemon: %s", err)
    }

    c, err := client.CreateContainer(CreateOptions(imageName))

    if err != nil {
        t.Fatalf("Cannot create Docker container: %s", err)
    }

    err = client.StartContainer(c.ID, &docker.HostConfig{})
    if err != nil {
        t.Fatalf("Cannot start Docker container: %s", err)
    }
    return client, c
}

func RemoveContainer(client *docker.Client, cID string, t *testing.T) {
    if err := client.RemoveContainer(docker.RemoveContainerOptions{
        ID:    cID,
        Force: true,
    }); err != nil {
        t.Fatalf("cannot remove container: %s", err)
    }
}

func WaitReachable(maxWait time.Duration, t *testing.T) {
    done := time.Now().Add(maxWait)

    for time.Now().Before(done) {
        err := datasource.SetupTestSql()
        if err == nil {
            datasource.Sql.Close()
            return
        }
        time.Sleep(1 * time.Second)
    }
    t.Fatalf("Couldn't reach PostgreSLQ server for testing, aborting.")
}

func WaitStarted(client *docker.Client, id string, maxWait time.Duration, t *testing.T) {
    done := time.Now().Add(maxWait)

    for time.Now().Before(done) {
        c, err := client.InspectContainer(id)
        if err != nil {
            break
        }
        if c.State.Running {
            return
        }
        time.Sleep(100 * time.Millisecond)
    }
    t.Fatalf("Couldn't start container")
}