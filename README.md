# Expense Tracker App

An expense tracker application with a powerful backend built using NestJS, Postgres, and a dynamic frontend crafted with Go, HTMX, Tailwind CSS, and Go ECharts.

## Features

- **Create Reports**: Easily create reports for both expenses and income.
- **Edit Reports**: Modify existing reports to keep your data accurate.
- **Delete Reports**: Remove reports that are no longer needed.
- **View Graphs**: Analyze your reports with interactive and insightful graphs.

## Tech Stack

### Backend

- **NestJS**: A progressive Node.js framework for building efficient and scalable server-side applications.
- **PostgreSQL**: A powerful, open source object-relational database system.
- **Redis**: An in-memory data structure store, used for rate limiting.
- **Docker**: Containerization platform for easy deployment.

### Frontend

- **Go**: Used for building the frontend application logic.
- **HTMX**: For dynamic HTML without using JavaScript.
- **Tailwind CSS**: A utility-first CSS framework for rapid UI development.
- **Go ECharts**: A Go library to integrate Apache ECharts for interactive charting and visualization.

## Installation

### Backend

1. Clone the repository:

   ```sh
   git clone https://github.com/Anuolu-2020/Expense-Calculator-App.git
   cd expense-tracker-app/server
   ```

2. Install dependencies:

   ```sh
   npm install
   ```

3. Configure the environment variables:

   ```sh
   cp .env.example .env
   ```

4. Start the server:
   ```sh
   npm run start
   ```

### Frontend

1. Navigate to the frontend directory:

   ```sh
   cd ../client
   ```

2. Install dependencies:

   ```sh
   go get
   ```

3. Start the frontend application:
   ```sh
   go run server.go
   ```

## Usage

1. Open your browser and navigate to `http://localhost:3000` for the backend and `http://localhost:8080` for the frontend.
2. Use the interface to create, edit, delete, and view reports.
3. Analyze your financial data with the interactive graphs provided by Go ECharts.

## Live Demo

Check out the live version of the app [here](https://expense-tracker-app-u3za.onrender.com).

## Example Graph

Below is an example of the interactive graph you can generate with this app:

![Example Graph](./assests/report-graph.png)

## Technologies

- [NestJS](https://nestjs.com/)
- [Go](https://golang.org/)
- [HTMX](https://htmx.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [Go ECharts](https://github.com/go-echarts/go-echarts)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)
- [Docker](https://www.docker.com/)

---

Feel free to explore and enjoy using this expense tracker app!
