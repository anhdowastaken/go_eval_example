package main

import (
    "strings"
    "fmt"

    "github.com/Knetic/govaluate"
    "github.com/spf13/viper"
)

type Config struct {
    Expression string
    Parameters []bool
}

func main() {
    v := viper.New()
    v.SetConfigFile("config.toml")
    v.SetConfigType("toml")
    err := v.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    var cfg Config
    err = v.UnmarshalKey("config", &cfg)
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
    fmt.Printf("%+v\n", cfg)

    exp_str := cfg.Expression
    for i := 0; i < len(cfg.Parameters); i++ {
        p := cfg.Parameters[i]
        old_str := fmt.Sprintf("{{p%d}}", i)
        new_str := fmt.Sprintf("%t", p)
        exp_str = strings.Replace(exp_str, old_str, new_str, -1)
    } 
    fmt.Printf("## %s\n", cfg.Expression)
    fmt.Printf("-> %s\n", exp_str)

    expression, err := govaluate.NewEvaluableExpression(exp_str)
    if err != nil {
        panic(fmt.Errorf("%s \n", err))
    }

    result, err := expression.Evaluate(nil)
    if err != nil {
        panic(fmt.Errorf("%s \n", err))
    }
    fmt.Printf("-> %t\n", result)
}
