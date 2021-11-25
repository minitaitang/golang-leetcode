package main

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/davecgh/go-spew/spew"
)

func HashObject(object interface{}) string {
	hf := fnv.New32()
	printer := spew.ConfigState{
		Indent:         " ",
		SortKeys:       true,
		DisableMethods: true,
		SpewKeys:       true,
	}

	_, _ = printer.Fprintf(hf, "%#v", object)

	return fmt.Sprint(hf.Sum32())
}

func main() {
	// data, err := ioutil.ReadFile("./values.json")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(HashObject("engine-staic-feature-db-engine-staic-feature-db-44419611625521336"))

	// tag := "tag"
	// fmt.Println(filepath.Join("ks/v1", "ndis", "engine/static-feature/v1"), fmt.Sprintf("/"+"engine/static-feature-%s/v1", tag))
	// res, err := simplejson.NewJson(data)
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// 	return
	// }
	// gpuUsageM := map[string]uint64{"T4": 9}
	// properties := res.Get("properties").Get("workers").Get("items").GetIndex(0).Get("properties")
	// if properties.Interface() != nil {
	// 	for typ, num := range gpuUsageM {
	// 		if num != 0 {
	// 			name := properties.Get("name")
	// 			name.Set("enum", []string{strings.ToLower(typ)})
	// 			name.Set("default", strings.ToLower(typ))
	// 			replicas := properties.Get("replicas")
	// 			replicas.Set("enum", []uint64{num})
	// 			replicas.Set("default", num)
	// 			break
	// 		}
	// 	}
	// }
	// b, _ := res.MarshalJSON()
	// fmt.Println(string(b))

	data, err := ioutil.ReadFile("./xfd.json")
	if err != nil {
		return
	}
	valueJson, err := simplejson.NewJson([]byte(data))
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(valueJson)

	// // gpuType, _ := schemaJson.Get("worker_sets").Get("items").GetIndex(0).Get("sets").Int()
	// arr, _ := valueJson.Get("worker_sets").Array()
	// sets, _ := valueJson.Get("worker_sets").GetIndex(0).Get("sets").Int()
	// replicas, _ := valueJson.Get("worker_sets").GetIndex(0).Get("replicas").Int()
	// for _, m := range arr {
	// 	fmt.Println(m)
	// }
	// fmt.Printf("%T\n", arr)
	// fmt.Println(arr, sets, replicas)

	gpuinfo := map[string]uint64{
		"p4": 0,
	}

	// var needGpu bool
	// for _, gpuNum := range gpuinfo {
	// 	if gpuNum > 0 {
	// 		needGpu = true
	// 		break
	// 	}
	// }

	workerSetsArr, _ := valueJson.Get("worker_sets").Array()
	valueGpuM := make(map[string]uint64)
	for i, _ := range workerSetsArr {
		name, _ := valueJson.Get("worker_sets").GetIndex(i).Get("name").String()
		sets, _ := valueJson.Get("worker_sets").GetIndex(i).Get("sets").Uint64()
		replicas, _ := valueJson.Get("worker_sets").GetIndex(i).Get("replicas").Uint64()
		valueGpuM[name] = sets * replicas

		fmt.Println(name, sets, replicas)
		for gpuType, gpuNum := range gpuinfo {
			if strings.ToLower(gpuType) == strings.ToLower(name) {
				if sets*replicas != gpuNum {
					fmt.Println("failed!!!!!!!")
					return
				}
			}
		}
	}
	fmt.Println("success")
	return
	// if needGpu && !CompareMapEqual(gpuinfo, valueGpuM, true) {
	// 	fmt.Println("not match not match")
	// 	return
	// }

	// return

	// m := map[string]uint64{
	// 	"p4": 10,
	// 	"T4": 0,
	// }
	// n := map[string]uint64{
	// 	"P4": 10,
	// 	"T4": 20,
	// }
	// fmt.Println(CompareMapEqual(m, n, true))
}

func CompareMapEqual(m map[string]uint64, n map[string]uint64, deleteZero bool) (isEqual bool) {
	mTemp := make(map[string]uint64)
	nTemp := make(map[string]uint64)
	for k, v := range m {
		if deleteZero && v == 0 {
			continue
		}
		mTemp[strings.ToLower(k)] = v
	}

	for k, v := range n {
		if deleteZero && v == 0 {
			continue
		}
		nTemp[strings.ToLower(k)] = v
	}

	if len(mTemp) != len(nTemp) {
		return false
	}

	for k, v := range mTemp {
		if _, ok := nTemp[k]; !ok {
			return false
		}
		if nTemp[k] != v {
			return false
		}
	}

	for k, v := range nTemp {
		if _, ok := mTemp[k]; !ok {
			return false
		}
		if mTemp[k] != v {
			return false
		}
	}

	return true
}
