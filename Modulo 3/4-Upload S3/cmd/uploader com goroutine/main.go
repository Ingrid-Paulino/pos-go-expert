package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//OBS: Não testei esse bloco de código, pois tinha que criar uma conta na AWS e não criei
//O objetivo desse código é fazer upload/subir/baixar  arquivos para o S3 na AWS

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

// func init rosa antes de main
func init() {
	//seção de credenciais para comunicar com AWS
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWS_ACCESS_KEY_ID"),
				os.Getenv("AWS_SECRET_ACCESS_KEY"), //senha
				"",                                 //token
			),
		})
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess) //cria um novo cliente S3/sessao
	s3Bucket = os.Getenv("S3_BUCKET_NAME")
}

func main() {
	dir, err := os.Open("./tmp") //abre o diretório
	if err != nil {
		panic(err)
	}
	defer dir.Close() //fecha o diretório

	//chanel para controlar a quantidade de upload ativos
	//posivel subir 100 arquivos ao mesmo tempo
	uploadControl := make(chan struct{}, 100) //cria um canal para controlar o upload de arquivos, esse canal permite fazer upload de até 100 arquivos ao mesmo tempo
	errorFileUpload := make(chan string, 10)  //controla os uploads que deram erro - faz ate 10 tentativas de forma simutanea

	//toda vez que chegar algum erro, vamor retentar o upload do arquivo
	go func() {
		for {
			select {
			case filename := <-errorFileUpload:
				uploadControl <- struct{}{} //preenche o canal
				wg.Add(1)
				go uploadFile(filename, uploadControl, errorFileUpload)
			}
		}
	}()

	for { //loop para ler todos os arquivos do diretório e fazer upload de cada um para o S3 na AWS
		files, err := dir.ReadDir(1) //lê o diretório
		if err != nil {
			if err == io.EOF { //verifica se chegou ao final do diretório
				break
			}
			fmt.Printf("Error reading directory: %s", err.Error())
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}                                    //depois que o canal estiver cheio, ele não permite fazer mais uploads até que um arquivo seja enviado com sucesso
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload) //sobe  varios arquivos ao mesmo tempo
	}
	wg.Wait()
}

// Retenta o upload de um arquivo caso de algum erro, assim não perco nenhum arquivo / controle de erro
// função para fazer upload/subir arquivo para o S3 na AWS
// faz o upload paralelo - varios arquivos ao mesmo tempo
func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)
	f, err := os.Open(completeFileName) //abre o arquivo
	if err != nil {
		fmt.Sprintf("Error opening file %s: %s", completeFileName, err.Error())
		<-uploadControl                     //esvazia o canal
		errorFileUpload <- completeFileName //envia para o canal o erro
		return
	}
	defer f.Close()                                 //fecha o arquivo
	_, err = s3Client.PutObject(&s3.PutObjectInput{ //sobe arquivo para AWS
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	}) //envia o arquivo para o S3 / faz upload do arquivo
	if err != nil {
		fmt.Sprintf("Error uploading file %s: %s", completeFileName, err.Error())
		<-uploadControl                     //esvazia o canal
		errorFileUpload <- completeFileName //envia para o canal o erro
		return
	}
	fmt.Sprintf("File %s uploaded successfully", completeFileName)
	<-uploadControl //esvazia o canal
}
