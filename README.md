# eResumen Go-Parser-Poc

Prueba de concepto para el generador de resumenes en pdf de eResumen

## Descripción de modulos:

- Commons: Helpers y dtos compartidos entre componentes
- Database-Service: Microservicio encargado de la persistencia de resumenes
- Email-Consumer: Microservicio encargado del envio de los resumenes por email
- Email-Mock-Server: Servidor smtp mockeado para guardar en redis los mails enviados
- File-Parser-Producer: Punto de entrada para la ejecución del proceso