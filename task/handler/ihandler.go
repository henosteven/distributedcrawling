package handler

import (
    "os"
)

type Handler interface {
    DoTask(task string) string, os.Error
}

