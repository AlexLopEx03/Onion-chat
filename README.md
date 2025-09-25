![EN](https://flagcdn.com/w20/gb.png) [English version](https://github.com/AlexLopEx03/Onion-chat/blob/main/README.en.md) of this readme

<div align="center">
  <h1>Onion-chat</h1>
</div>

Herramienta de línea de comandos para mantener un chat privado en la terminal mediante Tor.

***Proyecto personal de código abierto por AlexLopEx03 bajo licencia AGPLv3.0*** 📜

<br>

<div align="center">
  
| Golang | Tor |
|:-----:|:----:|
| <img src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg" width="120"/> | <img src="https://logo.svgcdn.com/l/tor.png" width="120"/> |
</div>

---

## Instalación

### Windows

```powershell
powershell -ExecutionPolicy Bypass -Command "iwr https://raw.githubusercontent.com/alexlopex03/Onion-chat/main/scripts/installer.ps1 | iex"
```

### Linux y MacOS (No disponible, aún en desarrollo)

```bash
# No disponible
curl "https://raw.githubusercontent.com/alexlopex03/Onion-chat/main/scripts/installer.sh"
```

---

## Guía de uso

> Se puede simular el chat entre dos usuarios con dos terminales diferentes desde un mismo equipo

- Un usuario ejecutará lo siguiente:

```cmd
onion-chat create
```

- Esto devolverá la dirección onion que deberá de recibir el otro usuario del chat, puede tardar aproximadamente 20-30 segundos en generar el servicio de Tor.

> Ejemplo de dirección onion: h37nkv26vmumhqzt536eo5zglfk5vkjf2ajid.onion

- El otro usuario ejecutará lo siguiente con la dirección recibida como parámetro:

```cmd
onion-chat connect example-url.onion
```

- Tras unos segundos, cuando se inicie la conexión entre ambos usuarios ya verán el chat final.

> [!WARNING]
> Onion-chat utiliza los puertos 80 y 3000, es probable que si los puertos estan ocupados pueda no funcionar alguno de los pasos previos sin ningún aviso.

---
<div align="center">
  
> ## Ejemplo final de chat en la terminal: 
  
  <img width="500" height="153" alt="onion-chat-preview-1" src="https://github.com/user-attachments/assets/e15ebda3-3c6b-47fc-8d82-dee813702212" />

<br>

> ### Así vería el chat el otro usuario:

  <img width="500" height="138" alt="onion-chat-preview-2" src="https://github.com/user-attachments/assets/28b04aba-5094-4293-9d47-72ca8e327ea1" />
  
</div>

---

## Acerca del proyecto

> [!IMPORTANT]
> Este proyecto utiliza Bine, una librería de Golang para abstraer de parte de la configuración de los servicios de Tor y el torrc (El archivo principal de configuración de Tor).
>
> Bine dejó de ser desarrollado y mantenido a mediados de 2023, tan solo actualiza periódicamente sus dependencias.
> <br>Esto no implica grandes problemas, la gran mayoría de la seguridad recae en el motor de Tor, Bine tan solo abstrae del control de puertos y arranque de procesos de Tor, no gestiona ninguna tarea crítica como cifrado de capas, enrutado, gestión de red... u otras tareas del motor Tor.
> 
> Enlace del [repositorio de Bine](https://github.com/cretz/bine)

> [!NOTE]
> Además este proyecto utiliza el motor tradicional de Tor desarrollado en C, el cual tiene una configuración altamente estable y conservadora que asegura compatibilidad futura.
>
> Como alternativa, existe Arti, una versión moderna desarrollada en Rust, donde actualmente se está reescribiendo todo el motor, es algo más inestable.

---

<div align="center">
    
## Roadmap de desarrollo

| ⚙️ Características pendientes de implementar |
| :-----------------------------------------------------------------|
| Desarrollo multiplataforma para Linux y MacOS                     |
| Persistencia de la misma dirección onion entre servicios creados  |
| Flag opcional para ocultar logs                                   |
| Comando update para actualizar la App y el motor de Tor           |
| Corrección de errores de la interfaz gráfica del chat             |
| Automatización de auditorias de seguridad mediante Github actions |
| Mejora del feedback de algunos comandos                           |

</div>

#### Cualquier duda o comentario acerca del proyecto puedes dirigirte a la sección de Discussions.
