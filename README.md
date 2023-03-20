# Progress Bar for Pomodoro Timer
Create a Basic Progress bar that would count down and show progress till completion based on Input time.

Features :
- Animated loading of the bar.
- Label for timer

Functionality similar to sleep , except with a graphic. Aim to create a Pomdoro Timer for studying settings

## Usage and Build

First Build using the following command:
```sh
go build -o Progress
./Progress <duration>
./Progress -n <name> <duration>
```
Duration can have the following valid time units : "ns", "us" (or "µs"), "ms", "s", "m", "h".
Default duration : seconds ("s").

## For Pomodoro Timer on Mac

Make sure to First install terminal-notifier using:
```
brew install terminal-notifier
```
Then copy 'zshrc' to local MAC '.zshrc' for pomodoro commands