# INICIAR LOG

```go
func main() {
    logger.StartZapLog()
    defer logger.Sync() // Importante para não perder logs em buffer

    logger.Info("Aplicação iniciada", zap.String("versao", "1.0.0"))
    logger.Infof("Usuário %s logado com sucesso", "João")
}
```