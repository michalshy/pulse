#!/bin/bash

# Terminal 1: Docker Compose
gnome-terminal --title="Docker Compose" -- bash -c "sudo docker compose up; exec bash" &

# Terminal 2: Dashboard
gnome-terminal --title="Dashboard" -- bash -c "cd dashboard && npm run dev; exec bash" &

# Terminal 3: 
gnome-terminal --title="Backend" -- bash -c "cd backend && go run .; exec bash" &
