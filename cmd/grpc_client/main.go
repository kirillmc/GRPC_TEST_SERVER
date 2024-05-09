package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	filler "github.com/kirillmc/grpc_test_server/internal/filler_pb"
	desc "github.com/kirillmc/grpc_test_server/pkg/program_v3"
)

const (
	address     = "localhost:50051"
	oldAddress  = "2.tcp.eu.ngrok.io:10883"
	messageSize = 1024 * 1024 * 1024
	avg         = 5.0
	sizeCh      = 1
)

// func DialogOptions(chains ...grpc.UnaryClientInterceptor) []grpc.DialOption {
func DialogOptions() []grpc.DialOption {
	//	chains = append(chains, userinfo.UnaryClientInterceptor())
	return []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(messageSize),
			grpc.MaxCallSendMsgSize(messageSize),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithChainUnaryInterceptor(chains...),
	}
}

func getNProgramsClient(n int64) (*desc.TrainPrograms, error) {
	conn, err := grpc.Dial(address,
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(messageSize)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect %v", err)
		return nil, err
	}

	defer conn.Close()

	c := desc.NewProgramV3Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 110*time.Second)
	defer cancel()

	//maxSizeOption := grpc.MaxCallRecvMsgSize(1024 * 1024 *1024)
	programs, err := c.Get(ctx, &desc.GetRequest{Count: n})

	if err != nil {
		log.Println(err)
	}

	return programs, nil
}

func postNProgramsClient(programs *desc.TrainPrograms) (*desc.Response, float64, error) {
	dataToSend, err := proto.Marshal(programs)
	if err != nil {
		return &desc.Response{Message: err.Error()}, float64(len(dataToSend)), err
	}

	conn, err := grpc.Dial(address,
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(messageSize)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("failed to connect %v", err)
		return &desc.Response{Message: err.Error()}, float64(len(dataToSend)), err
	}

	defer conn.Close()

	c := desc.NewProgramV3Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 110*time.Second)
	defer cancel()

	mess, err := c.Create(ctx, programs)

	return mess, float64(len(dataToSend)), nil
}

func updateNProgramsClient(programs *desc.TrainPrograms) (*desc.Response, float64, error) {
	dataToUpdate, err := proto.Marshal(programs)
	if err != nil {
		return &desc.Response{Message: err.Error()}, float64(len(dataToUpdate)), err
	}

	conn, err := grpc.Dial(address,
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(messageSize)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("failed to connect %v", err)
		return &desc.Response{Message: err.Error()}, float64(len(dataToUpdate)), err
	}

	defer conn.Close()

	c := desc.NewProgramV3Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 110*time.Second)
	defer cancel()

	mess, err := c.Update(ctx, programs)

	return mess, float64(len(dataToUpdate)), nil
}

func deleteNProgramsClient(req *desc.DeleteRequest) (*desc.Response, float64, error) {
	dataToDelete, err := proto.Marshal(req)
	if err != nil {
		return &desc.Response{Message: err.Error()}, float64(len(dataToDelete)), err
	}

	conn, err := grpc.Dial(address,
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(messageSize)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("failed to connect %v", err)
		return &desc.Response{Message: err.Error()}, float64(len(dataToDelete)), err
	}

	defer conn.Close()

	c := desc.NewProgramV3Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 110*time.Second)
	defer cancel()

	mess, err := c.Delete(ctx, req)

	return mess, float64(len(dataToDelete)), nil
}

// TODO В цикле при росте количества горутин - смотреть увеличение средней задержки, не увеличивая размер
// TODO В цикле при статическом количестве горутин - смотреть увеличение средней задержки, увеличивая размер
var count int

func main() {
	//var n int64 = 35
	//
	//launchFirstTestForm0ToN(n)

	launchThirdTestWithNProgramsAndWg(25, 11)

}

