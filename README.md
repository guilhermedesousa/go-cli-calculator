# 📌 CLI Calculator

A simple command-line calculator written in Go that supports **addition (`+`)**, **subtraction (`-`)**, **multiplication (`*`)**, and **division (`/`)**.

---

## Features

✅ Supports interactive and command-line modes  
✅ Handles invalid input and division by zero  
✅ Easy to use and extend  

---

## Installation & Usage

### 1️. Clone the Repository

```sh
git clone https://github.com/guilhermedesousa/go-cli-calculator
cd cli-calculator
```

### 2️. Run in Interactive Mode

```sh
go run main.go
```

#### Example:

```
Enter first number: 10
Enter an operator (+, -, *, /): *
Enter second number: 5
Result: 10.00 * 5.00 = 50.00
```

### 3️. Run with CLI Arguments

```sh
go run main.go 15 / 3
```

#### Example Output:

```
Result: 15.00 / 3.00 = 5.00
```