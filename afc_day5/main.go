package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/niemeyer/golang/src/pkg/container/vector"
)

type Map struct {
	DestStart int
	SrcStart  int
	Range     int
}

type SeedPair struct {
	Seed  int
	Range int
}

type SeedProperties struct {
	SeedToSoil            vector.Vector
	SoilToFertilizer      vector.Vector
	FertilizerToWater     vector.Vector
	WaterToLight          vector.Vector
	LightToTemperature    vector.Vector
	TemperatureToHumidity vector.Vector
	HumidityToLocation    vector.Vector
}

type ParseState int

const (
	ParseState_None                  ParseState = 0
	ParseState_SeedToSoil                       = 1
	ParseState_SoilToFertilizer                 = 2
	ParseState_FertilizerToWater                = 3
	ParseState_WaterToLight                     = 4
	ParseState_LightToTemperature               = 5
	ParseState_TemperatureToHumidity            = 6
	ParseState_HumidityToLocation               = 7
)

func setParseState(line string, parse_state *ParseState) {
	switch {
	case strings.Contains(line, "seed-to-soil"):
		*parse_state = ParseState_SeedToSoil
	case strings.Contains(line, "soil-to-fertilizer"):
		*parse_state = ParseState_SoilToFertilizer
	case strings.Contains(line, "fertilizer-to-water"):
		*parse_state = ParseState_FertilizerToWater
	case strings.Contains(line, "water-to-light"):
		*parse_state = ParseState_WaterToLight
	case strings.Contains(line, "light-to-temperature"):
		*parse_state = ParseState_LightToTemperature
	case strings.Contains(line, "temperature-to-humidity"):
		*parse_state = ParseState_TemperatureToHumidity
	case strings.Contains(line, "humidity-to-location"):
		*parse_state = ParseState_HumidityToLocation
	}
}

func populateProperties(line string, parse_state ParseState, seed_properties *SeedProperties) {
	if parse_state != ParseState_None {
		data := strings.Split(strings.Trim(line, " "), " ")
		if len(data) == 3 {
			dest_start, err := strconv.Atoi(data[0])
			if err != nil {
				panic(err)
			}
			src_start, err := strconv.Atoi(data[1])
			if err != nil {
				panic(err)
			}
			range_length, err := strconv.Atoi(data[2])
			if err != nil {
				panic(err)
			}
			m := Map{
				DestStart: dest_start,
				SrcStart:  src_start,
				Range:     range_length,
			}
			switch parse_state {
			case ParseState_SeedToSoil:
				seed_properties.SeedToSoil.Push(m)
			case ParseState_SoilToFertilizer:
				seed_properties.SoilToFertilizer.Push(m)
			case ParseState_FertilizerToWater:
				seed_properties.FertilizerToWater.Push(m)
			case ParseState_WaterToLight:
				seed_properties.WaterToLight.Push(m)
			case ParseState_LightToTemperature:
				seed_properties.LightToTemperature.Push(m)
			case ParseState_TemperatureToHumidity:
				seed_properties.TemperatureToHumidity.Push(m)
			case ParseState_HumidityToLocation:
				seed_properties.HumidityToLocation.Push(m)
			}
		}
	}
}

func getPropertyValue(prev int, property vector.Vector) int {
	for _, p := range property {
		if prev >= p.(Map).SrcStart && prev <= (p.(Map).SrcStart+p.(Map).Range) {
			diff := prev - p.(Map).SrcStart
			return p.(Map).DestStart + diff
		}
	}
	return prev
}

func printSeedProperty(property *vector.Vector, name string) {
	for i, v := range *property {
		log.Printf("%s[%d]: %v", name, i, v)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
		log.Println("File was successfully closed.")
	}()

	var seed_properties SeedProperties
	var parse_state ParseState = ParseState_None
	var seeds = vector.IntVector{} // {seed, range}
	var scanner *bufio.Scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "seeds") {
			data := strings.Trim(strings.Split(line, ":")[1], " ")
			for _, v := range strings.Split(data, " ") {
				n, err := strconv.Atoi(v)
				if err != nil {
					continue
				}
				seeds.Push(n)
			}
			continue
		}
		populateProperties(line, parse_state, &seed_properties)
		setParseState(line, &parse_state)
	}

	lowest_location := 0
	for i := 0; i < seeds.Len()/2-1; i++ {
		start_seed := seeds[i*2]
		seed_range := seeds[i*2+1]
		for j := start_seed; j <= (start_seed + seed_range); j++ {
			soil := getPropertyValue(j, seed_properties.SeedToSoil)
			fertilizer := getPropertyValue(soil, seed_properties.SoilToFertilizer)
			water := getPropertyValue(fertilizer, seed_properties.FertilizerToWater)
			light := getPropertyValue(water, seed_properties.WaterToLight)
			temperature := getPropertyValue(light, seed_properties.LightToTemperature)
			humidity := getPropertyValue(temperature, seed_properties.TemperatureToHumidity)
			location := getPropertyValue(humidity, seed_properties.HumidityToLocation)
			if i == 0 && j == start_seed {
				lowest_location = location
			} else {
				if location < lowest_location {
					lowest_location = location
				}
			}
		}
	}
	log.Println("lowest location: ", lowest_location)
}