func launchSecondTestWithNProgramsAndWg(n int64, wgGroupCount int64) {
	fmt.Printf("SIZE - IS ONLY SIZE OF DATA(body)[%d]", n)

	fmt.Printf("\nMETHOD GET from 1 to %d USERS:\n", wgGroupCount)
	methodWithConstAVGOfNGproutines(n, wgGroupCount, oneToGet)

	fmt.Printf("\nMETHOD POST from 0 to %d USERS:\n", wgGroupCount)
	methodWithConstAVGOfNGproutines(n, wgGroupCount, oneToPost)

	fmt.Printf("\nMETHOD UPDATE from 0 to %d USERS:\n", wgGroupCount)
	methodWithConstAVGOfNGproutines(n, wgGroupCount, oneToUpdate)

	fmt.Printf("\nMETHOD DELETE from 0 to %d USERS:\n", wgGroupCount)
	methodWithConstAVGOfNGproutines(n, wgGroupCount, oneToDelete)
}
func launchThirdTestWithNProgramsAndWg(n int64, wgGroupCount int64) {
	fmt.Printf("SIZE - IS ONLY SIZE OF DATA(body) [USERS: %d]", wgGroupCount)

	fmt.Printf("\nMETHOD GET from 1 to %d COUNTS:\n", n)
	methodFrom0ToNWithAVGOfNGproutines(n, wgGroupCount, oneToGet)

	fmt.Printf("\nMETHOD POST from 0 to %d COUNTS:\n", n)
	methodFrom0ToNWithAVGOfNGproutines(n, wgGroupCount, oneToPost)

	fmt.Printf("\nMETHOD UPDATE from 0 to %d COUNTS:\n", n)
	methodFrom0ToNWithAVGOfNGproutines(n, wgGroupCount, oneToUpdate)

	fmt.Printf("\nMETHOD DELETE from 0 to %d COUNTS:\n", n)
	methodFrom0ToNWithAVGOfNGproutines(n, wgGroupCount, oneToDelete)
}

// Будет увеличиваться объем данных от 0 до n, статично горутин wgGroupCount
func methodFrom0ToNWithAVGOfNGproutines(n int64, wgGroupCount int64, fun func(int64) (float64, float64)) {
	log.Printf("USERS;\tCOUNT;\tTIME(nanoS);\tSIZE(byte);\n")
	for i := int64(0); i <= n; i++ {
		printAvgOfGoroutines(i, wgGroupCount, fun)
	}
}

// Будет увеличиваться количетсово горутин от 0 до wgGroupCount, статично объем данных n
func methodWithConstAVGOfNGproutines(n int64, wgGroupCount int64, fun func(int64) (float64, float64)) {
	log.Printf("USERS;\tCOUNT;\tTIME(nanoS);\tSIZE(byte);\n")
	if wgGroupCount <= 0 {
		wgGroupCount = 1
	}
	for i := int64(1); i <= wgGroupCount; i++ {
		printAvgOfGoroutines(n, i, fun)
	}
}

func printAvgOfGoroutines(n int64, wgGroupCount int64, fun func(int64) (float64, float64)) {
	var durOfSize []float64
	var wg sync.WaitGroup
	wg.Add(int(wgGroupCount))
	// Создаем канал для передачи результатов из горутин
	resultСh := make(chan float64, wgGroupCount)
	sizeСhn := make(chan float64, wgGroupCount)

	// Создаем мьютекс для безопасного доступа к срезу result
	var mu sync.Mutex

	// Используем цикл для запуска горутин
	for i := int64(1); i <= wgGroupCount; i++ {
		go func(n int64) {
			defer wg.Done()
			dur, size := fun(n)
			resultСh <- dur
			sizeСhn <- size
		}(n)
	}

	wg.Wait()
	close(resultСh)
	close(sizeСhn)
	var size float64
	for res := range resultСh {
		mu.Lock()
		size = <-sizeСhn
		durOfSize = append(durOfSize, res)
		mu.Unlock()
	}

	log.Printf("\t%d;\t%d;\t%f;\t%f;\n", wgGroupCount, n, getAvgFromSlice(wgGroupCount, durOfSize), size)
}

func getAvgFromSlice(n int64, durOfSize []float64) float64 {
	var avgTime float64
	for i := int64(0); i < n; i++ {
		avgTime += durOfSize[i]
	}

	return avgTime / float64(n)
}

func launchFirstTestForm0ToN(n int64) {
	fmt.Printf("SIZE - IS ONLY SIZE OF DATA(body)")

	fmt.Printf("\nMETHOD GET from 0 to %d:\n", n)
	methodFrom0ToNWithAVG(n, oneToGet)

	fmt.Printf("\nMETHOD POST from 0 to %d:\n", n)
	methodFrom0ToNWithAVG(n, oneToPost)

	fmt.Printf("\nMETHOD UPDATE from 0 to %d:\n", n)
	methodFrom0ToNWithAVG(n, oneToUpdate)

	fmt.Printf("\nMETHOD DELETE from 0 to %d:\n", n)
	methodFrom0ToNWithAVG(n, oneToDelete)

}

