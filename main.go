package main

import "io.wizdumb/simplecalendar"

/**
Based on configuration enable remote user to select
events and they will be added to the user's calendar

--Event Types--
Display the event types that are available

--Free/Busy--
- What is the general daily schedule?

--Schedule Event--
- Given a selected EventType display the next available time slots
- If a date is specified begin the time slots at the date specified
- Else, find earliest time
*/
func main() {
	simplecalendar.Main()
}

