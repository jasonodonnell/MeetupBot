# Meetup Bot

Notifies Slack of Tech Lancaster meetups for the week.  Executed by a cron job 
on Monday mornings.

## Install

```bash
go get github.com/jasonodonnell/meetupbot
```

## Usage

```
export API_KEY='GOOGLE CALENDAR API KEY'
export CAL_ID='GOOGLE CALENDAR ID'
export SLACK_WEBHOOK='SLACK WEBHOOK URL'
meetupbot
```
