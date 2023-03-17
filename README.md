# Progress Bar for Pomodoro Timer
Create a Basic Progress bar that would count down and show progress till completion based on Input time.

Features :
- Animated loading of the bar.
- Label for timer

Functionality similar to sleep , except with a graphic. Aim to create a Pomdoro Timer for studying settings

## Usage

```sh
go run main.go <duration>
go run main.go -n <name> <duration>
```

It is possible to pass a time unit for `<duration>`.

Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
If no unit is passed, it defaults to seconds ("s").