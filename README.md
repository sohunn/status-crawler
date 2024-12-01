# <img src="https://user-images.githubusercontent.com/17984549/91302719-343a1d80-e7a7-11ea-8d6a-9448ef598420.png" height="25" /> StatusCrawler
This is a simple tool used to detect dead links on a website and summarize their HTTP statuses in a clear table, written in Golang.

## Features
- Supports and validates links using `http` and `https` schemes.
- Uses [playwright](https://pkg.go.dev/github.com/playwright-community/playwright-go) to perform efficient web scraping.
- Leverages the power go-routines with mutexes, wait groups and distributed locking mechanisms to increase performance and concurrency. 🚀
- Clean summary in a tabular format. 


## How to use
- Make sure you have the latest version of [go](https://go.dev/dl/) installed.

- Clone the repository using the following command.
```
git clone https://github.com/sohunn/status-crawler.git
```

- Install dependencies.
``` 
go mod tidy
```

- From the root of the project
```
go run ./ <URL>
```

## Example
```
go run ./ "https://sohunn.me"
```