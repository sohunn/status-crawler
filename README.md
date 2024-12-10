# <img src="https://user-images.githubusercontent.com/17984549/91302719-343a1d80-e7a7-11ea-8d6a-9448ef598420.png" height="25" /> StatusCrawler
This is a simple tool used to detect dead links on a website and summarize their HTTP statuses in a clear table, written in Golang.

## Features✨
- Supports and validates links using `http` and `https` schemes.
- Uses [playwright](https://pkg.go.dev/github.com/playwright-community/playwright-go) to perform efficient web scraping.
- Leverages the power of go-routines with mutexes, wait groups and distributed locking mechanisms to increase performance and concurrency 🚀
- Clean summary in a tabular format. 


## How to use❓
- Make sure you have the latest version of [go](https://go.dev/dl/) installed.

- Clone the repository using the following command:
```
git clone https://github.com/sohunn/status-crawler.git
```

- Install dependencies:
``` 
go mod tidy
```

- Make sure to install the browsers and OS dependencies:
```
go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps
```

- From the root of the project:
```
go run ./ <URL>
```

## Example
```
go run ./ "https://sohunn.me"
```

## Building 🛠️
Check your Go env variables (`GOOS` and `GOPATH`) to make sure you are building the executable for the right platform. Once verified, run:
```
go build -o crawler.exe ./
```

**Note:** You can call your executable whatever you want. I have specified `crawler` in the example

Once done, simply run the executable with the arguments like you normally would.

```
crawler.exe "https://sohunn.me"
```