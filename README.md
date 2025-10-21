# 📘 Astroid Shooter

> Astroid shooter game in go using raylib-go 

---

## 🧩 Overview

Game development combines logic, physics, and design to produce interactive experiences.
While high-performance engines like Unity (C#) or Unreal (C++) dominate, Go offers a compelling blend of simplicity and speed.
This project builds a retro-style asteroid shooter, focusing on game loops, collision detection, and rendering.

---

## 🎯 Objectives / Learning Goals

- 🔹 [Goal 1] — Learn Go development environment
- 🔹 [Goal 2] — Explore Go’s concurrency model and graphics libraries.  
- 🔹 [Goal 3] — Learn game loop structure (update, render, input).  

---

## ⚙️ Tech Stack

| Category | Tools / Languages |
|-----------|------------------|
| Language | `Go` |
| Libraries | `Raylib-go`|

---

## 🧪 Usage / Running the Project

### 🖥️ Setup

install go from: https://go.dev/dl/ - This is a go project so installing the language is essential.

install MSYS2 to get mingw-w64 https://www.msys2.org/ - this is needed for compiling the raylib-go library .dll file

After MSYS2 is installed, open the MSYS2 console and run the following command:

```nginx
pacman -Syu
```

update the base system and install the necessary dependencies, importantly the mingw-w64 package for compiling the raylib-go library.

After the dependencies are installed, close the window (if it hasnt automatically) and open  MSYS2 MinGW x64 terminal (important: NOT the plain MSYS one), then install GCC for 64-bit using the following command:

```nginx
pacman -Syu
pacman -S mingw-w64-x86_64-gcc
```

Once the install is complete, open the environment variables and add: 
```
C:\msys64\mingw64\bin to the path variable. 
```
(This will allow you to run the go command from the terminal.)
Restart the terminal/VSCode and check gcc has installed by running:
```
gcc --version`
```

With gcc now installed at in the path, you can run the go command from the terminal.

```
go env -w CGO_ENABLED=1
go env -w CC=gcc
```

This will point the go command to use the gcc compiler.
To check these were correctly set run:
```
go env CGO_ENABLED  # should be 1
go env CC           # should be gcc
```

install raylib-go https://github.com/gen2brain/raylib-go - raylib is a graphics library for go that is used for the graphical interface of this project.
```
go get -u github.com/gen2brain/raylib-go/raylib@latest
```


---

### ▶️ Run

```bash
go run main.go
```

---

## 📊 Results / Observations

- Key metrics or outcomes  
- Screenshots or performance graphs  
- Lessons learned, trade-offs, or insights  

---

## 🚀 Features / Implementation Plan

1. **Step 1 – Setup:** brief description (e.g., initialize repo, install deps)  
2. **Step 2 – Core Functionality:** what’s being built and how  
3. **Step 3 – Comparison / Benchmarking:** if relevant  
4. **Step 4 – Visualization / Output:** how results are displayed or tested  
5. **Step 5 – Deployment / Packaging:** optional (web, CLI, hardware, etc.)

> 📈 You can add diagrams or screenshots here  
> `![screenshot](docs/screenshot.png)`  

---

## 🔮 Future Improvements

- Add ship selection with differrent stats
- Add enemies
- Add powerups
- Add health bar

---

## 📚 References / Resources

- [Library or Framework Docs](https://github.com/gen2brain/raylib-go)
- [Tutorial or Blog Post](https://levelup.gitconnected.com/build-an-asteroids-game-with-raylib-go-4a92475b492c)
- [Assets](https://github.com/timlittle/blog-code/blob/main/go-asteroids/resources/space_background.png, https://github.com/timlittle/blog-code/blob/main/go-asteroids/resources/tilesheet.png)

---

## 🧑‍💻 Author

**Olly Bayley**  
GitHub: [@ombayley](https://github.com/ombayley)  

---

## 🪪 License

This project is licensed under the **GNU General Public License (GPL)** — See the [LICENSE](LICENSE) file for details.
The GPL License is a copyleft license, that requires any derivative work to also be released under the GPL License.
This means any derivative software that uses this code remains open-source and freely available to the public.

