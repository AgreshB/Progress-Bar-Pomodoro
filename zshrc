
# Requires below to be installed
# - https://github.com/julienXX/terminal-notifier

# Mac setup for pomo
# to avoid using ./ before Progress build
export PATH=$PATH:.


alias work="Progress 60m && terminal-notifier -message 'Pomodoro'\
        -title 'Work Timer is up! Take a Break 😊'\
        -appIcon '~/Pictures/Wallpaper-4.jpg'\
        -sound Crystal"
        
alias rest="Progress 10m && terminal-notifier -message 'Pomodoro'\
        -title 'Break is over! Get back to work 😬'\
        -appIcon '~/Pictures/Wallpaper-4.jpg'\
        -sound Crystal"
        
alias pomo="Progress 25m && terminal-notifier -message 'Pomodoro'\
        -title 'Time to focus is up! Take a Break - $ sbreak for 5m and $ lbreak for 15m 😊'\
        -sound Crystal"
        
alias sbreak="Progress 5m && terminal-notifier -message 'Pomodoro'\
        -title 'Short break is over! Get back to work - $ pomo for 25m focus time 😬'\
        -appIcon '~/Pictures/Wallpaper-4.jpg'\
        -sound Crystal"

alias lbreak="timer 5m && terminal-notifier -message 'Pomodoro'\
        -title 'Long break is over! Get back to work - $ pomo for 25m focus time 😬'\
        -appIcon '~/Pictures/Wallpaper-4.jpg'\
        -sound Crystal"