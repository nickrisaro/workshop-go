# Workshop de Go

## Cosas que necesitás
1. Go 1.13
2. Un poco de tiempo

## Por dónde empezar
1. Creá la carpeta para descargar este repo
```
mkdir -p $HOME/go/src/github/nickrisaro
```
2. Andá a esa carpeta
```
cd $HOME/go/src/github/nickrisaro
```
3. Cloná el repo
```
git clone git@github.com:nickrisaro/workshop-go.git
```
4. Compilá el proyecto 
```
go build
```
5. Ejecutá el programa
```
./workshop-go
```
6. Visitá la página http://localhost:8080/tweets

## Cómo seguir
Mirá las diapositivas, las podés leer con cualquier editor de textos o verlas más bonitas en el browser, para eso:
1. Instalá present
```
go get golang.org/x/tools/cmd/present
```
2. Andá al directorio de los slides
```
cd $HOME/go/src/github/nickrisaro/slides-resumido
```
3. Ejecutá present
```
present
```
Si no configuraste bien el PATH puede que no lo encuentre, probá co $HOME/go/bin/present

4. Mirá los slides en el browser
http://localhost:3999/01%20-%20introduccion%20a%20go.slide

Leé los TODOs que hay por el código e implementá la nueva funcionalidad

Este proyecto usa gin-gonic, https://github.com/gin-gonic/gin, testify https://github.com/stretchr/testify y go modules https://github.com/golang/go/wiki/Modules

En el archivo links-utiles.md hay un montón de sitios que fui leyendo para armar este workshop

En el archivo comunidad.md hay varias cuentas de twitter relacionadas con go y el link para entrar al slack oficial de go