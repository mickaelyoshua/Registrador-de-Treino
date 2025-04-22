#!/bin/bash

tmux new-session -d -t register
tmux new-window -n containers
tmux a -t register
