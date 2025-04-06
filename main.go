package main

import (
    "net/http"
    "os"

    "secure-go-api/internal/math"
    "secure-go-api/pkg/config"

    "github.com/gin-gonic/gin"
)

type Operacion struct {
    A int `json:"a"`
    B int `json:"b"`
}

func main() {
    args := os.Args
    color := "azul"     // valor por defecto
    size := "mediano"   // valor por defecto

    if len(args) > 1 {
        color = args[1]
    }
    if len(args) > 2 {
        size = args[2]
    }

    cfg, err := config.LoadConfig()
    if err != nil {
        panic(err)
    }

    sumaPath := os.Getenv("ENDPOINT_SUMA")
    if sumaPath == "" {
        sumaPath = "/sumar"
    }

    restaPath := os.Getenv("ENDPOINT_RESTA")
    if restaPath == "" {
        restaPath = "/restar"
    }

    r := gin.Default()

    r.GET("/color", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"color": cfg.Color})
    })

    r.POST(sumaPath, func(c *gin.Context) {
        var op Operacion
        if err := c.BindJSON(&op); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
            return
        }
        res := math.Sumar(op.A, op.B)
        c.JSON(http.StatusOK, gin.H{"resultado": res})
    })

    r.POST(restaPath, func(c *gin.Context) {
        var op Operacion
        if err := c.BindJSON(&op); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
            return
        }
        res := math.Restar(op.A, op.B)
        c.JSON(http.StatusOK, gin.H{"resultado": res})
    })

    r.GET("/props", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "color":  color,
            "tamaño": size,
        })
    })

    r.Run(":8080")
}