package generator

import (
    "bufio"
    "os"
    "io"
)

type FileGenerator struct {
    FilePath string
}

func (this FileGenerator) GeneratorInit(path interface{}) (simpleFileGen Generator){
    var fileGen FileGenerator
    filePath, ok := path.(string) 
    if !ok {
        panic("type change failed")
    }
    fileGen.FilePath = filePath
    return fileGen
}

func (FileGenerator) GetLatestTask() string {
    return "123"
}

func (this FileGenerator) GetAllTask() []string {
    var taskList []string
    inputFile, err := os.Open(this.FilePath)
    if err != nil {
        panic("cannot access task file")
    }
    defer inputFile.Close()
    inputReader := bufio.NewReader(inputFile)
    for {
        inputString, err := inputReader.ReadBytes('\n')
        if err == io.EOF {
            break
        }
        taskList = append(taskList, string(inputString))
    }
    return taskList
}
