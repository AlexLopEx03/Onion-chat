![ES](https://flagcdn.com/w20/es.png) [Versión en Español](https://github.com/AlexLopEx03/Onion-chat/blob/main/README.md) de este readme

<div align="center">
  <h1>Onion-chat</h1>
</div>

Command-line tool to maintain a private chat in the terminal via Tor.

***Personal open-source project by AlexLopEx03 under the AGPLv3.0 license*** 📜

<br>

<div align="center">
  
| Golang | Tor |
|:-----:|:----:|
| <img src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg" width="120"/> | <img src="https://logo.svgcdn.com/l/tor.png" width="120"/> |
</div>

---

## Installation

### Windows

```powershell
powershell -ExecutionPolicy Bypass -Command "iwr https://raw.githubusercontent.com/alexlopex03/Onion-chat/main/scripts/installer.ps1 | iex"
```

### Linux and MacOS (Not available, still in development)

```bash
# Not available
curl "https://raw.githubusercontent.com/alexlopex03/Onion-chat/main/scripts/installer.sh"
```

---

## Usage guide

> You can simulate the chat between two users with two different terminals on the same machine.

- One user will run the following:

```cmd
onion-chat create
```

- This will return the onion address that the other chat user must receive; it may take approximately 20-30 seconds to generate the Tor service.

> Example of an onion address: h37nkv26vmumhqzt536eo5zglfk5vkjf2ajid.onion

- The other user will run the following using the received address as a parameter:

```cmd
onion-chat connect example-url.onion
```

- After a few seconds, when the connection between both users starts, they will see the final chat.

> [!WARNING]
> Onion-chat uses ports 80 and 3000. It is likely that if the ports are occupied, some of the previous steps may not work without any warning.

---
<div align="center">
  
> ## Final example of a chat terminal: 
  
  <img width="500" height="153" alt="onion-chat-preview-1" src="https://github.com/user-attachments/assets/e15ebda3-3c6b-47fc-8d82-dee813702212" />

<br>

> ### How the other user would see it:

  <img width="500" height="138" alt="onion-chat-preview-2" src="https://github.com/user-attachments/assets/28b04aba-5094-4293-9d47-72ca8e327ea1" />
  
</div>

---

## Acerca del proyecto

> [!IMPORTANT]
> This project uses Bine, a Golang library to abstract part of the configuration of Tor services and the torrc (The main Tor configuration file).
>
> Bine stopped being actively developed and maintained in mid-2023; it only periodically updates its dependencies.
> <br>Although this does not imply major issues, the vast majority of security relies on the Tor engine. Bine only abstracts port control and Tor process startup; it does not handle any critical tasks such as layer encryption, routing, network management, or other Tor engine tasks.
> 
> Link to the [Bine repository](https://github.com/cretz/bine)

> [!NOTE]
> Additionally, this project uses the traditional Tor engine developed in C, which has a highly stable and conservative configuration that ensures future compatibility.
>
> As an alternative, there is Arti, a modern version developed in Rust, where the entire engine is currently being rewritten.

---

<div align="center">
    
## Development Roadmap

| ⚙️ Features pending implementation |
| :-----------------------------------------------------------------|
| Cross-platform developnebt for Linux and MacOS                    |
| Persistence of the same onion address across created services     |
| Optional flag to hide logs                                        |
| Update command to update the App and the Tor engine               |
| Fixes for chat interface bugs                                     |
| Automation of security audits via Github Actions                  |
| Improvement of feedback for some commands                         |

</div>

#### If you have any questions or comments about the project, you can direct them to the Discussions section.
