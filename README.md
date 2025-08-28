<p align="center"><h1 align="center">ECOM</h1></p>

<p align="center">
	<img src="https://img.shields.io/github/license/diegobbrito/ecom?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/diegobbrito/ecom?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/diegobbrito/ecom?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/diegobbrito/ecom?style=default&color=0080ff" alt="repo-language-count">
</p>
<br>

##  Table of Contents

- [ Overview](#-overview)
- [ Project Structure](#-project-structure)
  - [ Project Index](#-project-index)
- [ Getting Started](#-getting-started)
  - [ Prerequisites](#-prerequisites)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Testing](#-testing)
- [ License](#-license)

---

##  Overview

<code>❯ Project in Go</code>

---

##  Project Structure

```sh
└── ecom/
    ├── Dockerfile
    ├── LICENSE
    ├── Makefile
    ├── README.md
    ├── cmd
    │   ├── api
    │   ├── main.go
    │   └── migrate
    ├── config
    │   └── env.go
    ├── db
    │   └── db.go
    ├── docker-compose.yaml
    ├── go.mod
    ├── go.sum
    ├── service
    │   ├── auth
    │   ├── cart
    │   ├── order
    │   ├── product
    │   └── user
    ├── types
    │   └── types.go
    └── utils
        └── utils.go
```


###  Project Index
<details open>
	<summary><b><code>ECOM/</code></b></summary>
	<details> <!-- __root__ Submodule -->
		<summary><b>__root__</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/docker-compose.yaml'>docker-compose.yaml</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/go.mod'>go.mod</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/go.sum'>go.sum</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/Makefile'>Makefile</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/Dockerfile'>Dockerfile</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- types Submodule -->
		<summary><b>types</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/types/types.go'>types.go</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- config Submodule -->
		<summary><b>config</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/config/env.go'>env.go</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- utils Submodule -->
		<summary><b>utils</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/utils/utils.go'>utils.go</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- cmd Submodule -->
		<summary><b>cmd</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/cmd/main.go'>main.go</a></b></td>
			</tr>
			</table>
			<details>
				<summary><b>api</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/cmd/api/api.go'>api.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>migrate</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/cmd/migrate/main.go'>main.go</a></b></td>
					</tr>
					</table>
					<details>
						<summary><b>migrations</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/cmd/migrate/migrations/20250802224738_add-products-table.up.sql'>20250802224738_add-products-table.up.sql</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/cmd/migrate/migrations/20250802224759_add-orders-table.up.sql'>20250802224759_add-orders-table.up.sql</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/cmd/migrate/migrations/20250802224759_add-orders-table.down.sql'>20250802224759_add-orders-table.down.sql</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/cmd/migrate/migrations/20250801234913_add-user-table.up.sql'>20250801234913_add-user-table.up.sql</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/cmd/migrate/migrations/20250802225012_add-order-items-table.down.sql'>20250802225012_add-order-items-table.down.sql</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/cmd/migrate/migrations/20250802224738_add-products-table.down.sql'>20250802224738_add-products-table.down.sql</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/cmd/migrate/migrations/20250801234913_add-user-table.down.sql'>20250801234913_add-user-table.down.sql</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/cmd/migrate/migrations/20250802225012_add-order-items-table.up.sql'>20250802225012_add-order-items-table.up.sql</a></b></td>
							</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<details> <!-- service Submodule -->
		<summary><b>service</b></summary>
		<blockquote>
			<details>
				<summary><b>cart</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/service/cart/routes.go'>routes.go</a></b></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/service/cart/service.go'>service.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>auth</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/service/auth/jwt.go'>jwt.go</a></b></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/service/auth/password.go'>password.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>order</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/service/order/store.go'>store.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>product</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/service/product/routes.go'>routes.go</a></b></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/service/product/store.go'>store.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>user</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/service/user/routes.go'>routes.go</a></b></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/service/user/store.go'>store.go</a></b></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/service/user/routes_test.go'>routes_test.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<details> <!-- db Submodule -->
		<summary><b>db</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/ecom/blob/master/db/db.go'>db.go</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
</details>

---
##  Getting Started

###  Prerequisites

Before getting started with ecom, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go
- **Package Manager:** Go modules
- **Container Runtime:** Docker


###  Installation

Install ecom using one of the following methods:

**Build from source:**

1. Clone the ecom repository:
```sh
❯ git clone https://github.com/diegobbrito/ecom
```

2. Navigate to the project directory:
```sh
❯ cd ecom
```

3. Install the project dependencies:


**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go build
```


**Using `docker`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Docker-2CA5E0.svg?style={badge_style}&logo=docker&logoColor=white" />](https://www.docker.com/)

```sh
❯ docker build -t diegobbrito/ecom .
```




###  Usage
Run ecom using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run {entrypoint}
```


**Using `docker`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Docker-2CA5E0.svg?style={badge_style}&logo=docker&logoColor=white" />](https://www.docker.com/)

```sh
❯ docker run -it {image_name}
```


###  Testing
Run the test suite using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go test ./...
```
---

##  License

This project is protected under the [SELECT-A-LICENSE](https://choosealicense.com/licenses) License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/) file.

---