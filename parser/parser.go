package parser

import (
	"bufio"
	"os"
	"strings"
)

type KazanObject struct {
	Coordinates [2]string
	Name        string
	Extra       string
}

func Parser(file *os.File) []KazanObject {
	var KazanObjectsArray []KazanObject

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()

		KazanObjectData := strings.Split(str, ", ")

		lt := KazanObjectData[0]
		lg := KazanObjectData[1]
		coordinates := [2]string{lt, lg}

		name := KazanObjectData[2]

		var extra string
		if len(KazanObjectData) > 3 {
			extra = KazanObjectData[3]
		}

		KazanObject := &KazanObject{
			Coordinates: coordinates,
			Name:        name,
			Extra:       extra,
		}

		KazanObjectsArray = append(KazanObjectsArray, *KazanObject)
	}

	return KazanObjectsArray
}
