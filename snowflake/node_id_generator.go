package snowflake

import (
	"hash/fnv"
	"log"
	"math/rand"
	"net"
)

func nodeIDGenerator() {
	//Make sure to generate it once only ...
	if nodeID != 0 {
		return
	}

	//uses the system’s MAC address to create a unique identifier for the Node
	nodeIFAS, err := getMacAddr()
	if err != nil {
		log.Println("error generating nodeID::", err.Error())

		//In case it fails to get mac address, generate the random number,
		//gloabbly seeded with time.Now()
		nodeID = rand.Int()
		return
	}

	nodeID = hashCode(nodeIFAS) & maxNodeID
	if nodeID == -1 {
		nodeID = rand.Int()
	}

}

func getMacAddr() (string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var as string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as += a
		}
	}
	return as, nil
}

func hashCode(s string) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return int(h.Sum32())
}
