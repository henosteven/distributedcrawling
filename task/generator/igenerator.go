package generator

import (
)

type Generator interface {
    GeneratorInit(mixvar interface{})(gen Generator)
    GetLatestTask() string
    GetAllTask(taskch chan string)
}

