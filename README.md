# 🤖 Dimo

Dimo is a bot programmed in Go that allows you to play a word game chain on the Discord platform. 
It's the perfect game for you if you're looking to develop a sense of repair and improve your memory. 


## 🛠 Perequisities

To run this project

- Make sure you have Docker installed on your machine

## 📥 Installation

Want to add some magic code or cool features ✨ to the bot ? Bring the repository down to your local machine by following theses steps

- Clone this repository:

```bash
https://github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP01_BACK
cd PROJECT_FINAL_GROUP01_BACK
```

- Make sure you have an `.env` file that is matching the example

- Start application by typing

```bash
docker-compose up --build
```

## 🎨 Libraries

There are some external libraries used to build up this magnificent bot 🎖

[discordgo](https://gowalker.org/github.com/bwmarrin/discordgo)

## 🎯 Features


-  A bot that manages the game and plays the role of referee 👮
-  Multiplayer mode with friends on a server 🎎
-  Countdown timer for each round of the game ⌛

A web interface exists allowing ( accessible from the [repo](https://github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP01_FRONT)) :
- To have access to the statistics of the longest game ever played / or statistics of the last game played 📈
- To follow a game as a spectator 🍿


## 🥋 Contributing

Before contributing, please read the [contributing guidelines](https://github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP01_BACK/blob/main/COMMIT_CONVENTIONS.md)

## 👨‍💻 Development mode

1. Most of the time a ticket is assigned ,on the project board, to a team member. If it is not the case and you are told to do it yourself, assign the github ticket to you. 
2. When you start working on the ticket, move the concerned ticket to `In Progress`.
3. Create a branch specifically for this ticket with a name that follows the [conventions specified below](#branch-naming-convention).
4. Commit regularly at each significant step with unambiguous commit messages (see [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md) file).
5. Create a merge request that follows the [conventions specified below](#pull-requests-pr) to the develop branch.
6. On the project board, move the ticket to the status `In Review`
7. Request a review from another team member.
8. It may take some back and forth before your pull request is validated
9. Your pull request will then be merged into the develop branch and the concerned ticket will be moved to `Done`

## ❓ Any Questions

If you have any questions, feel free to open an issue. Please check the open issues before submitting a new one 😉



### 🛠 Continuous Integration (CI)

A CI pipeline is configured for this project.

The pipeline will run 3 different jobs:

- Dependencies check
- Linter
- Build

## 🏄‍♂️ Authors

<table align="center">
  <tr>
    <td align="center">
    <a href="https://github.com/myouuu">
      <img src="https://avatars.githubusercontent.com/u/60980138?v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Meriem MRABENT</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/gensjaak">
      <img src="https://avatars.githubusercontent.com/u/17094432?v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Jean-Jacques AKAKPO</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/FaycalTOURE">
      <img src="https://avatars.githubusercontent.com/u/19931625?v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Fayçal TOURÉ</b></sub>
    </a>
    </td>
    </td>
        <td align="center">
        <a href="https://github.com/acauchois">
          <img src="https://avatars.githubusercontent.com/u/15887111?v=4" width="100px;" alt=""/>
          <br />
          <sub><b>Alexis Cauchois</b></sub>
        </a>
        </td>
  </tr>
</table>