func methodFrom0ToNWithAVG(n int64, fun func(int64) (float64, float64)) {
	log.Printf("\tCOUNT;\tTIME(nanoS);\tSIZE(byte);\n")
	for i := int64(0); i <= n; i++ {
		printAvgOfConst(i, fun)
	}
}

func printAvgOfConst(n int64, fun func(int64) (float64, float64)) {
	var avgTime float64
	var avgSize float64
	for i := 1; i <= avg; i++ {
		avgTempTime, avgTempSize := fun(n)
		avgTime += avgTempTime
		avgSize += avgTempSize
	}

	log.Printf("\t%d;\t%f;\t%f;\n", n, avgTime/avg, avgSize/avg)
}

func oneToGet(n int64) (float64, float64) {
	start := time.Now()

	programs, err := getNProgramsClient(n)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()

	//log.Printf("GET: %v", programs)
	numOfSets, err := proto.Marshal(programs)
	if err != nil {
		fmt.Errorf("fail to get json: %v", err)
	}

	return float64(end.Sub(start).Nanoseconds()), float64(len(numOfSets))
}

func oneToPost(n int64) (float64, float64) {
	start := time.Now()
	programs := filler.CreateOwnSetOfPrograms(int(n))

	//ms, postMessSize, err := postNProgramsClient(programs)
	_, postMessSize, err := postNProgramsClient(programs)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()

	//log.Printf("POST: %v\n", ms)

	return float64(end.Sub(start).Nanoseconds()), postMessSize
}

func oneToUpdate(n int64) (float64, float64) {
	start := time.Now()
	programs := filler.CreateOwnSetOfPrograms(int(n))

	//ms, postMessSize, err := updateNProgramsClient(programs)
	_, postMessSize, err := updateNProgramsClient(programs)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()

	//log.Printf("UPDATE: %v\n", ms)

	return float64(end.Sub(start).Nanoseconds()), postMessSize
}

func oneToDelete(n int64) (float64, float64) {
	start := time.Now()
	req := &desc.DeleteRequest{Id: n}
	//ms, postMessSize, err := deleteNProgramsClient(req)
	_, postMessSize, err := deleteNProgramsClient(req)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()

	//	log.Printf("DELETE: %v\n", ms)

	return float64(end.Sub(start).Nanoseconds()), postMessSize
}

func oldPrint(n int64) {
	start := time.Now()
	programs, err := getNProgramsClient(n)
	if err != nil {
		log.Println(err)
	}
	end := time.Now()
	numOfSets, err := proto.Marshal(programs)
	if err != nil {
		fmt.Errorf("fail to get json: %v", err)
	}
	log.Printf("|\t\t\tGRPC INFO: SIZE[%d]\t\t\t|\n", n)
	log.Printf("|\tTOTAL TIME TO GET PROGRAMS:\t%v\t\t|\n", end.Sub(start))
	log.Printf("|\tSIZE OF PROGRAMS:\t\t%s\t|\n", getSizeInFormattedString(int64(len(numOfSets))))
}

func getSizeInFormattedString(byteSize int64) string {
	if byteSize < 1024 {
		return fmt.Sprintf("%.3f байт\t", float64(byteSize))
	}
	if byteSize < 1024*1024 {
		return fmt.Sprintf("%.3f килобайт\t", float64(byteSize)/1024)
	} else {
		return fmt.Sprintf("%.3f мегабайт\t", float64(byteSize)/(1024*1024))
	}
}

//func main() {
//	//var n int64 = 55
//	//log.Printf("\tCOUNT;\tTIME(nanoS);\tSIZE;\n")
//	//for i := int64(1); i <= n; i++ {
//	//	printAvgOfConst(i)
//	//}
//	start := time.Now()
//	for j := 0; j < 10; j++ {
//		var wg sync.WaitGroup
//		wg.Add(100)
//		for i := 0; i < 100; i++ {
//			go func() {
//				defer wg.Done()
//				oldPrint(21)
//				count++
//			}()
//		}
//		wg.Wait()
//	}
//
//	fmt.Printf("TIME: %v, COUNT:%d\n", time.Now().Sub(start), count)
//}
