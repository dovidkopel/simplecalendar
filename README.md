# Simple Calendar

An attempt to make a standalone tool that enables third parties to schedule appointments with you by using a URL.
There are many of these services available and most of those with free options. 
I was convinced that I should be able to implement this functionality in a serverless fashion using only the Google Calendar for the most part.
Some small settings and confs may be stored in s3, but ultimately this tool has a few components.

* Describe complex schedules 
* Support inheritable sensible defaults
* Be able to specify negative ranges and make inferences from that
* Support multiple event types
* Check against at least one Google Calendar
* Store added events to Google Calendar

 